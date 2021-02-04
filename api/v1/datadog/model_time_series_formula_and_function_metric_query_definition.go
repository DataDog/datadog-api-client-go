/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// TimeSeriesFormulaAndFunctionMetricQueryDefinition A timeseries formula and functions metrics query.
type TimeSeriesFormulaAndFunctionMetricQueryDefinition struct {
	Aggregator *FormulaAndFunctionMetricAggregation `json:"aggregator,omitempty"`
	DataSource FormulaAndFunctionMetricDataSource   `json:"data_source"`
	// Name of the query for use in formulas.
	Name *string `json:"name,omitempty"`
	// Metrics query definition.
	Query string `json:"query"`
}

// NewTimeSeriesFormulaAndFunctionMetricQueryDefinition instantiates a new TimeSeriesFormulaAndFunctionMetricQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesFormulaAndFunctionMetricQueryDefinition(dataSource FormulaAndFunctionMetricDataSource, query string) *TimeSeriesFormulaAndFunctionMetricQueryDefinition {
	this := TimeSeriesFormulaAndFunctionMetricQueryDefinition{}
	this.DataSource = dataSource
	this.Query = query
	return &this
}

// NewTimeSeriesFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new TimeSeriesFormulaAndFunctionMetricQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesFormulaAndFunctionMetricQueryDefinitionWithDefaults() *TimeSeriesFormulaAndFunctionMetricQueryDefinition {
	this := TimeSeriesFormulaAndFunctionMetricQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation {
	if o == nil || o.Aggregator == nil {
		var ret FormulaAndFunctionMetricAggregation
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) HasAggregator() bool {
	if o != nil && o.Aggregator != nil {
		return true
	}

	return false
}

// SetAggregator gets a reference to the given FormulaAndFunctionMetricAggregation and assigns it to the Aggregator field.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation) {
	o.Aggregator = &v
}

// GetDataSource returns the DataSource field value
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetDataSource() FormulaAndFunctionMetricDataSource {
	if o == nil {
		var ret FormulaAndFunctionMetricDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetDataSource(v FormulaAndFunctionMetricDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TimeSeriesFormulaAndFunctionMetricQueryDefinition) SetQuery(v string) {
	o.Query = v
}

func (o TimeSeriesFormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if true {
		toSerialize["data_source"] = o.DataSource
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition struct {
	value *TimeSeriesFormulaAndFunctionMetricQueryDefinition
	isSet bool
}

func (v NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) Get() *TimeSeriesFormulaAndFunctionMetricQueryDefinition {
	return v.value
}

func (v *NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) Set(val *TimeSeriesFormulaAndFunctionMetricQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesFormulaAndFunctionMetricQueryDefinition(val *TimeSeriesFormulaAndFunctionMetricQueryDefinition) *NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition {
	return &NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition{value: val, isSet: true}
}

func (v NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesFormulaAndFunctionMetricQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
