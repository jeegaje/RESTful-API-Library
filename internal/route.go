package internal

import (
	"crud-book/internal/controllers"
	"crud-book/internal/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(route *gin.Engine) {
	route.GET("/health", controllers.CheckHealth)
}

func BookRoute(route *gin.Engine) {
	route.GET("/books", controllers.GetAllBooks)
	route.GET("/books/:id", handlers.GetBookById)
}

func AuthorRoute(route *gin.Engine) {
	route.GET("authors", controllers.GetAllAuthors)
	route.GET("authors/:id", controllers.GetAuthorById)
}

func GenreRoute(route *gin.Engine) {
	route.GET("/genres", controllers.GetAllGenres)
	route.GET("/genres/:id", controllers.GetGenreById)
}
