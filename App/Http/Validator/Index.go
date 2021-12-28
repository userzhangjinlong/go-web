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
	/**
	这里根据路由结尾"名称"作为action做不同路由调度 ps自己太菜了无法实现其他更灵活类型php可变变量调用方法 go实现只能使用
	Reflect但是貌似不推荐 后期可以做研究扩展让自己代码更优雅
	**/
	pathSlice := strings.Split(path, "/")
	last := pathSlice[len(pathSlice)-1]
	switch last {
		case "index":
			(&Controller.Index{}).Index(Context)

	}

}