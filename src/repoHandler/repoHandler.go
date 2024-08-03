package repohandler

import (
	"context"
	"time"

	"github.com/pavr1/people/src/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoHandler struct {
	log    *log.Entry
	config *config.Config
	client *mongo.Client
}

func NewRepoHandler(log *log.Entry, config *config.Config) (*RepoHandler, error) {
	log.Info("Connecting to MongoDB...")

	client, err := connectToMongoDB(config)
	if err != nil {
		log.WithField("error", err).Error("Failed to connect to MongoDB")

		return nil, err
	}

	return &RepoHandler{
		log:    log,
		client: client,
	}, nil
}

func connectToMongoDB(config *config.Config) (*mongo.Client, error) {
	uri := config.MongoDB.Uri

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
