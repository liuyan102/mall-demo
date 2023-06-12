package rabbitMQ

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var MQ *amqp.Connection

// InitMQ 初始化消息队列连接
func InitMQ() {
	head := viper.GetString("rabbitmq.head")
	address := viper.GetString("rabbitmq.address")
	username := viper.GetString("rabbitmq.username")
	password := viper.GetString("rabbitmq.password")

	connString := fmt.Sprintf("%s://%s:%s@%s/", head, username, password, address)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
