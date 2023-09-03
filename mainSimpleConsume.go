package main

import "github.com/SarletSKY/go-iris-e-commerce/rabbitmq"

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("zwxSimple")
	rabbitmq.ConsumeSimple()
}
