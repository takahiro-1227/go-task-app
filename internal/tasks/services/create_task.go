package services

import (
	"go-task-app/internal/tasks/constants"
	"go-task-app/internal/tasks/repo"
	"go-task-app/internal/tasks/types"
)

type Repo struct {
	CreateTask repo.CreateTaskFunc
}

func CreateTask(taskInput *types.CreateTaskServiceInput, repoInput *Repo) (*types.Task, error) {
	if taskInput.Title == "" {
		return nil, constants.ErrTitleIsEmpty
	}

	newTask, err := repoInput.CreateTask(&types.CreateTaskRepoInput{
		Title:  taskInput.Title,
		UserId: taskInput.UserId,
	})

	if err != nil {
		return nil, constants.ErrCreateFailed
	}

	return newTask, nil
}
