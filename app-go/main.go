package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/rez-dev/backend-go/app-go/models"
	// "github.com/rez-dev/backend-go/app-go/controllers"
)

func main() {
	router := gin.Default()

	// router.GET("/terms", GetTerms)

	router.Run("localhost:8080")
}
