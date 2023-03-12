package models

import (
	"time"

	"github.com/google/uuid"
)

type Centre struct {
	ID          uuid.UUID
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
