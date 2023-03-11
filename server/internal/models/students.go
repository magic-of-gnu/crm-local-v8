package models

import "time"

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	UpdatedAt time.Time
	CreatedAt time.Time
}
