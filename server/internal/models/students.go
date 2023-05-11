package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name" validate:"required,min=3,max=50"`
	LastName  string    `json:"last_name" validate:"required,min=3,max=50"`
	Username  string    `json:"username" validate:"required,min=3,max=50"`
	UpdatedAt time.Time `json:"updated_at,omitempty" validate:"omitempty,datetime"`
	CreatedAt time.Time `json:"created_at,omitempty" validate:"omitempty,datetime"`
}
