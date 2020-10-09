package Redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"sync"
	Configs "web_go/Public/Utils"
)

type RedisConnectPool struct {

}

var (
	once sync.Once
	instance *RedisConnectPool
	pool *redis.Pool
	errDB error
	db int
)

//单例模式连接redis
func GetInstance() *RedisConnectPool  {
	once.Do(func() {
		instance = &RedisConnectPool{}
	})

	return instance
}

func (this *RedisConnectPool) InitRedisPool(redisDb int) (result bool)  {
	var config = Configs.GetContext().GetConfig()
	var redisAddress = config["redis"]["root"]+":"+config["redis"]["port"]
	if  redisDb != 0{
		db = redisDb
	}else{
		db,_ = strconv.Atoi(config["redis"]["db"])
	}
	pool = &redis.Pool{
		MaxIdle:10000,
		MaxActive:0,
		IdleTimeout:300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddress,redis.DialPassword(config["redis"]["auth"]), redis.DialDatabase(db))
		},
	}

	defer pool.Close()
	log.Println("redis：实例化连接成功")
	return true
}

//获取redis连接
func (this *RedisConnectPool) GetRedisConnect() (redisPool *redis.Pool, err error)  {
	return pool,err
}