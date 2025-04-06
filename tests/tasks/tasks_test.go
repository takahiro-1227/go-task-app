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

var accessToken string

func TestCreateTask(t *testing.T) {
	helpers.InitTest()

	signUpInput := &usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	}

	initUserResponse := helpers.InitUser(signUpInput)

	accessToken = initUserResponse.AccessToken

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

type GetTasksResponse struct {
	Tasks []tasksTypes.Task `json:"tasks"`
}

func TestGetTasks(t *testing.T) {
	httpRecorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", strings.NewReader(""))

	req.Header.Set("Authorization", "Bearer "+accessToken)

	globals.Router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	var tasksResponse GetTasksResponse

	err := json.NewDecoder(httpRecorder.Body).Decode(&tasksResponse)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, len(tasksResponse.Tasks), 1)
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク1")
}
