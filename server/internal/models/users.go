package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	FirstName  string `json:"first_name" validate:"required,min=3,max=50"`
	LastName   string `json:"last_name" validate:"required,min=3,max=50"`
	Username   string `json:"username" validate:"required,min=3,max=50"`
	Password   string
	LastLogin  time.Time
	IsAdmin    bool
	IsAdminInt int
	UserType   int
	UpdatedAt  time.Time `json:"updated_at,omitempty" validate:"omitempty,datetime"`
	CreatedAt  time.Time `json:"created_at,omitempty" validate:"omitempty,datetime"`
}

type UserGetAllResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	LastLogin time.Time `json:"last_login"`
	IsAdmin   bool      `json:"is_admin"`
	UserType  int       `json:"user_type"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UsersCreateOneRequest struct {
	FirstName  string `json:"first_name" binding:"required,min=3,max=50"`
	LastName   string `json:"last_name" binding:"required,min=3,max=50"`
	Username   string `json:"username" binding:"required,min=3,max=50"`
	Password   string `json:"password" binding:"required,min=8,max=50"`
	Password2  string `json:"password2" binding:"required,min=8,max=50"`
	IsAdmin    bool
	IsAdminInt int `json:"is_admin" binding:"required"`
	UserType   int `json:"user_type" binding:"required"`
}

type UsersDeleteByIDRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
