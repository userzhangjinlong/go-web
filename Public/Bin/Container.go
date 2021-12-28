package Container

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
	ConnectPoolFactory "web_go/Public/ConnectPool"
	Configs "web_go/Public/Utils"
	"web_go/Route"
)

type container interface {
	scanConfig()
}

type contain struct {
	config map[string]interface{} //Config目录配置
}

var (
	yamlPath string
	sMap sync.Map
)

/*
*初始化env配置
*/
func init()  {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	os.Getenv(".env")
	yamlPath = os.Getenv("YAML_PATH")

}


func Run()  {
	//map[string]string{}
	var c = contain{
		config: make(map[string]interface{}),
	}

	//容器扫描配置
	c.scanConfig()
	Configs.Instance().SetConfig(c.config)
	//获取配置文件
	Configs.Instance().GetConfig()
	appDebug := Configs.Instance().GetBool("appDebug")
	if appDebug == false{
		//关闭debug调试模式 启动等各种日志都输出记录到日志文件中

	}
	
	//加载mysql redis实例
	ConnectPoolFactory.NewMysql()
	ConnectPoolFactory.NewRedis()
	
	//扫描路由文件
	router := Route.RegisterRoutes()
	//扫描路由启动
	router.Run(Configs.Instance().GetString("proxy.port"))


}

//实现接口 扫描配置目录
func (this *contain) scanConfig() {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Panicln(err.Error())
	}

	resultMap := make(map[string]interface{})
	errYaml := yaml.Unmarshal(yamlFile, resultMap)
	if errYaml != nil {
		log.Panicln(errYaml.Error())
	}

	var configMap = this.config
	for _,v := range resultMap{
		for key, value := range v.(map[interface{}]interface{}){
			var tmpMap = make(map[string]interface{})
			switch value.(type) {
			case bool:
				tmpMap[key.(string)] = value.(bool)
				this.Set(key.(string), value.(bool))
			case interface{}:
				for k,val := range value.(map[interface{}]interface{}){
					tmpMap[k.(string)] = val.(string)
					this.Set(key.(string)+"."+k.(string), val.(string))
				}

			}

			configMap[key.(string)] = tmpMap
		}
	}
	this.Set("ymlConfig", configMap)

}

//容器设置值
func (this *contain) Set(key string, value interface{}) bool {
	sMap.Store(key, value)
	return true
}

//获取容器设置的值
func (this *contain) Get(key string) interface{} {
	val,ok := sMap.Load(key)
	if ok != false {
		return val
	}
	return ""
}
