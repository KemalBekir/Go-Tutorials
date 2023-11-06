package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	dbName string
}

func NewDatabase(mongoURI, dbName string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
		dbName: dbName,
	}, nil
}

func (db *Database) GetClient() *mongo.Client {
	return db.client
}

func (db *Database) GetDatabaseName() string {
	return db.dbName
}
