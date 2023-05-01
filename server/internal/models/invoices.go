package models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID              uuid.UUID     `json:"id"`
	CourseID        uuid.UUID     `json:"course_id" binding:"required"`
	StudentID       uuid.UUID     `json:"student_id" binding:"required"`
	StartDate       time.Time     `json:"start_date" binding:"required"`
	Price           int           `json:"price" binding:"required,gt=0"`
	PaymentStatusID uuid.UUID     `json:"payment_status_id" binding:"required"`
	LecturesNumber  int           `json:"lectures_number" binding:"required,gt=0"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Course          Course        `json:"course" binding:"omitempty"`
	Student         Student       `json:"student" binding:"omitempty"`
	PaymentStatus   PaymentStatus `json:"payment_status" binding:"omitempty"`
}

type InvoiceRequest struct {
	ID               uuid.UUID `json:"id"`
	CourseID         uuid.UUID `json:"course_id" binding:"required"`
	StudentID        uuid.UUID `json:"student_id" binding:"required"`
	StartDate        time.Time `json:"-"`
	StartDateRequest int64     `json:"start_date" binding:"required"`
	Price            int       `json:"price" binding:"required,gt=0"`
	PaymentStatusID  uuid.UUID `json:"payment_status_id" binding:"required"`
	LecturesNumber   int       `json:"lectures_number" binding:"required,gt=0"`
}
