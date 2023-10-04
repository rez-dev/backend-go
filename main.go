package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	// router.GET("/terms", GetTerms)

	router.Run("localhost:8080")
}
