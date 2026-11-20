// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageBackendDefinition Test-only generated schema for the dashboard versioning proof of concept.
type ImageBackendDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Caption *string `json:"caption,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasBackground *bool `json:"has_background,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HasBorder *bool `json:"has_border,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HorizontalAlign *ImageBackendDefinitionHorizontalAlign `json:"horizontal_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Margin *ImageMargin `json:"margin,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sizing *ImageBackendDefinition20270101Sizing `json:"sizing,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ImageBackendDefinitionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Url string `json:"url"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	UrlDarkTheme *string `json:"url_dark_theme,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	VerticalAlign *ImageBackendDefinitionVerticalAlign `json:"vertical_align,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewImageBackendDefinition instantiates a new ImageBackendDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImageBackendDefinition(typeVar ImageBackendDefinitionType, url string) *ImageBackendDefinition {
	this := ImageBackendDefinition{}
	this.Type = typeVar
	this.Url = url
	return &this
}

// NewImageBackendDefinitionWithDefaults instantiates a new ImageBackendDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImageBackendDefinitionWithDefaults() *ImageBackendDefinition {
	this := ImageBackendDefinition{}
	return &this
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasCaption() bool {
	return o != nil && o.Caption != nil
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *ImageBackendDefinition) SetCaption(v string) {
	o.Caption = &v
}

// GetHasBackground returns the HasBackground field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetHasBackground() bool {
	if o == nil || o.HasBackground == nil {
		var ret bool
		return ret
	}
	return *o.HasBackground
}

// GetHasBackgroundOk returns a tuple with the HasBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetHasBackgroundOk() (*bool, bool) {
	if o == nil || o.HasBackground == nil {
		return nil, false
	}
	return o.HasBackground, true
}

// HasHasBackground returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasHasBackground() bool {
	return o != nil && o.HasBackground != nil
}

// SetHasBackground gets a reference to the given bool and assigns it to the HasBackground field.
func (o *ImageBackendDefinition) SetHasBackground(v bool) {
	o.HasBackground = &v
}

// GetHasBorder returns the HasBorder field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetHasBorder() bool {
	if o == nil || o.HasBorder == nil {
		var ret bool
		return ret
	}
	return *o.HasBorder
}

// GetHasBorderOk returns a tuple with the HasBorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetHasBorderOk() (*bool, bool) {
	if o == nil || o.HasBorder == nil {
		return nil, false
	}
	return o.HasBorder, true
}

// HasHasBorder returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasHasBorder() bool {
	return o != nil && o.HasBorder != nil
}

// SetHasBorder gets a reference to the given bool and assigns it to the HasBorder field.
func (o *ImageBackendDefinition) SetHasBorder(v bool) {
	o.HasBorder = &v
}

// GetHorizontalAlign returns the HorizontalAlign field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetHorizontalAlign() ImageBackendDefinitionHorizontalAlign {
	if o == nil || o.HorizontalAlign == nil {
		var ret ImageBackendDefinitionHorizontalAlign
		return ret
	}
	return *o.HorizontalAlign
}

// GetHorizontalAlignOk returns a tuple with the HorizontalAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetHorizontalAlignOk() (*ImageBackendDefinitionHorizontalAlign, bool) {
	if o == nil || o.HorizontalAlign == nil {
		return nil, false
	}
	return o.HorizontalAlign, true
}

// HasHorizontalAlign returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasHorizontalAlign() bool {
	return o != nil && o.HorizontalAlign != nil
}

// SetHorizontalAlign gets a reference to the given ImageBackendDefinitionHorizontalAlign and assigns it to the HorizontalAlign field.
func (o *ImageBackendDefinition) SetHorizontalAlign(v ImageBackendDefinitionHorizontalAlign) {
	o.HorizontalAlign = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetMargin() ImageMargin {
	if o == nil || o.Margin == nil {
		var ret ImageMargin
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetMarginOk() (*ImageMargin, bool) {
	if o == nil || o.Margin == nil {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasMargin() bool {
	return o != nil && o.Margin != nil
}

// SetMargin gets a reference to the given ImageMargin and assigns it to the Margin field.
func (o *ImageBackendDefinition) SetMargin(v ImageMargin) {
	o.Margin = &v
}

// GetSizing returns the Sizing field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetSizing() ImageBackendDefinition20270101Sizing {
	if o == nil || o.Sizing == nil {
		var ret ImageBackendDefinition20270101Sizing
		return ret
	}
	return *o.Sizing
}

// GetSizingOk returns a tuple with the Sizing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetSizingOk() (*ImageBackendDefinition20270101Sizing, bool) {
	if o == nil || o.Sizing == nil {
		return nil, false
	}
	return o.Sizing, true
}

// HasSizing returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasSizing() bool {
	return o != nil && o.Sizing != nil
}

// SetSizing gets a reference to the given ImageBackendDefinition20270101Sizing and assigns it to the Sizing field.
func (o *ImageBackendDefinition) SetSizing(v ImageBackendDefinition20270101Sizing) {
	o.Sizing = &v
}

// GetType returns the Type field value.
func (o *ImageBackendDefinition) GetType() ImageBackendDefinitionType {
	if o == nil {
		var ret ImageBackendDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetTypeOk() (*ImageBackendDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ImageBackendDefinition) SetType(v ImageBackendDefinitionType) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *ImageBackendDefinition) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ImageBackendDefinition) SetUrl(v string) {
	o.Url = v
}

// GetUrlDarkTheme returns the UrlDarkTheme field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetUrlDarkTheme() string {
	if o == nil || o.UrlDarkTheme == nil {
		var ret string
		return ret
	}
	return *o.UrlDarkTheme
}

// GetUrlDarkThemeOk returns a tuple with the UrlDarkTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetUrlDarkThemeOk() (*string, bool) {
	if o == nil || o.UrlDarkTheme == nil {
		return nil, false
	}
	return o.UrlDarkTheme, true
}

