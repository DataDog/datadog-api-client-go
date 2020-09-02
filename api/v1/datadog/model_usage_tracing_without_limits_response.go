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

// UsageTracingWithoutLimitsResponse Response containing the tracing without limits usage for each hour for a given organization.
type UsageTracingWithoutLimitsResponse struct {
	// Get hourly usage for tracing without limits.
	Usage *[]UsageTracingWithoutLimitsHour `json:"usage,omitempty"`
}

// NewUsageTracingWithoutLimitsResponse instantiates a new UsageTracingWithoutLimitsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTracingWithoutLimitsResponse() *UsageTracingWithoutLimitsResponse {
	this := UsageTracingWithoutLimitsResponse{}
	return &this
}

// NewUsageTracingWithoutLimitsResponseWithDefaults instantiates a new UsageTracingWithoutLimitsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTracingWithoutLimitsResponseWithDefaults() *UsageTracingWithoutLimitsResponse {
	this := UsageTracingWithoutLimitsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageTracingWithoutLimitsResponse) GetUsage() []UsageTracingWithoutLimitsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageTracingWithoutLimitsHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageTracingWithoutLimitsResponse) GetUsageOk() (*[]UsageTracingWithoutLimitsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageTracingWithoutLimitsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageTracingWithoutLimitsHour and assigns it to the Usage field.
func (o *UsageTracingWithoutLimitsResponse) SetUsage(v []UsageTracingWithoutLimitsHour) {
	o.Usage = &v
}

func (o UsageTracingWithoutLimitsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageTracingWithoutLimitsResponse struct {
	value *UsageTracingWithoutLimitsResponse
	isSet bool
}

func (v NullableUsageTracingWithoutLimitsResponse) Get() *UsageTracingWithoutLimitsResponse {
	return v.value
}

func (v *NullableUsageTracingWithoutLimitsResponse) Set(val *UsageTracingWithoutLimitsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageTracingWithoutLimitsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageTracingWithoutLimitsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageTracingWithoutLimitsResponse(val *UsageTracingWithoutLimitsResponse) *NullableUsageTracingWithoutLimitsResponse {
	return &NullableUsageTracingWithoutLimitsResponse{value: val, isSet: true}
}

func (v NullableUsageTracingWithoutLimitsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageTracingWithoutLimitsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
