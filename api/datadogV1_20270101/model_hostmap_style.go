// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapStyle Test-only generated schema for the dashboard versioning proof of concept.
type HostmapStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	FillMax *float64 `json:"fill_max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FillMin *float64 `json:"fill_min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *HostmapStylePalette `json:"palette,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PaletteFlip *bool `json:"palette_flip,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapStyle instantiates a new HostmapStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapStyle() *HostmapStyle {
	this := HostmapStyle{}
	return &this
}

// NewHostmapStyleWithDefaults instantiates a new HostmapStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapStyleWithDefaults() *HostmapStyle {
	this := HostmapStyle{}
	return &this
}

// GetFillMax returns the FillMax field value if set, zero value otherwise.
func (o *HostmapStyle) GetFillMax() float64 {
	if o == nil || o.FillMax == nil {
		var ret float64
		return ret
	}
	return *o.FillMax
}

// GetFillMaxOk returns a tuple with the FillMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapStyle) GetFillMaxOk() (*float64, bool) {
	if o == nil || o.FillMax == nil {
		return nil, false
	}
	return o.FillMax, true
}

// HasFillMax returns a boolean if a field has been set.
func (o *HostmapStyle) HasFillMax() bool {
	return o != nil && o.FillMax != nil
}

// SetFillMax gets a reference to the given float64 and assigns it to the FillMax field.
func (o *HostmapStyle) SetFillMax(v float64) {
	o.FillMax = &v
}

// GetFillMin returns the FillMin field value if set, zero value otherwise.
func (o *HostmapStyle) GetFillMin() float64 {
	if o == nil || o.FillMin == nil {
		var ret float64
		return ret
	}
	return *o.FillMin
}

// GetFillMinOk returns a tuple with the FillMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapStyle) GetFillMinOk() (*float64, bool) {
	if o == nil || o.FillMin == nil {
		return nil, false
	}
	return o.FillMin, true
}

// HasFillMin returns a boolean if a field has been set.
func (o *HostmapStyle) HasFillMin() bool {
	return o != nil && o.FillMin != nil
}

// SetFillMin gets a reference to the given float64 and assigns it to the FillMin field.
func (o *HostmapStyle) SetFillMin(v float64) {
	o.FillMin = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *HostmapStyle) GetPalette() HostmapStylePalette {
	if o == nil || o.Palette == nil {
		var ret HostmapStylePalette
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapStyle) GetPaletteOk() (*HostmapStylePalette, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *HostmapStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given HostmapStylePalette and assigns it to the Palette field.
func (o *HostmapStyle) SetPalette(v HostmapStylePalette) {
	o.Palette = &v
}

// GetPaletteFlip returns the PaletteFlip field value if set, zero value otherwise.
func (o *HostmapStyle) GetPaletteFlip() bool {
	if o == nil || o.PaletteFlip == nil {
		var ret bool
		return ret
	}
	return *o.PaletteFlip
}

// GetPaletteFlipOk returns a tuple with the PaletteFlip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapStyle) GetPaletteFlipOk() (*bool, bool) {
	if o == nil || o.PaletteFlip == nil {
		return nil, false
	}
	return o.PaletteFlip, true
}

// HasPaletteFlip returns a boolean if a field has been set.
func (o *HostmapStyle) HasPaletteFlip() bool {
	return o != nil && o.PaletteFlip != nil
}

// SetPaletteFlip gets a reference to the given bool and assigns it to the PaletteFlip field.
func (o *HostmapStyle) SetPaletteFlip(v bool) {
	o.PaletteFlip = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FillMax != nil {
		toSerialize["fill_max"] = o.FillMax
	}
	if o.FillMin != nil {
		toSerialize["fill_min"] = o.FillMin
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.PaletteFlip != nil {
		toSerialize["palette_flip"] = o.PaletteFlip
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostmapStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FillMax     *float64             `json:"fill_max,omitempty"`
		FillMin     *float64             `json:"fill_min,omitempty"`
		Palette     *HostmapStylePalette `json:"palette,omitempty"`
		PaletteFlip *bool                `json:"palette_flip,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.FillMax = all.FillMax
	o.FillMin = all.FillMin
	if all.Palette != nil && !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = all.Palette
	}
	o.PaletteFlip = all.PaletteFlip

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
