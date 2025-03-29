package services

import (
	"errors"
	"go-task-app/internal/config"
	"go-task-app/internal/users/types"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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
		emptySlices = append(emptySlices, "パスワード")
	}

	return emptySlices
}

func validatePassword(password string) error {
	if len(password) < 10 {
		return errors.New("パスワードは10文字以上で入力してください。")
	}

	if len(password) > 33 {
		return errors.New("パスワードは32文字以内で入力してください。")
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
		return errors.New("パスワードは英数字記号を含める必要があります。")
	}

	return nil
}

func existsDuplicatedUserName(newUser *types.User) bool {
	err := config.DB.Where("name = ?", newUser.Name).First(&types.User{}).Error

	return err == nil
}

func hashPassword(newUser *types.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)

	return nil
}

func SignUp(newUser *types.User) (*types.UserResponse, int, error) {
	emptySlices := extractEmptyData(newUser)

	if len(emptySlices) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(emptySlices, ", ") + "を入力してください。")
	}

	if err := validatePassword(newUser.Password); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if existsDuplicatedUserName(newUser) {
		return nil, http.StatusBadRequest, errors.New("入力されたユーザー名は既に存在します。")
	}

	err := hashPassword(newUser)

	if err != nil {
		log.Println(err.Error())
		return nil, http.StatusInternalServerError, errors.New("エラーが発生しました。")
	}

	result := config.DB.Create(newUser)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, http.StatusInternalServerError, errors.New("ユーザーの作成に失敗しました。")
	}

	return &types.UserResponse{UserBase: newUser.UserBase}, http.StatusCreated, nil
}
