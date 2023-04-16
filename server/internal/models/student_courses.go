package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentCourses struct {
	ID            uuid.UUID `json:"id" validate:"omitempty,uuid4_rfc4122"`
	StudentID     uuid.UUID `json:"student_id" validate:"required"`
	CourseID      uuid.UUID `json:"course_id" validate:"required"`
	PaymentAmount int       `json:"payment_amount" validate:"required,number"`
	Description   string    `json:"description" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"omitempty,datetime"`
	UpdatedAt     time.Time `json:"updated_at" validate:"omitempty,datetime"`
	Student       Student   `json:"student" validate:"omitempty"`
	Course        Course    `json:"course"  validate:"omitempty"`
}

type StudentCoursesListResponse struct {
	ID               uuid.UUID `json:"id"`
	StudentID        uuid.UUID `json:"student_id"`
	CourseID         uuid.UUID `json:"course_id"`
	PaymentAmount    int       `json:"payment_amount"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	StudentFirstName string    `json:"student_first_name"`
	StudentLastName  string    `json:"student_last_name"`
	CourseName       string    `json:"course_name"`
}

type StudentCoursesRequest struct {
	ID            uuid.UUID `json:"id" validate:"omitempty,uuid4_rfc4122"`
	StudentID     uuid.UUID `json:"student_id" validate:"required"`
	CourseID      uuid.UUID `json:"course_id" validate:"required"`
	PaymentAmount int       `json:"payment_amount" validate:"required,number"`
	Description   string    `json:"description" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"omitempty,datetime"`
	UpdatedAt     time.Time `json:"updated_at" validate:"omitempty,datetime"`
}

type StudentCoursesCustomRequest struct {
	ID            uuid.UUID `json:"id" validate:"omitempty"`
	StudentID     uuid.UUID `json:"student_id" validate:"omitempty"`
	CourseID      uuid.UUID `json:"course_id" validate:"omitempty"`
	PaymentAmount int       `json:"payment_amount" validate:"omitempty,number"`
	Description   string    `json:"description" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"omitempty,datetime"`
	UpdatedAt     time.Time `json:"updated_at" validate:"omitempty,datetime"`
}

type StudentCourseDeleteByIDRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
