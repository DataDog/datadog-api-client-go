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

// HistoryServiceLevelObjectiveMetricsSeries A representation of `metric` based SLO time series for the provided queries. This is the same response type from `batch_query` endpoint.
type HistoryServiceLevelObjectiveMetricsSeries struct {
	// Count of submitted metrics
	Count    int64                                             `json:"count"`
	Metadata HistoryServiceLevelObjectiveMetricsSeriesMetadata `json:"metadata"`
	// Total Sum of the query
	Sum float64 `json:"sum"`
	// The query values
	Values []float64 `json:"values"`
}

// NewHistoryServiceLevelObjectiveMetricsSeries instantiates a new HistoryServiceLevelObjectiveMetricsSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryServiceLevelObjectiveMetricsSeries(count int64, metadata HistoryServiceLevelObjectiveMetricsSeriesMetadata, sum float64, values []float64) *HistoryServiceLevelObjectiveMetricsSeries {
	this := HistoryServiceLevelObjectiveMetricsSeries{}
	this.Count = count
	this.Metadata = metadata
	this.Sum = sum
	this.Values = values
	return &this
}

// NewHistoryServiceLevelObjectiveMetricsSeriesWithDefaults instantiates a new HistoryServiceLevelObjectiveMetricsSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryServiceLevelObjectiveMetricsSeriesWithDefaults() *HistoryServiceLevelObjectiveMetricsSeries {
	this := HistoryServiceLevelObjectiveMetricsSeries{}
	return &this
}

// GetCount returns the Count field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// SetCount sets field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetCount(v int64) {
	o.Count = v
}

// GetMetadata returns the Metadata field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetMetadata() HistoryServiceLevelObjectiveMetricsSeriesMetadata {
	if o == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeriesMetadata
		return ret
	}

	return o.Metadata
}

// SetMetadata sets field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetMetadata(v HistoryServiceLevelObjectiveMetricsSeriesMetadata) {
	o.Metadata = v
}

// GetSum returns the Sum field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetSum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sum
}

// SetSum sets field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetSum(v float64) {
	o.Sum = v
}

// GetValues returns the Values field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) GetValues() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Values
}

// SetValues sets field value
func (o *HistoryServiceLevelObjectiveMetricsSeries) SetValues(v []float64) {
	o.Values = v
}

func (o HistoryServiceLevelObjectiveMetricsSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["sum"] = o.Sum
	}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryServiceLevelObjectiveMetricsSeries struct {
	value *HistoryServiceLevelObjectiveMetricsSeries
	isSet bool
}

func (v NullableHistoryServiceLevelObjectiveMetricsSeries) Get() *HistoryServiceLevelObjectiveMetricsSeries {
	return v.value
}

func (v NullableHistoryServiceLevelObjectiveMetricsSeries) Set(val *HistoryServiceLevelObjectiveMetricsSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryServiceLevelObjectiveMetricsSeries) IsSet() bool {
	return v.isSet
}

func (v NullableHistoryServiceLevelObjectiveMetricsSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryServiceLevelObjectiveMetricsSeries(val *HistoryServiceLevelObjectiveMetricsSeries) *NullableHistoryServiceLevelObjectiveMetricsSeries {
	return &NullableHistoryServiceLevelObjectiveMetricsSeries{value: val, isSet: true}
}

func (v NullableHistoryServiceLevelObjectiveMetricsSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryServiceLevelObjectiveMetricsSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
