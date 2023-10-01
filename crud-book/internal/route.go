package internal

import (
	"crud-book/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(route *gin.Engine) {
	route.GET("/health", controllers.CheckHealth)
}

func BookRoute(route *gin.Engine) {
	route.GET("/books", controllers.GetAllBooks)
}

func AuthorRoute(route *gin.Engine) {
	route.GET("authors", controllers.GetAllAuthors)
}

func GenreRoute(route *gin.Engine) {
	route.GET("/genres", controllers.GetAllGenres)
}
