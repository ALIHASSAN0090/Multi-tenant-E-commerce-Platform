package models

import (
	"time"
)

type SignUpReq struct {
	UserName    string `json:"user_name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Role        string `json:"role" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
}

type Users struct {
	ID          int64     `json:"id" db:"id"`
	UserName    string    `json:"user_name" db:"user_name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Role        string    `json:"role" db:"role"`
	Token       string    `json:"token,omitempty"`
	Address     string    `json:"address,omitempty" db:"address"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
