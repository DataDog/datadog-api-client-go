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

// MetricTagConfigurationCreateData Object for a single metric to be configure tags on.
type MetricTagConfigurationCreateData struct {
	Attributes *MetricTagConfigurationCreateAttributes `json:"attributes,omitempty"`
	// The metric name for this resource.
	Id   string                     `json:"id"`
	Type MetricTagConfigurationType `json:"type"`
}

// NewMetricTagConfigurationCreateData instantiates a new MetricTagConfigurationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationCreateData(id string, type_ MetricTagConfigurationType) *MetricTagConfigurationCreateData {
	this := MetricTagConfigurationCreateData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewMetricTagConfigurationCreateDataWithDefaults instantiates a new MetricTagConfigurationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationCreateDataWithDefaults() *MetricTagConfigurationCreateData {
	this := MetricTagConfigurationCreateData{}
	var type_ MetricTagConfigurationType = "manage_tags"
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetricTagConfigurationCreateData) GetAttributes() MetricTagConfigurationCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetricTagConfigurationCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateData) GetAttributesOk() (*MetricTagConfigurationCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetricTagConfigurationCreateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MetricTagConfigurationCreateAttributes and assigns it to the Attributes field.
func (o *MetricTagConfigurationCreateData) SetAttributes(v MetricTagConfigurationCreateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *MetricTagConfigurationCreateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetricTagConfigurationCreateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MetricTagConfigurationCreateData) GetType() MetricTagConfigurationType {
	if o == nil {
		var ret MetricTagConfigurationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateData) GetTypeOk() (*MetricTagConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricTagConfigurationCreateData) SetType(v MetricTagConfigurationType) {
	o.Type = v
}

func (o MetricTagConfigurationCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMetricTagConfigurationCreateData struct {
	value *MetricTagConfigurationCreateData
	isSet bool
}

func (v NullableMetricTagConfigurationCreateData) Get() *MetricTagConfigurationCreateData {
	return v.value
}

func (v *NullableMetricTagConfigurationCreateData) Set(val *MetricTagConfigurationCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationCreateData(val *MetricTagConfigurationCreateData) *NullableMetricTagConfigurationCreateData {
	return &NullableMetricTagConfigurationCreateData{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
