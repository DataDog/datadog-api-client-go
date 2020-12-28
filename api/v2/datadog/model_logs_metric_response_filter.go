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

// LogsMetricResponseFilter The log-based metric filter. Logs matching this filter will be aggregated in this metric.
type LogsMetricResponseFilter struct {
	// The search query - following the log search syntax.
	Query *string `json:"query,omitempty"`
}

// NewLogsMetricResponseFilter instantiates a new LogsMetricResponseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsMetricResponseFilter() *LogsMetricResponseFilter {
	this := LogsMetricResponseFilter{}
	return &this
}

// NewLogsMetricResponseFilterWithDefaults instantiates a new LogsMetricResponseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsMetricResponseFilterWithDefaults() *LogsMetricResponseFilter {
	this := LogsMetricResponseFilter{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogsMetricResponseFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponseFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogsMetricResponseFilter) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogsMetricResponseFilter) SetQuery(v string) {
	o.Query = &v
}

func (o LogsMetricResponseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableLogsMetricResponseFilter struct {
	value *LogsMetricResponseFilter
	isSet bool
}

func (v NullableLogsMetricResponseFilter) Get() *LogsMetricResponseFilter {
	return v.value
}

func (v *NullableLogsMetricResponseFilter) Set(val *LogsMetricResponseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMetricResponseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMetricResponseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMetricResponseFilter(val *LogsMetricResponseFilter) *NullableLogsMetricResponseFilter {
	return &NullableLogsMetricResponseFilter{value: val, isSet: true}
}

func (v NullableLogsMetricResponseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMetricResponseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
