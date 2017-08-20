package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiBaseMount(app *App) {
	info := app.Controller.Info()
	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"controllerName":    info.Name,
			"motorChannelCount": info.MotorChannelCount,
		})
	})
}
