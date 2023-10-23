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
		"message": "Success Get Books",
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

func CreateBook(c *gin.Context) {
	var requestBody *models.Book

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:           requestBody.Title,
		Description:     requestBody.Description,
		BookLanguage:    requestBody.BookLanguage,
		PublicationYear: requestBody.PublicationYear,
		Publisher:       requestBody.Publisher,
		Isbn:            requestBody.Isbn,
		PageCount:       requestBody.PageCount,
		Price:           requestBody.Price,
		Availability:    requestBody.Availability,
		AverageRating:   requestBody.AverageRating,
		AuthorId:        requestBody.AuthorId,
		GenreId:         requestBody.GenreId,
	}

	err = db.DB.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Create Book Success",
		"data": models.Book{
			Title:           book.Title,
			Description:     book.Description,
			BookLanguage:    book.BookLanguage,
			PublicationYear: book.PublicationYear,
			Publisher:       book.Publisher,
			Isbn:            book.Isbn,
			PageCount:       book.PageCount,
			Price:           book.Price,
			Availability:    book.Availability,
			AverageRating:   book.AverageRating,
			AuthorId:        book.AuthorId,
			GenreId:         book.GenreId,
			CreatedAt:       book.CreatedAt,
			UpdatedAt:       book.UpdatedAt,
		},
	})
}

func DeleteBook(c *gin.Context) {
	err := db.DB.First(&models.Book{}, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = db.DB.Delete(&models.Book{}, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success delete Book",
	})
}
