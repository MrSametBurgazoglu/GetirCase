package database

import (
	"context"
	"getir_case/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDatabase struct {
	Client           *mongo.Client
	RecordCollection *mongo.Collection
}

func SetupMongoDatabase() *MongoDatabase {
	database := new(MongoDatabase)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	database.Client = client
	database.RecordCollection = database.Client.Database(config.DBName).Collection(config.RecordCollectionName)
	return database
}
