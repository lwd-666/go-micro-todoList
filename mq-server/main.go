package main

import (
	"mq-server/conf"
	"mq-server/service"
)

func main() {
	conf.Init()
	forerver := make(chan bool)
	service.CreateTask()
	<-forerver
}
