package main

//	@title			GETIR CASE API
//	@version		0.0.1
//	@description	Getir Case Api
//	@host		localhost:8080
//	@BasePath	/

import (
	"getir_case/api/handler"
	_ "getir_case/api/requests/MemoryRequests"
	_ "getir_case/api/requests/RecordRequests"
	_ "getir_case/api/responses/MemoryMapResponses"
	_ "getir_case/api/responses/RecordResponses"
	"getir_case/config"
	"getir_case/drivers/database"
	"getir_case/pkg/services"
	"log"
	"net/http"
)

func getPort() string {
	port := config.Port
	if port == "" {
		port = "8080"
	}
	return port
}

func setupRouter() {

	db := database.SetupDatabase()

	apiServices := services.SetupServices(db)

	handler.SetupHandlers(apiServices)

}

func SetupServer(handler http.Handler) error {
	port := getPort()

	setupRouter()

	fs := http.FileServer(http.Dir("./docs/swagger-ui/"))
	http.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	return http.ListenAndServe(":"+port, handler)

}

func main() {

	config.InitConfig()

	err := SetupServer(nil)
	if err != nil {
		log.Fatalln(err)
	}
}
