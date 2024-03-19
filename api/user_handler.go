package api

import (
	"fmt"

	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/abidkiller/hotel_reservation_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); err != nil {
		return err
	}
	user, err := types.NewUsersFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.CreateUser(c.Context(), user)

	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
	// return c.JSON(use)

}
func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return nil
	}
	return c.JSON(map[string]string{"msg": fmt.Sprintf("deleted id: %d", id)})
}
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	// := context.Background()
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)

}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		updateReq types.UpdateUserReq
		userID    = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&updateReq); err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, updateReq); err != nil {
		return err
	}
	return c.JSON(map[string]string{"udpated ": userID})
}
