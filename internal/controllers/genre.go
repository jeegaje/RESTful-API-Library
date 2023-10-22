package controllers

import (
	"crud-book/db"
	"crud-book/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllGenres(c *gin.Context) {
	var genres []*models.Genre
	var responses []map[string]any
	var codeStatus int

	err := db.DB.Find(&genres).Error
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
	for _, genre := range genres {
		response := map[string]any{
			"ID":   genre.ID,
			"name": genre.Name,
		}
		responses = append(responses, response)
	}
	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Genre",
		"data":    responses,
	})
}

func GetGenreById(c *gin.Context) {
	var genre *models.Genre
	var codeStatus int

	err := db.DB.Model(&genre).Where("id = ?", c.Param("id")).First(&genre).Error
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
		"ID":   genre.ID,
		"name": genre.Name,
	}

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Genre By Id",
		"data":    response,
	})
}

func CreateGenre(c *gin.Context) {
	var reqestBody *models.Genre

	err := c.ShouldBindJSON(&reqestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	genre := models.Genre{
		Name: reqestBody.Name,
	}

	err = db.DB.Create(&genre).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Create Genre Success",
		"data": models.Genre{
			Name: genre.Name,
			ID:   genre.ID,
		},
	})
}

func DeleteGenre(c *gin.Context) {

	err := db.DB.First(&models.Genre{}, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = db.DB.Delete(&models.Genre{}, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success delete genre",
	})
}

func UpdateGenre(c *gin.Context) {
	var payload *models.Genre

	err := db.DB.First(&models.Genre{}, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	updatedGenre := models.Genre{
		Name: payload.Name,
	}

	err = db.DB.Model(&models.Genre{}).Where("id = ?", c.Param("id")).Updates(updatedGenre).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success Update Genre",
	})
}
