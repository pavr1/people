package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID       string
	Name     string
	LastName string
	Age      int
}

func NewPerson(id string, name string, lastName string, age int) Person {
	return Person{
		ID:       id,
		Name:     name,
		LastName: lastName,
		Age:      age,
	}
}

func GetPersonList() ([]Person, error) {
	people := []Person{}
	// Set up the MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Connect to the MongoDB server
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Access the "people" collection
	collection := client.Database("Person").Collection("Person")

	// Find all documents in the collection
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	// Iterate over the results
	for cur.Next(ctx) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, person)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return people, nil
}
