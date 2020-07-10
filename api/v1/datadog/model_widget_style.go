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

// WidgetStyle Widget style definition.
type WidgetStyle struct {
	// Color palette to apply to the widget.
	Palette *string `json:"palette,omitempty"`
}

// NewWidgetStyle instantiates a new WidgetStyle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetStyle() *WidgetStyle {
	this := WidgetStyle{}
	return &this
}

// NewWidgetStyleWithDefaults instantiates a new WidgetStyle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetStyleWithDefaults() *WidgetStyle {
	this := WidgetStyle{}
	return &this
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *WidgetStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *WidgetStyle) HasPalette() bool {
	if o != nil && o.Palette != nil {
		return true
	}

	return false
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *WidgetStyle) SetPalette(v string) {
	o.Palette = &v
}

func (o WidgetStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetStyle struct {
	value *WidgetStyle
	isSet bool
}

func (v NullableWidgetStyle) Get() *WidgetStyle {
	return v.value
}

func (v *NullableWidgetStyle) Set(val *WidgetStyle) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetStyle(val *WidgetStyle) *NullableWidgetStyle {
	return &NullableWidgetStyle{value: val, isSet: true}
}

func (v NullableWidgetStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


