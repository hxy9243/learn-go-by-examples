package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hxy9243/openapi-example/server/gen/openapi"
	"github.com/hxy9243/openapi-example/server/service"
)

func main() {
	service := service.NewBookServiceImpl()

	handler := openapi.NewRouter(
		openapi.NewDefaultApiController(service))

	address := "0.0.0.0:8000"

	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	log.Printf("Starting server on %s", address)
	log.Panic(http.ListenAndServe(address, handler))
}
