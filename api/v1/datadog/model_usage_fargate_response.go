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

// UsageFargateResponse struct for UsageFargateResponse
type UsageFargateResponse struct {
	Usage *[]UsageFargateHour `json:"usage,omitempty"`
}

// NewUsageFargateResponse instantiates a new UsageFargateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageFargateResponse() *UsageFargateResponse {
	this := UsageFargateResponse{}
	return &this
}

// NewUsageFargateResponseWithDefaults instantiates a new UsageFargateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageFargateResponseWithDefaults() *UsageFargateResponse {
	this := UsageFargateResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageFargateResponse) GetUsage() []UsageFargateHour {
	if o == nil || o.Usage == nil {
		var ret []UsageFargateHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsageFargateResponse) GetUsageOk() ([]UsageFargateHour, bool) {
	if o == nil || o.Usage == nil {
		var ret []UsageFargateHour
		return ret, false
	}
	return *o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageFargateResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageFargateHour and assigns it to the Usage field.
func (o *UsageFargateResponse) SetUsage(v []UsageFargateHour) {
	o.Usage = &v
}

func (o UsageFargateResponse) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageFargateResponse struct {
	value *UsageFargateResponse
	isSet bool
}

func (v NullableUsageFargateResponse) Get() *UsageFargateResponse {
	return v.value
}

func (v NullableUsageFargateResponse) Set(val *UsageFargateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageFargateResponse) IsSet() bool {
	return v.isSet
}

func (v NullableUsageFargateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageFargateResponse(val *UsageFargateResponse) *NullableUsageFargateResponse {
	return &NullableUsageFargateResponse{value: val, isSet: true}
}

func (v NullableUsageFargateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageFargateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
