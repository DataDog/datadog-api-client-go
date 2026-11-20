// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapRequestStyle Test-only generated schema for the dashboard versioning proof of concept.
type GeomapRequestStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColorBy *string `json:"color_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *string `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapRequestStyle instantiates a new GeomapRequestStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapRequestStyle() *GeomapRequestStyle {
	this := GeomapRequestStyle{}
	return &this
}

// NewGeomapRequestStyleWithDefaults instantiates a new GeomapRequestStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapRequestStyleWithDefaults() *GeomapRequestStyle {
	this := GeomapRequestStyle{}
	return &this
}

// GetColorBy returns the ColorBy field value if set, zero value otherwise.
func (o *GeomapRequestStyle) GetColorBy() string {
	if o == nil || o.ColorBy == nil {
		var ret string
		return ret
	}
	return *o.ColorBy
}

// GetColorByOk returns a tuple with the ColorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapRequestStyle) GetColorByOk() (*string, bool) {
	if o == nil || o.ColorBy == nil {
		return nil, false
	}
	return o.ColorBy, true
}

// HasColorBy returns a boolean if a field has been set.
func (o *GeomapRequestStyle) HasColorBy() bool {
	return o != nil && o.ColorBy != nil
}

// SetColorBy gets a reference to the given string and assigns it to the ColorBy field.
func (o *GeomapRequestStyle) SetColorBy(v string) {
	o.ColorBy = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *GeomapRequestStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapRequestStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *GeomapRequestStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *GeomapRequestStyle) SetPalette(v string) {
	o.Palette = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapRequestStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColorBy != nil {
		toSerialize["color_by"] = o.ColorBy
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapRequestStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorBy *string `json:"color_by,omitempty"`
		Palette *string `json:"palette,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ColorBy = all.ColorBy
	o.Palette = all.Palette

	return nil
}
