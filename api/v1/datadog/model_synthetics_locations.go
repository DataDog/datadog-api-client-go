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

// SyntheticsLocations struct for SyntheticsLocations
type SyntheticsLocations struct {
	Locations *[]SyntheticsLocation `json:"locations,omitempty"`
}

// NewSyntheticsLocations instantiates a new SyntheticsLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsLocations() *SyntheticsLocations {
	this := SyntheticsLocations{}
	return &this
}

// NewSyntheticsLocationsWithDefaults instantiates a new SyntheticsLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsLocationsWithDefaults() *SyntheticsLocations {
	this := SyntheticsLocations{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsLocations) GetLocations() []SyntheticsLocation {
	if o == nil || o.Locations == nil {
		var ret []SyntheticsLocation
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsLocations) GetLocationsOk() ([]SyntheticsLocation, bool) {
	if o == nil || o.Locations == nil {
		var ret []SyntheticsLocation
		return ret, false
	}
	return *o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsLocations) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []SyntheticsLocation and assigns it to the Locations field.
func (o *SyntheticsLocations) SetLocations(v []SyntheticsLocation) {
	o.Locations = &v
}

func (o SyntheticsLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsLocations struct {
	value *SyntheticsLocations
	isSet bool
}

func (v NullableSyntheticsLocations) Get() *SyntheticsLocations {
	return v.value
}

func (v NullableSyntheticsLocations) Set(val *SyntheticsLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsLocations) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsLocations(val *SyntheticsLocations) *NullableSyntheticsLocations {
	return &NullableSyntheticsLocations{value: val, isSet: true}
}

func (v NullableSyntheticsLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
