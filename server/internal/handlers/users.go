package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetUsersAll(c *gin.Context) {
	title := "Users List"

	items, err := App.UsersService.GetAll()

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

func PostUsersCreateOne(c *gin.Context) {
	title := "Users CreateOne"

	var req *models.UsersCreateOneRequest

	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("body: ", c.Request.Body)
		fmt.Println("error during binding, err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	req.IsAdmin = false
	if req.IsAdminInt == 1 {
		req.IsAdmin = true
	}

	item, err := App.UsersService.CreateOne(
		req.FirstName,
		req.LastName,
		req.Username,
		req.Password,
		req.IsAdmin,
		req.UserType,
		App.HashCost,
	)
	if err != nil {
		fmt.Println("error during write")
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error during write",
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

func DeleteUsersByID(c *gin.Context) {
	title := "Users Delete"

	var req *models.UsersDeleteByIDRequest

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

	err = App.UsersService.DeleteOneByID(req.ID)
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
