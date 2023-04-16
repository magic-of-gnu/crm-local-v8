package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/app"
)

var App *app.App

func RouteHandlers(router *gin.Engine, methodNames map[string]string) {
	// centres form
	// r.GET("centres", GetCentres)
	// r.POST("centres/create_one", PostCentres)

	// login handlers
	router.POST("api/server/login", PostLogin)
	// var r router.Group("")*gin.RouterGroup
	var r *gin.RouterGroup

	if !App.Debug {
		r = router.Group("", App.AuthMiddleware.IsAuthorized())
		// r = r.Use(App.AuthMiddleware.IsAuthorized())
	} else {
		r = router.Group("")
	}

	// centres handlers
	r.GET("api/server/centres/list", GetAllCentres)
	r.POST("api/server/centres/create_one", PostCentres)

	// rooms handlers
	r.GET("api/server/rooms/list", GetAllRooms)
	r.POST("api/server/rooms/create_one", PostRooms)

	// employees handlers
	r.GET("api/server/employees/list", GetAllEmployees)
	r.POST("api/server/employees/create_one", PostEmployeeCreateOne)

	// students handlers
	r.GET("api/server/students/list", GetAllStudents)
	r.POST("api/server/students/create_one", PostStudentsCreateOne)

	// courses handlers
	r.GET("api/server/courses/list", GetAllCourses)
	r.POST("api/server/courses/create_one", PostCoursesCreateOne)

	// attendanceValues handlers
	r.GET("api/server/attendance_values/list", GetAllAttendanceValues)
	r.POST("api/server/attendance_values/create_one", PostAttendanceValuesCreateOne)

	// studentCourses handlers
	r.GET("api/server/student_courses/list", GetAllStudentCourses)
	r.POST("api/server/student_courses/create_one", PostStudentCoursesCreateOne)
	r.GET("api/server/student_courses", GetManyStudentCoursesByCustomID)

	// lectureCalendar handlers
	r.GET("api/server/lectures_calendar/list", GetLecturesCalendarAll)
	r.POST("api/server/lectures_calendar/create_one", PostLecturesCalendarCreateOne)
	r.DELETE("api/server/lectures_calendar", DeleteLecturesCalendarByID)
	r.GET("api/server/lectures_calendar/getByCourseID", GetManyLecturesCalendarByCourseID)
	r.GET("api/server/lectures_calendar/:id", GetOneLecturesCalendarByID)

	// lectureCalendar handlers
	r.GET("api/server/attendances/list", GetAttendancesAll)
	r.POST("api/server/attendances/create_one", PostAttendancesCreateOne)
	r.DELETE("api/server/attendances", DeleteAttendancesByID)

	// users handlers
	r.GET("api/server/users/list", GetUsersAll)
	r.POST("api/server/users/create_one", PostUsersCreateOne)
	r.DELETE("api/server/users", DeleteUsersByID)

	methodNames["GetCentresList"] = "centres/list"

}

func NewApp(app *app.App) {
	App = app
}

type S struct {
	Username string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func MiddlewareCheck(c *gin.Context) {
	fmt.Println("in middleware check")
}
