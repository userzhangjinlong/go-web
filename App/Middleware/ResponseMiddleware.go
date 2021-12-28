package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_go/Utils/Response"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		//http 响应code
		httpCode := context.Writer.Status()
		if httpCode != http.StatusOK {
			//todo::非业务异常捕获处理
			fmt.Println("响应码")
			fmt.Println(httpCode)
		}

		//http200 业务响应code
		res, err := context.Get("response")
		if res != nil && err != false {
			//todo::系统异常做日志文件切割记录 后续可以做消息投递 扩展ELK搜索
			code := res.(Response.Response).Code
			msg := res.(Response.Response).Message
			fmt.Println(code, msg)
		}
	}
}
