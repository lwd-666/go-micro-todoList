package model

import "github.com/streadway/amqp"

//rabbitmq的连接单例
var MQ *amqp.Connection

func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
