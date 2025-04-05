package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/types"

	"log"
)

func DeleteTask(taskId uint, userId int) error {
	if !helpers.IsTaskOwnedByUser(taskId, userId) {
		return constants.ErrInvalidDelete
	}

	result := config.DB.Where("id = ?", taskId).Delete(&types.Task{})

	if result.Error != nil {
		log.Println(result.Error)
		return constants.ErrDeleteFailed
	}

	return nil
}
