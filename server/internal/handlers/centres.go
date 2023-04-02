package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllCentres(c *gin.Context) {

	fmt.Println("binding start")
	centres, err := App.CentresService.GetAll()
	fmt.Println("binding finished; centres: ", centres)

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"title": "Centres List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	centres_data := make([]map[string]string, len(centres))

	for ii, centre := range centres {
		centres_data[ii] = map[string]string{
			"id":          centre.ID.String(),
			"name":        centre.Name,
			"description": centre.Description,
			"created_at":  centre.CreatedAt.Format(time.DateTime),
			"updated_at":  centre.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = centres_data
	fmt.Println("here")

	c.JSON(http.StatusOK, gin.H{
		"title": "Centres List",
		"data":  data,
	})
}

func GetCentres(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"title":       "Centres",
		"post_action": "centres",
	})
}

func PostCentres(c *gin.Context) {

	var centre *models.Centre
	if err := c.Bind(&centre); err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Centres",
		})
		return
	}

	centre, err := App.CentresService.CreateOne(centre.Name, centre.Description)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Centres",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": "Centres",
	})
}
