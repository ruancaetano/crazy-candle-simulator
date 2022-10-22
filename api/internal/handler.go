package internal

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func HandleNewCandleMessage(newCandleChannel chan Candle) ConsumerHandlerFunc {
	return func(message amqp.Delivery) {
		var candle Candle
		err := json.Unmarshal(message.Body, &candle)

		if err != nil {
			log.Println(err)
		}

		newCandleChannel <- candle

	}
}
