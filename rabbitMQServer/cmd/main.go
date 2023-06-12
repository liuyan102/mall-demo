package main

import (
	"mall-demo/config"
	"mall-demo/internal/cache"
	"mall-demo/internal/initialize"
	"mall-demo/internal/rabbitMQ"
	"mall-demo/rabbitMQServer/service"
)

func main() {
	config.InitConfig() // 初始化配置文件
	initialize.InitDB() // 初始化数据库
	cache.InitRedis()   // 初始化redis
	rabbitMQ.InitMQ()   // 初始化消息队列

	// 等待消费者消费结束
	forever := make(chan bool)
	service.SecKillService()
	<-forever
}
