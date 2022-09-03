package service

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hxy9243/openapi-example/server/gen/openapi"
)

type BookServiceImpl struct {
	*openapi.DefaultApiService

	Books []openapi.Book
}

func NewBookServiceImpl() *BookServiceImpl {
	books := []openapi.Book{
		{
			Id:    uuid.NewString(),
			Title: "The Call of the Wild",
			Isbn:  "9780241341490",
			Authors: []openapi.Author{
				{
					Name: "Jack London",
				},
			},
			Publisher: "Penguin Classics",
			Published: time.Date(2018, time.June, 07, 0, 0, 0, 0, time.UTC),
			Description: "First published in 1903, The Call of the Wild is regarded as Jack London's " +
				"masterpiece. Based on London's experiences as a gold prospector in the Canadian wilderness " +
				"and his ideas about nature and the struggle for existence, The Call of the Wild is a tale " +
				"about unbreakable spirit and the fight for survival in the frozen Alaskan Klondike.",
			Genre: "fiction",
			Tags:  []string{"classic", "fiction", "wilderness", "adventure", "animals", "literature"},
		},
	}

	return &BookServiceImpl{
		Books: books,

		DefaultApiService: &openapi.DefaultApiService{},
	}
}

// Override the default "unimplemented" response of GETBooks for the path /books
func (s *BookServiceImpl) GETBooks(ctx context.Context,
	isbn string, title string, author string, genre string,
) (openapi.ImplResponse, error) {

	return openapi.Response(http.StatusOK, s.Books), nil
}
