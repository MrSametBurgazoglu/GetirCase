package main

//	@title			GETIR CASE API
//	@version		0.0.1
//	@description	Getir Case Api
//	@host		getir-case-3p88.onrender.com
//	@schemes	http https
//	@BasePath	/

import (
	"context"
	"getir_case/api/handler"
	_ "getir_case/api/requests/MemoryRequests"
	_ "getir_case/api/requests/RecordRequests"
	_ "getir_case/api/responses/MemoryMapResponses"
	_ "getir_case/api/responses/RecordResponses"
	"getir_case/config"
	"getir_case/drivers/database"
	"getir_case/pkg/services"
	"io"
	"log"
	"net/http"
	"os/signal"
	"syscall"
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

func SetupServer(handler http.Handler) *http.Server {
	port := getPort()

	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,
	}

	setupRouter()

	fs := http.FileServer(http.Dir("./docs/swagger-ui/"))
	http.Handle("/swagger/", http.StripPrefix("/swagger", fs))
	homepage := `<html>Welcome to my website!<br>for swagger please move to <a href="https://getir-case-3p88.onrender.com/swagger/">Swagger</a></html>`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, homepage)
	})

	return &server

}

func main() {

	config.InitConfig()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := SetupServer(nil)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	<-ctx.Done()

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}

}
