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

// UsageNetworkFlowsResponse Response containing the number of netflow events indexed for each hour for a given organization.
type UsageNetworkFlowsResponse struct {
	// Get hourly usage for Network Flows.
	Usage *[]UsageNetworkFlowsHour `json:"usage,omitempty"`
}

// NewUsageNetworkFlowsResponse instantiates a new UsageNetworkFlowsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageNetworkFlowsResponse() *UsageNetworkFlowsResponse {
	this := UsageNetworkFlowsResponse{}
	return &this
}

// NewUsageNetworkFlowsResponseWithDefaults instantiates a new UsageNetworkFlowsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageNetworkFlowsResponseWithDefaults() *UsageNetworkFlowsResponse {
	this := UsageNetworkFlowsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageNetworkFlowsResponse) GetUsage() []UsageNetworkFlowsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageNetworkFlowsHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkFlowsResponse) GetUsageOk() (*[]UsageNetworkFlowsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageNetworkFlowsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageNetworkFlowsHour and assigns it to the Usage field.
func (o *UsageNetworkFlowsResponse) SetUsage(v []UsageNetworkFlowsHour) {
	o.Usage = &v
}

func (o UsageNetworkFlowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageNetworkFlowsResponse struct {
	value *UsageNetworkFlowsResponse
	isSet bool
}

func (v NullableUsageNetworkFlowsResponse) Get() *UsageNetworkFlowsResponse {
	return v.value
}

func (v *NullableUsageNetworkFlowsResponse) Set(val *UsageNetworkFlowsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageNetworkFlowsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageNetworkFlowsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageNetworkFlowsResponse(val *UsageNetworkFlowsResponse) *NullableUsageNetworkFlowsResponse {
	return &NullableUsageNetworkFlowsResponse{value: val, isSet: true}
}

func (v NullableUsageNetworkFlowsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageNetworkFlowsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


