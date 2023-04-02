package models

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `json:"id" validate:"omitempty,uuid4"`
	Name        string    `json:"name" validate:"required,min=3,max=50"`
	Description string    `json:"description" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at" validate:"omitempty,datetime"`
	UpdatedAt   time.Time `json:"update_at" validate:"omitempty,datetime"`
}
