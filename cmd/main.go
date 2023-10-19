package cmd

import (
	"getir_case/api/handler"
	"getir_case/drivers/database"
	"getir_case/pkg/services"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func setupRouter() {

	db := database.SetupMongoDatabase()

	apiServices := services.SetupServices(db)

	handler.SetupHandlers(apiServices)

}

func main() {

	setupRouter()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
