package types

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"type:varchar(255); unique; not null" json:"title"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserId    int       `gorm:"not null" json:"user_id"`
}

type TaskHandlerInput struct {
	Title string `json:"title"`
}

type CreateTaskServiceInput struct {
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}

type UpdateTaskServiceInput struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
}

type TaskServiceResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int       `json:"user_id"`
}
