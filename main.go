package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// App contains all state used throughout the webapp
type App struct {
	Router     *gin.Engine
	Controller IOController
}

func main() {
	usePicon := getEnvOr("USE_PICON", "no") != "no"

	var ioController IOController
	if usePicon {
		var err error
		ioController, err = NewPiconController()
		if err != nil {
			log.Panicln(err)
		}
	} else {
		ioController = NewFakeController("Fake controller with 2 motor channels", 2)
	}

	app := &App{
		Router:     gin.New(),
		Controller: ioController,
	}

	app.Router.Use(gin.Logger())
	app.Router.Use(gin.ErrorLogger())
	app.Router.Use(gin.Recovery())

	info := app.Controller.Info()

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"controllerName":    info.Name,
			"motorChannelCount": info.MotorChannelCount,
		})
	})

	app.Router.GET("/motors", func(c *gin.Context) {
		descs, err := motorDescGetAll(app.Controller)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, descs)
	})

	app.Router.GET("/motors/:id", func(c *gin.Context) {
		motorID, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

		// validation of motorID
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		desc, err := motorDescGetOne(app.Controller, uint(motorID))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, desc)
	})

	app.Router.PUT("/motors/:id", func(c *gin.Context) {
		motorID, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

		// validation of motorID
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var payload MotorDescInput
		if err := c.Bind(&payload); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		desc, err := motorDescSetOne(app.Controller, uint(motorID), payload.Speed)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, desc)
	})

	log.Fatal(endless.ListenAndServe("0.0.0.0:8080", app.Router))
}
