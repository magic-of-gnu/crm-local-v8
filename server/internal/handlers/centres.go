package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetCentres(c *gin.Context) {
	c.HTML(http.StatusOK, "centres.html", gin.H{
		"title":       "Centres",
		"post_action": "centres",
	})
}

func PostCentres(c *gin.Context) {

	var centre *models.Centre
	if err := c.Bind(&centre); err != nil {
		fmt.Println("error: ", err)
		c.HTML(http.StatusInternalServerError, "centres.html", gin.H{
			"title": "Centres",
		})
		return
	}

	centre, err := App.CentresService.CreateOne(centre.Name, centre.Description)
	if err != nil {
		fmt.Println("error: ", err)
		c.HTML(http.StatusInternalServerError, "centres.html", gin.H{
			"title": "Centres",
		})
		return
	}

	c.HTML(http.StatusOK, "centres.html", gin.H{
		"title": "Centres",
	})
}
