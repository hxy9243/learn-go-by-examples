/*
 * OpenAPI Library Demo App
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Copy struct {

	// id of the copy
	Id string `json:"id,omitempty"`

	// id of the book
	BookId string `json:"book-id,omitempty"`

	Status string `json:"status,omitempty"`
}

// AssertCopyRequired checks if the required fields are not zero-ed
func AssertCopyRequired(obj Copy) error {
	return nil
}

// AssertRecurseCopyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Copy (e.g. [][]Copy), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCopyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCopy, ok := obj.(Copy)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCopyRequired(aCopy)
	})
}