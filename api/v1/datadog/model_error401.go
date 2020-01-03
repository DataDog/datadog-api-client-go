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

// Error401 struct for Error401
type Error401 struct {
	Errors []string `json:"errors"`
}

// GetErrors returns the Errors field value
func (o *Error401) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// SetErrors sets field value
func (o *Error401) SetErrors(v []string) {
	o.Errors = v
}

type NullableError401 struct {
	Value        Error401
	ExplicitNull bool
}

func (v NullableError401) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableError401) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
