package controllers

import (
	"crud-book/db"
	"crud-book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGenres(c *gin.Context) {
	var genres []*models.Genre
	var responses []map[string]any

	db.DB.Find(&genres)
	for _, genre := range genres {
		response := map[string]any{
			"ID":   genre.ID,
			"name": genre.Name,
		}
		responses = append(responses, response)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    responses,
		"message": "Success get all genres",
	})
}

func GetGenreById(c *gin.Context) {
	var genre *models.Genre

	db.DB.Model(&genre).Where("id = ?", c.Param("id")).First(&genre)
	response := map[string]any{
		"ID":   genre.ID,
		"name": genre.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    response,
		"message": "Get route by id success",
	})
}
