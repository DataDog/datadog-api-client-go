// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapWidgetStyle The style for the points coming from this request.
type GeomapWidgetStyle struct {
	// Field used for coloring the points based on a value.
	ColorBy *string `json:"color_by,omitempty"`
	// Defines the color of all the points for this request.
	Palette *string `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGeomapWidgetStyle instantiates a new GeomapWidgetStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapWidgetStyle() *GeomapWidgetStyle {
	this := GeomapWidgetStyle{}
	return &this
}

// NewGeomapWidgetStyleWithDefaults instantiates a new GeomapWidgetStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapWidgetStyleWithDefaults() *GeomapWidgetStyle {
	this := GeomapWidgetStyle{}
	return &this
}

// GetColorBy returns the ColorBy field value if set, zero value otherwise.
func (o *GeomapWidgetStyle) GetColorBy() string {
	if o == nil || o.ColorBy == nil {
		var ret string
		return ret
	}
	return *o.ColorBy
}

// GetColorByOk returns a tuple with the ColorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetStyle) GetColorByOk() (*string, bool) {
	if o == nil || o.ColorBy == nil {
		return nil, false
	}
	return o.ColorBy, true
}

// HasColorBy returns a boolean if a field has been set.
func (o *GeomapWidgetStyle) HasColorBy() bool {
	return o != nil && o.ColorBy != nil
}

// SetColorBy gets a reference to the given string and assigns it to the ColorBy field.
func (o *GeomapWidgetStyle) SetColorBy(v string) {
	o.ColorBy = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *GeomapWidgetStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *GeomapWidgetStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *GeomapWidgetStyle) SetPalette(v string) {
	o.Palette = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapWidgetStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ColorBy != nil {
		toSerialize["color_by"] = o.ColorBy
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapWidgetStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorBy *string `json:"color_by,omitempty"`
		Palette *string `json:"palette,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"color_by", "palette"})
	} else {
		return err
	}
	o.ColorBy = all.ColorBy
	o.Palette = all.Palette

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
