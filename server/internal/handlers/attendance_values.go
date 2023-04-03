package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllAttendanceValues(c *gin.Context) {

	items, err := App.AttendanceValuesRepo.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title":   "Attendance Value List",
			"message": "error during getting data from postgres",
			"error":   err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	data["data"] = items

	c.JSON(http.StatusOK, gin.H{
		"title": "Attendance Values List",
		"data":  data,
	})
}

func PostAttendanceValuesCreateOne(c *gin.Context) {

	var item *models.AttendanceValues

	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "AttendanceValues",
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	fmt.Println("after bind, course: ", item)

	err := App.Validator.Struct(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "AttendanceValues",
			"message": "error during validation",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("after validate, course: ", item)

	item, err = App.AttendanceValuesService.CreateOne(
		item.Value,
		item.Name,
		item.Description,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "AttendanceValues",
			"error":   err.Error(),
			"message": "error during writing data to db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "AttendanceValues",
		"item":  item,
	})
}
