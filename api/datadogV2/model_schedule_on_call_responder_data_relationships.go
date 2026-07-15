// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallResponderDataRelationships Relationships for a single position's (previous, current, or next) responder group.
type ScheduleOnCallResponderDataRelationships struct {
	// Defines the list of shifts satisfying this responder group's position. Multiple shifts occur when a schedule has multiple concurrent on-call responders at that position.
	Shifts *ScheduleOnCallResponderDataRelationshipsShifts `json:"shifts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallResponderDataRelationships instantiates a new ScheduleOnCallResponderDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallResponderDataRelationships() *ScheduleOnCallResponderDataRelationships {
	this := ScheduleOnCallResponderDataRelationships{}
	return &this
}

// NewScheduleOnCallResponderDataRelationshipsWithDefaults instantiates a new ScheduleOnCallResponderDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallResponderDataRelationshipsWithDefaults() *ScheduleOnCallResponderDataRelationships {
	this := ScheduleOnCallResponderDataRelationships{}
	return &this
}

// GetShifts returns the Shifts field value if set, zero value otherwise.
func (o *ScheduleOnCallResponderDataRelationships) GetShifts() ScheduleOnCallResponderDataRelationshipsShifts {
	if o == nil || o.Shifts == nil {
		var ret ScheduleOnCallResponderDataRelationshipsShifts
		return ret
	}
	return *o.Shifts
}

// GetShiftsOk returns a tuple with the Shifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderDataRelationships) GetShiftsOk() (*ScheduleOnCallResponderDataRelationshipsShifts, bool) {
	if o == nil || o.Shifts == nil {
		return nil, false
	}
	return o.Shifts, true
}

// HasShifts returns a boolean if a field has been set.
func (o *ScheduleOnCallResponderDataRelationships) HasShifts() bool {
	return o != nil && o.Shifts != nil
}

// SetShifts gets a reference to the given ScheduleOnCallResponderDataRelationshipsShifts and assigns it to the Shifts field.
func (o *ScheduleOnCallResponderDataRelationships) SetShifts(v ScheduleOnCallResponderDataRelationshipsShifts) {
	o.Shifts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallResponderDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Shifts != nil {
		toSerialize["shifts"] = o.Shifts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallResponderDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Shifts *ScheduleOnCallResponderDataRelationshipsShifts `json:"shifts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"shifts"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Shifts != nil && all.Shifts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Shifts = all.Shifts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
