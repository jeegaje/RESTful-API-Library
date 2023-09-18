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
	r.GET("/api/movies/:id", moviecontroller.Show)
	r.POST("/api/movie", moviecontroller.Create)
	r.PUT("/api/movie/:id", moviecontroller.Update)
	r.DELETE("/api/movie", moviecontroller.Delete)

	r.Run(":8181")
}
