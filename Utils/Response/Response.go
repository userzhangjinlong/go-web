package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnJson(c *gin.Context, response Response) {
	if response.Code != 200 {
		c.Set("response", response)
	}
	c.JSON(http.StatusOK, response)
}

//成功
func Success(c *gin.Context, SuccessCode int, msg string, data interface{}) {
	var response = Response{Code: SuccessCode, Msg: msg, Data: data}
	ReturnJson(c, response)
}

//异常
func Error(c *gin.Context, ErrCode int, msg string, data interface{}) {
	var response = Response{Code: ErrCode, Msg: msg, Data: data}
	ReturnJson(c, response)
	//终止
	c.Abort()
}
