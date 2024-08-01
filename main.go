package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Todo struct {
	ID         int `json:"id" bson:"_id"`
	Completed bool `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello react & Go Application")


	err := godotenv.Load(".env");

	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client,err := mongo.Connect(context.Background(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(),nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your Application is connected to MONGODB")

	collection = client.Database("golang_db").collection("todos")

	app := fiber.New()

	//api_routes

	app.Get("/api/todos",getTodos)
	app.Post("/api/todos",createTodos)
	app.Patch("/api/todos/:id",updateTodos)
	app.Delete("/api/todos/:id",deleteTodos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

//api_controller

func getTodos( c *fiber.Ctx ) error {}
	


// func createTodos( c *fiber.Ctx ) error {}
// func updateTodos( c *fiber.Ctx ) error {}
// func deleteTodos( c *fiber.Ctx ) error {}



