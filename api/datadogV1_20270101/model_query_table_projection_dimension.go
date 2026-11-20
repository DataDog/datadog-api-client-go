// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjectionDimension Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjectionDimension struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode *AdvancedQueryTableCellDisplayMode `json:"cell_display_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats *ConditionalFormats `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DateFormat *QueryTableDateFormat `json:"date_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension QueryTableProjectionDimensionDimension `json:"dimension"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Link *QueryTableColumnLink `json:"link,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextAlign *WidgetTextAlign `json:"text_align,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats []TextFormattingRule `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableProjectionDimension instantiates a new QueryTableProjectionDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableProjectionDimension(column string, dimension QueryTableProjectionDimensionDimension) *QueryTableProjectionDimension {
	this := QueryTableProjectionDimension{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewQueryTableProjectionDimensionWithDefaults instantiates a new QueryTableProjectionDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableProjectionDimensionWithDefaults() *QueryTableProjectionDimension {
	this := QueryTableProjectionDimension{}
	return &this
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetCellDisplayMode() AdvancedQueryTableCellDisplayMode {
	if o == nil || o.CellDisplayMode == nil {
		var ret AdvancedQueryTableCellDisplayMode
		return ret
	}
	return *o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetCellDisplayModeOk() (*AdvancedQueryTableCellDisplayMode, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given AdvancedQueryTableCellDisplayMode and assigns it to the CellDisplayMode field.
func (o *QueryTableProjectionDimension) SetCellDisplayMode(v AdvancedQueryTableCellDisplayMode) {
	o.CellDisplayMode = &v
}

// GetColumn returns the Column field value.
func (o *QueryTableProjectionDimension) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *QueryTableProjectionDimension) SetColumn(v string) {
	o.Column = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetConditionalFormats() ConditionalFormats {
	if o == nil || o.ConditionalFormats == nil {
		var ret ConditionalFormats
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetConditionalFormatsOk() (*ConditionalFormats, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given ConditionalFormats and assigns it to the ConditionalFormats field.
func (o *QueryTableProjectionDimension) SetConditionalFormats(v ConditionalFormats) {
	o.ConditionalFormats = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetDateFormat() QueryTableDateFormat {
	if o == nil || o.DateFormat == nil {
		var ret QueryTableDateFormat
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetDateFormatOk() (*QueryTableDateFormat, bool) {
	if o == nil || o.DateFormat == nil {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasDateFormat() bool {
	return o != nil && o.DateFormat != nil
}

// SetDateFormat gets a reference to the given QueryTableDateFormat and assigns it to the DateFormat field.
func (o *QueryTableProjectionDimension) SetDateFormat(v QueryTableDateFormat) {
	o.DateFormat = &v
}

// GetDimension returns the Dimension field value.
func (o *QueryTableProjectionDimension) GetDimension() QueryTableProjectionDimensionDimension {
	if o == nil {
		var ret QueryTableProjectionDimensionDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetDimensionOk() (*QueryTableProjectionDimensionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *QueryTableProjectionDimension) SetDimension(v QueryTableProjectionDimensionDimension) {
	o.Dimension = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *QueryTableProjectionDimension) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetLink() QueryTableColumnLink {
	if o == nil || o.Link == nil {
		var ret QueryTableColumnLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetLinkOk() (*QueryTableColumnLink, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasLink() bool {
	return o != nil && o.Link != nil
}

// SetLink gets a reference to the given QueryTableColumnLink and assigns it to the Link field.
func (o *QueryTableProjectionDimension) SetLink(v QueryTableColumnLink) {
	o.Link = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *QueryTableProjectionDimension) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetTextAlign returns the TextAlign field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetTextAlign() WidgetTextAlign {
	if o == nil || o.TextAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TextAlign
}

// GetTextAlignOk returns a tuple with the TextAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetTextAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TextAlign == nil {
		return nil, false
	}
	return o.TextAlign, true
}

// HasTextAlign returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasTextAlign() bool {
	return o != nil && o.TextAlign != nil
}

// SetTextAlign gets a reference to the given WidgetTextAlign and assigns it to the TextAlign field.
func (o *QueryTableProjectionDimension) SetTextAlign(v WidgetTextAlign) {
	o.TextAlign = &v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *QueryTableProjectionDimension) GetTextFormats() []TextFormattingRule {
	if o == nil || o.TextFormats == nil {
		var ret []TextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjectionDimension) GetTextFormatsOk() (*[]TextFormattingRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *QueryTableProjectionDimension) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given []TextFormattingRule and assigns it to the TextFormats field.
func (o *QueryTableProjectionDimension) SetTextFormats(v []TextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableProjectionDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CellDisplayMode != nil {
		toSerialize["cell_display_mode"] = o.CellDisplayMode
	}
	toSerialize["column"] = o.Column
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.DateFormat != nil {
		toSerialize["date_format"] = o.DateFormat
	}
	toSerialize["dimension"] = o.Dimension
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	if o.TextAlign != nil {
		toSerialize["text_align"] = o.TextAlign
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableProjectionDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellDisplayMode    *AdvancedQueryTableCellDisplayMode      `json:"cell_display_mode,omitempty"`
		Column             *string                                 `json:"column"`
		ConditionalFormats *ConditionalFormats                     `json:"conditional_formats,omitempty"`
		DateFormat         *QueryTableDateFormat                   `json:"date_format,omitempty"`
		Dimension          *QueryTableProjectionDimensionDimension `json:"dimension"`
		IsHidden           *bool                                   `json:"is_hidden,omitempty"`
		Link               *QueryTableColumnLink                   `json:"link,omitempty"`
		NumberFormat       *WidgetNumberFormat                     `json:"number_format,omitempty"`
		TextAlign          *WidgetTextAlign                        `json:"text_align,omitempty"`
		TextFormats        []TextFormattingRule                    `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}

	hasInvalidField := false
	if all.CellDisplayMode != nil && !all.CellDisplayMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CellDisplayMode = all.CellDisplayMode
	}
	o.Column = *all.Column
	o.ConditionalFormats = all.ConditionalFormats
	o.DateFormat = all.DateFormat
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}
	o.IsHidden = all.IsHidden
	if all.Link != nil && all.Link.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Link = all.Link
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat
	if all.TextAlign != nil && !all.TextAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TextAlign = all.TextAlign
	}
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
