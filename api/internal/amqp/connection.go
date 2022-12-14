package amqp

import amqp "github.com/rabbitmq/amqp091-go"

type AmqpConnection struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
}

func NewAmqpConnection() *AmqpConnection {
	return &AmqpConnection{}
}

func (c *AmqpConnection) Connect() {
	conn, err := amqp.Dial("amqp://guest:guest@broker:5672/")
	if err != nil {
		panic(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	err = ch.ExchangeDeclare("candle.generated", "fanout", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	queue, err := ch.QueueDeclare("candle-generated-queue", true, false, false, false, nil)
	if err != nil {
		panic(err.Error())
	}

	err = ch.QueueBind(queue.Name, "", "candle.generated", false, nil)
	if err != nil {
		panic(err.Error())
	}

	c.Channel = ch
	c.Queue = queue
}

func (c *AmqpConnection) Disconnect() {
	c.Channel.Close()
}
