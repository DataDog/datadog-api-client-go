// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapFormula Test-only generated schema for the dashboard versioning proof of concept.
type HostmapFormula struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension HostmapDimension `json:"dimension"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formula string `json:"formula"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *HostmapFormulaLimit `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *WidgetFormulaStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHostmapFormula instantiates a new HostmapFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostmapFormula(dimension HostmapDimension, formula string) *HostmapFormula {
	this := HostmapFormula{}
	this.Dimension = dimension
	this.Formula = formula
	return &this
}

// NewHostmapFormulaWithDefaults instantiates a new HostmapFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostmapFormulaWithDefaults() *HostmapFormula {
	this := HostmapFormula{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *HostmapFormula) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *HostmapFormula) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *HostmapFormula) SetAlias(v string) {
	o.Alias = &v
}

// GetDimension returns the Dimension field value.
func (o *HostmapFormula) GetDimension() HostmapDimension {
	if o == nil {
		var ret HostmapDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetDimensionOk() (*HostmapDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *HostmapFormula) SetDimension(v HostmapDimension) {
	o.Dimension = v
}

// GetFormula returns the Formula field value.
func (o *HostmapFormula) GetFormula() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formula, true
}

// SetFormula sets field value.
func (o *HostmapFormula) SetFormula(v string) {
	o.Formula = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *HostmapFormula) GetLimit() HostmapFormulaLimit {
	if o == nil || o.Limit == nil {
		var ret HostmapFormulaLimit
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetLimitOk() (*HostmapFormulaLimit, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *HostmapFormula) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given HostmapFormulaLimit and assigns it to the Limit field.
func (o *HostmapFormula) SetLimit(v HostmapFormulaLimit) {
	o.Limit = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *HostmapFormula) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *HostmapFormula) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *HostmapFormula) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostmapFormula) GetStyle() WidgetFormulaStyle {
	if o == nil || o.Style == nil {
		var ret WidgetFormulaStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostmapFormula) GetStyleOk() (*WidgetFormulaStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostmapFormula) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetFormulaStyle and assigns it to the Style field.
func (o *HostmapFormula) SetStyle(v WidgetFormulaStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostmapFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["dimension"] = o.Dimension
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
func (o *HostmapFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias        *string              `json:"alias,omitempty"`
		Dimension    *HostmapDimension    `json:"dimension"`
		Formula      *string              `json:"formula"`
		Limit        *HostmapFormulaLimit `json:"limit,omitempty"`
		NumberFormat *WidgetNumberFormat  `json:"number_format,omitempty"`
		Style        *WidgetFormulaStyle  `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}
	if all.Formula == nil {
		return fmt.Errorf("required field formula missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}
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
