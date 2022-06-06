package service

import (
	"encoding/json"
	"log"
	"mq-server/model"
)

//从rabbitmq中接收信息，写入数据库
func CreateTask() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("task queue", true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//时刻处于监听状态，监听生产端的生产，所以要阻塞进程
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println()
			_ = d.Ack(false)
		}
	}()
}
