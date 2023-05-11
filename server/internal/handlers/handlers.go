package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/app"
)

var App *app.App

func RouteHandlers(router *gin.Engine) {
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
	r.GET("api/server/centres", GetAllCentres)
	r.GET("api/server/centres/:id", GetOneCentresByID)
	r.POST("api/server/centres", PostCentres)
	r.DELETE("api/server/centres/:id", DeleteCentresByID)
	r.PATCH("api/server/centres/:id", PatchCentresByID)

	// rooms handlers
	r.GET("api/server/rooms", GetAllRooms)
	r.POST("api/server/rooms", PostRooms)
	r.GET("api/server/rooms/:id", GetOneRoomsByID)
	r.DELETE("api/server/rooms/:id", DeleteRoomsByID)
	r.PATCH("api/server/rooms/:id", PatchRoomsByID)

	// employees handlers
	r.GET("api/server/employees", GetAllEmployees)
	r.POST("api/server/employees", PostEmployeeCreateOne)
	r.GET("api/server/employees/:id", GetOneEmployeeByID)
	r.DELETE("api/server/employees/:id", DeleteEmployeesByID)
	r.PATCH("api/server/employees/:id", PatchEmployeesByID)

	// students handlers
	r.GET("api/server/students", GetAllStudents)
	r.POST("api/server/students", PostStudentsCreateOne)
	r.GET("api/server/students/:id", GetOneStudentsByID)
	r.DELETE("api/server/students/:id", DeleteStudentsByID)
	r.PATCH("api/server/students/:id", PatchStudentsByID)

	// courses handlers
	r.GET("api/server/courses", GetAllCourses)
	r.POST("api/server/courses", PostCoursesCreateOne)
	r.GET("api/server/courses/:id", GetOneCoursesByID)
	r.DELETE("api/server/courses/:id", DeleteCoursesByID)
	r.PATCH("api/server/courses/:id", PatchCoursesByID)

	// attendanceValues handlers
	r.GET("api/server/attendance_values", GetAllAttendanceValues)
	r.POST("api/server/attendance_values", PostAttendanceValuesCreateOne)
	r.GET("api/server/attendance_values/:id", GetOneAttendanceValuesByID)
	r.DELETE("api/server/attendance_values/:id", DeleteAttendanceValuesByID)
	r.PATCH("api/server/attendance_values/:id", PatchAttendanceValuesByID)

	// studentCourses handlers
	r.GET("api/server/student_courses", GetAllStudentCourses)
	r.POST("api/server/student_courses", PostStudentCoursesCreateOne)
	r.GET("api/server/student_courses/custom_id", GetManyStudentCoursesByCustomID)
	// r.GET("api/server/rooms/:id", GetOneCentresByID)
	// r.DELETE("api/server/rooms/:id", DeleteCentresByID)
	// r.PATCH("api/server/rooms/:id", PatchCentresByID)

	// lectureCalendar handlers
	r.GET("api/server/lectures_calendar", GetLecturesCalendarAll)
	r.POST("api/server/lectures_calendar", PostLecturesCalendarCreateOne)
	r.DELETE("api/server/lectures_calendar", DeleteLecturesCalendarByID)
	r.GET("api/server/lectures_calendar/getByCourseID", GetManyLecturesCalendarByCourseID)
	r.GET("api/server/lectures_calendar/:id", GetOneLecturesCalendarByID)
	// r.GET("api/server/rooms/:id", GetOneCentresByID)
	// r.DELETE("api/server/rooms/:id", DeleteCentresByID)
	// r.PATCH("api/server/rooms/:id", PatchCentresByID)

	// attendances handlers
	r.GET("api/server/attendances", GetAttendancesAll)
	r.POST("api/server/attendances", PostAttendancesCreateOne)
	r.POST("api/server/attendances/create_many", PostAttendancesCreateMany)
	r.DELETE("api/server/attendances/:id", DeleteAttendancesByID)
	// r.GET("api/server/rooms/:id", GetOneCentresByID)
	// r.DELETE("api/server/rooms/:id", DeleteCentresByID)
	// r.PATCH("api/server/rooms/:id", PatchCentresByID)

	// users handlers
	r.GET("api/server/users", GetUsersAll)
	r.POST("api/server/users", PostUsersCreateOne)
	r.DELETE("api/server/users", DeleteUsersByID)
	r.GET("api/server/users/:id", GetOneUsersByID)
	r.DELETE("api/server/users/:id", DeleteUsersByID)
	// r.PATCH("api/server/users/:id", PatchCentresByID)

	// invoices handlers
	r.GET("api/server/invoices", GetInvoicesAll)
	r.POST("api/server/invoices", PostInvoicesCreateOne)
	r.DELETE("api/server/invoices/:id", DeleteInvoicesByID)
	// r.PATCH("api/server/invoices/:id", DeleteInvoicesByID)
	// r.GET("api/server/rooms/:id", GetOneCentresByID)
	// r.DELETE("api/server/rooms/:id", DeleteCentresByID)
	// r.PATCH("api/server/rooms/:id", PatchCentresByID)

	// r.GET("api/server/invoices/filtered", GetInvoicesFilteredAll)
	// payment statuses
	r.GET("api/server/payment_statuses", GetPaymentStatusesAll)
	r.GET("api/server/payment_statuses/:id", GetOnePaymentStatusByID)
	r.POST("api/server/payment_statuses", PostPaymentStatusesCreateOne)
	r.PATCH("api/server/payment_statuses/:id", PatchPaymentStatusesByID)
	r.DELETE("api/server/payment_statuses/:id", DeletePaymentStatusesByID)

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
