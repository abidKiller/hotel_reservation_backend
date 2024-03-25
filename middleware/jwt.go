package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := godotenv.Load(".env")
		token, ok := c.GetReqHeaders()["X-Api-Token"]

		if !ok {
			return fmt.Errorf("token not found")
		}
		fmt.Println("token: ", token[0])
		claims, err := parseToken(token[0])

		if err != nil {
			return fmt.Errorf("token could be validated found")

		}
		expires := claims["expires"].(float64)

		if time.Now().Unix() > int64(expires) {
			return c.Status(http.StatusUnauthorized).JSON(
				"token Expired",
			)
		}
		userID, ok := claims["id"].(string)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(
				"token error",
			)
		}
		user, err := userStore.GetUserById(c.Context(), userID)
		if err != nil {
			return err
		}

		c.Context().SetUserValue("user", *user)

		return c.Next()
	}

}

func parseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid singing method", token.Header["alg"])
			return nil, fmt.Errorf("Unautherized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println(secret)
		return []byte(secret), nil
	},
	)
	if err != nil {
		fmt.Println("Failed to parse JWT token")
		return nil, err
	}
	if !token.Valid {
		fmt.Println("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unautherised")
	}

	return claims, nil
}
