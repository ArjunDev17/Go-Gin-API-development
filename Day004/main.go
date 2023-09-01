/*
Loging in GIn
how default logging work in GIN
Define formate of the log of routes in gin
Define formate of the log with GIN
write logs to file in GIN
controlling log output colouring in  cosole with GIN
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hi ")
	router := gin.Default()
	router.GET("/getData", getData)
	router.Run()
}
func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "it is new ",
	})
}
