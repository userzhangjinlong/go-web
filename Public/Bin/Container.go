package Container

import (
	ConnectPoolFactory "web_go/Public/ConnectPool"
	"web_go/Route"
	"web_go/Utils/Config"
)

//Run 初始化容器入口
func Run() {
	var c Config.Config

	//容器扫描配置
	c.GetInstance().ScanConfig()
	//获取配置文件
	c.GetInstance().GetConfig()
	appDebug := c.GetInstance().Get("appDebug")
	if appDebug == false {
		//关闭debug调试模式 启动等各种日志都输出记录到日志文件中

	}

	//加载mysql redis实例
	ConnectPoolFactory.NewMysql()
	ConnectPoolFactory.NewRedis()

	//扫描路由文件
	router := Route.RegisterRoutes()
	//扫描路由启动
	router.Run(c.GetInstance().GetString("proxy.port"))

}
