package main

import (
	"log"
	"net/http"

	server "github.com/kyoh86/openapi-tryout/simple/gen/server"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := server.NewDefaultApiService()
	DefaultApiController := server.NewDefaultApiController(DefaultApiService)

	router := server.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8000", router))
}
