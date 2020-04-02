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

// UsageLogsHour struct for UsageLogsHour
type UsageLogsHour struct {
	// Contains the number of billable log bytes ingested.
	BillableIngestedBytes *int64 `json:"billable_ingested_bytes,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of log events indexed.
	IndexedEventsCount *int64 `json:"indexed_events_count,omitempty"`
	// Contains the number of log bytes ingested.
	IngestedEventsBytes *int64 `json:"ingested_events_bytes,omitempty"`
}

// NewUsageLogsHour instantiates a new UsageLogsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLogsHour() *UsageLogsHour {
	this := UsageLogsHour{}
	return &this
}

// NewUsageLogsHourWithDefaults instantiates a new UsageLogsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLogsHourWithDefaults() *UsageLogsHour {
	this := UsageLogsHour{}
	return &this
}

// GetBillableIngestedBytes returns the BillableIngestedBytes field value if set, zero value otherwise.
func (o *UsageLogsHour) GetBillableIngestedBytes() int64 {
	if o == nil || o.BillableIngestedBytes == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytes
}

// GetBillableIngestedBytesOk returns a tuple with the BillableIngestedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetBillableIngestedBytesOk() (*int64, bool) {
	if o == nil || o.BillableIngestedBytes == nil {
		return nil, false
	}
	return o.BillableIngestedBytes, true
}

// HasBillableIngestedBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasBillableIngestedBytes() bool {
	if o != nil && o.BillableIngestedBytes != nil {
		return true
	}

	return false
}

// SetBillableIngestedBytes gets a reference to the given int64 and assigns it to the BillableIngestedBytes field.
func (o *UsageLogsHour) SetBillableIngestedBytes(v int64) {
	o.BillableIngestedBytes = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageLogsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIndexedEventsCount returns the IndexedEventsCount field value if set, zero value otherwise.
func (o *UsageLogsHour) GetIndexedEventsCount() int64 {
	if o == nil || o.IndexedEventsCount == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCount
}

// GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetIndexedEventsCountOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCount == nil {
		return nil, false
	}
	return o.IndexedEventsCount, true
}

// HasIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageLogsHour) HasIndexedEventsCount() bool {
	if o != nil && o.IndexedEventsCount != nil {
		return true
	}

	return false
}

// SetIndexedEventsCount gets a reference to the given int64 and assigns it to the IndexedEventsCount field.
func (o *UsageLogsHour) SetIndexedEventsCount(v int64) {
	o.IndexedEventsCount = &v
}

// GetIngestedEventsBytes returns the IngestedEventsBytes field value if set, zero value otherwise.
func (o *UsageLogsHour) GetIngestedEventsBytes() int64 {
	if o == nil || o.IngestedEventsBytes == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytes
}

// GetIngestedEventsBytesOk returns a tuple with the IngestedEventsBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetIngestedEventsBytesOk() (*int64, bool) {
	if o == nil || o.IngestedEventsBytes == nil {
		return nil, false
	}
	return o.IngestedEventsBytes, true
}

// HasIngestedEventsBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasIngestedEventsBytes() bool {
	if o != nil && o.IngestedEventsBytes != nil {
		return true
	}

	return false
}

// SetIngestedEventsBytes gets a reference to the given int64 and assigns it to the IngestedEventsBytes field.
func (o *UsageLogsHour) SetIngestedEventsBytes(v int64) {
	o.IngestedEventsBytes = &v
}

func (o UsageLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillableIngestedBytes != nil {
		toSerialize["billable_ingested_bytes"] = o.BillableIngestedBytes
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.IndexedEventsCount != nil {
		toSerialize["indexed_events_count"] = o.IndexedEventsCount
	}
	if o.IngestedEventsBytes != nil {
		toSerialize["ingested_events_bytes"] = o.IngestedEventsBytes
	}
	return json.Marshal(toSerialize)
}

type NullableUsageLogsHour struct {
	value *UsageLogsHour
	isSet bool
}

func (v NullableUsageLogsHour) Get() *UsageLogsHour {
	return v.value
}

func (v *NullableUsageLogsHour) Set(val *UsageLogsHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLogsHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLogsHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLogsHour(val *UsageLogsHour) *NullableUsageLogsHour {
	return &NullableUsageLogsHour{value: val, isSet: true}
}

func (v NullableUsageLogsHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLogsHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
