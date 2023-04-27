package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "StudentCourses",
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	item, err := App.StudentCoursesService.CreateOne(
		req.StudentID,
		req.CourseID,
		req.PaymentAmount,
		req.Description,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "StudentCourses",
			"error":   err.Error(),
			"message": "error during writing data to db",
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "StudentCourses",
		"item":  item,
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}

func DeleteStudentCoursesByID(c *gin.Context) {
	title := "Student Courses Delete"

	var req *models.StudentCourseDeleteByIDRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	err := App.StudentCoursesService.DeleteOneByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}

func GetManyStudentCoursesByCustomID(c *gin.Context) {
	title := "Student Courses Get By Username"

	queryValues := c.Request.URL.Query()

	if len(queryValues) != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   "incorrect request",
			"message": "error during bidning data",
		})
		return
	}

	var column_name string
	var custom_id uuid.UUID
	var err error

	for key, val := range queryValues {
		column_name = key
		custom_id, err = uuid.Parse(val[0])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"title":   title,
				"error":   err.Error(),
				"message": "error during bidning data",
			})
			return
		}
	}

	items, err := App.StudentCoursesService.GetManyByCustomID(custom_id, column_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error during db",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"data":    items,
	})
}
