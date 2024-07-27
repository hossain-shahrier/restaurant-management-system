package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance creates and returns a MongoDB client instance
func DBinstance() *mongo.Client {
	os.Setenv("MONGODB_URI", "mongodb+srv://hossainshahriersh:PPGnQqmmS9BkHt64@restaurantmanagementsys.9c84kba.mongodb.net/?retryWrites=true&w=majority&appName=RestaurantManagementSystem")
	MongoDb := os.Getenv("MONGODB_URI")
	if MongoDb == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// Client is a MongoDB client instance
var Client *mongo.Client = DBinstance()

// OpenCollection opens a MongoDB collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
