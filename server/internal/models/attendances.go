package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID                 uuid.UUID        `json:"id"`
	LecturesCalendarID uuid.UUID        `json:"lectures_calendar_id"`
	StudentID          uuid.UUID        `json:"student_id"`
	AttendanceValueID  uuid.UUID        `json:"attendnace_value_id"`
	PaymentStatusID    uuid.UUID        `json:"payment_status_id"`
	InvoiceID          uuid.UUID        `json:"invoice_id"`
	Description        string           `json:"description"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	LectureCalendar    LectureCalendar  `json:"lecture_calendar"`
	Student            Student          `json:"student"`
	AttendanceValues   AttendanceValues `json:"attendance_values"`
	PaymentStatus      PaymentStatus    `json:"payment_status"`
	Invoice            Invoice          `json:"invoice"`
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
	PaymentStatusID       uuid.UUID `json:"payment_status_id"`
	PaymentStatusName     string    `json:"payment_status_name"`
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
