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

// CheckCanDeleteServiceLevelObjectiveResponse A service level objective response containing the requested object.
type CheckCanDeleteServiceLevelObjectiveResponse struct {
	Data CheckCanDeleteServiceLevelObjectiveResponseData `json:"data"`
	// A mapping of SLO id to it's current usages.
	Errors *map[string]string `json:"errors,omitempty"`
}

// GetData returns the Data field value
func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetData() CheckCanDeleteServiceLevelObjectiveResponseData {
	if o == nil {
		var ret CheckCanDeleteServiceLevelObjectiveResponseData
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *CheckCanDeleteServiceLevelObjectiveResponse) SetData(v CheckCanDeleteServiceLevelObjectiveResponseData) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetErrors() map[string]string {
	if o == nil || o.Errors == nil {
		var ret map[string]string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetErrorsOk() (map[string]string, bool) {
	if o == nil || o.Errors == nil {
		var ret map[string]string
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CheckCanDeleteServiceLevelObjectiveResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.
func (o *CheckCanDeleteServiceLevelObjectiveResponse) SetErrors(v map[string]string) {
	o.Errors = &v
}

type NullableCheckCanDeleteServiceLevelObjectiveResponse struct {
	Value        CheckCanDeleteServiceLevelObjectiveResponse
	ExplicitNull bool
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCheckCanDeleteServiceLevelObjectiveResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
