package moviecontroller

import (
	"crud-movie/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Index(c *gin.Context) {
	var movies []models.Movies

	models.DB.Find(&movies)
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func Show(c *gin.Context) {
	// var product models.Product

	// id := c.Param("id")
	// if err := models.DB.First(&product, id).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
	// 		return
	// 	default:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	// var product models.Product

	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// models.DB.Create(&product)
	// c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {
	// var product models.Product
	// id := c.Param("id")

	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update Data Failed"})
	// 	return

	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Data Success Update"})

}

func Delete(c *gin.Context) {
	// var product models.Product

	// // input := map[string]string{"id": "0"}
	// var input struct {
	// 	Id json.Number
	// }

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// id, _ := input.Id.Int64()
	// if models.DB.Delete(&product, id).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Delete Data failed"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Data Success Delete"})

}
