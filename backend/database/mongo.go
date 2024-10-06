package database

import (
	"context"
	"log"
	"todo-go/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectDB() {
    clientOptions := options.Client().ApplyURI(config.GetMongoURI())
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    Collection = client.Database("golang_db").Collection("todos")
    log.Println("Connected to MongoDB")
}
