package tasks

import (
	"encoding/json"
	"fmt"
	"go-task-app/internal/config"
	tasksTypes "go-task-app/internal/tasks/types"
	usersTypes "go-task-app/internal/users/types"
	"go-task-app/tests/globals"
	"go-task-app/tests/helpers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var accessTokenUser1 string
var accessTokenUser2 string

func setUpUsers() {
	initUserResponse := helpers.InitUser(&usersTypes.SignUpInput{
		Name:     "test1",
		Password: "test1234---2A",
	})
	accessTokenUser1 = initUserResponse.AccessToken

	initUserResponse = helpers.InitUser(&usersTypes.SignUpInput{
		Name:     "test2",
		Password: "test5678---3A",
	})
	accessTokenUser2 = initUserResponse.AccessToken
}

func createReaderFromStruct(arg any) io.Reader {
	jsonData, _ := json.Marshal(arg)
	return strings.NewReader(string(jsonData))
}

func TestCreateTask(t *testing.T) {
	helpers.InitTest()
	setUpUsers()

	httpRecorder := httptest.NewRecorder()

	createTaskInput := &tasksTypes.TaskInput{
		Title: "タスク1",
	}

	req, _ := http.NewRequest(http.MethodPost, "/task", createReaderFromStruct(createTaskInput))

	helpers.SetAuthHeader(req, accessTokenUser1)

	globals.Router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var tasks []tasksTypes.Task
	config.DB.Find(&tasks)
	assert.Equal(t, tasks[0].Title, "タスク1")
}

func createTaskUser2() {
	httpRecorder := httptest.NewRecorder()
	createTaskInput := &tasksTypes.TaskInput{
		Title: "タスク2",
	}

	req, _ := http.NewRequest(http.MethodPost, "/task", createReaderFromStruct(createTaskInput))
	helpers.SetAuthHeader(req, accessTokenUser2)
	globals.Router.ServeHTTP(httpRecorder, req)

	if httpRecorder.Code == http.StatusCreated {
		return
	}

	panic(httpRecorder)
}

type GetTasksResponse struct {
	Tasks []tasksTypes.Task `json:"tasks"`
}

func requestGetUsers(accessToken string) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", strings.NewReader(""))

	helpers.SetAuthHeader(req, accessToken)

	globals.Router.ServeHTTP(httpRecorder, req)

	return httpRecorder
}

func decodeGetUsersResponse(httpRecorder *httptest.ResponseRecorder) *GetTasksResponse {
	var tasksResponse GetTasksResponse
	err := json.NewDecoder(httpRecorder.Body).Decode(&tasksResponse)
	if err != nil {
		panic(err)
	}

	return &tasksResponse
}

func TestGetTasksUser1(t *testing.T) {
	createTaskUser2()

	httpRecorder := requestGetUsers(accessTokenUser1)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse := decodeGetUsersResponse(httpRecorder)

	assert.Equal(t, len(tasksResponse.Tasks), 1)
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク1")
}

func TestGetTasksUser2(t *testing.T) {
	httpRecorder := requestGetUsers(accessTokenUser2)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse := decodeGetUsersResponse(httpRecorder)

	assert.Equal(t, len(tasksResponse.Tasks), 1)
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク2")
}

func requestUpdateTask(accessToken string, taskId int, input *tasksTypes.TaskInput) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/task/%d", taskId), createReaderFromStruct(input))

	helpers.SetAuthHeader(req, accessToken)

	globals.Router.ServeHTTP(httpRecorder, req)

	return httpRecorder
}

func TestUpdateTask(t *testing.T) {
	httpRecorder := requestUpdateTask(accessTokenUser1, 1, &tasksTypes.TaskInput{
		Title: "タスク2",
	})

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	var task tasksTypes.Task
	config.DB.Where("id = ?", 1).First(&task)
	assert.Equal(t, task.Title, "タスク2")

	httpRecorder = requestUpdateTask(accessTokenUser2, 1, &tasksTypes.TaskInput{
		Title: "タスク1",
	})

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)

	config.DB.Where("id = ?", 1).First(&task)
	assert.Equal(t, task.Title, "タスク2")
}

func requestDeleteTask(accessToken string, taskId int) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/task/%d", taskId), strings.NewReader(""))

	helpers.SetAuthHeader(req, accessToken)

	globals.Router.ServeHTTP(httpRecorder, req)

	return httpRecorder
}
func TestDeleteTask(t *testing.T) {
	httpRecorder := requestDeleteTask(accessTokenUser2, 1)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)

	var task tasksTypes.Task
	config.DB.Where("id = ?", 1).First(&task)
	assert.Equal(t, task.Title, "タスク2")

	httpRecorder = requestDeleteTask(accessTokenUser1, 1)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	res := config.DB.Where("id = ?", 1).First(&task)
	assert.Equal(t, res.Error.Error(), "record not found")
}
