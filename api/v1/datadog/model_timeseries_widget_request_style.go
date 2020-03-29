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

// TimeseriesWidgetRequestStyle struct for TimeseriesWidgetRequestStyle
type TimeseriesWidgetRequestStyle struct {
	LineType  *WidgetLineType  `json:"line_type,omitempty"`
	LineWidth *WidgetLineWidth `json:"line_width,omitempty"`
	// Color palette to apply to the widget.
	Palette *string `json:"palette,omitempty"`
}

// NewTimeseriesWidgetRequestStyle instantiates a new TimeseriesWidgetRequestStyle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesWidgetRequestStyle() *TimeseriesWidgetRequestStyle {
	this := TimeseriesWidgetRequestStyle{}
	return &this
}

// NewTimeseriesWidgetRequestStyleWithDefaults instantiates a new TimeseriesWidgetRequestStyle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesWidgetRequestStyleWithDefaults() *TimeseriesWidgetRequestStyle {
	this := TimeseriesWidgetRequestStyle{}
	return &this
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequestStyle) GetLineType() WidgetLineType {
	if o == nil || o.LineType == nil {
		var ret WidgetLineType
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequestStyle) GetLineTypeOk() (*WidgetLineType, bool) {
	if o == nil || o.LineType == nil {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequestStyle) HasLineType() bool {
	if o != nil && o.LineType != nil {
		return true
	}

	return false
}

// SetLineType gets a reference to the given WidgetLineType and assigns it to the LineType field.
func (o *TimeseriesWidgetRequestStyle) SetLineType(v WidgetLineType) {
	o.LineType = &v
}

// GetLineWidth returns the LineWidth field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequestStyle) GetLineWidth() WidgetLineWidth {
	if o == nil || o.LineWidth == nil {
		var ret WidgetLineWidth
		return ret
	}
	return *o.LineWidth
}

// GetLineWidthOk returns a tuple with the LineWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequestStyle) GetLineWidthOk() (*WidgetLineWidth, bool) {
	if o == nil || o.LineWidth == nil {
		return nil, false
	}
	return o.LineWidth, true
}

// HasLineWidth returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequestStyle) HasLineWidth() bool {
	if o != nil && o.LineWidth != nil {
		return true
	}

	return false
}

// SetLineWidth gets a reference to the given WidgetLineWidth and assigns it to the LineWidth field.
func (o *TimeseriesWidgetRequestStyle) SetLineWidth(v WidgetLineWidth) {
	o.LineWidth = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequestStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequestStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequestStyle) HasPalette() bool {
	if o != nil && o.Palette != nil {
		return true
	}

	return false
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *TimeseriesWidgetRequestStyle) SetPalette(v string) {
	o.Palette = &v
}

func (o TimeseriesWidgetRequestStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LineType != nil {
		toSerialize["line_type"] = o.LineType
	}
	if o.LineWidth != nil {
		toSerialize["line_width"] = o.LineWidth
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	return json.Marshal(toSerialize)
}

type NullableTimeseriesWidgetRequestStyle struct {
	value *TimeseriesWidgetRequestStyle
	isSet bool
}

func (v NullableTimeseriesWidgetRequestStyle) Get() *TimeseriesWidgetRequestStyle {
	return v.value
}

func (v *NullableTimeseriesWidgetRequestStyle) Set(val *TimeseriesWidgetRequestStyle) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesWidgetRequestStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesWidgetRequestStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesWidgetRequestStyle(val *TimeseriesWidgetRequestStyle) *NullableTimeseriesWidgetRequestStyle {
	return &NullableTimeseriesWidgetRequestStyle{value: val, isSet: true}
}

func (v NullableTimeseriesWidgetRequestStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesWidgetRequestStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
