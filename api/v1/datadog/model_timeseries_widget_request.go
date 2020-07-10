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

// TimeseriesWidgetRequest Updated timeseries widget.
type TimeseriesWidgetRequest struct {
	ApmQuery *LogQueryDefinition `json:"apm_query,omitempty"`
	DisplayType *WidgetDisplayType `json:"display_type,omitempty"`
	EventQuery *EventQueryDefinition `json:"event_query,omitempty"`
	LogQuery *LogQueryDefinition `json:"log_query,omitempty"`
	// Used to define expression aliases.
	Metadata *[]TimeseriesWidgetRequestMetadata `json:"metadata,omitempty"`
	NetworkQuery *LogQueryDefinition `json:"network_query,omitempty"`
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// Widget query.
	Q *string `json:"q,omitempty"`
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
	Style *WidgetRequestStyle `json:"style,omitempty"`
}

// NewTimeseriesWidgetRequest instantiates a new TimeseriesWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesWidgetRequest() *TimeseriesWidgetRequest {
	this := TimeseriesWidgetRequest{}
	return &this
}

// NewTimeseriesWidgetRequestWithDefaults instantiates a new TimeseriesWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesWidgetRequestWithDefaults() *TimeseriesWidgetRequest {
	this := TimeseriesWidgetRequest{}
	return &this
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasApmQuery() bool {
	if o != nil && o.ApmQuery != nil {
		return true
	}

	return false
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *TimeseriesWidgetRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetDisplayType() WidgetDisplayType {
	if o == nil || o.DisplayType == nil {
		var ret WidgetDisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetDisplayTypeOk() (*WidgetDisplayType, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasDisplayType() bool {
	if o != nil && o.DisplayType != nil {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given WidgetDisplayType and assigns it to the DisplayType field.
func (o *TimeseriesWidgetRequest) SetDisplayType(v WidgetDisplayType) {
	o.DisplayType = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetEventQuery() EventQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret EventQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasEventQuery() bool {
	if o != nil && o.EventQuery != nil {
		return true
	}

	return false
}

// SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.
func (o *TimeseriesWidgetRequest) SetEventQuery(v EventQueryDefinition) {
	o.EventQuery = &v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasLogQuery() bool {
	if o != nil && o.LogQuery != nil {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *TimeseriesWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetMetadata() []TimeseriesWidgetRequestMetadata {
	if o == nil || o.Metadata == nil {
		var ret []TimeseriesWidgetRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetMetadataOk() (*[]TimeseriesWidgetRequestMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []TimeseriesWidgetRequestMetadata and assigns it to the Metadata field.
func (o *TimeseriesWidgetRequest) SetMetadata(v []TimeseriesWidgetRequestMetadata) {
	o.Metadata = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasNetworkQuery() bool {
	if o != nil && o.NetworkQuery != nil {
		return true
	}

	return false
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *TimeseriesWidgetRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasProcessQuery() bool {
	if o != nil && o.ProcessQuery != nil {
		return true
	}

	return false
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *TimeseriesWidgetRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TimeseriesWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasRumQuery() bool {
	if o != nil && o.RumQuery != nil {
		return true
	}

	return false
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *TimeseriesWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequest) GetStyle() WidgetRequestStyle {
	if o == nil || o.Style == nil {
		var ret WidgetRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequest) GetStyleOk() (*WidgetRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequest) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given WidgetRequestStyle and assigns it to the Style field.
func (o *TimeseriesWidgetRequest) SetStyle(v WidgetRequestStyle) {
	o.Style = &v
}

func (o TimeseriesWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
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

type NullableTimeseriesWidgetRequest struct {
	value *TimeseriesWidgetRequest
	isSet bool
}

func (v NullableTimeseriesWidgetRequest) Get() *TimeseriesWidgetRequest {
	return v.value
}

func (v *NullableTimeseriesWidgetRequest) Set(val *TimeseriesWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesWidgetRequest(val *TimeseriesWidgetRequest) *NullableTimeseriesWidgetRequest {
	return &NullableTimeseriesWidgetRequest{value: val, isSet: true}
}

func (v NullableTimeseriesWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


