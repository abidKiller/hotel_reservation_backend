package db

import (
	"context"

	"github.com/abidkiller/hotel_reservation_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "rooms"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection

	hotelStore HotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection(roomCollection),
		hotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)
	//update the parent hotel
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	if err := s.hotelStore.UpdateHotel(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	res, err := s.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err := res.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
