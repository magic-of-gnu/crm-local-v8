package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetPaymentStatusesAll(c *gin.Context) {
	title := "Payment Statuses List"

	items, err := App.PaymentStatusesService.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title":   title,
			"message": "error during getting data from postgres",
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
	data["data"] = items

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"data":  data,
	})
}

func PostPaymentStatusesCreateOne(c *gin.Context) {
	title := "Payment Status CreateOne"

	var req *models.PaymentStatus

	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("req: ", req, " err: ", err)
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

	item, err := App.PaymentStatusesService.CreateOne(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error db",
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

	c.JSON(http.StatusCreated, gin.H{
		"title": title,
		"data":  item,
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}

func GetOnePaymentStatusByID(c *gin.Context) {
	title := "Payment Status GetOne"

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

	item, err := App.PaymentStatusesRepo.GetOneByID(uid)
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

func DeletePaymentStatusesByID(c *gin.Context) {
	title := "Payment Status Delete"

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

	err = App.PaymentStatusesRepo.DeleteOneByID(uid)
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

func PatchPaymentStatusesByID(c *gin.Context) {
	title := "Payment Status Patch"

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

	req, err := App.PaymentStatusesService.GetOneByID(uid)
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

	err = App.PaymentStatusesRepo.UpdateOneByID(req)
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
