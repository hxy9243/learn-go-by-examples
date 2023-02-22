package main

import (
	"context"
	"log"

	client "github.com/hxy9243/openapi-example/client/gen/openapi"
)

func main() {
	apicli := client.NewAPIClient(&client.Configuration{
		Scheme: "http://",
		Servers: client.ServerConfigurations{
			client.ServerConfiguration{
				URL: "localhost:8000",
			},
		},
	})

	books, _, err := apicli.DefaultApi.GETBooks(context.TODO()).Execute()
	if err != nil {
		log.Fatalf("Error: getting books: %s", err)
	}

	for i, book := range books {
		log.Printf("Book %d info: %+v", i, book)
	}

	book, resp, err := apicli.DefaultApi.GETBook(context.TODO(), books[0].GetId()).
		BookDetails([]string{"book", "details", "example"}).
		BookAuthor("whatever").
		Execute()

	log.Printf("Getting book request: %+v", *resp.Request.URL)

	if err != nil {
		log.Fatalf("Error: getting book: %s", err)
	}

	log.Printf("Getting book info: %+v", book)

}
