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

	db.DB.Preload("Author").Preload("Genre").Find(&books)

	for _, book := range books {
		response := map[string]any{
			"title":      book.Title,
			"descripton": book.Description,
			"author":     book.Author.FirstName,
			"genre":      book.Genre.Name,
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     responses,
		"messsage": "Success get all books",
	})
}
