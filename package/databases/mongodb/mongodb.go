package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDbConnect(cfg configs.MongoDb) (*mongo.Client, error) {
	// Set time out for connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to mongoDb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:2717"))
	if err != nil || client == nil {
		return nil, fmt.Errorf("error, can't connect to mongodb with an error: %v", err)
	}
	defer log.Println("🍃mongodb has been connected, have a nice day!")
	return client, nil
}

func MongoDbDisconnect(client *mongo.Client) {
	ctx := context.Background()
	log.Println("🍃mongodb has been disconnected, good bye")

	// Disconnect from database
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
