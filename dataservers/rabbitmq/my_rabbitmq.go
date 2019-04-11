package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel  *amqp.Channel //链接
	Name     string        //queue的名字
	exchange string        //exchange的名字
}

//创建新的rabbitmq对象
func New(s string) *RabbitMQ {
	conn, e := amqp.Dial(s)
	if e != nil {
		panic(e)
	}

	ch, e := conn.Channel()
	if e != nil {
		panic(e)
	}
	q, e := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if e != nil {
		panic(e)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

//让rabbitmq绑定exchange
func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil)

	if e != nil {
		panic(e)
	}
	q.exchange = exchange
}

//生产者发送消息到rabbitmq的queue中
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish("",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}
}

//发送消息到相应的exchange中
func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(exchange,
		"",
		false,
		false, amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}
}

//定义消费者接收的方法
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if e != nil {
		panic(e)
	}
	return c
}

//rabbitmq关闭的方法
func (q *RabbitMQ) Close() {
	e := q.channel.Close()
	if e != nil {
		fmt.Println(e)
	}
}
