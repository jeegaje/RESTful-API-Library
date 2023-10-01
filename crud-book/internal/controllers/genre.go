package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGenres(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success connect to Author Controllers",
	})
}
