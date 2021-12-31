package Log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"web_go/Utils/Config"
	"web_go/Utils/File"
)

var configs Config.Config

func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	dir, _ := os.Getwd()
	//获取配置文件
	//todo::实现容器单列模式注入配置直接获取 不需要每次使用都扫描获取配置
	configs.GetInstance().ScanConfig()
	configs.GetInstance().GetConfig()

	filePath := dir + configs.GetInstance().GetString("log.path") + "/logs/"
	dirErr := File.CreateIfNotExistDir(filePath)
	if dirErr == false {
		//todo::目录创建失败异常操作
	}
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 天轮转一个新文件，保留最近 15 天半个月的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		filePath+"%Y%m%d%H%M"+".log",
		rotatelogs.WithLinkName(filePath+"/logs"),
		rotatelogs.WithMaxAge(time.Hour*24*15),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logrus.SetOutput(writer)

	//todo::后续考虑将日志投递到elk
}

//Info 级别日志
func Info(code int, data interface{}, msg string) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Info(msg)
}

//Warning 级别日志
func Warning(code int, data interface{}, msg string) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Warning(msg)
}

//Error 级别日志
func Error(code int, data interface{}, msg interface{}) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Error(msg)
}

func logToEs() {

}
