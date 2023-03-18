package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddApiController(e *gin.Engine) {
	api := e.Group("/api")
	{

		api.GET("/v1/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		AddBookHandlers(api)
		AddAuthorHandlers(api)
	}
}
