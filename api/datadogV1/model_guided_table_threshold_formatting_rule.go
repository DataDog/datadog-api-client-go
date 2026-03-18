// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableThresholdFormattingRule Conditional formatting rule that applies when a value satisfies a threshold comparator condition.
type GuidedTableThresholdFormattingRule struct {
	// Comparator to apply.
	Comparator WidgetComparator `json:"comparator"`
	// Custom background color. Used with `custom_bg` palette.
	CustomBgColor *string `json:"custom_bg_color,omitempty"`
	// Custom text color. Used with `custom_text` palette.
	CustomFgColor *string `json:"custom_fg_color,omitempty"`
	// Whether to hide the data value when this rule matches.
	HideValue *bool `json:"hide_value,omitempty"`
	// URL of an image to display as background.
	ImageUrl *string `json:"image_url,omitempty"`
	// Color palette used by threshold-based conditional formatting and text formatting rules in a guided table.
	Palette GuidedTableThresholdPalette `json:"palette"`
	// Threshold value to compare against.
	Value *GuidedTableThresholdFormattingRuleValue `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableThresholdFormattingRule instantiates a new GuidedTableThresholdFormattingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableThresholdFormattingRule(comparator WidgetComparator, palette GuidedTableThresholdPalette) *GuidedTableThresholdFormattingRule {
	this := GuidedTableThresholdFormattingRule{}
	this.Comparator = comparator
	this.Palette = palette
	return &this
}

// NewGuidedTableThresholdFormattingRuleWithDefaults instantiates a new GuidedTableThresholdFormattingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableThresholdFormattingRuleWithDefaults() *GuidedTableThresholdFormattingRule {
	this := GuidedTableThresholdFormattingRule{}
	return &this
}

// GetComparator returns the Comparator field value.
func (o *GuidedTableThresholdFormattingRule) GetComparator() WidgetComparator {
	if o == nil {
		var ret WidgetComparator
		return ret
	}
	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetComparatorOk() (*WidgetComparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value.
func (o *GuidedTableThresholdFormattingRule) SetComparator(v WidgetComparator) {
	o.Comparator = v
}

// GetCustomBgColor returns the CustomBgColor field value if set, zero value otherwise.
func (o *GuidedTableThresholdFormattingRule) GetCustomBgColor() string {
	if o == nil || o.CustomBgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomBgColor
}

// GetCustomBgColorOk returns a tuple with the CustomBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetCustomBgColorOk() (*string, bool) {
	if o == nil || o.CustomBgColor == nil {
		return nil, false
	}
	return o.CustomBgColor, true
}

// HasCustomBgColor returns a boolean if a field has been set.
func (o *GuidedTableThresholdFormattingRule) HasCustomBgColor() bool {
	return o != nil && o.CustomBgColor != nil
}

// SetCustomBgColor gets a reference to the given string and assigns it to the CustomBgColor field.
func (o *GuidedTableThresholdFormattingRule) SetCustomBgColor(v string) {
	o.CustomBgColor = &v
}

// GetCustomFgColor returns the CustomFgColor field value if set, zero value otherwise.
func (o *GuidedTableThresholdFormattingRule) GetCustomFgColor() string {
	if o == nil || o.CustomFgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomFgColor
}

// GetCustomFgColorOk returns a tuple with the CustomFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetCustomFgColorOk() (*string, bool) {
	if o == nil || o.CustomFgColor == nil {
		return nil, false
	}
	return o.CustomFgColor, true
}

// HasCustomFgColor returns a boolean if a field has been set.
func (o *GuidedTableThresholdFormattingRule) HasCustomFgColor() bool {
	return o != nil && o.CustomFgColor != nil
}

// SetCustomFgColor gets a reference to the given string and assigns it to the CustomFgColor field.
func (o *GuidedTableThresholdFormattingRule) SetCustomFgColor(v string) {
	o.CustomFgColor = &v
}

// GetHideValue returns the HideValue field value if set, zero value otherwise.
func (o *GuidedTableThresholdFormattingRule) GetHideValue() bool {
	if o == nil || o.HideValue == nil {
		var ret bool
		return ret
	}
	return *o.HideValue
}

// GetHideValueOk returns a tuple with the HideValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetHideValueOk() (*bool, bool) {
	if o == nil || o.HideValue == nil {
		return nil, false
	}
	return o.HideValue, true
}

// HasHideValue returns a boolean if a field has been set.
func (o *GuidedTableThresholdFormattingRule) HasHideValue() bool {
	return o != nil && o.HideValue != nil
}

// SetHideValue gets a reference to the given bool and assigns it to the HideValue field.
func (o *GuidedTableThresholdFormattingRule) SetHideValue(v bool) {
	o.HideValue = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GuidedTableThresholdFormattingRule) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GuidedTableThresholdFormattingRule) HasImageUrl() bool {
	return o != nil && o.ImageUrl != nil
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GuidedTableThresholdFormattingRule) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetPalette returns the Palette field value.
func (o *GuidedTableThresholdFormattingRule) GetPalette() GuidedTableThresholdPalette {
	if o == nil {
		var ret GuidedTableThresholdPalette
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetPaletteOk() (*GuidedTableThresholdPalette, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *GuidedTableThresholdFormattingRule) SetPalette(v GuidedTableThresholdPalette) {
	o.Palette = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GuidedTableThresholdFormattingRule) GetValue() GuidedTableThresholdFormattingRuleValue {
	if o == nil || o.Value == nil {
		var ret GuidedTableThresholdFormattingRuleValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableThresholdFormattingRule) GetValueOk() (*GuidedTableThresholdFormattingRuleValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GuidedTableThresholdFormattingRule) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given GuidedTableThresholdFormattingRuleValue and assigns it to the Value field.
func (o *GuidedTableThresholdFormattingRule) SetValue(v GuidedTableThresholdFormattingRuleValue) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableThresholdFormattingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["comparator"] = o.Comparator
	if o.CustomBgColor != nil {
		toSerialize["custom_bg_color"] = o.CustomBgColor
	}
	if o.CustomFgColor != nil {
		toSerialize["custom_fg_color"] = o.CustomFgColor
	}
	if o.HideValue != nil {
		toSerialize["hide_value"] = o.HideValue
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	toSerialize["palette"] = o.Palette
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableThresholdFormattingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comparator    *WidgetComparator                        `json:"comparator"`
		CustomBgColor *string                                  `json:"custom_bg_color,omitempty"`
		CustomFgColor *string                                  `json:"custom_fg_color,omitempty"`
		HideValue     *bool                                    `json:"hide_value,omitempty"`
		ImageUrl      *string                                  `json:"image_url,omitempty"`
		Palette       *GuidedTableThresholdPalette             `json:"palette"`
		Value         *GuidedTableThresholdFormattingRuleValue `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Comparator == nil {
		return fmt.Errorf("required field comparator missing")
	}
	if all.Palette == nil {
		return fmt.Errorf("required field palette missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comparator", "custom_bg_color", "custom_fg_color", "hide_value", "image_url", "palette", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Comparator.IsValid() {
		hasInvalidField = true
	} else {
		o.Comparator = *all.Comparator
	}
	o.CustomBgColor = all.CustomBgColor
	o.CustomFgColor = all.CustomFgColor
	o.HideValue = all.HideValue
	o.ImageUrl = all.ImageUrl
	if !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = *all.Palette
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
