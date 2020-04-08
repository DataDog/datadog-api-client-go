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

// DistributionWidgetRequest TODO.
type DistributionWidgetRequest struct {
	ApmQuery     *LogQueryDefinition     `json:"apm_query,omitempty"`
	EventQuery   *EventQueryDefinition   `json:"event_query,omitempty"`
	LogQuery     *LogQueryDefinition     `json:"log_query,omitempty"`
	NetworkQuery *LogQueryDefinition     `json:"network_query,omitempty"`
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// TODO.
	Q        *string             `json:"q,omitempty"`
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
	Style    *WidgetStyle        `json:"style,omitempty"`
}

// NewDistributionWidgetRequest instantiates a new DistributionWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionWidgetRequest() *DistributionWidgetRequest {
	this := DistributionWidgetRequest{}
	return &this
}

// NewDistributionWidgetRequestWithDefaults instantiates a new DistributionWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionWidgetRequestWithDefaults() *DistributionWidgetRequest {
	this := DistributionWidgetRequest{}
	return &this
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasApmQuery() bool {
	if o != nil && o.ApmQuery != nil {
		return true
	}

	return false
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *DistributionWidgetRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetEventQuery() EventQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret EventQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasEventQuery() bool {
	if o != nil && o.EventQuery != nil {
		return true
	}

	return false
}

// SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.
func (o *DistributionWidgetRequest) SetEventQuery(v EventQueryDefinition) {
	o.EventQuery = &v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasLogQuery() bool {
	if o != nil && o.LogQuery != nil {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *DistributionWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasNetworkQuery() bool {
	if o != nil && o.NetworkQuery != nil {
		return true
	}

	return false
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *DistributionWidgetRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasProcessQuery() bool {
	if o != nil && o.ProcessQuery != nil {
		return true
	}

	return false
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *DistributionWidgetRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *DistributionWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasRumQuery() bool {
	if o != nil && o.RumQuery != nil {
		return true
	}

	return false
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *DistributionWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionWidgetRequest) GetStyle() WidgetStyle {
	if o == nil || o.Style == nil {
		var ret WidgetStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetRequest) GetStyleOk() (*WidgetStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionWidgetRequest) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given WidgetStyle and assigns it to the Style field.
func (o *DistributionWidgetRequest) SetStyle(v WidgetStyle) {
	o.Style = &v
}

func (o DistributionWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.NetworkQuery != nil {
		toSerialize["network_query"] = o.NetworkQuery
	}
	if o.ProcessQuery != nil {
		toSerialize["process_query"] = o.ProcessQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.RumQuery != nil {
		toSerialize["rum_query"] = o.RumQuery
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionWidgetRequest struct {
	value *DistributionWidgetRequest
	isSet bool
}

func (v NullableDistributionWidgetRequest) Get() *DistributionWidgetRequest {
	return v.value
}

func (v *NullableDistributionWidgetRequest) Set(val *DistributionWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionWidgetRequest(val *DistributionWidgetRequest) *NullableDistributionWidgetRequest {
	return &NullableDistributionWidgetRequest{value: val, isSet: true}
}

func (v NullableDistributionWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
