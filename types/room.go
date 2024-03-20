package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

var RoomTypeString = map[RoomType]string{
	SingleRoomType:  "SingleRoom",
	DoubleRoomType:  "DoubleRoom",
	SeaSideRoomType: "SeaSideROom",
	DeluxeRoomType:  "DeluxeRoom",
}

type RoomCreationReq struct {
	Type      string  `json:"type"`
	BasePrice float64 `json:"basePrice"`
	HotelID   string  `json:"hotelID"`
	Size      string  `json:"size"` // small, normal, kingsize

}

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type      string             `bson:"type" json:"type"`
	Size      string             `bson:"size" json:"size"` // small, normal, kingsize
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	HotelID   primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

func (r *RoomCreationReq) ToRoom(hotelID primitive.ObjectID) *Room {
	return &Room{
		HotelID:   hotelID,
		Type:      r.Type,
		BasePrice: r.BasePrice,
		Size:      r.Size,
	}
}
