package repository

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nukahaha/car_store/src/internal/configuration"
)

type Database struct {
	Mongo *mongo.Client
}

func NewDatabase(databaseConfiguration *configuration.DatabaseConfiguration) (*Database, error) {
	db := &Database{}
	var err error

	db.Mongo, err = db.connectToMongoDB(databaseConfiguration)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Close() error {
	if d.Mongo != nil {
		err := d.Mongo.Disconnect(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Database) connectToMongoDB(databaseConfiguration *configuration.DatabaseConfiguration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := *databaseConfiguration.ConnectionURI
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
