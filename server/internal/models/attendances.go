package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID                 uuid.UUID
	LecturesCalendarID uuid.UUID
	StudentID          uuid.UUID
	AttendanceValueID  uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
