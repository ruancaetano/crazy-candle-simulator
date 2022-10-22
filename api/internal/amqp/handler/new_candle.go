package handler

import (
	"context"
	"encoding/json"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal"
	internalAmqp "githbub.com/ruancaetano/crazy-candle-simulator/api/internal/amqp"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func HandleNewCandleMessage(newCandleChannel chan entities.Candle, repository *internal.MongoRepository) internalAmqp.ConsumerHandlerFunc {

	return func(message amqp.Delivery) {
		var candle entities.Candle
		err := json.Unmarshal(message.Body, &candle)

		if err != nil {
			log.Println(err)
		}

		collection := repository.Client.Database("crazy-candles").Collection("candles")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		_, err = collection.InsertOne(ctx, candle)

		if err != nil {
			log.Println(err)
		}

		newCandleChannel <- candle

	}
}
