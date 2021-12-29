package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_go/Utils/Log"
	"web_go/Utils/ReqParam"
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
			request := make(map[string]interface{})
			request["Host"] = context.Request.Host
			request["Ip"] = context.ClientIP()
			request["Header"] = context.Request.Header
			param, err := ReqParam.GetFormParam(context)
			if err == nil {
				request["Param"] = param
			}

			//目前做日志记录和增量追加 后期可考虑添加es做日志分析使用
			Log.Error(
				res.(Response.Response).Code,
				request,
				res.(Response.Response).Message)
		}
	}
}
