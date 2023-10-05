package controllers

import (
	"crud-book/db"
	"crud-book/models"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) ([]map[string]any, error) {
	var books []*models.Book
	var responses []map[string]any

	err := db.DB.Preload("Author").Preload("Genre").Find(&books).Error
	if err != nil {
		return nil, err
	}

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

	return responses, nil
}

func GetBookById(c *gin.Context) (map[string]any, error) {
	var book *models.Book

	err := db.DB.Preload("Author").Preload("Genre").Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		return nil, err
	}

	response := map[string]any{
		"title":           book.Title,
		"description":     book.Description,
		"bookLanguage":    book.BookLanguage,
		"publicationYear": book.PublicationYear,
		"publisher":       book.Publisher,
		"isbn":            book.Isbn,
		"pageCount":       book.PageCount,
		"price":           book.Price,
		"availablity":     book.Availability,
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
			"Name": book.Genre.Name,
		},
	}

	return response, nil
}
