package main

import (
	"log"
	"net/http"

	"github.com/hxy9243/openapi-example/server/gen/openapi"
	"github.com/hxy9243/openapi-example/server/service"
)

func main() {
	service := service.NewBookServiceImpl()

	handler := openapi.NewRouter(
		openapi.NewDefaultApiController(service))

	log.Printf("Starting server on 0.0.0.0:8000")
	log.Panic(http.ListenAndServe("0.0.0.0:8000", handler))
}
