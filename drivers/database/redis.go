package database

import (
	"getir_case/config"
	"github.com/redis/go-redis/v9"
)

type RedisDatabase struct {
	Client *redis.Client
}

func SetupRedisDatabase() *RedisDatabase {
	database := new(RedisDatabase)
	rdb := redis.NewClient(&redis.Options{
		Username: config.MongoUserName,
		Addr:     config.RedisConnectionAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDatabaseNumber,
	})
	database.Client = rdb
	return database
}
