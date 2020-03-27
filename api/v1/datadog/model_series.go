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

// Series A metric to submit to Datadog. See [Datadog metrics](https://docs.datadoghq.com/developers/metrics/#custom-metrics-properties).
type Series struct {
	// The name of the host that produced the metric.
	Host *string `json:"host,omitempty"`
	// If the type of the metric is rate or count, define the corresponding interval.
	Interval NullableInt64 `json:"interval,omitempty"`
	// The name of the timeseries.
	Metric string      `json:"metric"`
	Points [][]float64 `json:"points"`
	// A list of tags associated with the metric.
	Tags *[]string `json:"tags,omitempty"`
	Type *string   `json:"type,omitempty"`
}

// NewSeries instantiates a new Series object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeries(metric string, points [][]float64) *Series {
	this := Series{}
	this.Metric = metric
	this.Points = points
	var type_ string = "gauge"
	this.Type = &type_
	return &this
}

// NewSeriesWithDefaults instantiates a new Series object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesWithDefaults() *Series {
	this := Series{}
	var type_ string = "gauge"
	this.Type = &type_
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Series) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Series) GetHostOk() (string, bool) {
	if o == nil || o.Host == nil {
		var ret string
		return ret, false
	}
	return *o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Series) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Series) SetHost(v string) {
	o.Host = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *Series) GetInterval() NullableInt64 {
	if o == nil {
		var ret NullableInt64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Series) GetIntervalOk() (NullableInt64, bool) {
	if o == nil {
		var ret NullableInt64
		return ret, false
	}
	return o.Interval, o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Series) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt64 and assigns it to the Interval field.
func (o *Series) SetInterval(v NullableInt64) {
	o.Interval = v
}

// GetMetric returns the Metric field value
func (o *Series) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// SetMetric sets field value
func (o *Series) SetMetric(v string) {
	o.Metric = v
}

// GetPoints returns the Points field value
func (o *Series) GetPoints() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Points
}

// SetPoints sets field value
func (o *Series) SetPoints(v [][]float64) {
	o.Points = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Series) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Series) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Series) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Series) SetTags(v []string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Series) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Series) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Series) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Series) SetType(v string) {
	o.Type = &v
}

func (o Series) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["points"] = o.Points
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSeries struct {
	value *Series
	isSet bool
}

func (v NullableSeries) Get() *Series {
	return v.value
}

func (v NullableSeries) Set(val *Series) {
	v.value = val
	v.isSet = true
}

func (v NullableSeries) IsSet() bool {
	return v.isSet
}

func (v NullableSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeries(val *Series) *NullableSeries {
	return &NullableSeries{value: val, isSet: true}
}

func (v NullableSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
