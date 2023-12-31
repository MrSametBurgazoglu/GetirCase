package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoConnectionString string
	DBName                string
	RecordCollectionName  string
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
}
