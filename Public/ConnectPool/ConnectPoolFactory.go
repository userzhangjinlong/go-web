package ConnectPoolFactory

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
	Configs "web_go/Public/Utils"
)

type Pool interface {
	GetInstance() *ConnectPool
	InitConnectPool() bool
	GetConnectLibrary() (interface{}, error)
}

var (
	once sync.Once
	instance *ConnectPool
	errDb error
	db *gorm.DB
	pool *redis.Pool
	redisDb int
	dbType string
)

type ConnectPool struct {

}

func (this *ConnectPool)GetInstance() *ConnectPool {
	once.Do(func() {
		instance = &ConnectPool{}
	})

	return instance
}

func (this *ConnectPool)InitConnectPool() (result bool) {
	var config = Configs.GetContext().GetConfig()
	switch dbType {
		case "mysql":
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
		case "redis":
			var redisAddress = config["redis"]["root"]+":"+config["redis"]["port"]
			pool = &redis.Pool{
				MaxIdle:10000,
				MaxActive:0,
				IdleTimeout:300,
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", redisAddress,redis.DialPassword(config["redis"]["auth"]), redis.DialDatabase(redisDb))
				},
			}

			defer pool.Close()
			log.Println("redis：实例化连接成功")


	}
	return true
}

func (this *ConnectPool)GetConnectLibrary() (res interface{},err error) {

	switch dbType {
		case "mysql":
			return db,err
		case "redis":
			return pool,err
		default:
			return db,err
	}
}


func NewConnect(connect string) *ConnectPool {
	dbType = connect
	return &ConnectPool{}
}