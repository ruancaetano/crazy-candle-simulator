package main

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal"
	internalAmqp "githbub.com/ruancaetano/crazy-candle-simulator/api/internal/amqp"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/amqp/handler"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	"log"
)

func main() {
	newCandleChannel := make(chan entities.Candle)

	amqpConnection := internalAmqp.NewAmqpConnection()
	amqpConnection.Connect()
	defer amqpConnection.Disconnect()

	mongoRepository := internal.NewMongoRepository()
	mongoRepository.Connect()
	defer mongoRepository.Disconnect()

	consumer := internalAmqp.NewAmqpConsumer(
		amqpConnection.Channel,
		&amqpConnection.Queue,
		handler.HandleNewCandleMessage(newCandleChannel, mongoRepository),
	)
	go consumer.Consume()

	server := internal.NewServer(newCandleChannel)
	err := server.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
