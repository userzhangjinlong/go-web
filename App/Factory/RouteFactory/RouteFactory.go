package RouteFactory

import (
	"github.com/gin-gonic/gin"
	"web_go/App/Http/Controller"
)

func CreateRoute(directory string, action string) func(context *gin.Context) {
	(&Controller.directory{}).action
}
