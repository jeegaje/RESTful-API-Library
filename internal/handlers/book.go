package handlers

import (
	"crud-book/internal/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var codeStatus int

func GetBookById(c *gin.Context) {

	data, err := controllers.GetBookById(c)
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

	codeStatus = http.StatusOK
	c.JSON(codeStatus, gin.H{
		"code":    codeStatus,
		"message": "Success Get Book By ID",
		"data":    data,
	})
}
