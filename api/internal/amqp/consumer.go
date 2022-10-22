package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type ConsumerHandlerFunc func(delivery amqp.Delivery)

type AmqpConsumer struct {
	channel *amqp.Channel
	queue   *amqp.Queue
	handler ConsumerHandlerFunc
}

func NewAmqpConsumer(channel *amqp.Channel, queue *amqp.Queue, handler ConsumerHandlerFunc) *AmqpConsumer {
	return &AmqpConsumer{
		channel,
		queue,
		handler,
	}
}

func (c *AmqpConsumer) Consume() {
	forever := make(chan bool)

	messages, err := c.channel.Consume(
		c.queue.Name, // amqp
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		log.Println(err.Error())
	}

	for msg := range messages {
		go c.handler(msg)
	}

	<-forever
}
