package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	MongoClient      *mongo.Client
	RecordCollection *mongo.Collection
	InMemoryDatabase *InMemoryDatabase
}

func SetupDatabase() *Database {
	database := new(Database)
	mongoDatabase := SetupMongoDatabase()
	database.MongoClient = mongoDatabase.Client
	database.RecordCollection = mongoDatabase.RecordCollection
	inMemoryDatabase := new(InMemoryDatabase)
	database.InMemoryDatabase = inMemoryDatabase
	return database
}
