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

// UsageSyntheticsResponse Response containing the number of Synthetics API tests run for each hour for a given organization.
type UsageSyntheticsResponse struct {
	// Array with the number of hourly Synthetics test run for a given organization.
	Usage *[]UsageSyntheticsHour `json:"usage,omitempty"`
}

// NewUsageSyntheticsResponse instantiates a new UsageSyntheticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSyntheticsResponse() *UsageSyntheticsResponse {
	this := UsageSyntheticsResponse{}
	return &this
}

// NewUsageSyntheticsResponseWithDefaults instantiates a new UsageSyntheticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSyntheticsResponseWithDefaults() *UsageSyntheticsResponse {
	this := UsageSyntheticsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSyntheticsResponse) GetUsage() []UsageSyntheticsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageSyntheticsHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsResponse) GetUsageOk() (*[]UsageSyntheticsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSyntheticsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageSyntheticsHour and assigns it to the Usage field.
func (o *UsageSyntheticsResponse) SetUsage(v []UsageSyntheticsHour) {
	o.Usage = &v
}

func (o UsageSyntheticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageSyntheticsResponse struct {
	value *UsageSyntheticsResponse
	isSet bool
}

func (v NullableUsageSyntheticsResponse) Get() *UsageSyntheticsResponse {
	return v.value
}

func (v *NullableUsageSyntheticsResponse) Set(val *UsageSyntheticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSyntheticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSyntheticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSyntheticsResponse(val *UsageSyntheticsResponse) *NullableUsageSyntheticsResponse {
	return &NullableUsageSyntheticsResponse{value: val, isSet: true}
}

func (v NullableUsageSyntheticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSyntheticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


