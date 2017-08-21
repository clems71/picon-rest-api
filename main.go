package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
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
	app.Router.Use(cors.Default())

	apiBaseMount(app)
	apiMotorMount(app)
	apiCameraMount(app)

	log.Fatal(endless.ListenAndServe("0.0.0.0:8080", app.Router))
}
