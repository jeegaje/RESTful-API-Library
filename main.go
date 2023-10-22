package main

import (
	"crud-book/db"
	"crud-book/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	route := gin.Default()
	internal.HealthRoute(route)
	internal.BookRoute(route)
	internal.AuthorRoute(route)
	internal.GenreRoute(route)

	route.Run(":8081")
}
