package models

import (
	"time"

	"github.com/google/uuid"
)

type LectureCalendar struct {
	ID         uuid.UUID     `json:"id"`
	RoomID     uuid.UUID     `json:"room_id"`
	CourseID   uuid.UUID     `json:"course_id"`
	EmployeeID uuid.UUID     `json:"employee_id"`
	Date       time.Time     `json:"date"`
	Duration   time.Duration `json:"duration"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	Room       Room          `json:"room"`
	Course     Course        `json:"course"`
	Employee   Employee      `json:"employee"`
}

type LectureCalendarDeleteByIDRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type DayAndTimeRequest struct {
	Day              int    `json:"day" binding:"required,min=1,max=7"`
	StartTimeRequest string `json:"start_time" binding:"required,len=4"`
	Duration         int    `json:"duration" binding:"required,min=0,max=120"` // in minutes
}

type LectureCalendarRequest struct {
	RoomID           uuid.UUID           `json:"room_id" binding:"required"`
	CourseID         uuid.UUID           `json:"course_id" binding:"required"`
	EmployeeID       uuid.UUID           `json:"employee_id" binding:"required"`
	StartDate        time.Time           `json:"-"`
	StartDateRequest int64               `json:"start_date" binding:"required"`
	EndDate          time.Time           `json:"-"`
	EndDateRequest   int64               `json:"end_date" binding:"required"`
	DayAndTimeList   []DayAndTimeRequest `json:"dates_and_times" binding:"required,dive,required"`
	CreatedAt        time.Time           `json:"created_at" binding:"omitempty,datetime"`
	UpdatedAt        time.Time           `json:"updated_at" binding:"omitempty,datetime"`
}

type LectureCalendarResponse struct {
	ID                uuid.UUID `json:"id"`
	RoomID            uuid.UUID `json:"room_id"`
	RoomName          string    `json:"room_name"`
	CourseID          uuid.UUID `json:"course_id"`
	CourseName        string    `json:"course_name"`
	EmployeeID        uuid.UUID `json:"employee_id"`
	EmployeeFirstName string    `json:"employee_first_name"`
	EmployeeLastName  string    `json:"employee_last_name"`
	Date              int64     `json:"lecture_date"`
	Duration          int       `json:"lecture_duration"`
	CreatedAt         time.Time `json:"created_at" validate:"datetime"`
	UpdatedAt         time.Time `json:"updated_at"`
}
