/*
 * OpenAPI Library Demo App
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"DELETEBookCopy",
			strings.ToUpper("Delete"),
			"/books/{book-id}/copies/{copy-id}",
			c.DELETEBookCopy,
		},
		{
			"DELETEBooks",
			strings.ToUpper("Delete"),
			"/books",
			c.DELETEBooks,
		},
		{
			"DELETEUser",
			strings.ToUpper("Delete"),
			"/users/{user-id}",
			c.DELETEUser,
		},
		{
			"GETBook",
			strings.ToUpper("Get"),
			"/books/{book-id}",
			c.GETBook,
		},
		{
			"GETBookCopies",
			strings.ToUpper("Get"),
			"/books/{book-id}/copies",
			c.GETBookCopies,
		},
		{
			"GETBookCopy",
			strings.ToUpper("Get"),
			"/books/{book-id}/copies/{copy-id}",
			c.GETBookCopy,
		},
		{
			"GETBooks",
			strings.ToUpper("Get"),
			"/books",
			c.GETBooks,
		},
		{
			"GETUser",
			strings.ToUpper("Get"),
			"/users/{user-id}",
			c.GETUser,
		},
		{
			"GETUserBorrow",
			strings.ToUpper("Get"),
			"/users/{user-id}/borrows/{borrow-id}",
			c.GETUserBorrow,
		},
		{
			"GETUserBorrows",
			strings.ToUpper("Get"),
			"/users/{user-id}/borrows",
			c.GETUserBorrows,
		},
		{
			"GETUsers",
			strings.ToUpper("Get"),
			"/users/",
			c.GETUsers,
		},
		{
			"PATCHBookCopy",
			strings.ToUpper("Patch"),
			"/books/{book-id}/copies/{copy-id}",
			c.PATCHBookCopy,
		},
		{
			"PATCHUser",
			strings.ToUpper("Patch"),
			"/users/{user-id}",
			c.PATCHUser,
		},
		{
			"PATCHUserBorrow",
			strings.ToUpper("Patch"),
			"/users/{user-id}/borrows/{borrow-id}",
			c.PATCHUserBorrow,
		},
		{
			"POSTBookCopies",
			strings.ToUpper("Post"),
			"/books/{book-id}/copies",
			c.POSTBookCopies,
		},
		{
			"POSTBooks",
			strings.ToUpper("Post"),
			"/books",
			c.POSTBooks,
		},
		{
			"POSTUserBorrows",
			strings.ToUpper("Post"),
			"/users/{user-id}/borrows",
			c.POSTUserBorrows,
		},
		{
			"POSTUsers",
			strings.ToUpper("Post"),
			"/users/",
			c.POSTUsers,
		},
	}
}

// DELETEBookCopy -
func (c *DefaultApiController) DELETEBookCopy(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DELETEBookCopy(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DELETEBooks -
func (c *DefaultApiController) DELETEBooks(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DELETEBooks(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DELETEUser -
func (c *DefaultApiController) DELETEUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DELETEUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETBook -
func (c *DefaultApiController) GETBook(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETBook(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETBookCopies -
func (c *DefaultApiController) GETBookCopies(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETBookCopies(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETBookCopy -
func (c *DefaultApiController) GETBookCopy(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETBookCopy(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETBooks -
func (c *DefaultApiController) GETBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	isbnParam := query.Get("isbn")
	titleParam := query.Get("title")
	authorParam := query.Get("author")
	genreParam := query.Get("genre")
	result, err := c.service.GETBooks(r.Context(), isbnParam, titleParam, authorParam, genreParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETUser -
func (c *DefaultApiController) GETUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETUserBorrow -
func (c *DefaultApiController) GETUserBorrow(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETUserBorrow(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETUserBorrows -
func (c *DefaultApiController) GETUserBorrows(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETUserBorrows(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GETUsers -
func (c *DefaultApiController) GETUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GETUsers(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PATCHBookCopy -
func (c *DefaultApiController) PATCHBookCopy(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.PATCHBookCopy(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PATCHUser -
func (c *DefaultApiController) PATCHUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.PATCHUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PATCHUserBorrow -
func (c *DefaultApiController) PATCHUserBorrow(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.PATCHUserBorrow(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// POSTBookCopies -
func (c *DefaultApiController) POSTBookCopies(w http.ResponseWriter, r *http.Request) {
	copyParam := Copy{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&copyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCopyRequired(copyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.POSTBookCopies(r.Context(), copyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// POSTBooks -
func (c *DefaultApiController) POSTBooks(w http.ResponseWriter, r *http.Request) {
	bookParam := Book{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBookRequired(bookParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.POSTBooks(r.Context(), bookParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// POSTUserBorrows -
func (c *DefaultApiController) POSTUserBorrows(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.POSTUserBorrows(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// POSTUsers -
func (c *DefaultApiController) POSTUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.POSTUsers(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}