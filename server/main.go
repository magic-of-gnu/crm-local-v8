package main

import (
	"context"
	"crypto"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/app"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/handlers"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/handlers/middlewares"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"

	"github.com/go-playground/validator/v10"
)

func main() {

	appConfig, err := app.NewAppConfig()
	if err != nil {
		return
	}

	privateKey, publicKey, err := initEd25119Keys(appConfig.PrivateKeyPath, appConfig.PublicKeyPath)
	if err != nil {
		return
	}

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

	// lectureCalendar
	lectureCalendarRepo := repos.NewLectureCalendarPostgresRepo(dbpool)
	lectureCalendarService := services.NewLectureCalendarService(lectureCalendarRepo)

	// attendances
	attendancesRepo := repos.NewAttendancesPostgresRepo(dbpool)
	attendancesService := services.NewAttendancesServiceService(attendancesRepo)

	// users
	usersRepo := repos.NewUsersPostgresRepo(dbpool)
	usersService := services.NewUsersService(usersRepo)

	// tokens
	tokensService := services.NewTokensService(models.TokenKeys{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	},
		appConfig.TokenSigningMethod,
	)

	loginService := services.NewLoginService(
		usersRepo,
		usersService,
		tokensService,
		appConfig.TokenSigningMethod,
		appConfig.HashCost,
	)

	authMiddleware := middlewares.NewAuthMiddleware(tokensService)

	// repos
	// repos := app.NewRepos(lectureCalendarRepo)
	// services := app.NewServices(lectureCalendarService)

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
		lectureCalendarRepo,
		lectureCalendarService,
		attendancesRepo,
		attendancesService,
		usersRepo,
		usersService,
		tokensService,
		loginService,
		authMiddleware,
		appConfig.HashCost,
		appConfig.Debug,
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

func initEd25119Keys(privateKeyPath, publicKeyPath string) (crypto.PrivateKey, crypto.PublicKey, error) {

	key, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("error occurred during reading public key from disk; keypath: ", publicKeyPath)
		return nil, nil, err
	}

	public_key, err := jwt.ParseEdPublicKeyFromPEM(key)
	if err != nil {
		fmt.Println("error occurred during parsing public key, err: ", err, " keypath: ", publicKeyPath)
		return nil, nil, err
	}

	key, err = os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("error occurred during reading private key from diskl keypath: ", privateKeyPath)
		return nil, nil, err
	}

	private_key, err := jwt.ParseEdPrivateKeyFromPEM(key)
	if err != nil {
		fmt.Println("error occurred during parsing private key, err: ", err, " keypath: ", privateKeyPath)
		return nil, nil, err
	}

	return private_key, public_key, nil

}
