package Log

import "github.com/sirupsen/logrus"

//Info 级别日志
func Info(code int, data interface{}, desc string) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Info(desc)
}

//Warning 级别日志
func Warning(code int, data interface{}, desc string) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Warning(desc)
}

//Error 级别日志
func Error(code int, data interface{}, desc string) {
	logrus.WithFields(logrus.Fields{
		"code": code,
		"data": data,
	}).Error(desc)
}
