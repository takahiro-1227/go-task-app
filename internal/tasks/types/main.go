package types

import (
	"errors"
	"time"
)

type Task struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(255); unique; not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

var (
	ErrInvalidInput   = errors.New("タスクのタイトルを入力してください")
	ErrDuplicateTitle = errors.New("タスクのタイトルが重複しています")
	ErrCreateFailed   = errors.New("タスクの作成に失敗しました")
)
