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

// ApplicationKeyListResponse struct for ApplicationKeyListResponse
type ApplicationKeyListResponse struct {
	ApplicationKeys *[]ApplicationKey `json:"application_keys,omitempty"`
}

// GetApplicationKeys returns the ApplicationKeys field value if set, zero value otherwise.
func (o *ApplicationKeyListResponse) GetApplicationKeys() []ApplicationKey {
	if o == nil || o.ApplicationKeys == nil {
		var ret []ApplicationKey
		return ret
	}
	return *o.ApplicationKeys
}

// GetApplicationKeysOk returns a tuple with the ApplicationKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyListResponse) GetApplicationKeysOk() ([]ApplicationKey, bool) {
	if o == nil || o.ApplicationKeys == nil {
		var ret []ApplicationKey
		return ret, false
	}
	return *o.ApplicationKeys, true
}

// HasApplicationKeys returns a boolean if a field has been set.
func (o *ApplicationKeyListResponse) HasApplicationKeys() bool {
	if o != nil && o.ApplicationKeys != nil {
		return true
	}

	return false
}

// SetApplicationKeys gets a reference to the given []ApplicationKey and assigns it to the ApplicationKeys field.
func (o *ApplicationKeyListResponse) SetApplicationKeys(v []ApplicationKey) {
	o.ApplicationKeys = &v
}

type NullableApplicationKeyListResponse struct {
	Value        ApplicationKeyListResponse
	ExplicitNull bool
}

func (v NullableApplicationKeyListResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationKeyListResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
