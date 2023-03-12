package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	UpdatedAt time.Time
	CreatedAt time.Time
}
