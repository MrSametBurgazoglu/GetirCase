package database

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	MongoClient      *mongo.Client
	RecordCollection *mongo.Collection
	RedisClient      *redis.Client
}

func SetupDatabase() *Database {
	database := new(Database)
	mongoDatabase := SetupMongoDatabase()
	database.MongoClient = mongoDatabase.Client
	database.RecordCollection = mongoDatabase.RecordCollection
	redisDatabase := SetupRedisDatabase()
	database.RedisClient = redisDatabase.Client
	return database
}
