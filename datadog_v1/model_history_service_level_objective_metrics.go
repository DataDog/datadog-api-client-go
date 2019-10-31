/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"encoding/json"
)

// HistoryServiceLevelObjectiveMetrics A `metric` based SLO history response.
type HistoryServiceLevelObjectiveMetrics struct {
	Denominator *HistoryServiceLevelObjectiveMetricsSeries `json:"denominator,omitempty"`

	// The aggregated query interval for the series data. It's implicit based on the query time window.
	Interval *int32 `json:"interval,omitempty"`

	// Optional message if there are specific query issues/warnings.
	Message *string `json:"message,omitempty"`

	Numerator *HistoryServiceLevelObjectiveMetricsSeries `json:"numerator,omitempty"`

	// The combined numerator && denominator query CSV.
	Query *string `json:"query,omitempty"`

	// The series result type. This mimics `batch_query` response type
	ResType *string `json:"res_type,omitempty"`

	// The series response version type. This mimics `batch_query` response type
	RespVersion *int32 `json:"resp_version,omitempty"`
}

// GetDenominator returns the Denominator field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetDenominator() HistoryServiceLevelObjectiveMetricsSeries {
	if o == nil || o.Denominator == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret
	}
	return *o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetDenominatorOk() (HistoryServiceLevelObjectiveMetricsSeries, bool) {
	if o == nil || o.Denominator == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret, false
	}
	return *o.Denominator, true
}

// HasDenominator returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasDenominator() bool {
	if o != nil && o.Denominator != nil {
		return true
	}

	return false
}

// SetDenominator gets a reference to the given HistoryServiceLevelObjectiveMetricsSeries and assigns it to the Denominator field.
func (o *HistoryServiceLevelObjectiveMetrics) SetDenominator(v HistoryServiceLevelObjectiveMetricsSeries) {
	o.Denominator = &v
}

// GetInterval returns the Interval field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetInterval() int32 {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetIntervalOk() (int32, bool) {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret, false
	}
	return *o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *HistoryServiceLevelObjectiveMetrics) SetInterval(v int32) {
	o.Interval = &v
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HistoryServiceLevelObjectiveMetrics) SetMessage(v string) {
	o.Message = &v
}

// GetNumerator returns the Numerator field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetNumerator() HistoryServiceLevelObjectiveMetricsSeries {
	if o == nil || o.Numerator == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret
	}
	return *o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetNumeratorOk() (HistoryServiceLevelObjectiveMetricsSeries, bool) {
	if o == nil || o.Numerator == nil {
		var ret HistoryServiceLevelObjectiveMetricsSeries
		return ret, false
	}
	return *o.Numerator, true
}

// HasNumerator returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasNumerator() bool {
	if o != nil && o.Numerator != nil {
		return true
	}

	return false
}

// SetNumerator gets a reference to the given HistoryServiceLevelObjectiveMetricsSeries and assigns it to the Numerator field.
func (o *HistoryServiceLevelObjectiveMetrics) SetNumerator(v HistoryServiceLevelObjectiveMetricsSeries) {
	o.Numerator = &v
}

// GetQuery returns the Query field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetQueryOk() (string, bool) {
	if o == nil || o.Query == nil {
		var ret string
		return ret, false
	}
	return *o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *HistoryServiceLevelObjectiveMetrics) SetQuery(v string) {
	o.Query = &v
}

// GetResType returns the ResType field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetResType() string {
	if o == nil || o.ResType == nil {
		var ret string
		return ret
	}
	return *o.ResType
}

// GetResTypeOk returns a tuple with the ResType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetResTypeOk() (string, bool) {
	if o == nil || o.ResType == nil {
		var ret string
		return ret, false
	}
	return *o.ResType, true
}

// HasResType returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasResType() bool {
	if o != nil && o.ResType != nil {
		return true
	}

	return false
}

// SetResType gets a reference to the given string and assigns it to the ResType field.
func (o *HistoryServiceLevelObjectiveMetrics) SetResType(v string) {
	o.ResType = &v
}

// GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.
func (o *HistoryServiceLevelObjectiveMetrics) GetRespVersion() int32 {
	if o == nil || o.RespVersion == nil {
		var ret int32
		return ret
	}
	return *o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveMetrics) GetRespVersionOk() (int32, bool) {
	if o == nil || o.RespVersion == nil {
		var ret int32
		return ret, false
	}
	return *o.RespVersion, true
}

// HasRespVersion returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveMetrics) HasRespVersion() bool {
	if o != nil && o.RespVersion != nil {
		return true
	}

	return false
}

// SetRespVersion gets a reference to the given int32 and assigns it to the RespVersion field.
func (o *HistoryServiceLevelObjectiveMetrics) SetRespVersion(v int32) {
	o.RespVersion = &v
}

// MarshalJSON returns the JSON representation of the model.
func (o HistoryServiceLevelObjectiveMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Denominator != nil {
		toSerialize["denominator"] = o.Denominator
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Numerator != nil {
		toSerialize["numerator"] = o.Numerator
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.ResType != nil {
		toSerialize["res_type"] = o.ResType
	}
	if o.RespVersion != nil {
		toSerialize["resp_version"] = o.RespVersion
	}
	return json.Marshal(toSerialize)
}
