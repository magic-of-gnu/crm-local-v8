package handlers

import (
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
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	item, err := App.AttendanceValuesService.CreateOne(
		item.Value,
		item.Name,
		item.Description,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   "AttendanceValues",
			"error":   err.Error(),
			"message": "error during writing data to db",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "AttendanceValues",
		"item":  item,
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}
