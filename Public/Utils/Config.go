package Configs

import (
	"sync"
)

type Context struct {
	config map[string]map[string]string
}

var (
	once sync.Once
	instance *Context
)

//单例获取上下文配置
func GetContext() *Context  {
	once.Do(func() {
		instance = &Context{
			map[string]map[string]string{},
		}
	})
	return instance
}

//设置配置
func (this *Context) SetConfig(conf map[string]map[string]string)  {
	this.config = conf
}

//获取配置
func (this *Context) GetConfig() map[string]map[string]string {
	return this.config
}
