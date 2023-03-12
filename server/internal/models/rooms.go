package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	CentreID  uuid.UUID
	Name      string
	NumSeats  int
	Info      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
