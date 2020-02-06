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

// AwsLogsAsyncResponseErrors struct for AwsLogsAsyncResponseErrors
type AwsLogsAsyncResponseErrors struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewAwsLogsAsyncResponseErrors instantiates a new AwsLogsAsyncResponseErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsLogsAsyncResponseErrors() *AwsLogsAsyncResponseErrors {
	this := AwsLogsAsyncResponseErrors{}
	return &this
}

// NewAwsLogsAsyncResponseErrorsWithDefaults instantiates a new AwsLogsAsyncResponseErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsLogsAsyncResponseErrorsWithDefaults() *AwsLogsAsyncResponseErrors {
	this := AwsLogsAsyncResponseErrors{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AwsLogsAsyncResponseErrors) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsAsyncResponseErrors) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AwsLogsAsyncResponseErrors) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AwsLogsAsyncResponseErrors) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AwsLogsAsyncResponseErrors) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AwsLogsAsyncResponseErrors) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AwsLogsAsyncResponseErrors) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AwsLogsAsyncResponseErrors) SetMessage(v string) {
	o.Message = &v
}

type NullableAwsLogsAsyncResponseErrors struct {
	Value        AwsLogsAsyncResponseErrors
	ExplicitNull bool
}

func (v NullableAwsLogsAsyncResponseErrors) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAwsLogsAsyncResponseErrors) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
