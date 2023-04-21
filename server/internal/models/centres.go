package models

import (
	"time"

	"github.com/google/uuid"
)

type Centre struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CentreDeleteByIDRequest struct {
	ID uuid.UUID `json:"id"`
}
