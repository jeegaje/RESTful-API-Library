package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success access authors",
	})
}
