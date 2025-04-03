package types

import (
	"time"
)

type UserBase struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255); unique; not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type User struct {
	UserBase
	Password string `gorm:"type:varchar(255); not null"`
}

type UserResponse struct {
	UserBase
}

type SignInInput struct {
	Name     string
	Password string
}

type SignInResponse struct {
	AccessToken string
	User        UserResponse
}
