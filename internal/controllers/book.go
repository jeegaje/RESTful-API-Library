package controllers

import (
	"crud-book/db"
	"crud-book/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBooks(c *gin.Context) {
	var books []*models.Book
	var responses []map[string]any
	var codeStatus int

	err := db.DB.Preload("Author").Preload("Genre").Find(&books).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			codeStatus = http.StatusNotFound
		default:
			codeStatus = http.StatusBadRequest
		}

		c.JSON(codeStatus, gin.H{
			"code":    codeStatus,
			"message": err.Error(),
		})
		return
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

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Book By ID",
		"data":    responses,
	})
}

func GetBookById(c *gin.Context) {
	var book *models.Book
	var codeStatus int

	err := db.DB.Preload("Author").Preload("Genre").Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			codeStatus = http.StatusNotFound
		default:
			codeStatus = http.StatusBadRequest
		}

		c.JSON(codeStatus, gin.H{
			"code":    codeStatus,
			"message": err.Error(),
		})
		return
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

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Book By ID",
		"data":    response,
	})
}
