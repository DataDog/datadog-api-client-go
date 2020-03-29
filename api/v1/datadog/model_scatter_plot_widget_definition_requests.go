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

// ScatterPlotWidgetDefinitionRequests struct for ScatterPlotWidgetDefinitionRequests
type ScatterPlotWidgetDefinitionRequests struct {
	X ScatterPlotRequest `json:"x"`
	Y ScatterPlotRequest `json:"y"`
}

// NewScatterPlotWidgetDefinitionRequests instantiates a new ScatterPlotWidgetDefinitionRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScatterPlotWidgetDefinitionRequests(x ScatterPlotRequest, y ScatterPlotRequest) *ScatterPlotWidgetDefinitionRequests {
	this := ScatterPlotWidgetDefinitionRequests{}
	this.X = x
	this.Y = y
	return &this
}

// NewScatterPlotWidgetDefinitionRequestsWithDefaults instantiates a new ScatterPlotWidgetDefinitionRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScatterPlotWidgetDefinitionRequestsWithDefaults() *ScatterPlotWidgetDefinitionRequests {
	this := ScatterPlotWidgetDefinitionRequests{}
	return &this
}

// GetX returns the X field value
func (o *ScatterPlotWidgetDefinitionRequests) GetX() ScatterPlotRequest {
	if o == nil {
		var ret ScatterPlotRequest
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ScatterPlotWidgetDefinitionRequests) GetXOk() (*ScatterPlotRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ScatterPlotWidgetDefinitionRequests) SetX(v ScatterPlotRequest) {
	o.X = v
}

// GetY returns the Y field value
func (o *ScatterPlotWidgetDefinitionRequests) GetY() ScatterPlotRequest {
	if o == nil {
		var ret ScatterPlotRequest
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ScatterPlotWidgetDefinitionRequests) GetYOk() (*ScatterPlotRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ScatterPlotWidgetDefinitionRequests) SetY(v ScatterPlotRequest) {
	o.Y = v
}

func (o ScatterPlotWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	return json.Marshal(toSerialize)
}

type NullableScatterPlotWidgetDefinitionRequests struct {
	value *ScatterPlotWidgetDefinitionRequests
	isSet bool
}

func (v NullableScatterPlotWidgetDefinitionRequests) Get() *ScatterPlotWidgetDefinitionRequests {
	return v.value
}

func (v *NullableScatterPlotWidgetDefinitionRequests) Set(val *ScatterPlotWidgetDefinitionRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterPlotWidgetDefinitionRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterPlotWidgetDefinitionRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterPlotWidgetDefinitionRequests(val *ScatterPlotWidgetDefinitionRequests) *NullableScatterPlotWidgetDefinitionRequests {
	return &NullableScatterPlotWidgetDefinitionRequests{value: val, isSet: true}
}

func (v NullableScatterPlotWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterPlotWidgetDefinitionRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
