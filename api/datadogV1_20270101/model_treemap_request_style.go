// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapRequestStyle Test-only generated schema for the dashboard versioning proof of concept.
type TreemapRequestStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *TreemapRequestStylePalette `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapRequestStyle instantiates a new TreemapRequestStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapRequestStyle() *TreemapRequestStyle {
	this := TreemapRequestStyle{}
	return &this
}

// NewTreemapRequestStyleWithDefaults instantiates a new TreemapRequestStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapRequestStyleWithDefaults() *TreemapRequestStyle {
	this := TreemapRequestStyle{}
	return &this
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *TreemapRequestStyle) GetPalette() TreemapRequestStylePalette {
	if o == nil || o.Palette == nil {
		var ret TreemapRequestStylePalette
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapRequestStyle) GetPaletteOk() (*TreemapRequestStylePalette, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *TreemapRequestStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given TreemapRequestStylePalette and assigns it to the Palette field.
func (o *TreemapRequestStyle) SetPalette(v TreemapRequestStylePalette) {
	o.Palette = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapRequestStyle) MarshalJSON() ([]byte, error) {
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
func (o *TreemapRequestStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette *TreemapRequestStylePalette `json:"palette,omitempty"`
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
