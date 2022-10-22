package main

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	newCandleChannel := make(chan internal.Candle)

	conn := connectAmqpServer()
	defer conn.Close()

	ch := getAmqpChannel(conn)
	defer ch.Close()

	queue := getQueue(ch)

	consumer := internal.NewAmqpConsumer(ch, &queue, internal.HandleNewCandleMessage(newCandleChannel))

	go consumer.Consume()

	server := internal.NewServer(newCandleChannel)
	err := server.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}

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
