package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//todo::执行中做某些事情
		fmt.Println("中间件开始执行了")
	}
}
