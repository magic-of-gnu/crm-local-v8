package repos

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresRepo interface {
	GetAll() ([]models.Centre, error)
	CreateOne(id uuid.UUID, name, description string, created_at, updated_at time.Time) (*models.Centre, error)
}

type RoomsRepo interface {
	GetAll() ([]models.Room, error)
	CreateOne(
		uid,
		centre_uid uuid.UUID,
		name string,
		num_seats int,
		info string,
		created_at, updated_at time.Time,
	) (*models.Room, error)
}

type EmployeesRepo interface {
	GetAll() ([]models.Employee, error)
	CreateOne(
		uid uuid.UUID,
		first_name,
		last_name,
		username,
		info string,
		created_at, updated_at time.Time,
	) (*models.Employee, error)
}

type StudentsRepo interface {
	GetAll() ([]models.Student, error)
	CreateOne(
		uid uuid.UUID,
		first_name,
		last_name,
		username string,
		created_at, updated_at time.Time,
	) (*models.Student, error)
}

type CoursesRepo interface {
	GetAll() ([]models.Course, error)
	CreateOne(
		uid uuid.UUID,
		name,
		description string,
		created_at, updated_at time.Time,
	) (*models.Course, error)
}

type AttendanceValues interface {
	GetAll() ([]models.AttendanceValues, error)
	CreateOne(
		uid uuid.UUID,
		value int,
		name,
		description string,
		created_at, updated_at time.Time,
	) (*models.AttendanceValues, error)
}

type StudentCoursesRepo interface {
	GetAll() ([]models.StudentCourses, error)
	CreateOne(
		uid,
		student_uid,
		course_uid uuid.UUID,
		payment_amount int,
		description string,
		created_at, updated_at time.Time,
	) (*models.StudentCourses, error)
}

type LectureCalendarRepo interface {
	GetAll() ([]models.LectureCalendar, error)
	CreateOne(
		uid,
		roomd_uid,
		course_uid,
		employee_uid uuid.UUID,
		date time.Time,
		duration time.Duration,
		created_at, updated_at time.Time,
	) (*models.LectureCalendar, error)
	DeleteOneByID(uid uuid.UUID) error
}

type AttendancesRepo interface {
	GetAll() ([]models.Attendance, error)
	CreateOne(
		uid,
		lecture_calendar_id,
		student_id,
		attendance_value_id uuid.UUID,
		description string,
		created_at, updated_at time.Time,
	) (*models.Attendance, error)
	DeleteOneByID(uid uuid.UUID) error
}

type UsersRepo interface {
	GetAll() ([]models.User, error)
	CreateOne(
		uid uuid.UUID,
		first_name,
		last_name,
		username,
		password string,
		last_login time.Time,
		is_admin bool,
		user_type int,
		created_at, updated_at time.Time,
	) (*models.User, error)
	DeleteOneByID(uid uuid.UUID) error
}
