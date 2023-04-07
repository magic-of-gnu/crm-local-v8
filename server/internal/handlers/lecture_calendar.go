package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetLecturesCalendarAll(c *gin.Context) {
	title := "Lecture Calendar list"

	items, err := App.LectureCalendarService.GetAll()

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

func PostLecturesCalendarCreateOne(c *gin.Context) {
	title := "Lecture Calendar CreateOne"

	var req *models.LectureCalendarRequest

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

	start_date := time.Unix(req.StartDateRequest, 0)
	end_date := time.Unix(req.EndDateRequest, 0)
	req.StartDate = time.Date(start_date.Year(), start_date.Month(), start_date.Day(), 0, 0, 0, 0, time.UTC)
	req.EndDate = time.Date(end_date.Year(), end_date.Month(), end_date.Day(), 0, 0, 0, 0, time.UTC)

	items, err := App.LectureCalendarService.CreateMany(
		req.RoomID,
		req.CourseID,
		req.EmployeeID,
		req.StartDate,
		req.EndDate,
		req.DayAndTimeList,
	)

	c.JSON(http.StatusCreated, gin.H{
		"title": title,
		"data":  items,
	})
}

func DeleteLecturesCalendarByID(c *gin.Context) {
	title := "Lectures Calendar Delete"

	var req *models.LectureCalendarDeleteByIDRequest

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

	err = App.LectureCalendarService.DeleteOneByID(req.ID)
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