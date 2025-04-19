package repo

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"
)

type CreateTaskFunc func(taskInput *types.CreateTaskRepoInput) (*types.Task, error)

var CreateTask CreateTaskFunc = func(taskInput *types.CreateTaskRepoInput) (*types.Task, error) {
	var newTask types.Task

	newTask.Title = taskInput.Title
	newTask.UserId = taskInput.UserId

	result := config.DB.Create(&newTask)

	return &newTask, result.Error
}
