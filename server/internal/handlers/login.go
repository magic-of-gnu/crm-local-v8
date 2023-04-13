package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func PostLogin(c *gin.Context) {
	title := "Login"

	var req *models.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		})
		return
	}

	fmt.Println("post students courses2")

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

	fmt.Println("post students courses4")

	item, isLoggedIn, err := App.LoginService.Login(
		req.Username,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during writing data to db",
		})
		return
	}

	if !isLoggedIn {
		c.JSON(http.StatusUnauthorized, gin.H{
			"title":   title,
			"message": "incorrect username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"data":  item,
	})

}
