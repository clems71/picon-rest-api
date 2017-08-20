package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func apiMotorMount(app *App) {
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
}
