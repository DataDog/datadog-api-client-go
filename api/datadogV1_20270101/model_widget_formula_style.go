// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormulaStyle Test-only generated schema for the dashboard versioning proof of concept.
type WidgetFormulaStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette string `json:"palette"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PaletteIndex *float64 `json:"palette_index,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetFormulaStyle instantiates a new WidgetFormulaStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFormulaStyle(palette string) *WidgetFormulaStyle {
	this := WidgetFormulaStyle{}
	this.Palette = palette
	return &this
}

// NewWidgetFormulaStyleWithDefaults instantiates a new WidgetFormulaStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFormulaStyleWithDefaults() *WidgetFormulaStyle {
	this := WidgetFormulaStyle{}
	return &this
}

// GetPalette returns the Palette field value.
func (o *WidgetFormulaStyle) GetPalette() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *WidgetFormulaStyle) GetPaletteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *WidgetFormulaStyle) SetPalette(v string) {
	o.Palette = v
}

// GetPaletteIndex returns the PaletteIndex field value if set, zero value otherwise.
func (o *WidgetFormulaStyle) GetPaletteIndex() float64 {
	if o == nil || o.PaletteIndex == nil {
		var ret float64
		return ret
	}
	return *o.PaletteIndex
}

// GetPaletteIndexOk returns a tuple with the PaletteIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormulaStyle) GetPaletteIndexOk() (*float64, bool) {
	if o == nil || o.PaletteIndex == nil {
		return nil, false
	}
	return o.PaletteIndex, true
}

// HasPaletteIndex returns a boolean if a field has been set.
func (o *WidgetFormulaStyle) HasPaletteIndex() bool {
	return o != nil && o.PaletteIndex != nil
}

// SetPaletteIndex gets a reference to the given float64 and assigns it to the PaletteIndex field.
func (o *WidgetFormulaStyle) SetPaletteIndex(v float64) {
	o.PaletteIndex = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFormulaStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["palette"] = o.Palette
	if o.PaletteIndex != nil {
		toSerialize["palette_index"] = o.PaletteIndex
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetFormulaStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette      *string  `json:"palette"`
		PaletteIndex *float64 `json:"palette_index,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}
	o.Palette = *all.Palette
	o.PaletteIndex = all.PaletteIndex

	return nil
}
