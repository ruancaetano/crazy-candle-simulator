package controllers

import (
	"context"
	"encoding/json"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

type GetCandlesController struct {
	repository *repositories.MongoRepository
}

func NewGetCandlesController(repository *repositories.MongoRepository) *GetCandlesController {
	return &GetCandlesController{
		repository,
	}
}

func (c *GetCandlesController) Execute(w http.ResponseWriter, r *http.Request) {
	collection := c.repository.Client.Database("crazy-candles").Collection("candles")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{
		"timestamp": bson.M{"$gte": time.Now().Add(-time.Second * 60)},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	result := []entities.Candle{}
	for cursor.Next(ctx) {
		var decodedCandle entities.Candle
		err := cursor.Decode(&decodedCandle)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, decodedCandle)
		// do something with result....
	}

	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
