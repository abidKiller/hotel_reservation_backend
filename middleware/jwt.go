package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("JWT Auth ...")
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}
	token, ok := c.GetReqHeaders()["X-Api-Token"]

	if !ok {
		return fmt.Errorf("token not found")
	}
	fmt.Println("token: ", token[0])
	secret := os.Getenv("JWT_SECRET")
	fmt.Println("secret ", secret)

	return c.Next()
}

func parseToken(tokenStr string) string {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {

			return 
		}
	)
}
