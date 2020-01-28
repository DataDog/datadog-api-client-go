/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// SyntheticsLocations struct for SyntheticsLocations
type SyntheticsLocations struct {
	Locations *[]SyntheticsLocation `json:"locations,omitempty"`
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

type NullableSyntheticsLocations struct {
	Value        SyntheticsLocations
	ExplicitNull bool
}

func (v NullableSyntheticsLocations) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsLocations) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
