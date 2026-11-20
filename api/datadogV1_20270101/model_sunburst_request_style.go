// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstRequestStyle Test-only generated schema for the dashboard versioning proof of concept.
type SunburstRequestStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *SunburstRequestStylePalette `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSunburstRequestStyle instantiates a new SunburstRequestStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstRequestStyle() *SunburstRequestStyle {
	this := SunburstRequestStyle{}
	return &this
}

// NewSunburstRequestStyleWithDefaults instantiates a new SunburstRequestStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstRequestStyleWithDefaults() *SunburstRequestStyle {
	this := SunburstRequestStyle{}
	return &this
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *SunburstRequestStyle) GetPalette() SunburstRequestStylePalette {
	if o == nil || o.Palette == nil {
		var ret SunburstRequestStylePalette
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstRequestStyle) GetPaletteOk() (*SunburstRequestStylePalette, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *SunburstRequestStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given SunburstRequestStylePalette and assigns it to the Palette field.
func (o *SunburstRequestStyle) SetPalette(v SunburstRequestStylePalette) {
	o.Palette = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstRequestStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SunburstRequestStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette *SunburstRequestStylePalette `json:"palette,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Palette != nil && !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = all.Palette
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
