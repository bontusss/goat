package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bontusss/goat/internal/goat/models"
	"github.com/bontusss/goat/internal/goat/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// user = goat pass=goat2024
	mongoURI := "mongodb+srv://goat:goat2024@cluster0.idf2ghr.mongodb.net/"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	fmt.Println("mongodb connected")

	// Ensure the client will be disconnected
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	dbName := "goat"
	collectionName := "users"
	userRepo, err := repository.NewMongoDBUserRepository(client, dbName, collectionName)
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}

	// create a new user
	user := &models.User{
		Email:    "test@example.com",
		Password: "password",
	}

	// Register the user
	if err := userRepo.Register(user); err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}
	log.Println("User registered successfully")
}
