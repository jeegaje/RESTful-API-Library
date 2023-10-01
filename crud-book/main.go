package main

import (
	"crud-book/db"
	"crud-book/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	r := gin.Default()
	internal.HealthRoute(r)
	internal.BookRoute(r)

	r.Run()
}
