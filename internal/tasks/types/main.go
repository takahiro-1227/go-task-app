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

type TaskInput struct {
	ID        uint
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
