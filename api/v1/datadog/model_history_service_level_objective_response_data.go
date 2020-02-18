/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// HistoryServiceLevelObjectiveResponseData struct for HistoryServiceLevelObjectiveResponseData
type HistoryServiceLevelObjectiveResponseData struct {
	// the `from` timestamp in epoch seconds
	FromTs  *int64                               `json:"from_ts,omitempty"`
	Groups  *HistoryServiceLevelObjectiveGroups  `json:"groups,omitempty"`
	Overall *HistoryServiceLevelObjectiveOverall `json:"overall,omitempty"`
	Series  *HistoryServiceLevelObjectiveMetrics `json:"series,omitempty"`
	// mapping of string timeframe to the SLO threshold.
	Thresholds *map[string]SloThreshold `json:"thresholds,omitempty"`
	// the `to` timestamp in epoch seconds
	ToTs *int64 `json:"to_ts,omitempty"`
}

// NewHistoryServiceLevelObjectiveResponseData instantiates a new HistoryServiceLevelObjectiveResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryServiceLevelObjectiveResponseData() *HistoryServiceLevelObjectiveResponseData {
	this := HistoryServiceLevelObjectiveResponseData{}
	return &this
}

// NewHistoryServiceLevelObjectiveResponseDataWithDefaults instantiates a new HistoryServiceLevelObjectiveResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryServiceLevelObjectiveResponseDataWithDefaults() *HistoryServiceLevelObjectiveResponseData {
	this := HistoryServiceLevelObjectiveResponseData{}
	return &this
}

// GetFromTs returns the FromTs field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetFromTs() int64 {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret
	}
	return *o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetFromTsOk() (int64, bool) {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret, false
	}
	return *o.FromTs, true
}

// HasFromTs returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasFromTs() bool {
	if o != nil && o.FromTs != nil {
		return true
	}

	return false
}

// SetFromTs gets a reference to the given int64 and assigns it to the FromTs field.
func (o *HistoryServiceLevelObjectiveResponseData) SetFromTs(v int64) {
	o.FromTs = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetGroups() HistoryServiceLevelObjectiveGroups {
	if o == nil || o.Groups == nil {
		var ret HistoryServiceLevelObjectiveGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetGroupsOk() (HistoryServiceLevelObjectiveGroups, bool) {
	if o == nil || o.Groups == nil {
		var ret HistoryServiceLevelObjectiveGroups
		return ret, false
	}
	return *o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given HistoryServiceLevelObjectiveGroups and assigns it to the Groups field.
func (o *HistoryServiceLevelObjectiveResponseData) SetGroups(v HistoryServiceLevelObjectiveGroups) {
	o.Groups = &v
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetOverall() HistoryServiceLevelObjectiveOverall {
	if o == nil || o.Overall == nil {
		var ret HistoryServiceLevelObjectiveOverall
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetOverallOk() (HistoryServiceLevelObjectiveOverall, bool) {
	if o == nil || o.Overall == nil {
		var ret HistoryServiceLevelObjectiveOverall
		return ret, false
	}
	return *o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasOverall() bool {
	if o != nil && o.Overall != nil {
		return true
	}

	return false
}

// SetOverall gets a reference to the given HistoryServiceLevelObjectiveOverall and assigns it to the Overall field.
func (o *HistoryServiceLevelObjectiveResponseData) SetOverall(v HistoryServiceLevelObjectiveOverall) {
	o.Overall = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetSeries() HistoryServiceLevelObjectiveMetrics {
	if o == nil || o.Series == nil {
		var ret HistoryServiceLevelObjectiveMetrics
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetSeriesOk() (HistoryServiceLevelObjectiveMetrics, bool) {
	if o == nil || o.Series == nil {
		var ret HistoryServiceLevelObjectiveMetrics
		return ret, false
	}
	return *o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given HistoryServiceLevelObjectiveMetrics and assigns it to the Series field.
func (o *HistoryServiceLevelObjectiveResponseData) SetSeries(v HistoryServiceLevelObjectiveMetrics) {
	o.Series = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetThresholds() map[string]SloThreshold {
	if o == nil || o.Thresholds == nil {
		var ret map[string]SloThreshold
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetThresholdsOk() (map[string]SloThreshold, bool) {
	if o == nil || o.Thresholds == nil {
		var ret map[string]SloThreshold
		return ret, false
	}
	return *o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasThresholds() bool {
	if o != nil && o.Thresholds != nil {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given map[string]SloThreshold and assigns it to the Thresholds field.
func (o *HistoryServiceLevelObjectiveResponseData) SetThresholds(v map[string]SloThreshold) {
	o.Thresholds = &v
}

// GetToTs returns the ToTs field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseData) GetToTs() int64 {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret
	}
	return *o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseData) GetToTsOk() (int64, bool) {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret, false
	}
	return *o.ToTs, true
}

// HasToTs returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseData) HasToTs() bool {
	if o != nil && o.ToTs != nil {
		return true
	}

	return false
}

// SetToTs gets a reference to the given int64 and assigns it to the ToTs field.
func (o *HistoryServiceLevelObjectiveResponseData) SetToTs(v int64) {
	o.ToTs = &v
}

type NullableHistoryServiceLevelObjectiveResponseData struct {
	Value        HistoryServiceLevelObjectiveResponseData
	ExplicitNull bool
}

func (v NullableHistoryServiceLevelObjectiveResponseData) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHistoryServiceLevelObjectiveResponseData) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
