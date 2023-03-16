package controllers

import "github.com/gin-gonic/gin"

func AddApiController(e *gin.Engine) {
	api := e.Group("/api")
	{
		AddBookHandlers(api)
	}
}
