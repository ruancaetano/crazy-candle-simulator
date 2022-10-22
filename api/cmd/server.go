package main

import (
	"log"

	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal"
	internalAmqp "githbub.com/ruancaetano/crazy-candle-simulator/api/internal/amqp"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/amqp/handler"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/repositories"
)

func main() {
	newCandleChannel := make(chan entities.Candle)

	amqpConnection := internalAmqp.NewAmqpConnection()
	amqpConnection.Connect()
	defer amqpConnection.Disconnect()

	mongoRepository := repositories.NewMongoRepository()
	mongoRepository.Connect()
	defer mongoRepository.Disconnect()

	consumer := internalAmqp.NewAmqpConsumer(
		amqpConnection.Channel,
		&amqpConnection.Queue,
		handler.HandleNewCandleMessage(newCandleChannel, mongoRepository),
	)
	go consumer.Consume()

	server := internal.NewServer(newCandleChannel, mongoRepository)
	err := server.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
