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

// LogsIndexListResponse struct for LogsIndexListResponse
type LogsIndexListResponse struct {
	Indexes *[]LogsIndex `json:"indexes,omitempty"`
}

// NewLogsIndexListResponse instantiates a new LogsIndexListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsIndexListResponse() *LogsIndexListResponse {
	this := LogsIndexListResponse{}
	return &this
}

// NewLogsIndexListResponseWithDefaults instantiates a new LogsIndexListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsIndexListResponseWithDefaults() *LogsIndexListResponse {
	this := LogsIndexListResponse{}
	return &this
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogsIndexListResponse) GetIndexes() []LogsIndex {
	if o == nil || o.Indexes == nil {
		var ret []LogsIndex
		return ret
	}
	return *o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsIndexListResponse) GetIndexesOk() ([]LogsIndex, bool) {
	if o == nil || o.Indexes == nil {
		var ret []LogsIndex
		return ret, false
	}
	return *o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogsIndexListResponse) HasIndexes() bool {
	if o != nil && o.Indexes != nil {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []LogsIndex and assigns it to the Indexes field.
func (o *LogsIndexListResponse) SetIndexes(v []LogsIndex) {
	o.Indexes = &v
}

func (o LogsIndexListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	return json.Marshal(toSerialize)
}

type NullableLogsIndexListResponse struct {
	value *LogsIndexListResponse
	isSet bool
}

func (v NullableLogsIndexListResponse) Get() *LogsIndexListResponse {
	return v.value
}

func (v NullableLogsIndexListResponse) Set(val *LogsIndexListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsIndexListResponse) IsSet() bool {
	return v.isSet
}

func (v NullableLogsIndexListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsIndexListResponse(val *LogsIndexListResponse) *NullableLogsIndexListResponse {
	return &NullableLogsIndexListResponse{value: val, isSet: true}
}

func (v NullableLogsIndexListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsIndexListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
