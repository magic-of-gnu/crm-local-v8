package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllCourses(c *gin.Context) {

	items, err := App.CoursesService.GetAll()

	if err != nil {

		fmt.Println("error here")
		c.JSON(http.StatusOK, gin.H{
			"title": "Students List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	items_data := make([]map[string]string, len(items))

	for ii, item := range items {
		items_data[ii] = map[string]string{
			"id":          item.ID.String(),
			"name":        item.Name,
			"description": item.Description,
			"created_at":  item.CreatedAt.Format(time.DateTime),
			"updated_at":  item.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = items_data

	c.JSON(http.StatusOK, gin.H{
		"title": "Courses List",
		"data":  data,
	})
}

func PostCoursesCreateOne(c *gin.Context) {

	fmt.Println("post courses")
	var item *models.Course

	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "Student",
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	fmt.Println("after bind, course: ", item)

	err := App.Validator.Struct(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "Course",
			"message": "error during validation",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("after validate, course: ", item)

	item, err = App.CoursesService.CreateOne(
		item.Name,
		item.Description,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "Courses",
			"error":   err.Error(),
			"message": "error during writing data to db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "Courses",
		"item":  item,
	})
}
