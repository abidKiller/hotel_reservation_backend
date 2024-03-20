package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/abidkiller/hotel_reservation_backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type Dropper interface {
	Drop(ctx context.Context) error
}
type UserStore interface {
	Dropper

	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, values types.UpdateUserReq) error
	//UpdateUsers()
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbName string) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(dbName).Collection(userCollection),
	}
}
func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("dropping collection")
	if err := s.collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID) // type assertion type cast inserted
	return user, nil
}
func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("No Document found with id: %v", id)
		}
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("No Document found with id: %v", id)
	}

	return nil
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, values types.UpdateUserReq) error {
	// res, err := s.collection.UpdateMany(ctx, filter, update)
	// if err != nil {
	// 	if errors.Is(err, mongo.ErrNoDocuments) {
	// 		return nil, fmt.Errorf("No Document found")
	// 	}
	// 	return nil, err
	// }
	// if res.ModifiedCount == 0 {
	// 	return nil, fmt.Errorf("No Document updated")
	// }
	// return res.UpsertedCount, nil
	update := bson.M{"$set": values.ToBson()}
	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil

}
