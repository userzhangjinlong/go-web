package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web_go/Utils/Response"
)

type Index struct {
}

func (class *Index) Index(Context *gin.Context) {
	zero := 0
	x := 3 / zero
	fmt.Println("x=", x)
	// 这里随便模拟一条数据返回
	Response.Error(Context, 404, "惜败", gin.H{
		"newsType": "newsType",
		"page":     1,
		"limit":    20,
		"userIp":   "127.0.0.1",
		"title":    "门户首页公司新闻标题001",
		"content":  "门户新闻内容001",
	})
}
