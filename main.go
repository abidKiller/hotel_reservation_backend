package main

import (
	"github.com/abidkiller/hotel_reservation_backend/api"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	app.Get("/hi", handleReq)
	apiv1.Get("/hi", handleReq)
	apiv1.Get("/user", api.HandleUser)
	app.Listen(":5000")
}

func handleReq(c *fiber.Ctx) error {
	return c.JSON(
		map[string]string{
			"msg": "working fine",
		},
	)

}
