package main

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/generator/internal"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn := connectAmqpServer()
	defer conn.Close()

	ch := getAmqpChannel(conn)
	defer ch.Close()

	queue := getQueue(ch)

	producer := internal.NewProducer("candle.generated", ch, &queue)

	generator := internal.NewCandleGenerator(time.Second, 0, 1000, producer)

	generator.Start()
}

func connectAmqpServer() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err.Error())
	}

	return conn

}

func getAmqpChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	return ch

}

func getQueue(ch *amqp.Channel) amqp.Queue {
	err := ch.ExchangeDeclare("candle.generated", "fanout", true, false, false, false, nil)
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

	return queue
}
