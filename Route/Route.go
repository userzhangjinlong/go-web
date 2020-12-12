package Route

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"os"
	"web_go/App/Http/Controller"
)

const (
	MethodGet 		=	"GET"
	MethodHead		=	"HEAD"
	MethodPost		=	"POST"
	MethodPut		=	"PUT"
	MethodDelete	=	"DELETE"
	MethodPatch		=	"PATCH"
	MethodConnect	=	"CONNECT"
	MethodTract		=	"TRACE"
)

type Route struct {
	Method string
	Pattern string
	Callback interface{}
}


func setWebRoute(context *gin.Context) map[string][]Route {

	routes := map[string][]Route{
		"v1":{
			{MethodGet, "/index", (&Controller.Index{}).Index(context)},
			{MethodGet, "/index1", Controller.Index{}},
		},
	}


	return routes
}


func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
	pprof.Register(router)

	//路由注入上下文内容
	var context *gin.Context

	webRoute := setWebRoute(context)

	for group,routes := range webRoute{
		group := router.Group(group)
		{
			for i := 0; i < len(routes); i++  {
				fmt.Println(routes[i])
				os.Exit(1)
				switch routes[i][0] {
					case MethodGet:
						//这里后续写增加回调方法的工厂方法调用指定位置的回调方法
						group.GET(routes[i][1])
				}
			}
		}


	}

	return router
}
