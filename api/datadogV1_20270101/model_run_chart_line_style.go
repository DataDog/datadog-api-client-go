// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartLineStyle Test-only generated schema for the dashboard versioning proof of concept.
type RunChartLineStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorOrder *TimeseriesColorOrder `json:"color_order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LineWidth *WidgetLineWidth `json:"line_width,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *string `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunChartLineStyle instantiates a new RunChartLineStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunChartLineStyle() *RunChartLineStyle {
	this := RunChartLineStyle{}
	return &this
}

// NewRunChartLineStyleWithDefaults instantiates a new RunChartLineStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunChartLineStyleWithDefaults() *RunChartLineStyle {
	this := RunChartLineStyle{}
	return &this
}

// GetColorOrder returns the ColorOrder field value if set, zero value otherwise.
func (o *RunChartLineStyle) GetColorOrder() TimeseriesColorOrder {
	if o == nil || o.ColorOrder == nil {
		var ret TimeseriesColorOrder
		return ret
	}
	return *o.ColorOrder
}

// GetColorOrderOk returns a tuple with the ColorOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartLineStyle) GetColorOrderOk() (*TimeseriesColorOrder, bool) {
	if o == nil || o.ColorOrder == nil {
		return nil, false
	}
	return o.ColorOrder, true
}

// HasColorOrder returns a boolean if a field has been set.
func (o *RunChartLineStyle) HasColorOrder() bool {
	return o != nil && o.ColorOrder != nil
}

// SetColorOrder gets a reference to the given TimeseriesColorOrder and assigns it to the ColorOrder field.
func (o *RunChartLineStyle) SetColorOrder(v TimeseriesColorOrder) {
	o.ColorOrder = &v
}

// GetLineWidth returns the LineWidth field value if set, zero value otherwise.
func (o *RunChartLineStyle) GetLineWidth() WidgetLineWidth {
	if o == nil || o.LineWidth == nil {
		var ret WidgetLineWidth
		return ret
	}
	return *o.LineWidth
}

// GetLineWidthOk returns a tuple with the LineWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartLineStyle) GetLineWidthOk() (*WidgetLineWidth, bool) {
	if o == nil || o.LineWidth == nil {
		return nil, false
	}
	return o.LineWidth, true
}

// HasLineWidth returns a boolean if a field has been set.
func (o *RunChartLineStyle) HasLineWidth() bool {
	return o != nil && o.LineWidth != nil
}

// SetLineWidth gets a reference to the given WidgetLineWidth and assigns it to the LineWidth field.
func (o *RunChartLineStyle) SetLineWidth(v WidgetLineWidth) {
	o.LineWidth = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *RunChartLineStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartLineStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *RunChartLineStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *RunChartLineStyle) SetPalette(v string) {
	o.Palette = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunChartLineStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorOrder != nil {
		toSerialize["color_order"] = o.ColorOrder
	}
	if o.LineWidth != nil {
		toSerialize["line_width"] = o.LineWidth
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunChartLineStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorOrder *TimeseriesColorOrder `json:"color_order,omitempty"`
		LineWidth  *WidgetLineWidth      `json:"line_width,omitempty"`
		Palette    *string               `json:"palette,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.ColorOrder != nil && !all.ColorOrder.IsValid() {
		hasInvalidField = true
	} else {
		o.ColorOrder = all.ColorOrder
	}
	if all.LineWidth != nil && !all.LineWidth.IsValid() {
		hasInvalidField = true
	} else {
		o.LineWidth = all.LineWidth
	}
	o.Palette = all.Palette

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
