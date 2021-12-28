package RouteFactory

import (
	"github.com/gin-gonic/gin"
	"web_go/App/Interface/RouteInterface"
)

func CreateRoute(directory interface{}) func(context *gin.Context) {
	if val, isOk := directory.(RouteInterface.Route); isOk {
		return val.CreateRoute
	}

	return nil
}
