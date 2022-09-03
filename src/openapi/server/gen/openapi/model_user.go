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
	"time"
)

type User struct {

	// user id
	Id string `json:"id,omitempty"`

	// date of birth for user
	Birthdate time.Time `json:"birthdate,omitempty"`

	// name of user
	Name string `json:"name,omitempty"`

	// user's name
	Occupation string `json:"occupation,omitempty"`

	// the date user become a member
	JoinedDate time.Time `json:"joined-date,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}
