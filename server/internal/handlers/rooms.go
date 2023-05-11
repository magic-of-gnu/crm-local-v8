package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"title": "Rooms",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
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
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"title": "Rooms",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": room,
	})
}

func DeleteRoomsByID(c *gin.Context) {
	title := "Room Delete"

	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
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

	err = App.RoomsService.DeleteOneByID(uid)
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
				"content": "Deleted",
				"color":   "success",
			},
		},
	})
}

func GetOneRoomsByID(c *gin.Context) {
	title := "Rooms GetOne"

	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
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

	item, err := App.RoomsRepo.GetOneByID(uid)
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

	data := make(map[string]interface{})
	data["data"] = item

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": data,
	})
}

func PatchRoomsByID(c *gin.Context) {
	title := "Rooms Patch"

	// get the item from db and update it by the request
	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	req, err := App.RoomsRepo.GetOneByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	err = App.RoomsRepo.UpdateOneByID(req)
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
