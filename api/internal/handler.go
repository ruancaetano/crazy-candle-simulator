package internal

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func HandleNewCandleMessage(newCandleChannel chan Candle, repository *MongoRepository) ConsumerHandlerFunc {

	return func(message amqp.Delivery) {
		var candle Candle
		err := json.Unmarshal(message.Body, &candle)

		if err != nil {
			log.Println(err)
		}

		collection := repository.client.Database("crazy-candles").Collection("candles")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		_, err = collection.InsertOne(ctx, candle)

		if err != nil {
			log.Println(err)
		}

		newCandleChannel <- candle

	}
}
