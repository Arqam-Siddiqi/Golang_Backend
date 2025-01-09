package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client
var Database *mongo.Database
var UserModel *mongo.Collection

// ConnectToMongo connects to a local MongoDB instance
func ConnectToMongo() {
	// Replace this with your MongoDB connection details
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/Golang")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Ping the MongoDB server to verify connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	MongoDB = client
	Database = client.Database("Golang")
	UserModel = Database.Collection("Users")

	fmt.Println("Successfully connected to MongoDB!")
}
