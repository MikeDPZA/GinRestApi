package controllers

import (
	"GinRestApi/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AddBookHandlers = func(router *gin.RouterGroup) {
	books := router.Group("/v1/Books")
	{
		books.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get all books",
			})
		})

		books.POST("", func(c *gin.Context) {
			var book dto.Book
			if c.Bind(&book) == nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad request",
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Create a new book",
				"book":    book,
			})
		})

		books.GET("/:id", func(c *gin.Context) {
			var id = c.Param("id")
			response := "Get a book with id " + id
			c.JSON(http.StatusOK, gin.H{
				"message": response,
			})
		})

		books.PUT("/:id", func(c *gin.Context) {
			var id = c.Param("id")
			var book dto.Book
			if c.Bind(&book) != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad request",
				})
				return
			}

			response := "Update a book with id " + id
			c.JSON(http.StatusOK, gin.H{
				"message": response,
				"id":      id,
				"book":    book,
			})
		})

		books.DELETE("/:id", func(c *gin.Context) {
			var id = c.Param("id")
			response := "Delete a book with id " + id
			c.JSON(http.StatusOK, gin.H{
				"message": response,
			})
		})
	}

}
