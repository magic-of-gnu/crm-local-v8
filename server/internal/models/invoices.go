package models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID              uuid.UUID     `json:"id"`
	CourseID        uuid.UUID     `json:"course_id"`
	StudentID       uuid.UUID     `json:"student_id"`
	StartDate       time.Time     `json:"start_date"`
	Price           int           `json:"price"`
	PaymentStatusID uuid.UUID     `json:"payment_status_id"`
	LectureNumber   int           `json:"lecture_number"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Course          Course        `json:"course"`
	Student         Student       `json:"student"`
	PaymentStatus   PaymentStatus `json:"payment_status"`
}
