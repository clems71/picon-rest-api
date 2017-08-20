package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InstallOutputsAPI sets API routes.
func InstallOutputsAPI(app App) {
	// OUTPUTS
	app.Router.GET("/outputs", func(c *gin.Context) {
		descs, err := outputDescGetAll(app.Controller)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, descs)
	})

	app.Router.GET("/outputs/:id", func(c *gin.Context) {
		outputID, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

		// validation of outputID
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		desc, err := outputDescGetOne(app.Controller, uint(outputID))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, desc)
	})

	app.Router.PUT("/outputs/:id", func(c *gin.Context) {
		outputID, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)

		// validation of outputID
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var payload OutputDescInput
		if err := c.Bind(&payload); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		desc, err := outputDescSetOne(app.Controller, uint(outputID), payload.Value)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, desc)
	})
}
