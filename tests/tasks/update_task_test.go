package tasks

import (
	"fmt"
	"go-task-app/internal/config"
	tasksTypes "go-task-app/internal/tasks/types"
	"go-task-app/tests/helpers"
	tasksTestHelpers "go-task-app/tests/tasks/helpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestUpdateTask(accessToken string, taskId int, input *tasksTypes.TaskInput) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/task/%d", taskId), tasksTestHelpers.CreateReaderFromStruct(input))

	return tasksTestHelpers.Request(req, accessToken)
}
func TestUpdateTask(t *testing.T) {
	helpers.InitIntegrationTest()
	accessTokenUser1, accessTokenUser2 := tasksTestHelpers.SetUpUsers()
	tasksTestHelpers.RequestCreateTask(accessTokenUser1, &tasksTypes.TaskInput{
		Title: "タスク1",
	})
	tasksTestHelpers.RequestCreateTask(accessTokenUser2, &tasksTypes.TaskInput{
		Title: "タスク2",
	})

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
