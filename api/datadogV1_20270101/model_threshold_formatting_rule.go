// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThresholdFormattingRule Test-only generated schema for the dashboard versioning proof of concept.
type ThresholdFormattingRule struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Comparator Comparator `json:"comparator"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomBgColor *string `json:"custom_bg_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomFgColor *string `json:"custom_fg_color,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideValue *bool `json:"hide_value,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	HighlightRow *bool `json:"highlight_row,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ImageUrl *string `json:"image_url,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric *string `json:"metric,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Palette string `json:"palette"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewThresholdFormattingRule instantiates a new ThresholdFormattingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewThresholdFormattingRule(comparator Comparator, palette string, value float64) *ThresholdFormattingRule {
	this := ThresholdFormattingRule{}
	this.Comparator = comparator
	this.Palette = palette
	this.Value = value
	return &this
}

// NewThresholdFormattingRuleWithDefaults instantiates a new ThresholdFormattingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewThresholdFormattingRuleWithDefaults() *ThresholdFormattingRule {
	this := ThresholdFormattingRule{}
	return &this
}

// GetComparator returns the Comparator field value.
func (o *ThresholdFormattingRule) GetComparator() Comparator {
	if o == nil {
		var ret Comparator
		return ret
	}
	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetComparatorOk() (*Comparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value.
func (o *ThresholdFormattingRule) SetComparator(v Comparator) {
	o.Comparator = v
}

// GetCustomBgColor returns the CustomBgColor field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetCustomBgColor() string {
	if o == nil || o.CustomBgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomBgColor
}

// GetCustomBgColorOk returns a tuple with the CustomBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetCustomBgColorOk() (*string, bool) {
	if o == nil || o.CustomBgColor == nil {
		return nil, false
	}
	return o.CustomBgColor, true
}

// HasCustomBgColor returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasCustomBgColor() bool {
	return o != nil && o.CustomBgColor != nil
}

// SetCustomBgColor gets a reference to the given string and assigns it to the CustomBgColor field.
func (o *ThresholdFormattingRule) SetCustomBgColor(v string) {
	o.CustomBgColor = &v
}

// GetCustomFgColor returns the CustomFgColor field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetCustomFgColor() string {
	if o == nil || o.CustomFgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomFgColor
}

// GetCustomFgColorOk returns a tuple with the CustomFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetCustomFgColorOk() (*string, bool) {
	if o == nil || o.CustomFgColor == nil {
		return nil, false
	}
	return o.CustomFgColor, true
}

// HasCustomFgColor returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasCustomFgColor() bool {
	return o != nil && o.CustomFgColor != nil
}

// SetCustomFgColor gets a reference to the given string and assigns it to the CustomFgColor field.
func (o *ThresholdFormattingRule) SetCustomFgColor(v string) {
	o.CustomFgColor = &v
}

// GetHideValue returns the HideValue field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetHideValue() bool {
	if o == nil || o.HideValue == nil {
		var ret bool
		return ret
	}
	return *o.HideValue
}

// GetHideValueOk returns a tuple with the HideValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetHideValueOk() (*bool, bool) {
	if o == nil || o.HideValue == nil {
		return nil, false
	}
	return o.HideValue, true
}

// HasHideValue returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasHideValue() bool {
	return o != nil && o.HideValue != nil
}

// SetHideValue gets a reference to the given bool and assigns it to the HideValue field.
func (o *ThresholdFormattingRule) SetHideValue(v bool) {
	o.HideValue = &v
}

// GetHighlightRow returns the HighlightRow field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetHighlightRow() bool {
	if o == nil || o.HighlightRow == nil {
		var ret bool
		return ret
	}
	return *o.HighlightRow
}

// GetHighlightRowOk returns a tuple with the HighlightRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetHighlightRowOk() (*bool, bool) {
	if o == nil || o.HighlightRow == nil {
		return nil, false
	}
	return o.HighlightRow, true
}

// HasHighlightRow returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasHighlightRow() bool {
	return o != nil && o.HighlightRow != nil
}

// SetHighlightRow gets a reference to the given bool and assigns it to the HighlightRow field.
func (o *ThresholdFormattingRule) SetHighlightRow(v bool) {
	o.HighlightRow = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasImageUrl() bool {
	return o != nil && o.ImageUrl != nil
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ThresholdFormattingRule) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ThresholdFormattingRule) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ThresholdFormattingRule) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ThresholdFormattingRule) SetMetric(v string) {
	o.Metric = &v
}

// GetPalette returns the Palette field value.
func (o *ThresholdFormattingRule) GetPalette() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetPaletteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Palette, true
}

// SetPalette sets field value.
func (o *ThresholdFormattingRule) SetPalette(v string) {
	o.Palette = v
}

// GetValue returns the Value field value.
func (o *ThresholdFormattingRule) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ThresholdFormattingRule) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ThresholdFormattingRule) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ThresholdFormattingRule) MarshalJSON() ([]byte, error) {
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
	if o.HighlightRow != nil {
		toSerialize["highlight_row"] = o.HighlightRow
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	toSerialize["palette"] = o.Palette
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ThresholdFormattingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comparator    *Comparator `json:"comparator"`
		CustomBgColor *string     `json:"custom_bg_color,omitempty"`
		CustomFgColor *string     `json:"custom_fg_color,omitempty"`
		HideValue     *bool       `json:"hide_value,omitempty"`
		HighlightRow  *bool       `json:"highlight_row,omitempty"`
		ImageUrl      *string     `json:"image_url,omitempty"`
		Metric        *string     `json:"metric,omitempty"`
		Palette       *string     `json:"palette"`
		Value         *float64    `json:"value"`
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
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
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
	o.HighlightRow = all.HighlightRow
	o.ImageUrl = all.ImageUrl
	o.Metric = all.Metric
	o.Palette = *all.Palette
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
