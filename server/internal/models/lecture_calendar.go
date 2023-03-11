package models

import "time"

type LectureCalendar struct {
	ID         int
	RoomID     int
	CourseID   int
	EmployeeID int
	Date       time.Time
	Duration   time.Duration
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
