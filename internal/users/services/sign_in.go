package services

import (
	"errors"
	"go-task-app/internal/config"
	"go-task-app/internal/users/constants"
	"go-task-app/internal/users/types"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func checkEmptyErr(signInInput types.SignInInput) error {
	var emptyLabels []string

	if len(signInInput.Name) == 0 {
		emptyLabels = append(emptyLabels, constants.UserLabels.Name)
	}

	if len(signInInput.Password) == 0 {
		emptyLabels = append(emptyLabels, constants.UserLabels.Password)
	}

	if len(emptyLabels) == 0 {
		return nil
	}

	return errors.New(strings.Join(emptyLabels, "、") + "の入力は必須です。")
}

func SignIn(signInInput types.SignInInput) (*types.SignInResponse, error) {
	err := checkEmptyErr(signInInput)

	if err != nil {
		return nil, err
	}

	var user types.User

	result := config.DB.Where("name = ?", signInInput.Name).First(&user)

	if result.Error != nil {
		return nil, constants.ErrSignIn
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInInput.Password))

	if err != nil {
		return nil, constants.ErrSignIn
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    user.ID,
		"expiresIn": time.Now().Add(time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.AuthSecret))

	if err != nil {
		return nil, err
	}

	userResponse := types.UserResponse{
		UserBase: types.UserBase{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		},
	}

	return &types.SignInResponse{
		AccessToken: signedToken,
		User:        userResponse,
	}, nil
}
