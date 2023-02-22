# Borrow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the borrow | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**CopyId** | Pointer to **string** |  | [optional] 
**BorrowTime** | Pointer to **time.Time** |  | [optional] 
**DueTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBorrow

`func NewBorrow() *Borrow`

NewBorrow instantiates a new Borrow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBorrowWithDefaults

`func NewBorrowWithDefaults() *Borrow`

NewBorrowWithDefaults instantiates a new Borrow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Borrow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Borrow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Borrow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Borrow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Borrow) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Borrow) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Borrow) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Borrow) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCopyId

`func (o *Borrow) GetCopyId() string`

GetCopyId returns the CopyId field if non-nil, zero value otherwise.

### GetCopyIdOk

`func (o *Borrow) GetCopyIdOk() (*string, bool)`

GetCopyIdOk returns a tuple with the CopyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyId

`func (o *Borrow) SetCopyId(v string)`

SetCopyId sets CopyId field to given value.

### HasCopyId

`func (o *Borrow) HasCopyId() bool`

HasCopyId returns a boolean if a field has been set.

### GetBorrowTime

`func (o *Borrow) GetBorrowTime() time.Time`

GetBorrowTime returns the BorrowTime field if non-nil, zero value otherwise.

### GetBorrowTimeOk

`func (o *Borrow) GetBorrowTimeOk() (*time.Time, bool)`

GetBorrowTimeOk returns a tuple with the BorrowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowTime

`func (o *Borrow) SetBorrowTime(v time.Time)`

SetBorrowTime sets BorrowTime field to given value.

### HasBorrowTime

`func (o *Borrow) HasBorrowTime() bool`

HasBorrowTime returns a boolean if a field has been set.

### GetDueTime

`func (o *Borrow) GetDueTime() time.Time`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *Borrow) GetDueTimeOk() (*time.Time, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *Borrow) SetDueTime(v time.Time)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *Borrow) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetStatus

`func (o *Borrow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Borrow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Borrow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Borrow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


