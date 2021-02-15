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

// SLOCorrectionCreateRequest An object that defines a correction to be applied to an SLO
type SLOCorrectionCreateRequest struct {
	Data *SLOCorrectionCreateData `json:"data,omitempty"`
}

// NewSLOCorrectionCreateRequest instantiates a new SLOCorrectionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOCorrectionCreateRequest() *SLOCorrectionCreateRequest {
	this := SLOCorrectionCreateRequest{}
	return &this
}

// NewSLOCorrectionCreateRequestWithDefaults instantiates a new SLOCorrectionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOCorrectionCreateRequestWithDefaults() *SLOCorrectionCreateRequest {
	this := SLOCorrectionCreateRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLOCorrectionCreateRequest) GetData() SLOCorrectionCreateData {
	if o == nil || o.Data == nil {
		var ret SLOCorrectionCreateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionCreateRequest) GetDataOk() (*SLOCorrectionCreateData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLOCorrectionCreateRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SLOCorrectionCreateData and assigns it to the Data field.
func (o *SLOCorrectionCreateRequest) SetData(v SLOCorrectionCreateData) {
	o.Data = &v
}

func (o SLOCorrectionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSLOCorrectionCreateRequest struct {
	value *SLOCorrectionCreateRequest
	isSet bool
}

func (v NullableSLOCorrectionCreateRequest) Get() *SLOCorrectionCreateRequest {
	return v.value
}

func (v *NullableSLOCorrectionCreateRequest) Set(val *SLOCorrectionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOCorrectionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOCorrectionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOCorrectionCreateRequest(val *SLOCorrectionCreateRequest) *NullableSLOCorrectionCreateRequest {
	return &NullableSLOCorrectionCreateRequest{value: val, isSet: true}
}

func (v NullableSLOCorrectionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOCorrectionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
