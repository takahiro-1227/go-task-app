package services

import (
	"errors"
	"go-task-app/internal/config"
	"go-task-app/internal/users/constants"
	"go-task-app/internal/users/types"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

type keyLabelStruct struct {
	Name     string
	Password string
}

var keyLabel = keyLabelStruct{Name: "名前", Password: "パスワード"}

func extractEmptyData(newUser *types.User) []string {
	var emptySlices []string
	if newUser.Name == "" {
		emptySlices = append(emptySlices, keyLabel.Name)
	}

	if newUser.Password == "" {
		emptySlices = append(emptySlices, keyLabel.Password)
	}

	return emptySlices
}

func validatePassword(password string) error {
	if len(password) <= constants.MinPasswordLength {
		return constants.ErrMorePasswordLength
	}

	if len(password) >= constants.MaxPasswordLength {
		return constants.ErrLessPasswordLength
	}

	hasLetter := false
	hasNumber := false
	hasSymbol := false
	symbols := "!@#$%^&*()_+{}[]:;\"'<>,.?/~`\\-"

	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasNumber = true
		case strings.ContainsRune(symbols, char):
			hasSymbol = true
		}
	}

	if !hasLetter || !hasNumber || !hasSymbol {
		return constants.ErrPasswordCharacterCategory
	}

	return nil
}

func existsDuplicatedUserName(newUser *types.User) bool {
	err := config.DB.Where("name = ?", newUser.Name).First(&types.User{}).Error

	return err == nil
}

func hashPassword(newUser *types.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), constants.HashCost)

	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)

	return nil
}

func SignUp(newUser *types.User) (*types.UserResponse, error) {
	emptySlices := extractEmptyData(newUser)

	if len(emptySlices) > 0 {
		return nil, errors.New(strings.Join(emptySlices, "、") + constants.ErrSuffixRequiredInput)
	}

	if err := validatePassword(newUser.Password); err != nil {
		return nil, err
	}

	if existsDuplicatedUserName(newUser) {
		return nil, constants.ErrDuplicatedUserName
	}

	err := hashPassword(newUser)

	if err != nil {
		return nil, err
	}

	result := config.DB.Create(newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &types.UserResponse{UserBase: newUser.UserBase}, nil
}
