package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllStudents(c *gin.Context) {

	fmt.Println("handler students get all")

	students, err := App.StudentsService.GetAll()

	if err != nil {

		fmt.Println("error here")
		c.JSON(http.StatusOK, gin.H{
			"title": "Students List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	students_data := make([]map[string]string, len(students))

	for ii, item := range students {
		students_data[ii] = map[string]string{
			"id":         item.ID.String(),
			"first_name": item.FirstName,
			"last_name":  item.LastName,
			"username":   item.Username,
			"created_at": item.CreatedAt.Format(time.DateTime),
			"updated_at": item.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = students_data

	c.JSON(http.StatusOK, gin.H{
		"title": "Students List",
		"data":  data,
	})
}

func PostStudentsCreateOne(c *gin.Context) {

	fmt.Println("post students")
	var student *models.Student

	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Student",
			"error": err.Error(),
		})
		return
	}

	fmt.Println("after bind, student: ", student)

	err := App.Validator.Struct(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "Student",
			"message": "error during validation",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("after validate, student: ", student)

	student, err = App.StudentsService.CreateOne(
		student.FirstName,
		student.LastName,
		student.Username,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Students",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "Students",
	})
}
