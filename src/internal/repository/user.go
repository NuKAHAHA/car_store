package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nukahaha/car_store/src/internal/model"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database, collectionName string) *UserRepository {
	collection := database.Collection(collectionName)

	return &UserRepository{
		Collection: collection,
	}
}

func (ur *UserRepository) Register(user *model.User) error {
	_, err := ur.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetByFieldMail(mail string) (*model.User, error) {
	var user model.User
	filter := bson.M{"mail": mail}
	err := ur.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByID(id string) (*model.User, error) {
	filter := bson.M{"_id": id}

	var user model.User
	err := ur.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
