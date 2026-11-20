// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormula Test-only generated schema for the dashboard versioning proof of concept.
type WidgetFormula struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *WidgetFormulaLimit `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *WidgetFormulaStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetFormula instantiates a new WidgetFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFormula(formula string) *WidgetFormula {
	this := WidgetFormula{}
	this.Formula = formula
	return &this
}

// NewWidgetFormulaWithDefaults instantiates a new WidgetFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFormulaWithDefaults() *WidgetFormula {
	this := WidgetFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *WidgetFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *WidgetFormula) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *WidgetFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetFormula returns the Formula field value.
func (o *WidgetFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *WidgetFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *WidgetFormula) GetLimit() WidgetFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret WidgetFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetLimitOk() (*WidgetFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *WidgetFormula) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given WidgetFormulaLimit and assigns it to the Limit field.
func (o *WidgetFormula) SetLimit(v WidgetFormulaLimit) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *WidgetFormula) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *WidgetFormula) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *WidgetFormula) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *WidgetFormula) GetStyle() WidgetFormulaStyle {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormula) GetStyleOk() (*WidgetFormulaStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *WidgetFormula) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle and assigns it to the Style field.
func (o *WidgetFormula) SetStyle(v WidgetFormulaStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
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
func (o *WidgetFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias        *string             `json:"alias,omitempty"`
		Formula      *string             `json:"formula"`
		Limit        *WidgetFormulaLimit `json:"limit,omitempty"`
		NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
		Style        *WidgetFormulaStyle `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
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
