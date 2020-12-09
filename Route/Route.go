package Route

import (
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
			{MethodGet, "/index", "IndexController@index"},
			{MethodGet, "/index1", "IndexController@index1"},
		},
	}

	return routes
}


func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	webRoute := getWebRoute()

	for group,routes := range webRoute{
		group := router.Group(group)
		{
			for i := 0; i < len(routes); i++  {
				switch routes[i][0] {
					case MethodGet:
						group.GET(routes[i][1], func(context *gin.Context) {
							context.Status(200)
						})
				}
			}
		}


	}

	return router
}
