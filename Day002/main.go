package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("It is Firast API")
	router := gin.Default()
	// router.GET("/getData", getData)
	// router.GET("/getQueryString", getQueryString)
	auth := gin.BasicAuth(gin.Accounts{
		"user":  "pass",
		"user2": "pass2",
		"user3": "pass3",
	})
	router.GET("/getUrlData/:name/:age", getUrlData)

	//router.Run() is equilent
	//http.ListenAndServe(":8080", router) //if we are using this way colon is important

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
		// no need to write router.GET(-----) admin is enough
	}
	client := router.Group("/client")
	{

		client.GET("/getQueryString", getQueryString)
	}
	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
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
		"data": "taking data fro URL ",
		"name": name,
		"age":  age,
	})
}
func postData(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := io.ReadAll(body) //ioutil.ReadAll(body) depricated method
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
