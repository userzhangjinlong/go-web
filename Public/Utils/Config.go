package Configs

import (
	"strings"
	"sync"
)

type Context struct {
	config map[string]interface{}
}

var (
	once sync.Once
	instance *Context
	sMap sync.Map
)

//单例获取上下文配置
func Instance() *Context  {
	once.Do(func() {
		instance = &Context{
			make(map[string]interface{}),
		}
	})
	return instance
}

//设置配置
func (this *Context) SetConfig(conf map[string]interface{})  {
	this.config = conf
}

//获取配置
func (this *Context) GetConfig() map[string]interface{} {
	return this.config
}

//字符串配置获取
func (this *Context) GetString(keyName string) string {
	returnString, ok := sMap.Load(keyName)

	if ok == false {
		keyNameArr := strings.Split(keyName, ".")
		for key, val := range this.config {
			if keyName == key && len(keyNameArr) == 1 {
				switch val.(type) {
				case string:
					returnString = val
				}
			} else {
				if len(keyNameArr) > 1 {
					if keyNameArr[0] == key {
						switch val.(type) {
						case map[string]string:
							for k, v := range val.(map[string]string) {
								if k == keyNameArr[1] {
									returnString = v
								}
							}
						case map[string]interface{}:
							for k, v := range val.(map[string]interface{}) {
								if k == keyNameArr[1] {
									returnString = v
								}
							}
						}
					}

				} else {
					switch val.(type) {
					case map[string]string:
						for k, v := range val.(map[string]string) {
							if k == keyName {
								returnString = v
							}
						}
					case map[string]interface{}:
						for k, v := range val.(map[string]interface{}) {
							if k == keyName {
								returnString = v
							}
						}
					}
				}

			}
		}
	}


	return returnString.(string)
}

//布尔类型获取
func (this *Context) GetBool(keyName string) bool{
	returnBool, ok := sMap.Load(keyName)
	if ok == false {
		keyNameArr := strings.Split(keyName, ".")
		for key, val := range this.config {
			if keyName == key && len(keyNameArr) == 1 {
				switch val.(type) {
				case bool:
					returnBool = val
				case map[string]interface{}:
					for k, v := range val.(map[string]interface{}) {
						if k == keyName {
							returnBool = v
						}
					}

				}
			} else {
				if len(keyNameArr) > 1 {
					if keyNameArr[0] == key {
						switch val.(type) {
						case map[string]string:
							for k, v := range val.(map[string]string) {
								if k == keyNameArr[1] {
									returnBool = v
								}
							}
						case map[string]interface{}:
							for k, v := range val.(map[string]interface{}) {
								if k == keyNameArr[1] {
									returnBool = v
								}
							}
						}
					}

				} else {
					switch val.(type) {
					case map[string]string:
						for k, v := range val.(map[string]string) {
							if k == keyName {
								returnBool = v
							}
						}
					case map[string]interface{}:
						for k, v := range val.(map[string]interface{}) {
							if k == keyName {
								returnBool = v
							}
						}
					}
				}

			}
		}
	}


	return returnBool.(bool)
}
