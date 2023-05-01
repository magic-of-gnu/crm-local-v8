package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{
		"title": "Employees",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}
