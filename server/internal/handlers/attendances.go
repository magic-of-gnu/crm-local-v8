package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAttendancesAll(c *gin.Context) {
	title := "Attendances List"

	items, err := App.AttendancesServices.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title":   title,
			"message": "error during getting data from postgres",
			"error":   err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	data["data"] = items

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"data":  data,
	})
}

func PostAttendancesCreateOne(c *gin.Context) {
	title := "Attendances CreateOne"

	var req *models.AttendanceCreateOneRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	err := App.Validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		})
		fmt.Println("error during validation; err: ", err)
		return
	}

	item, err := App.AttendancesServices.CreateOne(
		req.LecturesCalendarID,
		req.StudentID,
		req.AttendanceValueID,
		req.Description,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		})
		fmt.Println("error during validation; err: ", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"title": title,
		"data":  item,
	})
}

func DeleteAttendancesByID(c *gin.Context) {
	title := "Lectures Calendar Delete"

	var req *models.AttendanceDeleteByIDRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	err := App.Validator.Struct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		})
		return
	}

	err = App.AttendancesServices.DeleteOneByID(req.ID)
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
