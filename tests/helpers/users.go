package helpers

import (
	usersServices "go-task-app/internal/users/services"
	usersTypes "go-task-app/internal/users/types"
	"log"
)

func InitUser(input *usersTypes.SignUpInput) *usersTypes.SignInResponse {
	_, err := usersServices.SignUp(input)

	if err != nil {
		panic(err)
	}

	res, err := usersServices.SignIn(usersTypes.SignInInput{
		Name:     input.Name,
		Password: input.Password,
	})

	if err != nil {
		panic(err)
	}

	log.Println("SignIn!")

	return res
}
