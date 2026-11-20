// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapStyle Test-only generated schema for the dashboard versioning proof of concept.
type GeomapStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette GeomapStylePalette `json:"palette"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PaletteFlip bool `json:"palette_flip"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapStyle instantiates a new GeomapStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapStyle(palette GeomapStylePalette, paletteFlip bool) *GeomapStyle {
	this := GeomapStyle{}
	this.Palette = palette
	this.PaletteFlip = paletteFlip
	return &this
}

// NewGeomapStyleWithDefaults instantiates a new GeomapStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapStyleWithDefaults() *GeomapStyle {
	this := GeomapStyle{}
	return &this
}

// GetPalette returns the Palette field value.
func (o *GeomapStyle) GetPalette() GeomapStylePalette {
	if o == nil {
		var ret GeomapStylePalette
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *GeomapStyle) GetPaletteOk() (*GeomapStylePalette, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *GeomapStyle) SetPalette(v GeomapStylePalette) {
	o.Palette = v
}

// GetPaletteFlip returns the PaletteFlip field value.
func (o *GeomapStyle) GetPaletteFlip() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PaletteFlip
}

// GetPaletteFlipOk returns a tuple with the PaletteFlip field value
// and a boolean to check if the value has been set.
func (o *GeomapStyle) GetPaletteFlipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaletteFlip, true
}

// SetPaletteFlip sets field value.
func (o *GeomapStyle) SetPaletteFlip(v bool) {
	o.PaletteFlip = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["palette"] = o.Palette
	toSerialize["palette_flip"] = o.PaletteFlip
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette     *GeomapStylePalette `json:"palette"`
		PaletteFlip *bool               `json:"palette_flip"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}
	if all.PaletteFlip == nil {
		return fmt.Errorf("required field palette_flip missing")
	}

	hasInvalidField := false
	if !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = *all.Palette
	}
	o.PaletteFlip = *all.PaletteFlip

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
