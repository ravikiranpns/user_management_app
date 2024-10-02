package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Postgres *sql.DB
var Mongo *mongo.Client

// InitPostgres - Initialize PostgreSQL connection
func InitPostgres() {
	var err error
	Postgres, err = sql.Open("postgres", "user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = Postgres.Ping(); err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}
	log.Println("Postgres connected")
}

// InitMongo - Initialize MongoDB connection
func InitMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Mongo = client
	log.Println("MongoDB connected")
}
