package Validator

import (
	"github.com/gin-gonic/gin"
	"web_go/App/Http/Controller"
)

type Index struct {

}

func (class *Index) CreateRoute(Context *gin.Context){
	//这里通过上下文传入的 路由名称 实现调用不同的方法
	(&Controller.Index{}).Index(Context)
}