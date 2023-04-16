package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllStudentCourses(c *gin.Context) {

	items, err := App.StudentCoursesService.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title":   "Student Courses List",
			"message": "error during getting data from postgres",
			"error":   err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	data["data"] = items

	c.JSON(http.StatusOK, gin.H{
		"title": "Student Courses List",
		"data":  data,
	})
}

func PostStudentCoursesCreateOne(c *gin.Context) {

	var req *models.StudentCoursesRequest
	fmt.Println("post students courses")

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "StudentCourses",
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	fmt.Println("post students courses2")

	err := App.Validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "StudentCourses",
			"message": "error during validation",
			"error":   err.Error(),
		})
		fmt.Println("error during validation; err: ", err)
		return
	}

	fmt.Println("post students courses4")

	item, err := App.StudentCoursesService.CreateOne(
		req.StudentID,
		req.CourseID,
		req.PaymentAmount,
		req.Description,
	)

	fmt.Println("post students courses5")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "StudentCourses",
			"error":   err.Error(),
			"message": "error during writing data to db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "StudentCourses",
		"item":  item,
	})
}

func DeleteStudentCoursesByID(c *gin.Context) {
	title := "Student Courses Delete"

	var req *models.StudentCourseDeleteByIDRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	err := App.StudentCoursesService.DeleteOneByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
	})
}
