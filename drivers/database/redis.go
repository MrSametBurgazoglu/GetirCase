package database

import (
	"context"
	"fmt"
	"getir_case/config"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisDatabase struct {
	Client *redis.Client
}

func SetupRedisDatabase() *RedisDatabase {
	database := new(RedisDatabase)
	ctx := context.TODO()
	var opt *redis.Options
	var err error
	if config.RedisConnectionUrl != "" {
		opt, err = redis.ParseURL(config.RedisConnectionUrl)
	} else {
		opt = &redis.Options{
			Addr:     config.RedisConnectionAddr,
			Password: config.RedisPassword,
			DB:       config.RedisDatabaseNumber,
		}
	}

	rdb := redis.NewClient(opt)
	database.Client = rdb
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(pong, err)

	return database
}
