package internal

import (
	"crud-book/internal/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(route *gin.Engine) {
	route.GET("/health", controllers.CheckHealth)
}