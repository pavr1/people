package models

import (
	"context"

	"github.com/pavr1/people/config"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	config   *config.Config
	ID       string
	Name     string
	LastName string
	Age      int
}

func NewPerson(config *config.Config) Person {
	return Person{
		config: config,
	}
}

func (p *Person) Populate(name string, lastName string, age int) {
	p.Name = name
	p.LastName = lastName
	p.Age = age
}

func (p *Person) GetPersonList() ([]Person, error) {
	people := []Person{}
	// Set up the MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(p.config.MongoDB.Uri))
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to MongoDB")

		return nil, err
	}

	// Get a handle to the collection
	collection := client.Database(p.config.MongoDB.Database).Collection(p.config.MongoDB.Collection)

	// Find all documents in the collection
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.WithError(err).Fatal("Failed to find documents in MongoDB")

		return nil, err
	}

	defer cur.Close(context.Background())

	// Iterate over the documents and print their contents
	for cur.Next(context.Background()) {
		var doc bson.M
		err := cur.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, Person{
			ID:       doc["_id"].(string),
			Name:     doc["name"].(string),
			LastName: doc["lastName"].(string),
			Age:      doc["age"].(int),
		})
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return people, nil
}
