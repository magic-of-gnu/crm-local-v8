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
	Description        string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	LectureCalendar    LectureCalendar
	Student            Student
	AttendanceValues   AttendanceValues
}

type AttendanceResponse struct {
	ID                    uuid.UUID `json:"id"`
	LecturesCalendarID    uuid.UUID `json:"lectures_calendar_id"`
	StudentID             uuid.UUID `json:"student_id"`
	AttendanceValueID     uuid.UUID `json:"attendance_value_id"`
	Description           string    `json:"description"` // ie why student was late
	CreatedAt             time.Time `json:"attendance_created_at"`
	UpdatedAt             time.Time `json:"attendance_updated_at"`
	AttendanceValue       int       `json:"attendance_value"`
	AttendanceName        string    `json:"attendance_name"`
	AttendanceDescription string    `json:"attendance_description"`
	StudentFirstName      string    `json:"student_first_name"`
	StudentLastName       string    `json:"student_last_name"`
	StudentUsername       string    `json:"student_username"`
	RoomID                uuid.UUID `json:"room_id"`
	CourseID              uuid.UUID `json:"course_id"`
	EmployeeID            uuid.UUID `json:"employee_id"`
	LectureDate           int64     `json:"lecture_date"`
	LectureDuration       int       `json:"lecture_duration"`
	RoomName              string    `json:"room_name"`
	CourseName            string    `json:"course_name"`
	EmployeeFirstName     string    `json:"employee_first_name"`
	EmployeeLastName      string    `json:"employee_last_name"`
	EmployeeUsername      string    `json:"employee_username"`
}

type AttendanceCreateOneRequest struct {
	LecturesCalendarID uuid.UUID `json:"lectures_calendar_id" validate:"required"`
	StudentID          uuid.UUID `json:"student_id" validate:"required"`
	AttendanceValueID  uuid.UUID `json:"attendance_value_id" validate:"required"`
	Description        string    `json:"description" validate:"omitempty"`
}

type AttendanceDeleteByIDRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
