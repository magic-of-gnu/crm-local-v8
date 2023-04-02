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
