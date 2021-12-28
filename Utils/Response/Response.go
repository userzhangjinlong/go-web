package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	ErrCode     int64 = 400
	SuccessCode int64 = 200
)

func ReturnJson(c *gin.Context, response Response) {
	if response.Code != 200 {
		c.Set("response", response)
	}
	c.JSON(http.StatusOK, response)
}

//成功
func Success(c *gin.Context, SuccessCode int, msg string, data interface{}) {
	var response = Response{Code: SuccessCode, Message: msg, Data: data}
	ReturnJson(c, response)
}

//异常
func Error(c *gin.Context, ErrCode int, msg string, data interface{}) {
	var response = Response{Code: ErrCode, Message: msg, Data: data}
	ReturnJson(c, response)
	//终止
	c.Abort()
}
