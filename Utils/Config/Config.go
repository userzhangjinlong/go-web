package Config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Config struct {
	conf map[string]interface{}
}

var (
	yamlPath string
	sMap     sync.Map
	once     sync.Once
	instance *Config
)

/*
*初始化env配置
 */
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	dir, _ := os.Getwd()
	os.Getenv(dir + "/.env")
	yamlPath = os.Getenv("YAML_PATH")

}

func (this *Config) GetInstance() *Config {
	once.Do(func() {
		instance = &Config{
			make(map[string]interface{}),
		}
	})

	return instance
}

func (this *Config) SetConfig(configMap map[string]interface{}) (result bool) {
	this.conf = configMap
	return true
}

func (this *Config) GetConfig() map[string]interface{} {
	return this.conf
}

//scanConfig 实现接口 扫描配置目录
func (this *Config) ScanConfig() {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Panicln(err.Error())
	}

	resultMap := make(map[string]interface{})
	errYaml := yaml.Unmarshal(yamlFile, resultMap)
	if errYaml != nil {
		log.Panicln(errYaml.Error())
	}

	for _, v := range resultMap {
		for key, value := range v.(map[interface{}]interface{}) {
			switch value.(type) {
			case bool:
				this.conf[key.(string)] = value.(bool)
				this.Set(key.(string), value.(bool))
			case interface{}:
				for k, val := range value.(map[interface{}]interface{}) {
					this.conf[k.(string)] = val.(string)
					this.Set(key.(string)+"."+k.(string), val.(string))
				}

			}

		}
	}
	this.Set("ymlConfig", this.conf)

}

//容器设置值
func (this *Config) Set(key string, value interface{}) bool {
	sMap.Store(key, value)
	return true
}

//获取容器设置的值
func (this *Config) Get(key string) interface{} {
	val, ok := sMap.Load(key)
	if ok != false {
		return val
	}
	return ""
}

//Get string获取容器设置的值
func (this *Config) GetString(key string) string {
	val, ok := sMap.Load(key)
	if ok != false {
		return val.(string)
	}
	return ""
}
