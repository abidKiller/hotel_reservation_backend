package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type HotelCreationReq struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Rooms    []RoomCreationReq `json:"rooms"`
	Rating   int               `json:"rating"`
}

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

func (h *HotelCreationReq) ToHotel() *Hotel {
	return &Hotel{
		Name:     h.Name,
		Location: h.Location,
		Rooms:    []primitive.ObjectID{},
		Rating:   h.Rating,
	}
}
