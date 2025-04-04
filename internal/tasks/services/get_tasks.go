package services

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/types"
	"log"
)

func GetTasks() (*[]types.Task, error) {
	var tasks []types.Task

	result := config.DB.Find(&tasks)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, constants.ErrGetTasks
	}

	return &tasks, nil
}
