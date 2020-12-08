package Container

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	ConnectPoolFactory "web_go/Public/ConnectPool"
	Configs "web_go/Public/Utils"
)

type container interface {
	scanConfig()
}

type contain struct {
	config map[string]map[string]string //Config目录配置
}

var yamlPath string

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
	run()
}

func run()  {
	var c = contain{
		config: map[string]map[string]string{},
	}

	//容器扫描配置
	c.scanConfig()
	Configs.GetContext().SetConfig(c.config)
	//加载mysql redis实例
	ConnectPoolFactory.NewMysql()
	ConnectPoolFactory.NewRedis()

	//扫描路由文件


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
			var tmpMap = make(map[string]string)
			for k,val := range value.(map[interface{}]interface{}){
				tmpMap[k.(string)] = val.(string)
			}
			configMap[key.(string)] = tmpMap
		}
	}

}
