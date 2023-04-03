package services

import (
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresService interface {
	GetAll() ([]models.Centre, error)
	CreateOne(name, description string) (*models.Centre, error)
}

type RoomsService interface {
	GetAll() ([]models.Room, error)
	CreateOne(
		centre_uid uuid.UUID,
		name string,
		num_seats int,
		info string,
	) (*models.Room, error)
}

type EmployeesService interface {
	GetAll() ([]models.Employee, error)
	CreateOne(
		first_name,
		last_name,
		username,
		info string,
	) (*models.Employee, error)
}

type StudentsService interface {
	GetAll() ([]models.Student, error)
	CreateOne(
		first_name,
		last_name,
		username string,
	) (*models.Student, error)
}

type CoursesService interface {
	GetAll() ([]models.Course, error)
	CreateOne(
		name,
		description string,
	) (*models.Course, error)
}

type AttendanceValuesService interface {
	GetAll() ([]models.AttendanceValues, error)
	CreateOne(
		value int,
		name,
		description string,
	) (*models.AttendanceValues, error)
}

type StudentCoursesService interface {
	GetAll() ([]models.StudentCoursesListResponse, error)
	CreateOne(
		student_uid,
		course_uid uuid.UUID,
		payment_amount int,
		description string,
	) (*models.StudentCourses, error)
}
