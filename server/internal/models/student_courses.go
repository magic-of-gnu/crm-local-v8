package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentCourses struct {
	ID            uuid.UUID
	StudentID     uuid.UUID
	CourseID      uuid.UUID
	PaymentAmount int
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
