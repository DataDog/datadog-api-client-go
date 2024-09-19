// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// WidgetStyle Widget style definition.
type WidgetStyle struct {
	// Color palette to apply to the widget.
	Palette *string `json:"palette,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewWidgetStyle instantiates a new WidgetStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetStyle() *WidgetStyle {
	this := WidgetStyle{}
	return &this
}

// NewWidgetStyleWithDefaults instantiates a new WidgetStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
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
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *WidgetStyle) SetPalette(v string) {
	o.Palette = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o WidgetStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Palette *string `json:"palette,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "palette",  })
	} else {
		return err
	}
	o.Palette = all.Palette

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
