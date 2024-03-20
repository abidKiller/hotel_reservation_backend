package db

import (
	"context"

	"github.com/abidkiller/hotel_reservation_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelCollection = "hotels"

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(context.Context, bson.M, bson.M) error
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)
	GetHotelByID(context.Context, primitive.ObjectID) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(dbname).Collection(hotelCollection),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	res, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, err

}

func (s *MongoHotelStore) GetHotelByID(ctx context.Context, oid primitive.ObjectID) (*types.Hotel, error) {
	//, err := s.collection.Find(ctx, filter)
	res := s.collection.FindOne(ctx, bson.M{"_id": oid})
	var hotel *types.Hotel
	if err := res.Decode(&hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}
