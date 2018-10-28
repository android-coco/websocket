package cache

import (
	"websocket/common"
	"websocket/module"
	"github.com/go-redis/redis"
)

var redisDb *redis.Client

func InitRedis() module.Error {
	redisConfig := common.Config.Section("redis")
	redisDb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Key("addr").MustString("localhost:6379"),
		Password: redisConfig.Key("passwd").MustString(""),
		DB:       redisConfig.Key("db").MustInt(0),
	})
	if _, err := redisDb.Ping().Result(); err != nil {
		return module.Error{ErrCode: common.ERROR_REDIS_COMMAND_PING, ErrMsg: err}
	}
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

// GetRedisDb 获取Redis连接实例
func GetRedisDb() *redis.Client {
	return redisDb
}
