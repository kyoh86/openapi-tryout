/*
 * openapi-tryout
 *
 * Try OpenAPI 3
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	server "github.com/GIT_USER_ID/GIT_REPO_ID/server"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := server.NewDefaultApiService()
	DefaultApiController := server.NewDefaultApiController(DefaultApiService)

	router := server.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8000", router))
}