// HasUrlDarkTheme returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasUrlDarkTheme() bool {
	return o != nil && o.UrlDarkTheme != nil
}

// SetUrlDarkTheme gets a reference to the given string and assigns it to the UrlDarkTheme field.
func (o *ImageBackendDefinition) SetUrlDarkTheme(v string) {
	o.UrlDarkTheme = &v
}

// GetVerticalAlign returns the VerticalAlign field value if set, zero value otherwise.
func (o *ImageBackendDefinition) GetVerticalAlign() ImageBackendDefinitionVerticalAlign {
	if o == nil || o.VerticalAlign == nil {
		var ret ImageBackendDefinitionVerticalAlign
		return ret
	}
	return *o.VerticalAlign
}

// GetVerticalAlignOk returns a tuple with the VerticalAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBackendDefinition) GetVerticalAlignOk() (*ImageBackendDefinitionVerticalAlign, bool) {
	if o == nil || o.VerticalAlign == nil {
		return nil, false
	}
	return o.VerticalAlign, true
}

// HasVerticalAlign returns a boolean if a field has been set.
func (o *ImageBackendDefinition) HasVerticalAlign() bool {
	return o != nil && o.VerticalAlign != nil
}

// SetVerticalAlign gets a reference to the given ImageBackendDefinitionVerticalAlign and assigns it to the VerticalAlign field.
func (o *ImageBackendDefinition) SetVerticalAlign(v ImageBackendDefinitionVerticalAlign) {
	o.VerticalAlign = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImageBackendDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Caption != nil {
		toSerialize["caption"] = o.Caption
	}
	if o.HasBackground != nil {
		toSerialize["has_background"] = o.HasBackground
	}
	if o.HasBorder != nil {
		toSerialize["has_border"] = o.HasBorder
	}
	if o.HorizontalAlign != nil {
		toSerialize["horizontal_align"] = o.HorizontalAlign
	}
	if o.Margin != nil {
		toSerialize["margin"] = o.Margin
	}
	if o.Sizing != nil {
		toSerialize["sizing"] = o.Sizing
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if o.UrlDarkTheme != nil {
		toSerialize["url_dark_theme"] = o.UrlDarkTheme
	}
	if o.VerticalAlign != nil {
		toSerialize["vertical_align"] = o.VerticalAlign
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImageBackendDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Caption         *string                                `json:"caption,omitempty"`
		HasBackground   *bool                                  `json:"has_background,omitempty"`
		HasBorder       *bool                                  `json:"has_border,omitempty"`
		HorizontalAlign *ImageBackendDefinitionHorizontalAlign `json:"horizontal_align,omitempty"`
		Margin          *ImageMargin                           `json:"margin,omitempty"`
		Sizing          *ImageBackendDefinition20270101Sizing  `json:"sizing,omitempty"`
		Type            *ImageBackendDefinitionType            `json:"type"`
		Url             *string                                `json:"url"`
		UrlDarkTheme    *string                                `json:"url_dark_theme,omitempty"`
		VerticalAlign   *ImageBackendDefinitionVerticalAlign   `json:"vertical_align,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}

	hasInvalidField := false
	o.Caption = all.Caption
	o.HasBackground = all.HasBackground
	o.HasBorder = all.HasBorder
	if all.HorizontalAlign != nil && !all.HorizontalAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.HorizontalAlign = all.HorizontalAlign
	}
	if all.Margin != nil && !all.Margin.IsValid() {
		hasInvalidField = true
	} else {
		o.Margin = all.Margin
	}
	o.Sizing = all.Sizing
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Url = *all.Url
	o.UrlDarkTheme = all.UrlDarkTheme
	if all.VerticalAlign != nil && !all.VerticalAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.VerticalAlign = all.VerticalAlign
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
