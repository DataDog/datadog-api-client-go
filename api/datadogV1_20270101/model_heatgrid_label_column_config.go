// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridLabelColumnConfig Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridLabelColumnConfig struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats []TextReplacementRule `json:"text_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width *HeatgridLabelColumnWidth `json:"width,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridLabelColumnConfig instantiates a new HeatgridLabelColumnConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridLabelColumnConfig() *HeatgridLabelColumnConfig {
	this := HeatgridLabelColumnConfig{}
	return &this
}

// NewHeatgridLabelColumnConfigWithDefaults instantiates a new HeatgridLabelColumnConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridLabelColumnConfigWithDefaults() *HeatgridLabelColumnConfig {
	this := HeatgridLabelColumnConfig{}
	return &this
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *HeatgridLabelColumnConfig) GetTextFormats() []TextReplacementRule {
	if o == nil || o.TextFormats == nil {
		var ret []TextReplacementRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridLabelColumnConfig) GetTextFormatsOk() (*[]TextReplacementRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *HeatgridLabelColumnConfig) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given []TextReplacementRule and assigns it to the TextFormats field.
func (o *HeatgridLabelColumnConfig) SetTextFormats(v []TextReplacementRule) {
	o.TextFormats = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *HeatgridLabelColumnConfig) GetWidth() HeatgridLabelColumnWidth {
	if o == nil || o.Width == nil {
		var ret HeatgridLabelColumnWidth
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridLabelColumnConfig) GetWidthOk() (*HeatgridLabelColumnWidth, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *HeatgridLabelColumnConfig) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given HeatgridLabelColumnWidth and assigns it to the Width field.
func (o *HeatgridLabelColumnConfig) SetWidth(v HeatgridLabelColumnWidth) {
	o.Width = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridLabelColumnConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridLabelColumnConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TextFormats []TextReplacementRule     `json:"text_formats,omitempty"`
		Width       *HeatgridLabelColumnWidth `json:"width,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.TextFormats = all.TextFormats
	if all.Width != nil && !all.Width.IsValid() {
		hasInvalidField = true
	} else {
		o.Width = all.Width
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
