package repohandler

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoHandler struct {
	client *mongo.Client
}

func NewRepoHandler() (*RepoHandler, error) {
	log.Info("Connecting to MongoDB...")

	client, err := connectToMongoDB()
	if err != nil {
		log.WithField("error", err).Error("Failed to connect to MongoDB")

		return nil, err
	}

	return &RepoHandler{
		client: client,
	}, nil
}

func connectToMongoDB() (*mongo.Client, error) {
	uri := "mongodb://localhost:27017" // Replace with your MongoDB connection URI

	log.WithField("uri", uri).Info("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")

		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.WithError(err).Fatal("Failed to ping MongoDB")
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client, nil
}
