package handlers

import (
	"fmt"
	"net/http"
	"time"

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
	empls_data := make([]map[string]string, len(empls))

	for ii, empl := range empls {
		empls_data[ii] = map[string]string{
			"id":         empl.ID.String(),
			"first_name": empl.FirstName,
			"last_name":  empl.LastName,
			"username":   empl.Username,
			"info":       empl.Info,
			"created_at": empl.CreatedAt.Format(time.DateTime),
			"updated_at": empl.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = empls_data

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
		fmt.Println("service on binding")

		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Employees",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "Employees",
	})
}
