package database

import (
	"github.com/redis/go-redis/v9"
)

type RedisDatabase struct {
	Client *redis.Client
}

func SetupRedisDatabase() *RedisDatabase {
	database := new(RedisDatabase)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	database.Client = rdb
	return database
}
