package RouteInterface

import "github.com/gin-gonic/gin"

type Route interface {
	CreateRoute(handleFunc *gin.Context)
}
