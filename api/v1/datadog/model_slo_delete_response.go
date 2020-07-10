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

// SLODeleteResponse A response list of all service level objective deleted.
type SLODeleteResponse struct {
	// An array containing the ID of the deleted service level objective object.
	Data *[]string `json:"data,omitempty"`
	// An dictionary containing the ID of the SLO as key and a deletion error as value.
	Errors *map[string]string `json:"errors,omitempty"`
}

// NewSLODeleteResponse instantiates a new SLODeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLODeleteResponse() *SLODeleteResponse {
	this := SLODeleteResponse{}
	return &this
}

// NewSLODeleteResponseWithDefaults instantiates a new SLODeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLODeleteResponseWithDefaults() *SLODeleteResponse {
	this := SLODeleteResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLODeleteResponse) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLODeleteResponse) GetDataOk() (*[]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLODeleteResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *SLODeleteResponse) SetData(v []string) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SLODeleteResponse) GetErrors() map[string]string {
	if o == nil || o.Errors == nil {
		var ret map[string]string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLODeleteResponse) GetErrorsOk() (*map[string]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLODeleteResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.
func (o *SLODeleteResponse) SetErrors(v map[string]string) {
	o.Errors = &v
}

func (o SLODeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableSLODeleteResponse struct {
	value *SLODeleteResponse
	isSet bool
}

func (v NullableSLODeleteResponse) Get() *SLODeleteResponse {
	return v.value
}

func (v *NullableSLODeleteResponse) Set(val *SLODeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSLODeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSLODeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLODeleteResponse(val *SLODeleteResponse) *NullableSLODeleteResponse {
	return &NullableSLODeleteResponse{value: val, isSet: true}
}

func (v NullableSLODeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLODeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


