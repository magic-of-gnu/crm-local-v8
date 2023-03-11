package models

import "time"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Info      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
