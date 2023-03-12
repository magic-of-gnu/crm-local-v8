package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Info      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
