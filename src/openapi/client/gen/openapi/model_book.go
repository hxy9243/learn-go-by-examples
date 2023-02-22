/*
OpenAPI Library Demo App

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Book struct for Book
type Book struct {
	// the id of the book
	Id *string `json:"id,omitempty"`
	// the id of the book
	Title *string `json:"title,omitempty"`
	// ISBN book number
	Isbn *string `json:"isbn,omitempty"`
	// the id of the book
	Authors []Author `json:"authors,omitempty"`
	// name of the publisher
	Publisher *string `json:"publisher,omitempty"`
	// time the book is published
	Published *time.Time `json:"published,omitempty"`
	// the description of the book
	Description *string `json:"description,omitempty"`
	// the genre of this book
	Genre *string `json:"genre,omitempty"`
	// the tags of the book
	Tags []string `json:"tags,omitempty"`
}

// NewBook instantiates a new Book object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBook() *Book {
	this := Book{}
	return &this
}

// NewBookWithDefaults instantiates a new Book object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookWithDefaults() *Book {
	this := Book{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Book) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Book) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Book) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Book) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Book) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Book) SetTitle(v string) {
	o.Title = &v
}

// GetIsbn returns the Isbn field value if set, zero value otherwise.
func (o *Book) GetIsbn() string {
	if o == nil || isNil(o.Isbn) {
		var ret string
		return ret
	}
	return *o.Isbn
}

// GetIsbnOk returns a tuple with the Isbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetIsbnOk() (*string, bool) {
	if o == nil || isNil(o.Isbn) {
		return nil, false
	}
	return o.Isbn, true
}

// HasIsbn returns a boolean if a field has been set.
func (o *Book) HasIsbn() bool {
	if o != nil && !isNil(o.Isbn) {
		return true
	}

	return false
}

// SetIsbn gets a reference to the given string and assigns it to the Isbn field.
func (o *Book) SetIsbn(v string) {
	o.Isbn = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *Book) GetAuthors() []Author {
	if o == nil || isNil(o.Authors) {
		var ret []Author
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetAuthorsOk() ([]Author, bool) {
	if o == nil || isNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *Book) HasAuthors() bool {
	if o != nil && !isNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []Author and assigns it to the Authors field.
func (o *Book) SetAuthors(v []Author) {
	o.Authors = v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *Book) GetPublisher() string {
	if o == nil || isNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetPublisherOk() (*string, bool) {
	if o == nil || isNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *Book) HasPublisher() bool {
	if o != nil && !isNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *Book) SetPublisher(v string) {
	o.Publisher = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *Book) GetPublished() time.Time {
	if o == nil || isNil(o.Published) {
		var ret time.Time
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetPublishedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *Book) HasPublished() bool {
	if o != nil && !isNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given time.Time and assigns it to the Published field.
func (o *Book) SetPublished(v time.Time) {
	o.Published = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Book) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Book) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Book) SetDescription(v string) {
	o.Description = &v
}

// GetGenre returns the Genre field value if set, zero value otherwise.
func (o *Book) GetGenre() string {
	if o == nil || isNil(o.Genre) {
		var ret string
		return ret
	}
	return *o.Genre
}

// GetGenreOk returns a tuple with the Genre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetGenreOk() (*string, bool) {
	if o == nil || isNil(o.Genre) {
		return nil, false
	}
	return o.Genre, true
}

// HasGenre returns a boolean if a field has been set.
func (o *Book) HasGenre() bool {
	if o != nil && !isNil(o.Genre) {
		return true
	}

	return false
}

// SetGenre gets a reference to the given string and assigns it to the Genre field.
func (o *Book) SetGenre(v string) {
	o.Genre = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Book) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Book) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Book) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Book) SetTags(v []string) {
	o.Tags = v
}

func (o Book) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Isbn) {
		toSerialize["isbn"] = o.Isbn
	}
	if !isNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !isNil(o.Publisher) {
		toSerialize["publisher"] = o.Publisher
	}
	if !isNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Genre) {
		toSerialize["genre"] = o.Genre
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableBook struct {
	value *Book
	isSet bool
}

func (v NullableBook) Get() *Book {
	return v.value
}

func (v *NullableBook) Set(val *Book) {
	v.value = val
	v.isSet = true
}

func (v NullableBook) IsSet() bool {
	return v.isSet
}

func (v *NullableBook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBook(val *Book) *NullableBook {
	return &NullableBook{value: val, isSet: true}
}

func (v NullableBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
