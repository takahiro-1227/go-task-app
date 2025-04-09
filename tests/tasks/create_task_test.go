package tasks

import (
	"go-task-app/internal/config"
	tasksTypes "go-task-app/internal/tasks/types"
	"go-task-app/tests/helpers"
	tasksTestHelpers "go-task-app/tests/tasks/helpers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	helpers.InitIntegrationTest()
	accessTokenUser1, _ := tasksTestHelpers.SetUpUsers()

	httpRecorder := tasksTestHelpers.RequestCreateTask(accessTokenUser1, &tasksTypes.TaskInput{
		Title: "タスク1",
	})

	assert.Equal(t, http.StatusCreated, httpRecorder.Code)

	var tasks []tasksTypes.Task
	config.DB.Find(&tasks)
	assert.Equal(t, tasks[0].Title, "タスク1")
}
