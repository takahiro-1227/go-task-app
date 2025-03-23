package services

import (
	"errors"
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"

	"log"
)

func DeleteTask(id int) error {
	result := config.DB.Where("id = ?", id).Delete(&types.Task{})

	if result.Error != nil {
		log.Println(result.Error)
		return errors.New("タスクの削除に失敗しました")
	}

	return nil
}
