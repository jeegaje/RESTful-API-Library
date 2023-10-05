package controllers

import (
	"crud-book/db"
	"crud-book/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {
	var authors []*models.Author
	var responses []map[string]any

	db.DB.Find(&authors)

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

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    responses,
		"message": "Success get all authors",
	})
}

func GetAuthorById(c *gin.Context) {
	var author *models.Author
	fmt.Println(c.Param("id"))
	db.DB.Model(&author).Where("id = ?", c.Param("id")).First(&author)
	response := map[string]any{
		"firstName":   author.FirstName,
		"lastName":    author.LastName,
		"birthDate":   author.BirthDate,
		"nationality": author.Nationality,
		"email":       author.Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    response,
		"message": "Get author by id success",
	})
}
