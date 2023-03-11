package models

import "time"

type AttendanceValues struct {
	ID          int
	Value       int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
