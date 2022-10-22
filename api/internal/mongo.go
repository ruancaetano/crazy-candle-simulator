package internal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository() *MongoRepository {
	return &MongoRepository{}
}

func (r *MongoRepository) Connect() {
	log.Println("Connecting to mongodb")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:development@127.0.0.1:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r.client = client
}

func (r *MongoRepository) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return r.client.Disconnect(ctx)
}
