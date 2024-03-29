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

// User struct for User
type User struct {
	// user id
	Id *string `json:"id,omitempty"`
	// date of birth for user
	Birthdate *time.Time `json:"birthdate,omitempty"`
	// name of user
	Name *string `json:"name,omitempty"`
	// user's name
	Occupation *string `json:"occupation,omitempty"`
	// the date user become a member
	JoinedDate *time.Time `json:"joined-date,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *User) GetBirthdate() time.Time {
	if o == nil || isNil(o.Birthdate) {
		var ret time.Time
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBirthdateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Birthdate) {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *User) HasBirthdate() bool {
	if o != nil && !isNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given time.Time and assigns it to the Birthdate field.
func (o *User) SetBirthdate(v time.Time) {
	o.Birthdate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetOccupation returns the Occupation field value if set, zero value otherwise.
func (o *User) GetOccupation() string {
	if o == nil || isNil(o.Occupation) {
		var ret string
		return ret
	}
	return *o.Occupation
}

// GetOccupationOk returns a tuple with the Occupation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOccupationOk() (*string, bool) {
	if o == nil || isNil(o.Occupation) {
		return nil, false
	}
	return o.Occupation, true
}

// HasOccupation returns a boolean if a field has been set.
func (o *User) HasOccupation() bool {
	if o != nil && !isNil(o.Occupation) {
		return true
	}

	return false
}

// SetOccupation gets a reference to the given string and assigns it to the Occupation field.
func (o *User) SetOccupation(v string) {
	o.Occupation = &v
}

// GetJoinedDate returns the JoinedDate field value if set, zero value otherwise.
func (o *User) GetJoinedDate() time.Time {
	if o == nil || isNil(o.JoinedDate) {
		var ret time.Time
		return ret
	}
	return *o.JoinedDate
}

// GetJoinedDateOk returns a tuple with the JoinedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetJoinedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.JoinedDate) {
		return nil, false
	}
	return o.JoinedDate, true
}

// HasJoinedDate returns a boolean if a field has been set.
func (o *User) HasJoinedDate() bool {
	if o != nil && !isNil(o.JoinedDate) {
		return true
	}

	return false
}

// SetJoinedDate gets a reference to the given time.Time and assigns it to the JoinedDate field.
func (o *User) SetJoinedDate(v time.Time) {
	o.JoinedDate = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Occupation) {
		toSerialize["occupation"] = o.Occupation
	}
	if !isNil(o.JoinedDate) {
		toSerialize["joined-date"] = o.JoinedDate
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
