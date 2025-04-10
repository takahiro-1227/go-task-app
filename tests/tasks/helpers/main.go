package helpers

import (
	tasksTypes "go-task-app/internal/tasks/types"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
)

func SetUpUsers() (string, string) {
	initUserResponse := helpers.InitUser(&usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	})
	accessTokenUser1 := initUserResponse.AccessToken

	initUserResponse = helpers.InitUser(&usersTypes.SignUpInput{
		Name:     "test2",
		Password: "test5678---3A",
	})
	accessTokenUser2 := initUserResponse.AccessToken

	return accessTokenUser1, accessTokenUser2
}

func RequestCreateTask(accessToken string, input *tasksTypes.TaskInput) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/task", helpers.CreateReaderFromStruct(input))

	return helpers.Request(req, &accessToken)
}
