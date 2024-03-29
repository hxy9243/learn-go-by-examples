/*
OpenAPI Library Demo App

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Copy struct for Copy
type Copy struct {
	// id of the copy
	Id *string `json:"id,omitempty"`
	// id of the book
	BookId *string `json:"book-id,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCopy instantiates a new Copy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopy() *Copy {
	this := Copy{}
	return &this
}

// NewCopyWithDefaults instantiates a new Copy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyWithDefaults() *Copy {
	this := Copy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Copy) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Copy) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Copy) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Copy) SetId(v string) {
	o.Id = &v
}

// GetBookId returns the BookId field value if set, zero value otherwise.
func (o *Copy) GetBookId() string {
	if o == nil || isNil(o.BookId) {
		var ret string
		return ret
	}
	return *o.BookId
}

// GetBookIdOk returns a tuple with the BookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Copy) GetBookIdOk() (*string, bool) {
	if o == nil || isNil(o.BookId) {
		return nil, false
	}
	return o.BookId, true
}

// HasBookId returns a boolean if a field has been set.
func (o *Copy) HasBookId() bool {
	if o != nil && !isNil(o.BookId) {
		return true
	}

	return false
}

// SetBookId gets a reference to the given string and assigns it to the BookId field.
func (o *Copy) SetBookId(v string) {
	o.BookId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Copy) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Copy) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Copy) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Copy) SetStatus(v string) {
	o.Status = &v
}

func (o Copy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BookId) {
		toSerialize["book-id"] = o.BookId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCopy struct {
	value *Copy
	isSet bool
}

func (v NullableCopy) Get() *Copy {
	return v.value
}

func (v *NullableCopy) Set(val *Copy) {
	v.value = val
	v.isSet = true
}

func (v NullableCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopy(val *Copy) *NullableCopy {
	return &NullableCopy{value: val, isSet: true}
}

func (v NullableCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
