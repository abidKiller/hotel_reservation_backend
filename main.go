package main

import (
	"context"
	"flag"
	"log"

	"github.com/abidkiller/hotel_reservation_backend/api"
	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://localhost:27017"
const DBNAME = "hotel_reservation"
const userCollection = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	userHanlder := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user/:id", userHanlder.HandleGetUser)
	apiv1.Get("/user", userHanlder.HandleGetUsers)

	app.Listen(*listenAddr)

}
