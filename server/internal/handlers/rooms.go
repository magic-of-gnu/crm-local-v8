package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllRooms(c *gin.Context) {

	rooms, err := App.RoomsService.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title": "Rooms List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	rooms_data := make([]map[string]string, len(rooms))

	for ii, room := range rooms {
		rooms_data[ii] = map[string]string{
			"id":         room.ID.String(),
			"centre_id":  room.CentreID.String(),
			"name":       room.Name,
			"num_seats":  strconv.Itoa(room.NumSeats),
			"info":       room.Info,
			"created_at": room.CreatedAt.Format(time.DateTime),
			"updated_at": room.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = rooms_data

	c.JSON(http.StatusOK, gin.H{
		"title": "Rooms List",
		"data":  data,
	})
}

func PostRooms(c *gin.Context) {

	var room *models.Room

	if err := c.ShouldBind(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Rooms",
		})
		return
	}

	room, err := App.RoomsService.CreateOne(
		room.CentreID,
		room.Name,
		room.NumSeats,
		room.Info,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Rooms",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "Rooms",
	})
}