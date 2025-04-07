package services

import (
	"errors"
	"go-task-app/internal/config"
	"go-task-app/internal/users/constants"
	"go-task-app/internal/users/types"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type keyLabelStruct struct {
	Name     string
	Password string
}

var keyLabel = keyLabelStruct{Name: "名前", Password: "パスワード"}

func extractEmptyData(signUpInput *types.SignUpInput) []string {
	var emptySlices []string
	if signUpInput.Name == "" {
		emptySlices = append(emptySlices, keyLabel.Name)
	}

	if signUpInput.Password == "" {
		emptySlices = append(emptySlices, keyLabel.Password)
	}

	return emptySlices
}

func ValidatePassword(password string) error {
	if len(password) < constants.MinPasswordLength {
		return constants.ErrLessPasswordLength
	}

	if len(password) > constants.MaxPasswordLength {
		return constants.ErrOverPasswordLength
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

func existsDuplicatedUserName(name string) bool {
	err := config.DB.Where("name = ?", name).First(&types.User{}).Error

	return err == nil
}

func hashPassword(signUpInput *types.SignUpInput) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpInput.Password), constants.HashCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func SignUp(signUpInput *types.SignUpInput) (*types.UserResponse, error) {
	emptySlices := extractEmptyData(signUpInput)

	if len(emptySlices) > 0 {
		return nil, errors.New(strings.Join(emptySlices, "、") + constants.ErrSuffixRequiredInput)
	}

	if err := ValidatePassword(signUpInput.Password); err != nil {
		return nil, err
	}

	if existsDuplicatedUserName(signUpInput.Name) {
		return nil, constants.ErrDuplicatedUserName
	}

	hashedPassword, err := hashPassword(signUpInput)

	if err != nil {
		return nil, err
	}

	var newUser types.User

	newUser.Name = signUpInput.Name
	newUser.Password = hashedPassword

	result := config.DB.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &types.UserResponse{UserBase: newUser.UserBase}, nil
}
