package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func PostLogin(c *gin.Context) {
	title := "Login"

	var req *models.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
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

	item, isLoggedIn, err := App.LoginService.Login(
		req.Username,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during writing data to db",
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
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
