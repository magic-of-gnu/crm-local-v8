package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/app"
)

var App *app.App

func RouteHandlers(r *gin.Engine, methodNames map[string]string) {
	r.GET("centres", GetCentres)
	r.POST("centres", PostCentres)

	methodNames["GetCentres"] = "centres"
	methodNames["PostCentres"] = "centres"

}

func NewApp(app *app.App) {
	App = app
}
