package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/types"
	"time"
)

func UpdateTask(taskInput *types.UpdateTaskServiceInput) (*types.TaskServiceResponse, error) {
	if taskInput.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	if !helpers.IsTaskOwnedByUser(taskInput.ID, taskInput.UserId) {
		return nil, constants.ErrInvalidUpdate
	}

	updatedAt := time.Now()

	result := config.DB.Where("id = ?", taskInput.ID).Model(&types.Task{}).Updates(types.Task{Title: taskInput.Title, UpdatedAt: updatedAt})

	if result.Error != nil {
		return nil, constants.ErrUpdateFailed
	}

	return &types.TaskServiceResponse{
		ID:        taskInput.ID,
		Title:     taskInput.Title,
		UpdatedAt: updatedAt,
		UserId:    taskInput.UserId,
	}, nil
}
