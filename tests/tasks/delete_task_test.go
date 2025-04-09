package tasks

import (
	"fmt"
	"go-task-app/internal/config"
	tasksTypes "go-task-app/internal/tasks/types"
	"go-task-app/tests/helpers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestDeleteTask(accessToken string, taskId int) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/task/%d", taskId), strings.NewReader(""))

	return Request(req, accessToken)
}
func TestDeleteTask(t *testing.T) {
	helpers.InitIntegrationTest()
	accessTokenUser1, accessTokenUser2 := SetUpUsers()
	RequestCreateTask(accessTokenUser1, &tasksTypes.TaskInput{
		Title: "タスク1",
	})
	RequestCreateTask(accessTokenUser2, &tasksTypes.TaskInput{
		Title: "タスク2",
	})

	httpRecorder := requestDeleteTask(accessTokenUser2, 1)

	assert.Equal(t, http.StatusForbidden, httpRecorder.Code)

	var task tasksTypes.Task
	config.DB.Where("id = ?", 1).First(&task)

	assert.Equal(t, task.Title, "タスク1")

	httpRecorder = requestDeleteTask(accessTokenUser1, 1)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)

	res := config.DB.Where("id = ?", 1).First(&task)

	assert.Equal(t, res.Error.Error(), "record not found")
}
