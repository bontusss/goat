package repository

import (
	"context"

	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// MongoDBUserRepository is a struct for MongoDB operations, encapsulating client and collection information.
type MongoDBUserRepository struct {
	Client     *mongo.Client     // MongoDB client for database access.
	collection *mongo.Collection // MongoDB collection for user documents.
}

// NewMongoDBUserRepository initializes a new MongoDBUserRepository with a given MongoDB client, database name, and collection name.
// It also ensures that an index on the email field is created to enforce uniqueness.
func NewMongoDBUserRepository(client *mongo.Client, dbName, collectionName string) (*MongoDBUserRepository, error) {
	ctx := context.Background()
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	// Create a unique index on the email field to ensure no duplicate emails are registered.
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // Index key
		Options: options.Index().SetUnique(true), // Enforce uniqueness
	})

	if err != nil {
		return nil, err
	}
	return &MongoDBUserRepository{Client: client, collection: collection}, nil
}

// Register adds a new user to the MongoDB collection. It hashes the user's password before saving.
func (r *MongoDBUserRepository) Register(user *models.User) error {
	ctx := context.Background()
	user.ID = uint(uuid.New().ID()) // Generate a unique ID for the user.

	// Hash the user's password for secure storage.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the new user document into the MongoDB collection.
	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// Login checks a user's credentials against the stored values in the MongoDB collection.
// If the credentials are valid, it returns the user object; otherwise, it returns an error.
func (r *MongoDBUserRepository) Login(email, password string) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}

	// Attempt to find the user by email.
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no document is found, return an invalid credentials error.
			return nil, goat.ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify the password against the hashed password stored in the database.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// If the password does not match, return an invalid credentials error.
			return nil, goat.ErrInvalidCredentials
		}
		return nil, err
	}
	return user, nil
}

func (r *MongoDBUserRepository) DeleteUser(id uint) error {
	ctx := context.Background()
	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBUserRepository) ResetPassword(email string, newPassword string) error {
	ctx := context.Background()
	_, err := r.collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": newPassword}})
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBUserRepository) UpdateUser(user *models.User) error {
	ctx := context.Background()
	_, err := r.collection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBUserRepository) GetUserByID(id uint) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

