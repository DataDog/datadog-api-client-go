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

// MetricIngestedIndexedVolumeAttributes Object containing the definition of a metric's ingested and indexed volume.
type MetricIngestedIndexedVolumeAttributes struct {
	// Indexed volume for the given metric.
	IndexedVolume *int64 `json:"indexed_volume,omitempty"`
	// Ingested volume for the given metric.
	IngestedVolume *int64 `json:"ingested_volume,omitempty"`
}

// NewMetricIngestedIndexedVolumeAttributes instantiates a new MetricIngestedIndexedVolumeAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIngestedIndexedVolumeAttributes() *MetricIngestedIndexedVolumeAttributes {
	this := MetricIngestedIndexedVolumeAttributes{}
	return &this
}

// NewMetricIngestedIndexedVolumeAttributesWithDefaults instantiates a new MetricIngestedIndexedVolumeAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricIngestedIndexedVolumeAttributesWithDefaults() *MetricIngestedIndexedVolumeAttributes {
	this := MetricIngestedIndexedVolumeAttributes{}
	return &this
}

// GetIndexedVolume returns the IndexedVolume field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolume() int64 {
	if o == nil || o.IndexedVolume == nil {
		var ret int64
		return ret
	}
	return *o.IndexedVolume
}

// GetIndexedVolumeOk returns a tuple with the IndexedVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIndexedVolumeOk() (*int64, bool) {
	if o == nil || o.IndexedVolume == nil {
		return nil, false
	}
	return o.IndexedVolume, true
}

// HasIndexedVolume returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIndexedVolume() bool {
	if o != nil && o.IndexedVolume != nil {
		return true
	}

	return false
}

// SetIndexedVolume gets a reference to the given int64 and assigns it to the IndexedVolume field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIndexedVolume(v int64) {
	o.IndexedVolume = &v
}

// GetIngestedVolume returns the IngestedVolume field value if set, zero value otherwise.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolume() int64 {
	if o == nil || o.IngestedVolume == nil {
		var ret int64
		return ret
	}
	return *o.IngestedVolume
}

// GetIngestedVolumeOk returns a tuple with the IngestedVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestedIndexedVolumeAttributes) GetIngestedVolumeOk() (*int64, bool) {
	if o == nil || o.IngestedVolume == nil {
		return nil, false
	}
	return o.IngestedVolume, true
}

// HasIngestedVolume returns a boolean if a field has been set.
func (o *MetricIngestedIndexedVolumeAttributes) HasIngestedVolume() bool {
	if o != nil && o.IngestedVolume != nil {
		return true
	}

	return false
}

// SetIngestedVolume gets a reference to the given int64 and assigns it to the IngestedVolume field.
func (o *MetricIngestedIndexedVolumeAttributes) SetIngestedVolume(v int64) {
	o.IngestedVolume = &v
}

func (o MetricIngestedIndexedVolumeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IndexedVolume != nil {
		toSerialize["indexed_volume"] = o.IndexedVolume
	}
	if o.IngestedVolume != nil {
		toSerialize["ingested_volume"] = o.IngestedVolume
	}
	return json.Marshal(toSerialize)
}

type NullableMetricIngestedIndexedVolumeAttributes struct {
	value *MetricIngestedIndexedVolumeAttributes
	isSet bool
}

func (v NullableMetricIngestedIndexedVolumeAttributes) Get() *MetricIngestedIndexedVolumeAttributes {
	return v.value
}

func (v *NullableMetricIngestedIndexedVolumeAttributes) Set(val *MetricIngestedIndexedVolumeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricIngestedIndexedVolumeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricIngestedIndexedVolumeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricIngestedIndexedVolumeAttributes(val *MetricIngestedIndexedVolumeAttributes) *NullableMetricIngestedIndexedVolumeAttributes {
	return &NullableMetricIngestedIndexedVolumeAttributes{value: val, isSet: true}
}

func (v NullableMetricIngestedIndexedVolumeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricIngestedIndexedVolumeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
