// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreeTextBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type FreeTextBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Color *string `json:"color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FontSize *string `json:"font_size,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Text string `json:"text"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextAlign *FreeTextBackendDefinitionTextAlign `json:"text_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type FreeTextBackendDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFreeTextBackendDefinition instantiates a new FreeTextBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFreeTextBackendDefinition(text string, typeVar FreeTextBackendDefinitionType) *FreeTextBackendDefinition {
	this := FreeTextBackendDefinition{}
	this.Text = text
	this.Type = typeVar
	return &this
}

// NewFreeTextBackendDefinitionWithDefaults instantiates a new FreeTextBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFreeTextBackendDefinitionWithDefaults() *FreeTextBackendDefinition {
	this := FreeTextBackendDefinition{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *FreeTextBackendDefinition) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeTextBackendDefinition) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *FreeTextBackendDefinition) HasColor() bool {
	return o != nil && o.Color != nil
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *FreeTextBackendDefinition) SetColor(v string) {
	o.Color = &v
}

// GetFontSize returns the FontSize field value if set, zero value otherwise.
func (o *FreeTextBackendDefinition) GetFontSize() string {
	if o == nil || o.FontSize == nil {
		var ret string
		return ret
	}
	return *o.FontSize
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeTextBackendDefinition) GetFontSizeOk() (*string, bool) {
	if o == nil || o.FontSize == nil {
		return nil, false
	}
	return o.FontSize, true
}

// HasFontSize returns a boolean if a field has been set.
func (o *FreeTextBackendDefinition) HasFontSize() bool {
	return o != nil && o.FontSize != nil
}

// SetFontSize gets a reference to the given string and assigns it to the FontSize field.
func (o *FreeTextBackendDefinition) SetFontSize(v string) {
	o.FontSize = &v
}

// GetText returns the Text field value.
func (o *FreeTextBackendDefinition) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *FreeTextBackendDefinition) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *FreeTextBackendDefinition) SetText(v string) {
	o.Text = v
}

// GetTextAlign returns the TextAlign field value if set, zero value otherwise.
func (o *FreeTextBackendDefinition) GetTextAlign() FreeTextBackendDefinitionTextAlign {
	if o == nil || o.TextAlign == nil {
		var ret FreeTextBackendDefinitionTextAlign
		return ret
	}
	return *o.TextAlign
}

// GetTextAlignOk returns a tuple with the TextAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeTextBackendDefinition) GetTextAlignOk() (*FreeTextBackendDefinitionTextAlign, bool) {
	if o == nil || o.TextAlign == nil {
		return nil, false
	}
	return o.TextAlign, true
}

// HasTextAlign returns a boolean if a field has been set.
func (o *FreeTextBackendDefinition) HasTextAlign() bool {
	return o != nil && o.TextAlign != nil
}

// SetTextAlign gets a reference to the given FreeTextBackendDefinitionTextAlign and assigns it to the TextAlign field.
func (o *FreeTextBackendDefinition) SetTextAlign(v FreeTextBackendDefinitionTextAlign) {
	o.TextAlign = &v
}

// GetType returns the Type field value.
func (o *FreeTextBackendDefinition) GetType() FreeTextBackendDefinitionType {
	if o == nil {
		var ret FreeTextBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeTextBackendDefinition) GetTypeOk() (*FreeTextBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FreeTextBackendDefinition) SetType(v FreeTextBackendDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FreeTextBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.FontSize != nil {
		toSerialize["font_size"] = o.FontSize
	}
	toSerialize["text"] = o.Text
	if o.TextAlign != nil {
		toSerialize["text_align"] = o.TextAlign
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FreeTextBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color     *string                             `json:"color,omitempty"`
		FontSize  *string                             `json:"font_size,omitempty"`
		Text      *string                             `json:"text"`
		TextAlign *FreeTextBackendDefinitionTextAlign `json:"text_align,omitempty"`
		Type      *FreeTextBackendDefinitionType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Color = all.Color
	o.FontSize = all.FontSize
	o.Text = *all.Text
	if all.TextAlign != nil && !all.TextAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TextAlign = all.TextAlign
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
