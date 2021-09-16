package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	username string
	password string
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/demo")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("demo").Collection("users")

	user := User{"consult", "dash1kj2h3hu23h"}

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}
