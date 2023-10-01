package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]any{
			"title":       "test title",
			"description": "test description",
		},
	})
}
