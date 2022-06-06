package handlers

import (
	"errors"
	"github.com/micro/go-micro/v2/util/log"
)

//包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err := errors.New("userService--" + err.Error())
		log.Info(err)
		panic(err)
	}
}

//包装错误
func PanicIfTaskError(err error) {
	if err != nil {
		err := errors.New("taskService--" + err.Error())
		log.Info(err)
		panic(err)
	}
}
