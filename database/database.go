package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoTimeout = 10

func GetDB(client *mongo.Client) *mongo.Database {
	// Send patch MongoDB driver to include an method to get the Default Database.
	const database = "pismo-demo"

	return client.Database(database)
}

func Connect() (*mongo.Client, error) {
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		log.Fatal("Please set variable MONGO_URI.")
	}

	log.Println("Connecting to database.")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout*time.Second)

	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Println(err)

		return nil, fmt.Errorf("%w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)

		return nil, fmt.Errorf("%w", err)
	}

	log.Println("Database connected.")

	return client, nil
}

func Disconnect(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Println(err)

		return
	}
}
