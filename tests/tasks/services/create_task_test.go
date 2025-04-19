package services

import (
	"go-task-app/internal/tasks/constants"
	tasksRepo "go-task-app/internal/tasks/repo"
	tasksServices "go-task-app/internal/tasks/services"
	tasksTypes "go-task-app/internal/tasks/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fixedTime = time.Date(2024, 4, 1, 12, 0, 0, 0, time.UTC)

var createTaskFunc tasksRepo.CreateTaskFunc = func(taskInput *tasksTypes.CreateTaskRepoInput) (*tasksTypes.Task, error) {
	return &tasksTypes.Task{
		ID:        1,
		Title:     taskInput.Title,
		UserId:    taskInput.UserId,
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}, nil
}

var createTaskRepo = &tasksServices.Repo{
	CreateTask: createTaskFunc,
}

func TestCreateTask(t *testing.T) {
	result, _ := tasksServices.CreateTask(&tasksTypes.CreateTaskServiceInput{
		Title:  "タスク1",
		UserId: 1,
	}, createTaskRepo)

	assert.Equal(t, result, &tasksTypes.Task{
		ID:        1,
		Title:     "タスク1",
		UserId:    1,
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	})
}

func TestCreateTaskEmptyTitle(t *testing.T) {
	result, err := tasksServices.CreateTask(&tasksTypes.CreateTaskServiceInput{
		Title:  "",
		UserId: 1,
	}, createTaskRepo)

	assert.Nil(t, result)
	assert.Equal(t, err, constants.ErrTitleIsEmpty)
}
