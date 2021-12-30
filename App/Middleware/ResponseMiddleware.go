package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"web_go/App/ErrCode"
	"web_go/Utils/Log"
	"web_go/Utils/ReqParam"
	"web_go/Utils/Response"
)

type requestData struct {
	Host   string
	Ip     string
	Header interface{}
	Method string
	Param  interface{}
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求日志参数
		var request requestData
		request.Host = context.Request.Host
		request.Ip = context.ClientIP()
		request.Header = context.Request.Header
		request.Method = context.Request.Method
		param, err := ReqParam.GetFormParam(context)
		if err == nil {
			request.Param = param
		}

		defer func() {
			if sysErr := recover(); sysErr != nil {
				errMsg := fmt.Sprintf("致命错误：%s", sysErr)
				stack := debug.Stack()
				Log.Error(ErrCode.SystemError, request, errMsg+"\n"+string(stack))
				var response Response.Response
				response.Code = ErrCode.SystemError
				response.Msg = ErrCode.SystemErrorMsg
				response.Data = ""
				context.JSON(ErrCode.SystemError, response)
			}
		}()
		context.Next()
		//http200 业务响应code
		res, resErr := context.Get("response")
		if res != nil && resErr != false {
			//todo::系统异常做日志文件切割记录 后续可以做消息投递 扩展ELK搜索
			//目前做日志记录和增量追加 后期可考虑添加es做日志分析使用
			Log.Error(
				res.(Response.Response).Code,
				request,
				res.(Response.Response).Msg)
		}
	}
}
