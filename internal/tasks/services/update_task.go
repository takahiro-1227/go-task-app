package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/types"
	"log"
	"time"
)

func UpdateTask(newTask *types.Task) (*types.Task, error) {
	if newTask.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	err := config.DB.Where("title = ?", newTask.Title).First(&types.Task{}).Error
	if err == nil {
		return nil, constants.ErrDuplicatedTitle
	}

	result := config.DB.Where("id = ?", newTask.ID).Model(&types.Task{}).Updates(types.Task{Title: newTask.Title, UpdatedAt: time.Now()})
	if result.Error != nil {
		log.Println(result.Error)
		return nil, constants.ErrCreateFailed
	}

	return newTask, nil
}
