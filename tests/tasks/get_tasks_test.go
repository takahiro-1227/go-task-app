package tasks

import (
	"encoding/json"
	tasksTypes "go-task-app/internal/tasks/types"
	"go-task-app/tests/helpers"
	tasksTestHelpers "go-task-app/tests/tasks/helpers"
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

	return helpers.Request(req, &accessToken)
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
	accessTokenUser1, accessTokenUser2 := tasksTestHelpers.SetUpUsers()
	tasksTestHelpers.RequestCreateTask(accessTokenUser1, &tasksTypes.TaskHandlerInput{
		Title: "タスク1",
	})
	tasksTestHelpers.RequestCreateTask(accessTokenUser2, &tasksTypes.TaskHandlerInput{
		Title: "タスク2",
	})

	httpRecorder := requestGetUsers(accessTokenUser1)
	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse := decodeGetUsersResponse(httpRecorder)
	assert.Equal(t, 1, len(tasksResponse.Tasks))
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク1")

	httpRecorder = requestGetUsers(accessTokenUser2)
	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	tasksResponse = decodeGetUsersResponse(httpRecorder)
	assert.Equal(t, 1, len(tasksResponse.Tasks))
	assert.Equal(t, tasksResponse.Tasks[0].Title, "タスク2")
}
