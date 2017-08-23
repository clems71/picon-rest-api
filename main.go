package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// App contains all state used throughout the webapp
type App struct {
	Router     *gin.Engine
	Controller IOController
	// Camera     CameraFrameProvider
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
		ioController = NewFakeController()
	}

	app := &App{
		Router:     gin.New(),
		Controller: ioController,
	}

	app.Router.Use(gin.Logger())
	app.Router.Use(gin.ErrorLogger())
	app.Router.Use(gin.Recovery())
	app.Router.Use(static.Serve("/", static.LocalFile("./public", true)))
	app.Router.Use(cors.Default())

	apiMotorMount(app)
	apiCameraMount(app)

	log.Fatal(endless.ListenAndServe("0.0.0.0:8080", app.Router))
}
