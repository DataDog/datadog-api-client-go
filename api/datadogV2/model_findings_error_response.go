// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingsErrorResponse API error response.
type FindingsErrorResponse struct {
	// A list of errors.
	Errors []FindingsErrorItem `json:"errors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewFindingsErrorResponse instantiates a new FindingsErrorResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingsErrorResponse(errors []FindingsErrorItem) *FindingsErrorResponse {
	this := FindingsErrorResponse{}
	this.Errors = errors
	return &this
}

// NewFindingsErrorResponseWithDefaults instantiates a new FindingsErrorResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingsErrorResponseWithDefaults() *FindingsErrorResponse {
	this := FindingsErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value.
func (o *FindingsErrorResponse) GetErrors() []FindingsErrorItem {
	if o == nil {
		var ret []FindingsErrorItem
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *FindingsErrorResponse) GetErrorsOk() (*[]FindingsErrorItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value.
func (o *FindingsErrorResponse) SetErrors(v []FindingsErrorItem) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingsErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingsErrorResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Errors *[]FindingsErrorItem `json:"errors"`
	}{}
	all := struct {
		Errors []FindingsErrorItem `json:"errors"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Errors == nil {
		return fmt.Errorf("required field errors missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors"})
	} else {
		return err
	}
	o.Errors = all.Errors
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
