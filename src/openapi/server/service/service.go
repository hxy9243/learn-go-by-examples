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
			Description: "The biting cold and the aching silence of the far North " +
				"become an unforgettable backdrop for Jack London's vivid, " +
				"rousing, superbly realistic wilderness adventure stories " +
				"featuring the author's unique knowledge of the Yukon and " +
				"the behavior of humans and animals facing nature at its cruelest.",
			Genre: "fiction",
			Tags:  []string{"fiction", "wilderness"},
		},
	}

	return &BookServiceImpl{
		Books: books,

		DefaultApiService: &openapi.DefaultApiService{},
	}
}

func (s *BookServiceImpl) GETBooks(ctx context.Context,
	isbn string, title string, author string, genre string,
) (openapi.ImplResponse, error) {

	return openapi.Response(http.StatusOK, s.Books), nil
}
