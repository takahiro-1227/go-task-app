package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/types"
	"time"
)

func UpdateTask(newTask *types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	if !helpers.IsTaskOwnedByUser(newTask.ID, newTask.UserId) {
		return nil, constants.ErrInvalidUpdate
	}

	result := config.DB.Where("id = ?", newTask.ID).Model(&types.Task{}).Updates(types.Task{Title: newTask.Title, UpdatedAt: time.Now()})

	if result.Error != nil {
		return nil, constants.ErrUpdateFailed
	}

	return newTask, nil
}
