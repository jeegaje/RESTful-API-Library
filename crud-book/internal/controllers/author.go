package controllers

import (
	"crud-book/db"
	"crud-book/models"
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
		"code": 200,
		"data": responses,
	})
}
