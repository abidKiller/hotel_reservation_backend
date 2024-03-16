package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()
	app.Get("/hi", handleReq)
	app.Listen(":5000")
}

func handleReq(c *fiber.Ctx) error {
	return c.JSON(
		map[string]string{
			"msg": "working fine",
		},
	)

}
