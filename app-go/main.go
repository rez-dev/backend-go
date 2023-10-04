package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rez-dev/backend-go/app-go/models"
)

var term = models.Term{}
var terms = models.Terms{}

func main() {
	router := gin.Default()

	// router.GET("/terms", GetTerms)

	router.Run("localhost:8080")
}
