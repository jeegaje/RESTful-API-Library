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

func CreateAuthor(c *gin.Context) {
	var reqestBody *models.Author

	err := c.ShouldBindJSON(&reqestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	author := models.Author{
		FirstName:   reqestBody.FirstName,
		LastName:    reqestBody.LastName,
		Email:       reqestBody.Email,
		Nationality: reqestBody.Nationality,
	}

	err = db.DB.Create(&author).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Create Author Success",
		"data": models.Author{
			FirstName:   author.FirstName,
			LastName:    author.LastName,
			Email:       author.Email,
			Nationality: author.Nationality,
		},
	})
}

func DeleteAuthor(c *gin.Context) {
	err := db.DB.First(&models.Author{}, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = db.DB.Delete(&models.Author{}, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success delete Author",
	})
}
