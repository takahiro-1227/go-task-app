package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"
	"time"

	"log"
)

func UpdateTask(newTask *types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, types.ErrInvalidInput
	}

	err := config.DB.Where("title = ?", newTask.Title).First(&types.Task{}).Error
	if err == nil {
		return nil, types.ErrDuplicateTitle
	}

	result := config.DB.Where("id = ?", newTask.ID).Model(&types.Task{}).Updates(types.Task{Title: newTask.Title, UpdatedAt: time.Now()})
	if result.Error != nil {
		log.Println(result.Error)
		return nil, types.ErrCreateFailed
	}

	return newTask, nil
}
