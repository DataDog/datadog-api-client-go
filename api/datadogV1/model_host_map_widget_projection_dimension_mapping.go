// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetProjectionDimensionMapping Maps a dataset column to a host map visual dimension.
type HostMapWidgetProjectionDimensionMapping struct {
	// Alias used to label the column instead of its name.
	Alias *string `json:"alias,omitempty"`
	// Source column name from the dataset.
	Column string `json:"column"`
	// Visual dimension for the host map widget. Used both by infrastructure-backed formulas and by DDSQL projection columns; `group` is only meaningful for DDSQL projection columns, where repeated entries define the grouping hierarchy.
	Dimension HostMapWidgetDimension `json:"dimension"`
	// Number format options for the widget.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetProjectionDimensionMapping instantiates a new HostMapWidgetProjectionDimensionMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetProjectionDimensionMapping(column string, dimension HostMapWidgetDimension) *HostMapWidgetProjectionDimensionMapping {
	this := HostMapWidgetProjectionDimensionMapping{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewHostMapWidgetProjectionDimensionMappingWithDefaults instantiates a new HostMapWidgetProjectionDimensionMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetProjectionDimensionMappingWithDefaults() *HostMapWidgetProjectionDimensionMapping {
	this := HostMapWidgetProjectionDimensionMapping{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *HostMapWidgetProjectionDimensionMapping) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjectionDimensionMapping) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *HostMapWidgetProjectionDimensionMapping) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *HostMapWidgetProjectionDimensionMapping) SetAlias(v string) {
	o.Alias = &v
}

// GetColumn returns the Column field value.
func (o *HostMapWidgetProjectionDimensionMapping) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjectionDimensionMapping) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *HostMapWidgetProjectionDimensionMapping) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *HostMapWidgetProjectionDimensionMapping) GetDimension() HostMapWidgetDimension {
	if o == nil {
		var ret HostMapWidgetDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjectionDimensionMapping) GetDimensionOk() (*HostMapWidgetDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *HostMapWidgetProjectionDimensionMapping) SetDimension(v HostMapWidgetDimension) {
	o.Dimension = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *HostMapWidgetProjectionDimensionMapping) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjectionDimensionMapping) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *HostMapWidgetProjectionDimensionMapping) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *HostMapWidgetProjectionDimensionMapping) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetProjectionDimensionMapping) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetProjectionDimensionMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias        *string                 `json:"alias,omitempty"`
		Column       *string                 `json:"column"`
		Dimension    *HostMapWidgetDimension `json:"dimension"`
		NumberFormat *WidgetNumberFormat     `json:"number_format,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "column", "dimension", "number_format"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
