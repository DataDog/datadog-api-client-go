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

// Error408 struct for Error408
type Error408 struct {
	Errors []string `json:"errors"`
}

// NewError408 instantiates a new Error408 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError408(errors []string) *Error408 {
	this := Error408{}
	this.Errors = errors
	return &this
}

// NewError408WithDefaults instantiates a new Error408 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError408WithDefaults() *Error408 {
	this := Error408{}
	return &this
}

// GetErrors returns the Errors field value
func (o *Error408) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// SetErrors sets field value
func (o *Error408) SetErrors(v []string) {
	o.Errors = v
}

type NullableError408 struct {
	Value        Error408
	ExplicitNull bool
}

func (v NullableError408) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableError408) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
