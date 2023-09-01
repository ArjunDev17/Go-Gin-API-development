// Define formate of the log with GIN
/*
Loging in GIn
how default logging work in GIN
Define formate of the log of routes in gin
Define formate of the log with GIN
write logs to file in GIN
controlling log output colouring in  cosole with GIN
*/
// 11:12 -->
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"api_dev4.1/logger"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", absolutePath, httpMethod, handlerName, nuHandlers)
	}
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	writers, _ := os.Create("ginLogging.log")
	//gin.DefaultWriter = io.MultiWriter(writers)
	gin.DefaultWriter = io.MultiWriter(writers, os.Stdout)
	fmt.Println("above line will create file along with console also")

	router.Use(gin.LoggerWithFormatter(logger.FormatLogs))

	router.GET("/getData", getData)
	router.Run()
}
func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "it is new ",
	})
}
