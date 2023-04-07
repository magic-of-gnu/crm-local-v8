package app

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
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

	MethodNames map[string]string

	Validator *validator.Validate
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
		MethodNames:             methodNames,
		Validator:               validator,
	}
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
