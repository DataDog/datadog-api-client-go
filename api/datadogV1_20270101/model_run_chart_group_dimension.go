// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartGroupDimension Test-only generated schema for the dashboard versioning proof of concept.
type RunChartGroupDimension struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension RunChartGroupDimensionDimension `json:"dimension"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunChartGroupDimension instantiates a new RunChartGroupDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunChartGroupDimension(column string, dimension RunChartGroupDimensionDimension) *RunChartGroupDimension {
	this := RunChartGroupDimension{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewRunChartGroupDimensionWithDefaults instantiates a new RunChartGroupDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunChartGroupDimensionWithDefaults() *RunChartGroupDimension {
	this := RunChartGroupDimension{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *RunChartGroupDimension) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunChartGroupDimension) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *RunChartGroupDimension) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *RunChartGroupDimension) SetAlias(v string) {
	o.Alias = &v
}

// GetColumn returns the Column field value.
func (o *RunChartGroupDimension) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *RunChartGroupDimension) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *RunChartGroupDimension) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *RunChartGroupDimension) GetDimension() RunChartGroupDimensionDimension {
	if o == nil {
		var ret RunChartGroupDimensionDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *RunChartGroupDimension) GetDimensionOk() (*RunChartGroupDimensionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *RunChartGroupDimension) SetDimension(v RunChartGroupDimensionDimension) {
	o.Dimension = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunChartGroupDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunChartGroupDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias     *string                          `json:"alias,omitempty"`
		Column    *string                          `json:"column"`
		Dimension *RunChartGroupDimensionDimension `json:"dimension"`
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
