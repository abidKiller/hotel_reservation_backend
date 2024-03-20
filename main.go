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

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	var (
		userStore  = db.NewMongoUserStore(client, db.DBNAME)
		hotelStore = db.NewMongoHotelStore(client, db.DBNAME)
		roomStore  = db.NewMongoRoomStore(client, db.DBNAME, hotelStore)

		userHanlder  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(hotelStore, roomStore)

		app = fiber.New(config)

		apiv1 = app.Group("/api/v1")
	)
	//user handlers
	apiv1.Post("/user", userHanlder.HandleCreateUser)
	apiv1.Get("/user/:id", userHanlder.HandleGetUser)
	apiv1.Get("/user", userHanlder.HandleGetUsers)
	apiv1.Delete("/user/:id", userHanlder.HandleDeleteUser)
	apiv1.Put("/user/:id", userHanlder.HandlePutUser)

	//hotel handlers
	apiv1.Post("/hotel", hotelHandler.HanldeCreateHotel)
	apiv1.Get("/hotel", hotelHandler.HanldeGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	app.Listen(*listenAddr)

}
