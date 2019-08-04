package main

import (
	"fmt"
	"go-container/container"
	"go-container/controller"
)

func main() {
	container.Init()
	controller := container.Bean(controller.BeanName).(controller.Controller)
	fmt.Println(controller.HTTPGet())
}
