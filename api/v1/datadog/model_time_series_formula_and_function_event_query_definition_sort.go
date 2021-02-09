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

// TimeSeriesFormulaAndFunctionEventQueryDefinitionSort Options for sorting group by results.
type TimeSeriesFormulaAndFunctionEventQueryDefinitionSort struct {
	Aggregation FormulaAndFunctionEventAggregation `json:"aggregation"`
	// Metric used for sorting group by results.
	Metric *string         `json:"metric,omitempty"`
	Order  *QuerySortOrder `json:"order,omitempty"`
}

// NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSort instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSort(aggregation FormulaAndFunctionEventAggregation) *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort {
	this := TimeSeriesFormulaAndFunctionEventQueryDefinitionSort{}
	this.Aggregation = aggregation
	var order QuerySortOrder = "desc"
	this.Order = &order
	return &this
}

// NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSortWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSortWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort {
	this := TimeSeriesFormulaAndFunctionEventQueryDefinitionSort{}
	var order QuerySortOrder = "desc"
	this.Order = &order
	return &this
}

// GetAggregation returns the Aggregation field value
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetAggregation() FormulaAndFunctionEventAggregation {
	if o == nil {
		var ret FormulaAndFunctionEventAggregation
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetAggregation(v FormulaAndFunctionEventAggregation) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

func (o TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort struct {
	value *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort
	isSet bool
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) Get() *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort {
	return v.value
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) Set(val *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort(val *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort {
	return &NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort{value: val, isSet: true}
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
