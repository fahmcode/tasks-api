package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	godotenv.Load(".env")

	db_uri := os.Getenv("MONGODB_URI")
	if len(db_uri) == 0 {
		db_uri = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	app.Listen(":" + port)
}
