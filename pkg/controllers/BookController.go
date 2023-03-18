package controllers

import (
	"GinRestApi/pkg/dto"
	"GinRestApi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var AddBookHandlers = func(router *gin.RouterGroup) {
	books := router.Group("/v1/Books")
	{
		books.GET("", func(c *gin.Context) {
			book := models.Book{}
			result := book.FindAll()
			c.JSON(http.StatusOK, gin.H{
				"message": "Get all books",
				"data":    result,
			})
		})

		books.POST("", func(c *gin.Context) {
			var request dto.Book
			if err := c.Bind(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad request",
				})
			}

			book := models.Book{
				Name:     request.Name,
				AuthorID: request.AuthorId,
			}
			book.Create()

			c.JSON(http.StatusOK, gin.H{
				"message": "Create a new book",
				"data":    book,
			})
		})

		books.GET("/:id", func(c *gin.Context) {
			var param = c.Param("id")
			id, _ := strconv.ParseInt(param, 10, 64)
			book := models.Book{}
			book.FindById(id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Get book by id",
				"data":    book,
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
			var param = c.Param("id")
			id, _ := strconv.ParseInt(param, 10, 64)
			book := models.Book{}
			book.Delete(id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Delete a book by id",
				"data":    book,
			})
		})
	}

}
