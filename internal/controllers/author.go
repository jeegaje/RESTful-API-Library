package controllers

import (
	"crud-book/db"
	"crud-book/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllAuthors(c *gin.Context) {
	var authors []*models.Author
	var responses []map[string]any
	var codeStatus int

	err := db.DB.Find(&authors).Error
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

	for _, author := range authors {
		response := map[string]any{
			"firstName":   author.FirstName,
			"lastName":    author.LastName,
			"birthDate":   author.BirthDate,
			"nationality": author.Nationality,
			"email":       author.Email,
		}
		responses = append(responses, response)
	}

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Auhtor",
		"data":    responses,
	})
}

func GetAuthorById(c *gin.Context) {
	var author *models.Author
	var codeStatus int

	err := db.DB.Model(&author).Where("id = ?", c.Param("id")).First(&author).Error
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
		"firstName":   author.FirstName,
		"lastName":    author.LastName,
		"birthDate":   author.BirthDate,
		"nationality": author.Nationality,
		"email":       author.Email,
	}

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Author By ID",
		"data":    response,
	})
}
