package tasks

import (
	"encoding/json"
	tasksTypes "go-task-app/internal/tasks/types"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetTasksResponse struct {
	Tasks []tasksTypes.Task `json:"tasks"`
}

func requestGetUsers(accessToken string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", strings.NewReader(""))

	return Request(req, accessToken)
}

func decodeGetUsersResponse(httpRecorder *httptest.ResponseRecorder) *GetTasksResponse {
	var tasksResponse GetTasksResponse
	err := json.NewDecoder(httpRecorder.Body).Decode(&tasksResponse)
	if err != nil {
		panic(err)
	}

	return &tasksResponse
}
func TestGetTasks(t *testing.T) {
	helpers.InitIntegrationTest()
	accessTokenUser1, accessTokenUser2 := SetUpUsers()

	RequestCreateTask(accessTokenUser1, &tasksTypes.TaskInput{
		Title: "タスク1",
	})
	RequestCreateTask(accessTokenUser1, &tasksTypes.TaskInput{
		Title: "タスク2",
	})

	httpRecorder := requestGetUsers(accessTokenUser1)
	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse := decodeGetUsersResponse(httpRecorder)
	assert.Equal(t, len(tasksResponse.Tasks), 1)
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク1")

	httpRecorder = requestGetUsers(accessTokenUser2)
	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse = decodeGetUsersResponse(httpRecorder)
	assert.Equal(t, len(tasksResponse.Tasks), 1)
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク2")
}
