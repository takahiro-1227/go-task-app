package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/types"

	"log"
)

func CreateTask(newTask *types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	err := config.DB.Where("title = ?", newTask.Title).First(&types.Task{}).Error
	if err == nil {
		return nil, constants.ErrDuplicatedTitle
	}

	result := config.DB.Create(&newTask)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, constants.ErrCreateFailed
	}

	return newTask, nil
}
