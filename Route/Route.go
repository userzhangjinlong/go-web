package Route

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"web_go/App/Factory/RouteFactory"
	"web_go/App/Http/Validator"
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
	MethodAny		=	"ANY"
)

type Route struct {
	Method string
	Pattern string
	Callback interface{}
}


func setWebRoute() map[string][]Route {

	routes := map[string][]Route{
		"v1":{
			//RouteFactory.CreateRoute(Controller.Index{})
			{MethodGet, "/index", RouteFactory.CreateRoute(&Validator.Index{})},
			//{MethodGet, "/index1", Controller.Index{}},
		},
	}


	return routes
}


func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
	pprof.Register(router)

	//路由注入上下文内容
	//var context *gin.Context

	webRoute := setWebRoute()

	for group,routes := range webRoute{
		group := router.Group(group)
		{
			for i := 0; i < len(routes); i++  {
				switch routes[i].Method {
					case MethodGet:
						//这里后续写增加回调方法的工厂方法调用指定位置的回调方法
						group.GET(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					case MethodPost:
						group.POST(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					case MethodPut:
						group.PUT(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					case MethodDelete:
						group.DELETE(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					case MethodAny:
						group.Any(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					default:
						//这里使用全局异常抛出500异常未定义方法不能请求路由
				}
			}
		}


	}

	return router
}
