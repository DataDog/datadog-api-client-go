/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// LogsListRequestTime Timeframe to retrieve the log from.
type LogsListRequestTime struct {
	// Minimum timestamp for requested logs.
	From time.Time `json:"from"`
	// Timezone can be specified both as an offset (e.g. \"UTC+03:00\") or a regional zone (e.g. \"Europe/Paris\").
	Timezone *string `json:"timezone,omitempty"`
	// Maximum timestamp for requested logs.
	To time.Time `json:"to"`
}

// NewLogsListRequestTime instantiates a new LogsListRequestTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsListRequestTime(from time.Time, to time.Time) *LogsListRequestTime {
	this := LogsListRequestTime{}
	this.From = from
	this.To = to
	return &this
}

// NewLogsListRequestTimeWithDefaults instantiates a new LogsListRequestTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsListRequestTimeWithDefaults() *LogsListRequestTime {
	this := LogsListRequestTime{}
	return &this
}

// GetFrom returns the From field value
func (o *LogsListRequestTime) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *LogsListRequestTime) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *LogsListRequestTime) SetFrom(v time.Time) {
	o.From = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *LogsListRequestTime) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListRequestTime) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *LogsListRequestTime) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *LogsListRequestTime) SetTimezone(v string) {
	o.Timezone = &v
}

// GetTo returns the To field value
func (o *LogsListRequestTime) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *LogsListRequestTime) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *LogsListRequestTime) SetTo(v time.Time) {
	o.To = v
}

func (o LogsListRequestTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableLogsListRequestTime struct {
	value *LogsListRequestTime
	isSet bool
}

func (v NullableLogsListRequestTime) Get() *LogsListRequestTime {
	return v.value
}

func (v *NullableLogsListRequestTime) Set(val *LogsListRequestTime) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsListRequestTime) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsListRequestTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsListRequestTime(val *LogsListRequestTime) *NullableLogsListRequestTime {
	return &NullableLogsListRequestTime{value: val, isSet: true}
}

func (v NullableLogsListRequestTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsListRequestTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
