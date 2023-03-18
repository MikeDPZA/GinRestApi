package controllers

import (
	"GinRestApi/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

var AddAuthorHandlers = func(e *gin.RouterGroup) {
	authors := e.Group("/v1/Authors")
	{
		authors.GET("", func(c *gin.Context) {
			a := models.Author{}
			result := a.FindAll()
			c.JSON(200, gin.H{
				"message": "Get all authors",
				"data":    result,
			})
		})

		authors.POST("", func(c *gin.Context) {
			var model models.Author
			if err := c.Bind(&model); err != nil {
				c.JSON(400, gin.H{
					"message": "Bad request",
				})
				return
			}
			model.Create()
			c.JSON(200, gin.H{
				"message": "Create a new author",
				"data":    model,
			})
		})

		authors.GET("/:id", func(c *gin.Context) {
			var param = c.Param("id")
			id, _ := strconv.ParseInt(param, 10, 64)

			author := models.Author{}
			author.FindById(id)

			c.JSON(200, gin.H{
				"message": "Get author by id",
				"data":    author,
			})
		})

		authors.PUT("/:id", func(c *gin.Context) {
			var id = c.Param("id")
			response := "Update an author with id " + id
			c.JSON(200, gin.H{
				"message": response,
			})
		})

		authors.DELETE("/:id", func(c *gin.Context) {
			var param = c.Param("id")
			id, _ := strconv.ParseInt(param, 10, 64)
			author := models.Author{}
			author.Delete(id)
			c.JSON(200, gin.H{
				"message": "Delete an author by id ",
				"data":    author,
			})
		})
	}
}
