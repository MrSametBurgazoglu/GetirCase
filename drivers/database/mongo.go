package database

import (
	"context"
	"getir_case/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Database struct {
	Client           *mongo.Client
	RecordCollection *mongo.Collection
}

func SetupDatabase() *Database {
	database := new(Database)
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	database.Client = client
	database.RecordCollection = database.Client.Database(config.DBName).Collection(config.RecordCollectionName)
	return database
}
