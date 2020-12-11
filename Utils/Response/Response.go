package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_go/Config/StatusCode"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{})  {
	c.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg": msg,
		"data": data,
	})
}


//成功
func Success(c *gin.Context, msg string, data interface{})  {
	ReturnJson(c, http.StatusOK, StatusCode.SUCCESS,msg, data);
}

//异常
func Error(c *gin.Context, msg string, data interface{})  {
	ReturnJson(c, http.StatusOK, StatusCode.ERROR, msg, data)
	//终止
	c.Abort()
}