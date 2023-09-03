package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// url格式：amqp://用户名:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://zwxuser:zwxuser@127.0.0.1:5672/zwx"

// 1.创建结构体
type RabbitMQ struct {
	conn      *amqp.Connection // 连接
	channel   *amqp.Channel    // 通道
	MQurl     string           // 队列设置url
	QueueName string           // 队列名称
	Key       string
	Exchange  string // 交换机
}

// 2. 新建rabbitmq队列结构体
func NewRabbitMQ(queueName, key, exchange string) *RabbitMQ {
	// 建立连接
	var err error
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Key:       key,
		Exchange:  exchange,
		MQurl:     MQURL,
	}
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MQurl)
	rabbitmq.failOnErr(err, "连接出错")
	// 建立通道
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "通道出错")
	return rabbitmq
}

// 3.停止队伍
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 4.创建错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 5. 创建simple模式队列
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 通道申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false,
		false, //是否具有排他性
		false, // 是否堵塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 在通道中发布消息
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // 如果为true，根据exchange类型，若没有找到丰富和条件的队列，将转发回发送者
		false, // 如果为true，当exchange发送消息队列没有绑定消费者，则发还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 简单模式下消费代码
func (r *RabbitMQ) ConsumeSimple() {
	// 通道申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false,
		false, //是否具有排他性
		false, // 是否堵塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 在队列中进行消费
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    //用来区分多个消费者
		true,  // 是否自动应答
		false, // 是否具有排他性
		false, // 若设置为true，则可在相同的connection中发送消息给这个消费者
		false, //是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 处理队列消息
	exit := make(chan bool)
	go func() {
		// 循环执行消息
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()
	// 退出清执行ctrl+c
	log.Printf("[*] Waiting for messages,Please To exit press Ctrl + c")
	<-exit
}
