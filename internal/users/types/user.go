package types

import (
	"time"
)

type UserBase struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255); unique; not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type User struct {
	UserBase
	Password string `gorm:"type:varchar(255); not null" json:"password"`
}

type UserResponse struct {
	UserBase
}

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
