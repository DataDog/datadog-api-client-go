// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridWidgetFormula A formula for a heatgrid widget request.
type HeatgridWidgetFormula struct {
	// Expression alias.
	Alias *string `json:"alias,omitempty"`
	// Conditional formatting rules. These rules do not affect heatgrid rendering.
	// Use the widget-level `color` configuration to control cell colors.
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	// String expression built from queries, formulas, and functions.
	Formula string `json:"formula"`
	// Options for limiting results returned.
	Limit *WidgetFormulaLimit `json:"limit,omitempty"`
	// Number format options for the widget.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Styling options for widget formulas.
	Style *WidgetFormulaStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridWidgetFormula instantiates a new HeatgridWidgetFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridWidgetFormula(formula string) *HeatgridWidgetFormula {
	this := HeatgridWidgetFormula{}
	this.Formula = formula
	return &this
}

// NewHeatgridWidgetFormulaWithDefaults instantiates a new HeatgridWidgetFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridWidgetFormulaWithDefaults() *HeatgridWidgetFormula {
	this := HeatgridWidgetFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *HeatgridWidgetFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *HeatgridWidgetFormula) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *HeatgridWidgetFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *HeatgridWidgetFormula) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *HeatgridWidgetFormula) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *HeatgridWidgetFormula) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = v
}

// GetFormula returns the Formula field value.
func (o *HeatgridWidgetFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *HeatgridWidgetFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *HeatgridWidgetFormula) GetLimit() WidgetFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret WidgetFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetLimitOk() (*WidgetFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *HeatgridWidgetFormula) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given WidgetFormulaLimit and assigns it to the Limit field.
func (o *HeatgridWidgetFormula) SetLimit(v WidgetFormulaLimit) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *HeatgridWidgetFormula) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *HeatgridWidgetFormula) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *HeatgridWidgetFormula) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HeatgridWidgetFormula) GetStyle() WidgetFormulaStyle {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridWidgetFormula) GetStyleOk() (*WidgetFormulaStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HeatgridWidgetFormula) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle and assigns it to the Style field.
func (o *HeatgridWidgetFormula) SetStyle(v WidgetFormulaStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridWidgetFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
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
func (o *HeatgridWidgetFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias              *string                   `json:"alias,omitempty"`
		ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
		Formula            *string                   `json:"formula"`
		Limit              *WidgetFormulaLimit       `json:"limit,omitempty"`
		NumberFormat       *WidgetNumberFormat       `json:"number_format,omitempty"`
		Style              *WidgetFormulaStyle       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	o.ConditionalFormats = all.ConditionalFormats
	o.Formula = *all.Formula
	if all.Limit != nil && all.Limit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limit = all.Limit
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
