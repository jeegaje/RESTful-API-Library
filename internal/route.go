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
	route.GET("/books/:id", controllers.GetBookById)
}

func AuthorRoute(route *gin.Engine) {
	route.GET("authors", controllers.GetAllAuthors)
	route.GET("authors/:id", controllers.GetAuthorById)
	route.POST("author", controllers.CreateAuthor)
}

func GenreRoute(route *gin.Engine) {
	route.GET("/genres", controllers.GetAllGenres)
	route.GET("/genres/:id", controllers.GetGenreById)
	route.POST("/genre", controllers.CreateGenre)
}
