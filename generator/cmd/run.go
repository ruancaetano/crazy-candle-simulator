package main

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/generator/internal"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	defer ch.Close()

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

	producer := internal.NewProducer("candle.generated", ch, &queue)

	generator := internal.NewCandleGenerator(time.Second, 0, 1000, producer)

	generator.Start()
}
