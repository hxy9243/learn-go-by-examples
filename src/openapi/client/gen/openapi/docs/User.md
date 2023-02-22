# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | user id | [optional] 
**Birthdate** | Pointer to **time.Time** | date of birth for user | [optional] 
**Name** | Pointer to **string** | name of user | [optional] 
**Occupation** | Pointer to **string** | user&#39;s name | [optional] 
**JoinedDate** | Pointer to **time.Time** | the date user become a member | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBirthdate

`func (o *User) GetBirthdate() time.Time`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *User) GetBirthdateOk() (*time.Time, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *User) SetBirthdate(v time.Time)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *User) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOccupation

`func (o *User) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *User) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *User) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *User) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetJoinedDate

`func (o *User) GetJoinedDate() time.Time`

GetJoinedDate returns the JoinedDate field if non-nil, zero value otherwise.

### GetJoinedDateOk

`func (o *User) GetJoinedDateOk() (*time.Time, bool)`

GetJoinedDateOk returns a tuple with the JoinedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedDate

`func (o *User) SetJoinedDate(v time.Time)`

SetJoinedDate sets JoinedDate field to given value.

### HasJoinedDate

`func (o *User) HasJoinedDate() bool`

HasJoinedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


