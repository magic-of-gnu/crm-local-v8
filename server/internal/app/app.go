package app

import (
	"path/filepath"

	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"

	"github.com/gin-contrib/multitemplate"

	"github.com/go-playground/validator/v10"
)

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
) *App {
	return &App{
		MethodNames: methodNames,
		Validator:   validator,

		CentresRepo:    centresRepo,
		CentresService: centresService,

		RoomsRepo:    roomsRepo,
		RoomsService: roomsService,

		EmployeesRepo:    employeesRepo,
		EmployeesService: employeesService,

		StudentsRepo:    studentsRepo,
		StudentsService: studentsService,

		CoursesRepo:    coursesRepo,
		CoursesService: coursesService,

		AttendanceValuesRepo:    attendanceValuesRepo,
		AttendanceValuesService: attendanceValuesServices,
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
