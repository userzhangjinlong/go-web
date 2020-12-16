package Validator

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web_go/App/Http/Controller"
)

type Index struct {

}

/**
路由最终定义处理完成
 */
func (class *Index) CreateRoute(Context *gin.Context){
	path := Context.FullPath()
	//这里根据路由结尾作为action做不同路由调度
	pathSlice := strings.Split(path, "/")
	last := pathSlice[len(pathSlice)-1]
	switch last {
		case "index":
			(&Controller.Index{}).Index(Context)

	}

}