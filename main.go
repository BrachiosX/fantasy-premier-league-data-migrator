package main

import (
	"context"
	"os"

	datamigrator "github.com/BrachiosX/fantasy-premier-league-data-migrator/data-migrator"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	loadEnvironmentVariables()

	client := connectToMongoDB()

	database := getDatabase(client)

	datamigrator.Migrate(database)

	defer client.Disconnect(context.Background())
}

func loadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("No .env file found.")
	}
}

func connectToMongoDB() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		panic("No MONGO_URI environment variable found.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	return client
}

func getDatabase(client *mongo.Client) *mongo.Database {
	databaseName := os.Getenv("MONGODB_DATABASE")

	if databaseName == "" {
		panic("No MONGODB_DATABASE environment variable found.")
	}

	return client.Database(databaseName)
}
