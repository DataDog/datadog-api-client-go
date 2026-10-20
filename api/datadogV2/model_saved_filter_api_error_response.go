// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterAPIErrorResponse A saved filter authorization error containing at least one error message.
type SavedFilterAPIErrorResponse struct {
	// The authorization error messages.
	Errors []string `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSavedFilterAPIErrorResponse instantiates a new SavedFilterAPIErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterAPIErrorResponse(errors []string) *SavedFilterAPIErrorResponse {
	this := SavedFilterAPIErrorResponse{}
	this.Errors = errors
	return &this
}

// NewSavedFilterAPIErrorResponseWithDefaults instantiates a new SavedFilterAPIErrorResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterAPIErrorResponseWithDefaults() *SavedFilterAPIErrorResponse {
	this := SavedFilterAPIErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *SavedFilterAPIErrorResponse) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *SavedFilterAPIErrorResponse) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *SavedFilterAPIErrorResponse) SetErrors(v []string) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterAPIErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterAPIErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors *[]string `json:"errors"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	o.Errors = *all.Errors

	return nil
}
