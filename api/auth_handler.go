package api

import (
	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/abidkiller/hotel_reservation_backend/types"
	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthResponse struct {
	User  types.User `json:"user"`
	Token string     `json:"token"`
}

type AuthHandler struct {
	userStore db.UserStore
}

func (a *AuthHandler) NewAuthHandler(userStore db.UserStore) *AuthHandler {

}

func (a *AuthHandler) HandlerAuthenticate(c *fiber.Ctx) error {

}
