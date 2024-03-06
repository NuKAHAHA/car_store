// repository/wishlist_repository.go
package repository

import (
	"context"
	"errors"
	"github.com/nukahaha/car_store/src/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WishlistRepository struct {
	Collection *mongo.Collection
}

func NewWishlistRepository(database *mongo.Database, collectionName string) *WishlistRepository {
	collection := database.Collection(collectionName)
	return &WishlistRepository{
		Collection: collection,
	}
}

func (r *WishlistRepository) AddToWishlist(wishlist *model.Wishlist) error {
	_, err := r.Collection.InsertOne(context.TODO(), wishlist)
	return err
}

func (r *WishlistRepository) RemoveFromWishlist(userID, carID string) error {
	filter := bson.M{"user_id": userID, "car_id": carID}
	result, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": filter})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("wishlist item not found")
	}

	return nil
}

func (r *WishlistRepository) GetWishlistByUserID(userID string) ([]model.Wishlist, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var wishlistItems []model.Wishlist
	if err := cursor.All(context.TODO(), &wishlistItems); err != nil {
		return nil, err
	}

	return wishlistItems, nil
}
