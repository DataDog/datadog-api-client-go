// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarShiftSplitConfig Test-only generated schema for the dashboard versioning proof of concept.
type CalendarShiftSplitConfig struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Shifts []CalendarShift `json:"shifts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCalendarShiftSplitConfig instantiates a new CalendarShiftSplitConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarShiftSplitConfig(shifts []CalendarShift) *CalendarShiftSplitConfig {
	this := CalendarShiftSplitConfig{}
	this.Shifts = shifts
	return &this
}

// NewCalendarShiftSplitConfigWithDefaults instantiates a new CalendarShiftSplitConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarShiftSplitConfigWithDefaults() *CalendarShiftSplitConfig {
	this := CalendarShiftSplitConfig{}
	return &this
}

// GetShifts returns the Shifts field value.
func (o *CalendarShiftSplitConfig) GetShifts() []CalendarShift {
	if o == nil {
		var ret []CalendarShift
		return ret
	}
	return o.Shifts
}

// GetShiftsOk returns a tuple with the Shifts field value
// and a boolean to check if the value has been set.
func (o *CalendarShiftSplitConfig) GetShiftsOk() (*[]CalendarShift, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shifts, true
}

// SetShifts sets field value.
func (o *CalendarShiftSplitConfig) SetShifts(v []CalendarShift) {
	o.Shifts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarShiftSplitConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["shifts"] = o.Shifts
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalendarShiftSplitConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Shifts *[]CalendarShift `json:"shifts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Shifts == nil {
		return fmt.Errorf("required field shifts missing")
	}
	o.Shifts = *all.Shifts

	return nil
}
