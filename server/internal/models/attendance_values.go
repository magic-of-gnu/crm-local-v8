package models

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceValues struct {
	ID          uuid.UUID `json:"id" validate:"omitempty,uuid4"`
	Value       int       `json:"value" validate:"required,number"`
	Name        string    `json:"name" validate:"required,min=3,max=50"`
	Description string    `json:"description" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at" validate:"omitempty,datetime"`
	UpdatedAt   time.Time `json:"updated_at" validate:"omitempty,datetime"`
}
