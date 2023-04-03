package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/app"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/handlers"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"

	"github.com/go-playground/validator/v10"
)

// var DATABASE_URL string = "dbname=local_v8 user=postgres password=postgres host=127.0.0.1 port=5432"

func main() {

	appConfig := app.NewAppConfig()

	methodNames := make(map[string]string)

	database_url := app.CreateDabaseURL(appConfig)

	dbpool, err := pgxpool.New(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// validator
	_Validator := validator.New()

	// centres
	centresRepo := repos.NewCentresPostgresRepo(dbpool)
	centresService := services.NewCentresService(centresRepo)

	// rooms
	roomsRepo := repos.NewRoomsPostgresRepo(dbpool)
	roomsService := services.NewRoomsService(roomsRepo)

	// employees
	employeesRepo := repos.NewEmployeesPostgresRepo(dbpool)
	employeesService := services.NewEmployeesService(employeesRepo)

	// students
	studentsRepo := repos.NewStudentssPostgresRepo(dbpool)
	studentsService := services.NewStudentService(studentsRepo)

	// courses
	coursesRepo := repos.NewCoursesPostgresRepo(dbpool)
	coursesServices := services.NewCoursesService(coursesRepo)

	// attendanceValues
	attendanceValuesRepo := repos.NewAttendanceValuesPostgresRepo(dbpool)
	attendanceValuesService := services.NewAttendanceValuesService(attendanceValuesRepo)

	// studentCourses
	studentCoursesRepo := repos.NewStudentCoursesPostgresRepo(dbpool)
	studentCoursesService := services.NewStudentCoursesService(studentCoursesRepo)

	App := app.NewApp(
		methodNames,
		_Validator,
		centresRepo,
		centresService,
		roomsRepo,
		roomsService,
		employeesRepo,
		employeesService,
		studentsRepo,
		studentsService,
		coursesRepo,
		coursesServices,
		attendanceValuesRepo,
		attendanceValuesService,
		studentCoursesRepo,
		studentCoursesService,
	)
	handlers.NewApp(App)

	// get all centres
	// centres, _ := centresService.GetAll()
	// fmt.Println(centres)

	// create one new centre
	// uid, _ := uuid.NewRandom()
	// centre, _ := centresRepo.CreateOne(uid, "eassy academy", "no description here", time.Now(), time.Now())

	// centre, _ := app.CentresService.CreateOne("aqwe2", "description")
	// fmt.Println("created centre: ", centre)

	// fmt.Println("service: ", centresService)
	fmt.Println("app: ", App)

	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	r.HTMLRender = app.LoadTemplates("./templates")
	r.Static("/static", "./static")

	fmt.Println("before methodNames: ", methodNames)
	handlers.RouteHandlers(r, methodNames)
	fmt.Println("after methodNames: ", methodNames)
	// r.GET("/centres", handlers.GetCentres)

	r.GET("/", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "pong",
		// })
		c.JSON(http.StatusOK, gin.H{
			"title": "Home Page",
		})
	})

	url_string := fmt.Sprintf(":%s", appConfig.Port)
	fmt.Println("port: ", appConfig.Port)

	fmt.Println("routes: ", r.Routes())
	r.Run(url_string)
}
