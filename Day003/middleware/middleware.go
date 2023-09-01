package middleware

import "github.com/gin-gonic/gin"

// it is a way to create Middleware

func Authenticate(ctx *gin.Context) {
	if !(ctx.Request.Header.Get("Token") == "auth") {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "Token Not Present",
		})
		return
	}
	ctx.Next()
}

//it is a another way to establish the connection with
// func Authenticate() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if !(ctx.Request.Header.Get("Token") == "auth") {
// 			ctx.AbortWithStatusJSON(500, gin.H{
// 				"message": "Token not",
// 			})
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
func AddHeder(ctx *gin.Context) {
	ctx.Writer.Header().Set("key", "Value")
	ctx.Next()
}
