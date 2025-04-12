package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/types"
)

func CreateTask(taskInput *types.CreateTaskServiceInput) (*types.Task, error) {
	if taskInput.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	var newTask types.Task

	newTask.Title = taskInput.Title
	newTask.UserId = taskInput.UserId

	result := config.DB.Create(&newTask)
	if result.Error != nil {
		return nil, constants.ErrCreateFailed
	}

	return &newTask, nil
}
