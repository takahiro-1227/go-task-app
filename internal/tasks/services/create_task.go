package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"

	"log"
)

func CreateTaskService(newTask types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, types.ErrInvalidInput
	}

	err := config.DB.Where("title = ?", newTask.Title).First(&types.Task{}).Error
	if err == nil {
		return nil, types.ErrDuplicateTitle
	}

	result := config.DB.Create(&newTask)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, types.ErrCreateFailed
	}

	return &newTask, nil
}
