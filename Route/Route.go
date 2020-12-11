package Route

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
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


func getWebRoute() map[string][][]string{
	routes := map[string][][]string{
		"v1":{
			{MethodGet, "/index", "Index.Index"},
			{MethodGet, "/index1", "Index.Index1"},
		},
	}

	return routes
}


func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
	pprof.Register(router)

	webRoute := getWebRoute()

	for group,routes := range webRoute{
		group := router.Group(group)
		{
			for i := 0; i < len(routes); i++  {
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
