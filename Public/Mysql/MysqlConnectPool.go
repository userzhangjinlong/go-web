package Mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
	Configs "web_go/Public/Utils"
)

type MysqlConnectPool struct {
}

var (
	instance *MysqlConnectPool
	once sync.Once
	db *gorm.DB
	errDb error
)


/**
数据库单例连接池封装获取
*/
func GetInstance() *MysqlConnectPool {
	once.Do(func() {
		instance = &MysqlConnectPool{}
	})

	return instance
}

//加载mysql配置
func (this *MysqlConnectPool) InitDbPool() (result bool) {
	var config = Configs.GetContext().GetConfig()
	db, errDb = gorm.Open("mysql", config["database_mysql"]["source"])

	if errDb != nil {
		log.Fatal(errDb.Error())
		return false
	}
	//默认表名加s配置去掉
	db.SingularTable(true)
	//关闭数据库连接，db会自动被多个goroutine共享，可以不调用
	defer db.Close()
	log.Println("mysql:初始化连接成功")
	return true
}

//mysql外部开发端口
func (this *MysqlConnectPool) GetMysqlDb() (dbPool *gorm.DB, err error)  {
	return db,err
}
