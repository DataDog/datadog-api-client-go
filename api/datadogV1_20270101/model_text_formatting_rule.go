// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextFormattingRule Test-only generated schema for the dashboard versioning proof of concept.
type TextFormattingRule struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomBgColor *string `json:"custom_bg_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomFgColor *string `json:"custom_fg_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HighlightRow *bool `json:"highlight_row,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Match TextMatch `json:"match"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette *ThresholdPaletteWithCustom `json:"palette,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Replace *TextFormattingRule20270101Replace `json:"replace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTextFormattingRule instantiates a new TextFormattingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTextFormattingRule(match TextMatch) *TextFormattingRule {
	this := TextFormattingRule{}
	this.Match = match
	return &this
}

// NewTextFormattingRuleWithDefaults instantiates a new TextFormattingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTextFormattingRuleWithDefaults() *TextFormattingRule {
	this := TextFormattingRule{}
	return &this
}

// GetCustomBgColor returns the CustomBgColor field value if set, zero value otherwise.
func (o *TextFormattingRule) GetCustomBgColor() string {
	if o == nil || o.CustomBgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomBgColor
}

// GetCustomBgColorOk returns a tuple with the CustomBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetCustomBgColorOk() (*string, bool) {
	if o == nil || o.CustomBgColor == nil {
		return nil, false
	}
	return o.CustomBgColor, true
}

// HasCustomBgColor returns a boolean if a field has been set.
func (o *TextFormattingRule) HasCustomBgColor() bool {
	return o != nil && o.CustomBgColor != nil
}

// SetCustomBgColor gets a reference to the given string and assigns it to the CustomBgColor field.
func (o *TextFormattingRule) SetCustomBgColor(v string) {
	o.CustomBgColor = &v
}

// GetCustomFgColor returns the CustomFgColor field value if set, zero value otherwise.
func (o *TextFormattingRule) GetCustomFgColor() string {
	if o == nil || o.CustomFgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomFgColor
}

// GetCustomFgColorOk returns a tuple with the CustomFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetCustomFgColorOk() (*string, bool) {
	if o == nil || o.CustomFgColor == nil {
		return nil, false
	}
	return o.CustomFgColor, true
}

// HasCustomFgColor returns a boolean if a field has been set.
func (o *TextFormattingRule) HasCustomFgColor() bool {
	return o != nil && o.CustomFgColor != nil
}

// SetCustomFgColor gets a reference to the given string and assigns it to the CustomFgColor field.
func (o *TextFormattingRule) SetCustomFgColor(v string) {
	o.CustomFgColor = &v
}

// GetHighlightRow returns the HighlightRow field value if set, zero value otherwise.
func (o *TextFormattingRule) GetHighlightRow() bool {
	if o == nil || o.HighlightRow == nil {
		var ret bool
		return ret
	}
	return *o.HighlightRow
}

// GetHighlightRowOk returns a tuple with the HighlightRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetHighlightRowOk() (*bool, bool) {
	if o == nil || o.HighlightRow == nil {
		return nil, false
	}
	return o.HighlightRow, true
}

// HasHighlightRow returns a boolean if a field has been set.
func (o *TextFormattingRule) HasHighlightRow() bool {
	return o != nil && o.HighlightRow != nil
}

// SetHighlightRow gets a reference to the given bool and assigns it to the HighlightRow field.
func (o *TextFormattingRule) SetHighlightRow(v bool) {
	o.HighlightRow = &v
}

// GetMatch returns the Match field value.
func (o *TextFormattingRule) GetMatch() TextMatch {
	if o == nil {
		var ret TextMatch
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetMatchOk() (*TextMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *TextFormattingRule) SetMatch(v TextMatch) {
	o.Match = v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *TextFormattingRule) GetPalette() ThresholdPaletteWithCustom {
	if o == nil || o.Palette == nil {
		var ret ThresholdPaletteWithCustom
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetPaletteOk() (*ThresholdPaletteWithCustom, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *TextFormattingRule) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given ThresholdPaletteWithCustom and assigns it to the Palette field.
func (o *TextFormattingRule) SetPalette(v ThresholdPaletteWithCustom) {
	o.Palette = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TextFormattingRule) GetReplace() TextFormattingRule20270101Replace {
	if o == nil || o.Replace == nil {
		var ret TextFormattingRule20270101Replace
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextFormattingRule) GetReplaceOk() (*TextFormattingRule20270101Replace, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TextFormattingRule) HasReplace() bool {
	return o != nil && o.Replace != nil
}

// SetReplace gets a reference to the given TextFormattingRule20270101Replace and assigns it to the Replace field.
func (o *TextFormattingRule) SetReplace(v TextFormattingRule20270101Replace) {
	o.Replace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TextFormattingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomBgColor != nil {
		toSerialize["custom_bg_color"] = o.CustomBgColor
	}
	if o.CustomFgColor != nil {
		toSerialize["custom_fg_color"] = o.CustomFgColor
	}
	if o.HighlightRow != nil {
		toSerialize["highlight_row"] = o.HighlightRow
	}
	toSerialize["match"] = o.Match
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TextFormattingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomBgColor *string                            `json:"custom_bg_color,omitempty"`
		CustomFgColor *string                            `json:"custom_fg_color,omitempty"`
		HighlightRow  *bool                              `json:"highlight_row,omitempty"`
		Match         *TextMatch                         `json:"match"`
		Palette       *ThresholdPaletteWithCustom        `json:"palette,omitempty"`
		Replace       *TextFormattingRule20270101Replace `json:"replace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}
	o.CustomBgColor = all.CustomBgColor
	o.CustomFgColor = all.CustomFgColor
	o.HighlightRow = all.HighlightRow
	o.Match = *all.Match
	o.Palette = all.Palette
	o.Replace = all.Replace

	return nil
}
