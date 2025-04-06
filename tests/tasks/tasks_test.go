package tasks

import (
	"encoding/json"
	"go-task-app/internal/config"
	tasksTypes "go-task-app/internal/tasks/types"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/globals"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	helpers.InitTest()

	signUpInput := &usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	}

	initUserResponse := helpers.InitUser(signUpInput)

	accessToken := initUserResponse.AccessToken

	httpRecorder := httptest.NewRecorder()

	createTaskInput := &tasksTypes.TaskInput{
		Title: "タスク1",
	}

	createTaskInputJson, _ := json.Marshal(createTaskInput)

	req, _ := http.NewRequest(http.MethodPost, "/task", strings.NewReader(string(createTaskInputJson)))

	req.Header.Set("Authorization", "Bearer "+accessToken)

	globals.Router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var tasks []tasksTypes.Task

	config.DB.Find(&tasks)

	assert.Equal(t, tasks[0].Title, "タスク1")
}
