// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyle Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyle struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Display *ToplistBackendDefinitionStyle20270101Display `json:"display,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *string `json:"palette,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scaling *ToplistBackendDefinitionStyleScaling `json:"scaling,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewToplistBackendDefinitionStyle instantiates a new ToplistBackendDefinitionStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistBackendDefinitionStyle() *ToplistBackendDefinitionStyle {
	this := ToplistBackendDefinitionStyle{}
	return &this
}

// NewToplistBackendDefinitionStyleWithDefaults instantiates a new ToplistBackendDefinitionStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistBackendDefinitionStyleWithDefaults() *ToplistBackendDefinitionStyle {
	this := ToplistBackendDefinitionStyle{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ToplistBackendDefinitionStyle) GetDisplay() ToplistBackendDefinitionStyle20270101Display {
	if o == nil || o.Display == nil {
		var ret ToplistBackendDefinitionStyle20270101Display
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle) GetDisplayOk() (*ToplistBackendDefinitionStyle20270101Display, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ToplistBackendDefinitionStyle) HasDisplay() bool {
	return o != nil && o.Display != nil
}

// SetDisplay gets a reference to the given ToplistBackendDefinitionStyle20270101Display and assigns it to the Display field.
func (o *ToplistBackendDefinitionStyle) SetDisplay(v ToplistBackendDefinitionStyle20270101Display) {
	o.Display = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *ToplistBackendDefinitionStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *ToplistBackendDefinitionStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *ToplistBackendDefinitionStyle) SetPalette(v string) {
	o.Palette = &v
}

// GetScaling returns the Scaling field value if set, zero value otherwise.
func (o *ToplistBackendDefinitionStyle) GetScaling() ToplistBackendDefinitionStyleScaling {
	if o == nil || o.Scaling == nil {
		var ret ToplistBackendDefinitionStyleScaling
		return ret
	}
	return *o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistBackendDefinitionStyle) GetScalingOk() (*ToplistBackendDefinitionStyleScaling, bool) {
	if o == nil || o.Scaling == nil {
		return nil, false
	}
	return o.Scaling, true
}

// HasScaling returns a boolean if a field has been set.
func (o *ToplistBackendDefinitionStyle) HasScaling() bool {
	return o != nil && o.Scaling != nil
}

// SetScaling gets a reference to the given ToplistBackendDefinitionStyleScaling and assigns it to the Scaling field.
func (o *ToplistBackendDefinitionStyle) SetScaling(v ToplistBackendDefinitionStyleScaling) {
	o.Scaling = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistBackendDefinitionStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.Scaling != nil {
		toSerialize["scaling"] = o.Scaling
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ToplistBackendDefinitionStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Display *ToplistBackendDefinitionStyle20270101Display `json:"display,omitempty"`
		Palette *string                                       `json:"palette,omitempty"`
		Scaling *ToplistBackendDefinitionStyleScaling         `json:"scaling,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Display = all.Display
	o.Palette = all.Palette
	if all.Scaling != nil && !all.Scaling.IsValid() {
		hasInvalidField = true
	} else {
		o.Scaling = all.Scaling
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
