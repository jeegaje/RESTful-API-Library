package main

import (
	moviecontroller "crud-movie/controllers/moviecontrollers"
	"crud-movie/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/movies", moviecontroller.Index)
	// r.GET("/api/products/:id", productcontroller.Show)
	// r.POST("/api/product", productcontroller.Create)
	// r.PUT("/api/product/:id", productcontroller.Update)
	// r.DELETE("/api/product", productcontroller.Delete)

	r.Run(":8181")
}
