package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus struct {
	ID          uuid.UUID `json:"id" binding:"omitempty"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaymentStatusUpdateRequest struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"omitempty"`
	Description string    `json:"description" binding:"omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
