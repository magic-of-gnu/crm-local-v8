package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Info      string    `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
