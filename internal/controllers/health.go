package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Sistem berjalan dengan baik",
	})
}
