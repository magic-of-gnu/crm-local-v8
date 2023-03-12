package models

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceValues struct {
	ID          uuid.UUID
	Value       int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
