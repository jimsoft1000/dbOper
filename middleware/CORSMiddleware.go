package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//CORS Middleware 是 Gin 框架内置的一个中间件，可以用来处理跨域请求问题。当浏览器请求来自不同源的资源时，如果服务器没有正确地设置CORS响应头，则这些请求将被浏览器拒绝。
//可以使用CORS Middleware来设置CORS响应头，从而允许跨域请求。
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
