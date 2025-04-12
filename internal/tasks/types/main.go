package types

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(255); unique; not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UserId    int       `gorm:"not null"`
}

type TaskHandlerInput struct {
	Title string
}

type CreateTaskServiceInput struct {
	Title  string
	UserId int
}

type UpdateTaskServiceInput struct {
	ID     uint
	Title  string
	UserId int
}

type TaskServiceResponse struct {
	ID        uint
	Title     string
	UpdatedAt time.Time
	UserId    int
}
