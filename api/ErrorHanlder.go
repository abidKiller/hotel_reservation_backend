package api

import (
	"net/http"

	"github.com/abidkiller/hotel_reservation_backend/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// known error
	if apiError, ok := err.(errors.Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}

	// unknown error
	apiError := errors.NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}
