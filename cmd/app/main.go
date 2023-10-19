package main

//	@title			GETIR CASE API
//	@version		0.0.1
//	@description	Getir Case Api
//	@host		localhost:8080
//	@BasePath	/

import (
	"fmt"
	"getir_case/api/handler"
	_ "getir_case/api/requests/RecordRequests"
	"getir_case/drivers/database"
	"getir_case/pkg/services"
	"html"
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

	db := database.SetupDatabase()

	apiServices := services.SetupServices(db)

	handler.SetupHandlers(apiServices)

}

func main() {

	port := getPort()

	setupRouter()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fs := http.FileServer(http.Dir("./docs/swagger-ui/"))
	http.Handle("/swagger/", http.StripPrefix("/swagger", fs))
	println("we handled it")
	println("port:", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
