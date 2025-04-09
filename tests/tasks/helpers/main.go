package helpers

import (
	"encoding/json"
	tasksTypes "go-task-app/internal/tasks/types"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/globals"
	"go-task-app/tests/helpers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func CreateReaderFromStruct(arg any) io.Reader {
	jsonData, _ := json.Marshal(arg)
	return strings.NewReader(string(jsonData))
}

func Request(req *http.Request, accessToken string) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()

	helpers.SetAuthHeader(req, accessToken)
	globals.Router.ServeHTTP(httpRecorder, req)

	return httpRecorder
}

func RequestCreateTask(accessToken string, input *tasksTypes.TaskInput) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/task", CreateReaderFromStruct(input))

	return Request(req, accessToken)
}
