package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetInvoicesAll(c *gin.Context) {
	title := "Invoices List"

	items, err := App.InvoicesService.GetAll()

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

func PostInvoicesCreateOne(c *gin.Context) {
	title := "Invoices CreateOne"

	var req *models.Invoice

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

	item, err := App.InvoicesService.CreateOne(req)
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

func DeleteInvoicesByID(c *gin.Context) {
	title := "Invoice Delete"

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

	err = App.InvoicesService.DeleteOneByID(uid)
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

func GetInvoicesFilteredAll(c *gin.Context) {
	title := "Invoices List"

	course_id, err := uuid.Parse("8d287eea-038b-4610-a7db-7b00f7b03fea")
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
	}

	start_date := time.Unix(1676923230, 0)

	// items, err := App.LectureCalendarRepo.
	fmt.Println("course_id: ", course_id)
	fmt.Println("start_date: ", start_date)

	items, err := App.LectureCalendarRepo.GetManyFilteredByCourseIDAndDate(course_id, start_date)
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
