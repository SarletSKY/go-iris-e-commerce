package main

import (
	"fmt"

	"github.com/SarletSKY/go-iris-e-commerce/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("zwxSimple")
	rabbitmq.PublishSimple("Hello ! ")
	fmt.Println("发送成功")
}
