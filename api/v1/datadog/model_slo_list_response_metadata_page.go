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

// SLOListResponseMetadataPage The object containing information about the pages of the list of SLOs.
type SLOListResponseMetadataPage struct {
	// The total number of resources that could be retrieved ignoring the parameters and filters in the request.
	TotalCount *int64 `json:"total_count,omitempty"`
	// The total number of resources that match the parameters and filters in the request. This attribute can be used by a client to determine the total number of pages.
	TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
}

// NewSLOListResponseMetadataPage instantiates a new SLOListResponseMetadataPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOListResponseMetadataPage() *SLOListResponseMetadataPage {
	this := SLOListResponseMetadataPage{}
	return &this
}

// NewSLOListResponseMetadataPageWithDefaults instantiates a new SLOListResponseMetadataPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOListResponseMetadataPageWithDefaults() *SLOListResponseMetadataPage {
	this := SLOListResponseMetadataPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SLOListResponseMetadataPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOListResponseMetadataPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SLOListResponseMetadataPage) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *SLOListResponseMetadataPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *SLOListResponseMetadataPage) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOListResponseMetadataPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *SLOListResponseMetadataPage) HasTotalFilteredCount() bool {
	if o != nil && o.TotalFilteredCount != nil {
		return true
	}

	return false
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *SLOListResponseMetadataPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

func (o SLOListResponseMetadataPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.TotalFilteredCount != nil {
		toSerialize["total_filtered_count"] = o.TotalFilteredCount
	}
	return json.Marshal(toSerialize)
}

type NullableSLOListResponseMetadataPage struct {
	value *SLOListResponseMetadataPage
	isSet bool
}

func (v NullableSLOListResponseMetadataPage) Get() *SLOListResponseMetadataPage {
	return v.value
}

func (v *NullableSLOListResponseMetadataPage) Set(val *SLOListResponseMetadataPage) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOListResponseMetadataPage) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOListResponseMetadataPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOListResponseMetadataPage(val *SLOListResponseMetadataPage) *NullableSLOListResponseMetadataPage {
	return &NullableSLOListResponseMetadataPage{value: val, isSet: true}
}

func (v NullableSLOListResponseMetadataPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOListResponseMetadataPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
