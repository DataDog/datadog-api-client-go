// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableFormula20270101OneOf4462669760 Test-only generated schema for the dashboard versioning proof of concept.
type TableFormula20270101OneOf4462669760 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101 `json:"cell_display_mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayModeOptions *TableTrendCellDisplayModeOptions_20270101 `json:"cell_display_mode_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats *ConditionalFormats_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *TableFormulaAnyOfAnyOfItemLimit4452334816_20270101 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat_20270101 `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *WidgetFormulaStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableFormula20270101OneOf4462669760 instantiates a new TableFormula20270101OneOf4462669760 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableFormula20270101OneOf4462669760(cellDisplayMode TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101, formula string) *TableFormula20270101OneOf4462669760 {
	this := TableFormula20270101OneOf4462669760{}
	this.CellDisplayMode = cellDisplayMode
	this.Formula = formula
	return &this
}

// NewTableFormula20270101OneOf4462669760WithDefaults instantiates a new TableFormula20270101OneOf4462669760 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableFormula20270101OneOf4462669760WithDefaults() *TableFormula20270101OneOf4462669760 {
	this := TableFormula20270101OneOf4462669760{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *TableFormula20270101OneOf4462669760) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value.
func (o *TableFormula20270101OneOf4462669760) GetCellDisplayMode() TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101 {
	if o == nil {
		var ret TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101
		return ret
	}
	return o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetCellDisplayModeOk() (*TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellDisplayMode, true
}

// SetCellDisplayMode sets field value.
func (o *TableFormula20270101OneOf4462669760) SetCellDisplayMode(v TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101) {
	o.CellDisplayMode = v
}

// GetCellDisplayModeOptions returns the CellDisplayModeOptions field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetCellDisplayModeOptions() TableTrendCellDisplayModeOptions_20270101 {
	if o == nil || o.CellDisplayModeOptions == nil {
		var ret TableTrendCellDisplayModeOptions_20270101
		return ret
	}
	return *o.CellDisplayModeOptions
}

// GetCellDisplayModeOptionsOk returns a tuple with the CellDisplayModeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetCellDisplayModeOptionsOk() (*TableTrendCellDisplayModeOptions_20270101, bool) {
	if o == nil || o.CellDisplayModeOptions == nil {
		return nil, false
	}
	return o.CellDisplayModeOptions, true
}

// HasCellDisplayModeOptions returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasCellDisplayModeOptions() bool {
	return o != nil && o.CellDisplayModeOptions != nil
}

// SetCellDisplayModeOptions gets a reference to the given TableTrendCellDisplayModeOptions_20270101 and assigns it to the CellDisplayModeOptions field.
func (o *TableFormula20270101OneOf4462669760) SetCellDisplayModeOptions(v TableTrendCellDisplayModeOptions_20270101) {
	o.CellDisplayModeOptions = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetConditionalFormats() ConditionalFormats_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret ConditionalFormats_20270101
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetConditionalFormatsOk() (*ConditionalFormats_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given ConditionalFormats_20270101 and assigns it to the ConditionalFormats field.
func (o *TableFormula20270101OneOf4462669760) SetConditionalFormats(v ConditionalFormats_20270101) {
	o.ConditionalFormats = &v
}

// GetFormula returns the Formula field value.
func (o *TableFormula20270101OneOf4462669760) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *TableFormula20270101OneOf4462669760) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetLimit() TableFormulaAnyOfAnyOfItemLimit4452334816_20270101 {
	if o == nil || o.Limit == nil {
		var ret TableFormulaAnyOfAnyOfItemLimit4452334816_20270101
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetLimitOk() (*TableFormulaAnyOfAnyOfItemLimit4452334816_20270101, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given TableFormulaAnyOfAnyOfItemLimit4452334816_20270101 and assigns it to the Limit field.
func (o *TableFormula20270101OneOf4462669760) SetLimit(v TableFormulaAnyOfAnyOfItemLimit4452334816_20270101) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetNumberFormat() WidgetNumberFormat_20270101 {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat_20270101
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetNumberFormatOk() (*WidgetNumberFormat_20270101, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat_20270101 and assigns it to the NumberFormat field.
func (o *TableFormula20270101OneOf4462669760) SetNumberFormat(v WidgetNumberFormat_20270101) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TableFormula20270101OneOf4462669760) GetStyle() WidgetFormulaStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableFormula20270101OneOf4462669760) GetStyleOk() (*WidgetFormulaStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TableFormula20270101OneOf4462669760) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle_20270101 and assigns it to the Style field.
func (o *TableFormula20270101OneOf4462669760) SetStyle(v WidgetFormulaStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableFormula20270101OneOf4462669760) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["cell_display_mode"] = o.CellDisplayMode
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
func (o *TableFormula20270101OneOf4462669760) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias                  *string                                                       `json:"alias,omitempty"`
		CellDisplayMode        *TableFormulaAnyOfAnyOfItemCellDisplayMode4452532448_20270101 `json:"cell_display_mode"`
		CellDisplayModeOptions *TableTrendCellDisplayModeOptions_20270101                    `json:"cell_display_mode_options,omitempty"`
		ConditionalFormats     *ConditionalFormats_20270101                                  `json:"conditional_formats,omitempty"`
		Formula                *string                                                       `json:"formula"`
		Limit                  *TableFormulaAnyOfAnyOfItemLimit4452334816_20270101           `json:"limit,omitempty"`
		NumberFormat           *WidgetNumberFormat_20270101                                  `json:"number_format,omitempty"`
		Style                  *WidgetFormulaStyle_20270101                                  `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CellDisplayMode == nil {
		return fmt.Errorf("required field cell_display_mode missing")
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.CellDisplayMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CellDisplayMode = *all.CellDisplayMode
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
