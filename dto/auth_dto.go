package dto

import "github.com/abidkiller/hotel_reservation_backend/types"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthResponse struct {
	User  types.User `json:"user"`
	Token string     `json:"token"`
}
