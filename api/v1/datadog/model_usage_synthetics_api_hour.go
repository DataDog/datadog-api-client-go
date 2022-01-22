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

// UsageSyntheticsAPIHour Number of Synthetics API tests run for each hour for a given organization.
type UsageSyntheticsAPIHour struct {
	// Contains the number of Synthetics API tests run.
	CheckCallsCount *int64 `json:"check_calls_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewUsageSyntheticsAPIHour instantiates a new UsageSyntheticsAPIHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSyntheticsAPIHour() *UsageSyntheticsAPIHour {
	this := UsageSyntheticsAPIHour{}
	return &this
}

// NewUsageSyntheticsAPIHourWithDefaults instantiates a new UsageSyntheticsAPIHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSyntheticsAPIHourWithDefaults() *UsageSyntheticsAPIHour {
	this := UsageSyntheticsAPIHour{}
	return &this
}

// GetCheckCallsCount returns the CheckCallsCount field value if set, zero value otherwise.
func (o *UsageSyntheticsAPIHour) GetCheckCallsCount() int64 {
	if o == nil || o.CheckCallsCount == nil {
		var ret int64
		return ret
	}
	return *o.CheckCallsCount
}

// GetCheckCallsCountOk returns a tuple with the CheckCallsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsAPIHour) GetCheckCallsCountOk() (*int64, bool) {
	if o == nil || o.CheckCallsCount == nil {
		return nil, false
	}
	return o.CheckCallsCount, true
}

// HasCheckCallsCount returns a boolean if a field has been set.
func (o *UsageSyntheticsAPIHour) HasCheckCallsCount() bool {
	if o != nil && o.CheckCallsCount != nil {
		return true
	}

	return false
}

// SetCheckCallsCount gets a reference to the given int64 and assigns it to the CheckCallsCount field.
func (o *UsageSyntheticsAPIHour) SetCheckCallsCount(v int64) {
	o.CheckCallsCount = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageSyntheticsAPIHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsAPIHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageSyntheticsAPIHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageSyntheticsAPIHour) SetHour(v time.Time) {
	o.Hour = &v
}

func (o UsageSyntheticsAPIHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CheckCallsCount != nil {
		toSerialize["check_calls_count"] = o.CheckCallsCount
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	return json.Marshal(toSerialize)
}

func (o *UsageSyntheticsAPIHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CheckCallsCount *int64     `json:"check_calls_count,omitempty"`
		Hour            *time.Time `json:"hour,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.CheckCallsCount = all.CheckCallsCount
	o.Hour = all.Hour
	return nil
}
