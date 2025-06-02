package db

import (
	"context"
	"log"
	"time"

	"github.com/Gitong23/go-fiber-hex-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB     *mongo.Database
	MongoClient *mongo.Client
)

func Connect() {
	cfg := config.AppConfig
	if cfg == nil {
		log.Fatal("Config not loaded")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	MongoClient = client
	MongoDB = client.Database(cfg.Mongo.Database)
	log.Println("Connected to MongoDB successfully")
}

func GetDatabase() *mongo.Database {
	return MongoDB
}

func Disconnect() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}
}
