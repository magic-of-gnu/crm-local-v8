package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

var WeekDays []string = []string{"", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Satyrday", "Sunday"}

type CentresService interface {
	CreateOne(name, description string) (*models.Centre, error)
	DeleteOneByID(uid uuid.UUID) error
	GetAll() ([]models.Centre, error)
	GetOneByID(uid uuid.UUID) (*models.Centre, error)
	UpdateOneByID(updateItem *models.Centre) error
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
	DeleteOneByID(uid uuid.UUID) error
	GetManyByCustomID(uid uuid.UUID, column_name string) ([]models.StudentCourses, error)
}

type LectureCalendarService interface {
	GetAll() ([]models.LectureCalendarResponse, error)
	CreateOne(
		room_uid,
		course_uid,
		employee_uid uuid.UUID,
		date time.Time,
		duration time.Duration,
	) (*models.LectureCalendar, error)
	CreateMany(
		room_uid,
		course_uid,
		employee_uid uuid.UUID,
		start_date time.Time,
		end_date time.Time,
		days_and_times []models.DayAndTimeRequest,
		default_attendance_value_id uuid.UUID,
	) ([]models.LectureCalendar, error)
	DeleteOneByID(uid uuid.UUID) error
	DeleteManyByID(lectures_calendar []models.LectureCalendar) error
	GetManyByCourseID(course_id uuid.UUID) ([]models.LectureCalendar, error)
	GetOneByID(uid uuid.UUID) (*models.LectureCalendar, error)
}

type AttendancesService interface {
	GetAll() ([]models.AttendanceResponse, error)
	CreateOne(
		lecture_calendar_id,
		student_id,
		attendance_value_id uuid.UUID,
		description string,
	) (*models.Attendance, error)
	DeleteOneByID(uid uuid.UUID) error
	CreateMany(
		attendances []models.AttendanceCreateOneRequest,
	) ([]models.Attendance, error)
}

type UsersService interface {
	GetAll() ([]models.UserGetAllResponse, error)
	CreateOne(
		first_name,
		last_name,
		username,
		password string,
		is_admin bool,
		user_type int,
		hash_cost int,
	) (*models.User, error)
	DeleteOneByID(uid uuid.UUID) error
	GetOneByUsername(username string) (*models.User, error)
}

type TokensService interface {
	SignTokenWithCustomClaims(token *jwt.Token) (string, error)
	ParseSignedToken(signedToken string) (*jwt.Token, error)
}

type LoginService interface {
	Login(username string, password string) (*models.LoginResponse, bool, error)
}

type PaymentStatusesService interface {
	CreateOne(createItem *models.PaymentStatus) (*models.PaymentStatus, error)
	DeleteOneByID(uid uuid.UUID) error
	GetAll() ([]models.PaymentStatus, error)
	GetOneByID(uid uuid.UUID) (*models.PaymentStatus, error)
	UpdateOneByID(updateItem *models.PaymentStatus) error
}
