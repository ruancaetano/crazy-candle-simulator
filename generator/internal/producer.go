package internal

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type AmqpProducer struct {
	channel  *amqp.Channel
	queue    *amqp.Queue
	exchange string
}

func NewProducer(exchange string, channel *amqp.Channel, queue *amqp.Queue) *AmqpProducer {

	return &AmqpProducer{
		channel,
		queue,
		exchange,
	}
}

func (p *AmqpProducer) SendMessage(body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.channel.PublishWithContext(ctx, p.exchange, p.queue.Name, false, false, amqp.Publishing{
		Timestamp:   time.Now(),
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return err
	}

	return nil
}
