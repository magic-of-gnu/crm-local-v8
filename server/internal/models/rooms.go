package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `json:"id"`
	CentreID  uuid.UUID `json:"centre_id"`
	Name      string    `json:"name"`
	NumSeats  int       `json:"num_seats"`
	Info      string    `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
