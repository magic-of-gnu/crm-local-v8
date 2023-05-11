package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllEmployees(c *gin.Context) {

	empls, err := App.EmployeesService.GetAll()

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title": "Employees List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	data["data"] = empls

	c.JSON(http.StatusOK, gin.H{
		"title": "Rooms List",
		"data":  data,
	})
}

func PostEmployeeCreateOne(c *gin.Context) {

	var empl *models.Employee
	fmt.Println("body: ", c.Request.Body)
	if err := c.ShouldBind(&empl); err != nil {
		fmt.Println("error on binding")
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Employees",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	empl, err := App.EmployeesService.CreateOne(
		empl.FirstName,
		empl.LastName,
		empl.Username,
		empl.Info,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Employees",
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
	data["data"] = empl

	c.JSON(http.StatusCreated, gin.H{
		"title": "Employees",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": empl,
	})
}

func GetOneEmployeeByID(c *gin.Context) {
	title := "Employee GetOne"

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

	item, err := App.EmployeesRepo.GetOneByID(uid)
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

func DeleteEmployeesByID(c *gin.Context) {
	title := "Employees Delete"

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

	err = App.EmployeesRepo.DeleteOneByID(uid)
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

func PatchEmployeesByID(c *gin.Context) {
	title := "Employees Patch"

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

	req, err := App.EmployeesRepo.GetOneByID(uid)
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

	err = App.EmployeesRepo.UpdateOneByID(req)
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
