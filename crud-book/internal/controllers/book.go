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
			"title":           book.Title,
			"descripton":      book.Description,
			"bookLanguage":    book.BookLanguage,
			"publicationYear": book.PublicationYear,
			"publisher":       book.Publisher,
			"isbn":            book.Isbn,
			"pageCount":       book.PageCount,
			"price":           book.Price,
			"availability":    book.Availability,
			"averageRating":   book.AverageRating,
			"author": map[string]any{
				"firstName":   book.Author.FirstName,
				"lastName":    book.Author.LastName,
				"birthDate":   book.Author.BirthDate,
				"nationality": book.Author.Nationality,
				"email":       book.Author.Email,
			},
			"genre": map[string]any{
				"ID":   book.Genre.ID,
				"name": book.Genre.Name,
			},
		}
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     responses,
		"messsage": "Success get all books",
	})
}
