package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rez-dev/backend-go/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/terms", controllers.GetTerms)
	router.GET("/users", controllers.GetUsers)
	router.GET("/tests", controllers.GetTests)
	router.Run("localhost:8080")
}
