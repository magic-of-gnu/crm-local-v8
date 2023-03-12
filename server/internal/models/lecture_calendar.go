package models

import (
	"time"

	"github.com/google/uuid"
)

type LectureCalendar struct {
	ID         uuid.UUID
	RoomID     int
	CourseID   int
	EmployeeID int
	Date       time.Time
	Duration   time.Duration
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
