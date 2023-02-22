# Book

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the book | [optional] 
**Title** | Pointer to **string** | the id of the book | [optional] 
**Isbn** | Pointer to **string** | ISBN book number | [optional] 
**Authors** | Pointer to [**[]Author**](Author.md) | the id of the book | [optional] 
**Publisher** | Pointer to **string** | name of the publisher | [optional] 
**Published** | Pointer to **time.Time** | time the book is published | [optional] 
**Description** | Pointer to **string** | the description of the book | [optional] 
**Genre** | Pointer to **string** | the genre of this book | [optional] 
**Tags** | Pointer to **[]string** | the tags of the book | [optional] 

## Methods

### NewBook

`func NewBook() *Book`

NewBook instantiates a new Book object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookWithDefaults

`func NewBookWithDefaults() *Book`

NewBookWithDefaults instantiates a new Book object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Book) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Book) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Book) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Book) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Book) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Book) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Book) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Book) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsbn

`func (o *Book) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *Book) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *Book) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *Book) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetAuthors

`func (o *Book) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Book) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Book) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *Book) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetPublisher

`func (o *Book) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *Book) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *Book) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *Book) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetPublished

`func (o *Book) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Book) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Book) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Book) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetDescription

`func (o *Book) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Book) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Book) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Book) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenre

`func (o *Book) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *Book) GetGenreOk() (*string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *Book) SetGenre(v string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *Book) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetTags

`func (o *Book) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Book) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Book) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Book) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


