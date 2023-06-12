package cache

import (
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisDB *redis.Client

// InitRedis 初始化redis连接
func InitRedis() {
	// redis配置信息
	redisAddr := viper.GetString("redis.address")
	redisPwd := viper.GetString("redis.password")
	redisDBName := viper.GetString("redis.dbName")
	db, _ := strconv.Atoi(redisDBName)
	// 新建连接
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPwd,
		DB:       db,
	})
	RedisDB = client
}
