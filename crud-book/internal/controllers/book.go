package controllers

import (
	"net/http"

	"crud-book/db"
	"crud-book/models"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []*models.Book
	var responses []map[string]any

	db.DB.Model(&models.Book{}).Find(&books)

	for _, book := range books {
		response := map[string]any{
			"title": book.Title,
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": responses,
	})
}
