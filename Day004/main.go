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
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hi ")
	router := gin.Default()
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", absolutePath, httpMethod, handlerName, nuHandlers)
	// }

	//2023/09/01 10:17:14 endpoint /getData GET main.getData 3

	writers, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(writers)
	// gin.DefaultWriter = io.MultiWriter(writers,os.Stdout)
	fmt.Println("above line will create file along with console also")

	//because of above two line one new file will be create and ou console
	//will be clear
	//now output in console like
	/*
			hi
		[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

		[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		 - using env:   export GIN_MODE=release
		 - using code:  gin.SetMode(gin.ReleaseMode)
	*/
	/*before ginLogging.log file-->
		hi
	[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

	[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
	 - using env:   export GIN_MODE=release
	 - using code:  gin.SetMode(gin.ReleaseMode)
	rs)
	[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
	[GIN-debug] Listening and serving HTTP on :8080

	*/
	router.GET("/getData", getData)
	router.Run()
}
func getData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "it is new ",
	})
}
