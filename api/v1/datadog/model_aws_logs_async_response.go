/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// AwsLogsAsyncResponse struct for AwsLogsAsyncResponse
type AwsLogsAsyncResponse struct {
	Errors *[]AwsLogsAsyncResponseErrors `json:"errors,omitempty"`
	Status *string                       `json:"status,omitempty"`
}

// NewAwsLogsAsyncResponse instantiates a new AwsLogsAsyncResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsLogsAsyncResponse() *AwsLogsAsyncResponse {
	this := AwsLogsAsyncResponse{}
	return &this
}

// NewAwsLogsAsyncResponseWithDefaults instantiates a new AwsLogsAsyncResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsLogsAsyncResponseWithDefaults() *AwsLogsAsyncResponse {
	this := AwsLogsAsyncResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AwsLogsAsyncResponse) GetErrors() []AwsLogsAsyncResponseErrors {
	if o == nil || o.Errors == nil {
		var ret []AwsLogsAsyncResponseErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsAsyncResponse) GetErrorsOk() ([]AwsLogsAsyncResponseErrors, bool) {
	if o == nil || o.Errors == nil {
		var ret []AwsLogsAsyncResponseErrors
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AwsLogsAsyncResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AwsLogsAsyncResponseErrors and assigns it to the Errors field.
func (o *AwsLogsAsyncResponse) SetErrors(v []AwsLogsAsyncResponseErrors) {
	o.Errors = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsLogsAsyncResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsAsyncResponse) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsLogsAsyncResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AwsLogsAsyncResponse) SetStatus(v string) {
	o.Status = &v
}

type NullableAwsLogsAsyncResponse struct {
	Value        AwsLogsAsyncResponse
	ExplicitNull bool
}

func (v NullableAwsLogsAsyncResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsLogsAsyncResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
