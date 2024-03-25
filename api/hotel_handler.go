package api

import (
	"fmt"

	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/abidkiller/hotel_reservation_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		roomStore:  rs,
		hotelStore: hs,
	}
}

func (h *HotelHandler) HanldeCreateHotel(c *fiber.Ctx) error {
	var reqBody types.HotelCreationReq
	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	hotel := reqBody.ToHotel()

	hotel, err := h.hotelStore.InsertHotel(c.Context(), hotel)

	if err != nil {
		return err
	}

	for _, roomReq := range reqBody.Rooms {
		room := roomReq.ToRoom(hotel.ID)
		room, err = h.roomStore.InsertRoom(c.Context(), room)
		if err != nil {
			return err
		}
		hotel.Rooms = append(hotel.Rooms, room.ID)

	}

	return c.JSON(hotel)
}

func (h *HotelHandler) HanldeGetHotels(c *fiber.Ctx) error {
	// var queryParams struct {
	// 	Room   bool
	// 	Rating int
	// }
	// if err := c.QueryParser(&queryParams); err != nil {
	// 	return err
	// }

	// fmt.Println(queryParams)
	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(hotels)

}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.roomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(c.Context().Value("user").(types.User))

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	hotel, err := h.hotelStore.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}
