package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/types"
)

func CreateTask(newTask *types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	result := config.DB.Create(&newTask)
	if result.Error != nil {
		return nil, constants.ErrCreateFailed
	}

	return newTask, nil
}
