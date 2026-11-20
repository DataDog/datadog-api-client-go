// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableFormula20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type TableFormula20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode *TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101 `json:"cell_display_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayModeOptions interface{} `json:"cell_display_mode_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats *ConditionalFormats_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *TableFormulaAnyOfAnyOfItemLimit_20270101 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat_20270101 `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *WidgetFormulaStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableFormula20270101OneOf instantiates a new TableFormula20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableFormula20270101OneOf(formula string) *TableFormula20270101OneOf {
	this := TableFormula20270101OneOf{}
	this.Formula = formula
	return &this
}

// NewTableFormula20270101OneOfWithDefaults instantiates a new TableFormula20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableFormula20270101OneOfWithDefaults() *TableFormula20270101OneOf {
	this := TableFormula20270101OneOf{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *TableFormula20270101OneOf) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetCellDisplayMode() TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101 {
	if o == nil || o.CellDisplayMode == nil {
		var ret TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101
		return ret
	}
	return *o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetCellDisplayModeOk() (*TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101 and assigns it to the CellDisplayMode field.
func (o *TableFormula20270101OneOf) SetCellDisplayMode(v TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101) {
	o.CellDisplayMode = &v
}

// GetCellDisplayModeOptions returns the CellDisplayModeOptions field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetCellDisplayModeOptions() interface{} {
	if o == nil || o.CellDisplayModeOptions == nil {
		var ret interface{}
		return ret
	}
	return o.CellDisplayModeOptions
}

// GetCellDisplayModeOptionsOk returns a tuple with the CellDisplayModeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetCellDisplayModeOptionsOk() (*interface{}, bool) {
	if o == nil || o.CellDisplayModeOptions == nil {
		return nil, false
	}
	return &o.CellDisplayModeOptions, true
}

// HasCellDisplayModeOptions returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasCellDisplayModeOptions() bool {
	return o != nil && o.CellDisplayModeOptions != nil
}

// SetCellDisplayModeOptions gets a reference to the given interface{} and assigns it to the CellDisplayModeOptions field.
func (o *TableFormula20270101OneOf) SetCellDisplayModeOptions(v interface{}) {
	o.CellDisplayModeOptions = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetConditionalFormats() ConditionalFormats_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret ConditionalFormats_20270101
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetConditionalFormatsOk() (*ConditionalFormats_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given ConditionalFormats_20270101 and assigns it to the ConditionalFormats field.
func (o *TableFormula20270101OneOf) SetConditionalFormats(v ConditionalFormats_20270101) {
	o.ConditionalFormats = &v
}

// GetFormula returns the Formula field value.
func (o *TableFormula20270101OneOf) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *TableFormula20270101OneOf) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetLimit() TableFormulaAnyOfAnyOfItemLimit_20270101 {
	if o == nil || o.Limit == nil {
		var ret TableFormulaAnyOfAnyOfItemLimit_20270101
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetLimitOk() (*TableFormulaAnyOfAnyOfItemLimit_20270101, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given TableFormulaAnyOfAnyOfItemLimit_20270101 and assigns it to the Limit field.
func (o *TableFormula20270101OneOf) SetLimit(v TableFormulaAnyOfAnyOfItemLimit_20270101) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetNumberFormat() WidgetNumberFormat_20270101 {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat_20270101
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetNumberFormatOk() (*WidgetNumberFormat_20270101, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat_20270101 and assigns it to the NumberFormat field.
func (o *TableFormula20270101OneOf) SetNumberFormat(v WidgetNumberFormat_20270101) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf) GetStyle() WidgetFormulaStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf) GetStyleOk() (*WidgetFormulaStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle_20270101 and assigns it to the Style field.
func (o *TableFormula20270101OneOf) SetStyle(v WidgetFormulaStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableFormula20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.CellDisplayMode != nil {
		toSerialize["cell_display_mode"] = o.CellDisplayMode
	}
	if o.CellDisplayModeOptions != nil {
		toSerialize["cell_display_mode_options"] = o.CellDisplayModeOptions
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["formula"] = o.Formula
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableFormula20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias                  *string                                             `json:"alias,omitempty"`
		CellDisplayMode        *TableFormulaAnyOfAnyOfItemCellDisplayMode_20270101 `json:"cell_display_mode,omitempty"`
		CellDisplayModeOptions interface{}                                         `json:"cell_display_mode_options,omitempty"`
		ConditionalFormats     *ConditionalFormats_20270101                        `json:"conditional_formats,omitempty"`
		Formula                *string                                             `json:"formula"`
		Limit                  *TableFormulaAnyOfAnyOfItemLimit_20270101           `json:"limit,omitempty"`
		NumberFormat           *WidgetNumberFormat_20270101                        `json:"number_format,omitempty"`
		Style                  *WidgetFormulaStyle_20270101                        `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if all.CellDisplayMode != nil && !all.CellDisplayMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CellDisplayMode = all.CellDisplayMode
	}
	o.CellDisplayModeOptions = all.CellDisplayModeOptions
	o.ConditionalFormats = all.ConditionalFormats
	o.Formula = *all.Formula
	o.Limit = all.Limit
	o.NumberFormat = all.NumberFormat
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
