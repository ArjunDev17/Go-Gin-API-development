package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("It is Firast API")
	router := gin.Default()
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.POST("/postdata", postData)
	router.PATCH("/patchData", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "i am PAtch function for ",
		})
	})
	router.Run()
}

// http://localehost:8080/getUrlData/name/Ram/age/18
func getUrlData(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.JSON(200, gin.H{
		"data": "getQrry ",
		"name": name,
		"age":  age,
	})
}

// http://localehost:8080/getQuerryString?name=Ram&age=18
func getQueryString(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")
	ctx.JSON(200, gin.H{
		"data": "getQrry ",
		"name": name,
		"age":  age,
	})
}
func postData(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := io.ReadAll(body) //ioutil.ReadAll(body)
	ctx.JSON(200, gin.H{
		"data":     "I am Post method",
		"bodyData": string(value),
	})
}
func getData(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": "Hi i am gin Framework",
	})
}
