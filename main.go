package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello react & Go Application");

	app := fiber.New();

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file: ", err) 
	}

	PORT := os.Getenv("PORT")

	log.Fatal(app.Listen(":"+PORT));
}