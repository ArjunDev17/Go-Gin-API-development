package main

import (
	"api_3/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New() // Create a new Gin router

	//router.Use(middleware.Authenticate)
	// admin := router.Group("/admin", middleware.Authenticate)
	// {
	// 	admin.GET("/getData", getData)
	// 	admin.GET("/getData1", getData1)
	// }

	router.GET("/getData", middleware.Authenticate, middleware.AddHeder, getData)
	router.GET("/getData1", getData1) // Changed the endpoint path
	router.GET("/getData2", getData2)

	router.Run(":8080") // Start the server on port 8080
}

func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hi, I am GetData Method",
	})
}

func getData1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hi, I am GetData Method 1",
	})
}

func getData2(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hi, I am GetData Method 2",
	})
}
