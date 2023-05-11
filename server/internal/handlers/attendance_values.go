package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	c.JSON(http.StatusCreated, gin.H{
		"title": "AttendanceValues",
		"item":  item,
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": item,
	})
}

func DeleteAttendanceValuesByID(c *gin.Context) {
	title := "Attendance Values Delete"

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

	err = App.AttendanceValuesService.DeleteOneByID(uid)
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

func GetOneAttendanceValuesByID(c *gin.Context) {
	title := "Attendance Values GetOne"

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

	item, err := App.AttendanceValuesService.GetOneByID(uid)
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

func PatchAttendanceValuesByID(c *gin.Context) {
	title := "Attendance Values Patch"

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

	req, err := App.AttendanceValuesService.GetOneByID(uid)
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

	err = App.AttendanceValuesService.UpdateOneByID(req)
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
