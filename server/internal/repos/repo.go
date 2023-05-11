package repos

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresRepo interface {
	CreateOne(id uuid.UUID, name, description string, created_at, updated_at time.Time) (*models.Centre, error)
	GetAll() ([]models.Centre, error)
	GetOneByID(uid uuid.UUID) (*models.Centre, error)
	DeleteOneByID(uid uuid.UUID) error
	UpdateOneByID(createItem *models.Centre) error
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
	DeleteOneByID(uid uuid.UUID) error
	GetAll() ([]models.Room, error)
	GetOneByID(uid uuid.UUID) (*models.Room, error)
	UpdateOneByID(createItem *models.Room) error
}

type EmployeesRepo interface {
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
	DeleteOneByID(uid uuid.UUID) error
	GetManyByCustomID(custom_id uuid.UUID, column_name string) ([]models.StudentCourses, error)
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
	GetManyByCourseID(course_id uuid.UUID) ([]models.LectureCalendar, error)
	GetOneByID(uid uuid.UUID) (*models.LectureCalendar, error)
	GetManyFilteredByCourseIDAndDate(course_id uuid.UUID, start_date time.Time, n int) ([]models.LectureCalendar, error)
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
	GetOneFilteredByStudentIDAndLectureCalendarID(lectureCalendarID, studentID uuid.UUID) (*models.Attendance, error)
	UpdateOnePaymentStatuses(attendanceNew *models.Attendance) (*models.Attendance, error)
	DeleteManyByID(uids []uuid.UUID) (int, error)
	RevertPaymentStatusesAndNullifyInvoiceID(attendanceNew *models.Attendance) error
	UpdatePaymentStatusMany(
		uid,
		lecture_calendar_id,
		student_id,
		attendance_value_id uuid.UUID,
		description string,
		created_at, updated_at time.Time,
	) (*models.Attendance, error)
	UpdateOneAtendanceWithLecturesCalendarIDAndStudentID(
		lecturesCalendarID,
		studentID,
		invoiceID uuid.UUID,
		time_now time.Time,
	) error
	GetOneAtendanceWithLecturesCalendarIDAndStudentID(
		lecturesCalendarID,
		studentID uuid.UUID,
	) (*models.Attendance, error)
	UpdateInvoiceIDOnOneAttendance(
		attendanceID,
		invoiceID uuid.UUID,
		updated_at time.Time,
	) error
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
	GetOneByUsername(username string) (*models.User, error)
}

type InvoicesRepo interface {
	GetAll() ([]models.Invoice, error)
	CreateOne(*models.Invoice) (*models.Invoice, error)
	DeleteOneByID(uid uuid.UUID) error
	UpdateOneByID(*models.Invoice) error
}

type PaymentStatusRepo interface {
	CreateOne(item *models.PaymentStatus) (*models.PaymentStatus, error)
	DeleteOneByID(uid uuid.UUID) error
	GetAll() ([]models.PaymentStatus, error)
	GetOneByID(uid uuid.UUID) (*models.PaymentStatus, error)
	UpdateOneByID(createItem *models.PaymentStatus) error
}
