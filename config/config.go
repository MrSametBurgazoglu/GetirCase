package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	MongoConnectionString string
	DBName                string
	RecordCollectionName  string
	RedisConnectionAddr   string
	RedisConnectionUrl    string
	RedisPassword         string
	RedisDatabaseNumber   int
	Port                  string
)

func InitConfig() {
	err := godotenv.Load("local.env")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("Error loading .env file")
		} else {
			log.Println("No .env file found, loading variables from environment")
		}
	}

	Port = os.Getenv("Port")
	MongoConnectionString = os.Getenv("MONGO_CONNECTION_STRING")
	DBName = os.Getenv("DB_NAME")
	RecordCollectionName = os.Getenv("RECORD_COLLECTION_NAME")
	RedisConnectionAddr = os.Getenv("REDIS_CONNECTION_ADDRESS")
	RedisConnectionUrl = os.Getenv("REDIS_CONNECTION_URL")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDatabaseNumber, err = strconv.Atoi(os.Getenv("REDIS_DATABASE_NO"))
	if err != nil {
		log.Println("REDIS_DATABASE_NO not set using default 0")
	}
}
