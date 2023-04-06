package main

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := NewApplication(connectToMongo())
	handler := NewHandler(app, mux.NewRouter())
	server := NewServer(handler)
	server.RegisterRoutes()
	server.Serve()
}

func connectToMongo() *mongo.Client {
	// create connection options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	log.Println("Mongo connected with success!")

	return c
}

