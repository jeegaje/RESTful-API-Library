package moviecontroller

import (
	"crud-movie/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

func Index(c *gin.Context) {
	var movies []models.Movies

	models.DB.Find(&movies)
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func Show(c *gin.Context) {
	var movies models.Movies

	id := c.Param("id")
	if err := models.DB.First(&movies, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	fmt.Println(movies)
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func Create(c *gin.Context) {
	var movie models.Movies

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func Update(c *gin.Context) {
	var movie models.Movies
	id := c.Param("id")

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&movie).Where("id = ?", id).Updates(&movie).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update Data Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Data Success"})

}

func Delete(c *gin.Context) {
	var movie models.Movies

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&movie, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete Data Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Delete Success"})

}
