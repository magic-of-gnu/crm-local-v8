package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	CentreID  uuid.UUID `json:"centre_id"`
	Name      string    `json:"name"`
	NumSeats  int       `json:"num_seats"`
	Info      string    `json:"info"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
