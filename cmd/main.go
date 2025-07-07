package main

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRoutes()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
