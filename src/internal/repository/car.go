package repository

import (
	"context"
	"log"

	// "fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nukahaha/car_store/src/internal/model"
)

type CarRepository struct {
	Collection *mongo.Collection
}

func NewCarRepository(database *mongo.Database, collectionName string) *CarRepository {
	collection := database.Collection(collectionName)

	return &CarRepository{
		Collection: collection,
	}
}

func (cr *CarRepository) GetAllCars() ([]model.Car, error) {
	var cars []model.Car
	cursor, err := cr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching cars from the database:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &cars); err != nil {
		log.Println("Error decoding cars:", err)
		return nil, err
	}

	log.Println("Fetched cars from the database:", cars)
	return cars, nil
}
func (cr *CarRepository) PostCarPage() ([]model.Car, error) {
	var cars []model.Car
	cursor, err := cr.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching cars from the database:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &cars); err != nil {
		log.Println("Error decoding cars:", err)
		return nil, err
	}

	log.Println("Fetched cars from the database:", cars)
	return cars, nil
}
func (cr *CarRepository) GetCarsSortedByYear() ([]model.Car, error) {
	var cars []model.Car
	opts := options.Find().SetSort(bson.D{{Key: "year", Value: -1}})
	cursor, err := cr.Collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &cars); err != nil {
		return nil, err
	}
	return cars, nil
}

func (cr *CarRepository) GetCarsSortedByModel() ([]model.Car, error) {
	var cars []model.Car
	opts := options.Find().SetSort(bson.D{{Key: "model_car", Value: 1}})
	cursor, err := cr.Collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &cars); err != nil {
		return nil, err
	}
	return cars, nil
}

func (cr *CarRepository) AddCar(car *model.Car) error {
	_, err := cr.Collection.InsertOne(context.Background(), car)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CarRepository) GetCarByID(carID string) (*model.Car, error) {
	// Преобразуем строку в ObjectID
	objectID, err := primitive.ObjectIDFromHex(carID)
	if err != nil {
		return nil, err
	}

	var car model.Car
	filter := bson.M{"_id": objectID}
	err = cr.Collection.FindOne(context.Background(), filter).Decode(&car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (cr *CarRepository) UpdateCar(car *model.Car) error {
	filter := bson.M{"_id": car.ID}
	update := bson.M{"$set": bson.M{
		"model":       car.Model_Car,
		"year":        car.Year,
		"cost":        car.Cost,
		"description": car.Description,
		"image":       car.Image,
		"brand":       car.Brand,
	}}

	_, err := cr.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CarRepository) DeleteCar(carID string) error {
	objectID, err := primitive.ObjectIDFromHex(carID)
	if err != nil {
		return err
	}

	_, err = cr.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
func (r *CarRepository) GetCarsByUserID(userID string) ([]model.Car, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var cars []model.Car
	for cursor.Next(context.TODO()) {
		var car model.Car
		if err := cursor.Decode(&car); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}
