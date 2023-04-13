package app

import (
	"github.com/magic-of-gnu/crm-local-v8/server/internal/handlers/middlewares"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"

	"github.com/go-playground/validator/v10"
)

type Repos struct {
	LectureCalendarRepo repos.LectureCalendarRepo
}

type Services struct {
	LectureCalendarService services.LectureCalendarService
}

func NewRepos(lectureCalendarRepo repos.LectureCalendarRepo) *Repos {
	return &Repos{
		LectureCalendarRepo: lectureCalendarRepo,
	}
}

func NewServices(lectureCalendarService services.LectureCalendarService) *Services {
	return &Services{
		LectureCalendarService: lectureCalendarService,
	}
}

type App struct {
	// centres
	CentresRepo    repos.CentresRepo
	CentresService services.CentresService

	// rooms
	RoomsRepo    repos.RoomsRepo
	RoomsService services.RoomsService

	// employees
	EmployeesRepo    repos.EmployeesRepo
	EmployeesService services.EmployeesService

	// students
	StudentsRepo    repos.StudentsRepo
	StudentsService services.StudentsService

	// courses
	CoursesRepo    repos.CoursesRepo
	CoursesService services.CoursesService

	// attendance_values
	AttendanceValuesRepo    repos.AttendanceValues
	AttendanceValuesService services.AttendanceValuesService

	// student_courses
	StudentCoursesRepo    repos.StudentCoursesRepo
	StudentCoursesService services.StudentCoursesService

	// lecture_calendar
	LectureCalendarRepo    repos.LectureCalendarRepo
	LectureCalendarService services.LectureCalendarService

	// attendances
	AttendancesRepo     repos.AttendancesRepo
	AttendancesServices services.AttendancesService

	// users
	UsersRepo    repos.UsersRepo
	UsersService services.UsersService

	// token
	TokensService services.TokensService

	// login
	LoginService services.LoginService

	// authMiddleware
	AuthMiddleware middlewares.AuthMiddleware

	MethodNames map[string]string

	Validator *validator.Validate

	HashCost int
}

func NewApp(
	methodNames map[string]string,
	validator *validator.Validate,
	centresRepo repos.CentresRepo,
	centresService services.CentresService,
	roomsRepo repos.RoomsRepo,
	roomsService services.RoomsService,
	employeesRepo repos.EmployeesRepo,
	employeesService services.EmployeesService,
	studentsRepo repos.StudentsRepo,
	studentsService services.StudentsService,
	coursesRepo repos.CoursesRepo,
	coursesService services.CoursesService,
	attendanceValuesRepo repos.AttendanceValues,
	attendanceValuesServices services.AttendanceValuesService,
	studentCoursesRepo repos.StudentCoursesRepo,
	studentCoursesService services.StudentCoursesService,
	lectureCalendarRepo repos.LectureCalendarRepo,
	lectureCalendarServicei services.LectureCalendarService,
	attendancesRepo repos.AttendancesRepo,
	attendancesServices services.AttendancesService,
	usersRepo repos.UsersRepo,
	usersService services.UsersService,
	tokensService services.TokensService,
	loginService services.LoginService,
	authMiddleware middlewares.AuthMiddleware,
	hashCost int,
) *App {
	return &App{
		CentresRepo:             centresRepo,
		CentresService:          centresService,
		RoomsRepo:               roomsRepo,
		RoomsService:            roomsService,
		EmployeesRepo:           employeesRepo,
		EmployeesService:        employeesService,
		StudentsRepo:            studentsRepo,
		StudentsService:         studentsService,
		CoursesRepo:             coursesRepo,
		CoursesService:          coursesService,
		AttendanceValuesRepo:    attendanceValuesRepo,
		AttendanceValuesService: attendanceValuesServices,
		StudentCoursesRepo:      studentCoursesRepo,
		StudentCoursesService:   studentCoursesService,
		LectureCalendarRepo:     lectureCalendarRepo,
		LectureCalendarService:  lectureCalendarServicei,
		AttendancesRepo:         attendancesRepo,
		AttendancesServices:     attendancesServices,
		UsersRepo:               usersRepo,
		UsersService:            usersService,
		TokensService:           tokensService,
		LoginService:            loginService,
		AuthMiddleware:          authMiddleware,
		MethodNames:             methodNames,
		Validator:               validator,
		HashCost:                hashCost,
	}
}
