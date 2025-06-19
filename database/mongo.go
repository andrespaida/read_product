package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ProductCollection *mongo.Collection

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Missing MONGODB_URI in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	dbName := os.Getenv("DB_NAME")
    ProductCollection = client.Database(dbName).Collection("products")
	log.Println("✅ Connected to MongoDB!")
}