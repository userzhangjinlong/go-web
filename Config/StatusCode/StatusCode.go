package StatusCode

import "fmt"

const (
	SUCCESS int = 200
	ERROR	int = 400
)

//type Code map[int]interface{}
type StatusCode struct {
	code map[int]map[string]string
}

func (c *StatusCode) init() {
	c.code = map[int]map[string]string{
		200:{"en":"success", "zh":"成功"},
		400:{"en":"error", "zh":"服务器返回异常"},
	}
}

func GetCode(code int) (int,string) {
	fmt.Println(code)
	return 1,"test"
}