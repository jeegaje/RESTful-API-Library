package main

import (
	"crud-book/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test health",
		})
	})

	r.Run()
}
