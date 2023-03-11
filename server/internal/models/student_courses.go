package models

import "time"

type StudentCourses struct {
	ID            int
	StudentID     int
	CourseID      int
	PaymentAmount int
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
