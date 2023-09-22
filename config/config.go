package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func InitializeDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URL")))
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/golangDB"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database(os.Getenv("DB_NAME")).Collection("users")

	fmt.Println("Connected to MongoDB")
	return client
}
