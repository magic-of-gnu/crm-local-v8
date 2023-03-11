package models

import "time"

type Room struct {
	ID        int
	CentreID  int
	Name      string
	NumSeats  int
	Info      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
