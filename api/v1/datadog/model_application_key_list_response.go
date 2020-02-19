/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// ApplicationKeyListResponse struct for ApplicationKeyListResponse
type ApplicationKeyListResponse struct {
	ApplicationKeys *[]ApplicationKey `json:"application_keys,omitempty"`
}

// NewApplicationKeyListResponse instantiates a new ApplicationKeyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyListResponse() *ApplicationKeyListResponse {
	this := ApplicationKeyListResponse{}
	return &this
}

// NewApplicationKeyListResponseWithDefaults instantiates a new ApplicationKeyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyListResponseWithDefaults() *ApplicationKeyListResponse {
	this := ApplicationKeyListResponse{}
	return &this
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

func (o ApplicationKeyListResponse) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.ApplicationKeys != nil {
		toSerialize["application_keys"] = o.ApplicationKeys
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationKeyListResponse struct {
	value *ApplicationKeyListResponse
	isSet bool
}

func (v NullableApplicationKeyListResponse) Get() *ApplicationKeyListResponse {
	return v.value
}

func (v NullableApplicationKeyListResponse) Set(val *ApplicationKeyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyListResponse) IsSet() bool {
	return v.isSet
}

func (v NullableApplicationKeyListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyListResponse(val *ApplicationKeyListResponse) *NullableApplicationKeyListResponse {
	return &NullableApplicationKeyListResponse{value: val, isSet: true}
}

func (v NullableApplicationKeyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
