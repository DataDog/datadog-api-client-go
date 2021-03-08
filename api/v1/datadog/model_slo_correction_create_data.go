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

// SLOCorrectionCreateData The data object associated with the SLO correction to be created
type SLOCorrectionCreateData struct {
	Attributes *SLOCorrectionCreateRequestAttributes `json:"attributes,omitempty"`
	Type       *SLOCorrectionType                    `json:"type,omitempty"`
}

// NewSLOCorrectionCreateData instantiates a new SLOCorrectionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOCorrectionCreateData() *SLOCorrectionCreateData {
	this := SLOCorrectionCreateData{}
	var type_ SLOCorrectionType = SLOCORRECTIONTYPE_CORRECTION
	this.Type = &type_
	return &this
}

// NewSLOCorrectionCreateDataWithDefaults instantiates a new SLOCorrectionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOCorrectionCreateDataWithDefaults() *SLOCorrectionCreateData {
	this := SLOCorrectionCreateData{}
	var type_ SLOCorrectionType = SLOCORRECTIONTYPE_CORRECTION
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SLOCorrectionCreateData) GetAttributes() SLOCorrectionCreateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret SLOCorrectionCreateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateData) GetAttributesOk() (*SLOCorrectionCreateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SLOCorrectionCreateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SLOCorrectionCreateRequestAttributes and assigns it to the Attributes field.
func (o *SLOCorrectionCreateData) SetAttributes(v SLOCorrectionCreateRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SLOCorrectionCreateData) GetType() SLOCorrectionType {
	if o == nil || o.Type == nil {
		var ret SLOCorrectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateData) GetTypeOk() (*SLOCorrectionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SLOCorrectionCreateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SLOCorrectionType and assigns it to the Type field.
func (o *SLOCorrectionCreateData) SetType(v SLOCorrectionType) {
	o.Type = &v
}

func (o SLOCorrectionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSLOCorrectionCreateData struct {
	value *SLOCorrectionCreateData
	isSet bool
}

func (v NullableSLOCorrectionCreateData) Get() *SLOCorrectionCreateData {
	return v.value
}

func (v *NullableSLOCorrectionCreateData) Set(val *SLOCorrectionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOCorrectionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOCorrectionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOCorrectionCreateData(val *SLOCorrectionCreateData) *NullableSLOCorrectionCreateData {
	return &NullableSLOCorrectionCreateData{value: val, isSet: true}
}

func (v NullableSLOCorrectionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOCorrectionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
