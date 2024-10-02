package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MongoClient *mongo.Client

func ConnectDB() {
	// PostgreSQL connection
	dsn := "host=localhost user=postgres password=postgres dbname=user_mgmt_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL!")
	}
	log.Println("PostgreSQL connected successfully!")
}

func ConnectMongoDB() {
	// MongoDB connection
	uri := "mongodb://localhost:27017" // Update with your MongoDB URI
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB!")
	}

	// Ping MongoDB to ensure the connection is established
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed!")
	}

	log.Println("MongoDB connected successfully!")
}
