// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartValueDimension Test-only generated schema for the dashboard versioning proof of concept.
type RunChartValueDimension struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension RunChartValueDimensionDimension `json:"dimension"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RunChartLineStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunChartValueDimension instantiates a new RunChartValueDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunChartValueDimension(column string, dimension RunChartValueDimensionDimension) *RunChartValueDimension {
	this := RunChartValueDimension{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewRunChartValueDimensionWithDefaults instantiates a new RunChartValueDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunChartValueDimensionWithDefaults() *RunChartValueDimension {
	this := RunChartValueDimension{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *RunChartValueDimension) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartValueDimension) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *RunChartValueDimension) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *RunChartValueDimension) SetAlias(v string) {
	o.Alias = &v
}

// GetColumn returns the Column field value.
func (o *RunChartValueDimension) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *RunChartValueDimension) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *RunChartValueDimension) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *RunChartValueDimension) GetDimension() RunChartValueDimensionDimension {
	if o == nil {
		var ret RunChartValueDimensionDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *RunChartValueDimension) GetDimensionOk() (*RunChartValueDimensionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *RunChartValueDimension) SetDimension(v RunChartValueDimensionDimension) {
	o.Dimension = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *RunChartValueDimension) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartValueDimension) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *RunChartValueDimension) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *RunChartValueDimension) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *RunChartValueDimension) GetStyle() RunChartLineStyle {
	if o == nil || o.Style == nil {
		var ret RunChartLineStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartValueDimension) GetStyleOk() (*RunChartLineStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *RunChartValueDimension) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RunChartLineStyle and assigns it to the Style field.
func (o *RunChartValueDimension) SetStyle(v RunChartLineStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunChartValueDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunChartValueDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias        *string                          `json:"alias,omitempty"`
		Column       *string                          `json:"column"`
		Dimension    *RunChartValueDimensionDimension `json:"dimension"`
		NumberFormat *WidgetNumberFormat              `json:"number_format,omitempty"`
		Style        *RunChartLineStyle               `json:"style,omitempty"`
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
	o.Alias = all.Alias
	o.Column = *all.Column
	if !all.Dimension.IsValid() {
		hasInvalidField = true
	} else {
		o.Dimension = *all.Dimension
	}
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
