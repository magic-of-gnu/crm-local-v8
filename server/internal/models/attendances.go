package models

import "time"

type Attendance struct {
	ID                 int
	LecturesCalendarID int
	StudentID          int
	AttendanceValueID  int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
