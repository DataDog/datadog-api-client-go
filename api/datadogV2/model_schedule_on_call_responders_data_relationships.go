// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersDataRelationships Relationships for a schedule's on-call responders lookup, including the schedule and its responder groups.
type ScheduleOnCallRespondersDataRelationships struct {
	// Defines the list of per-position (previous, current, next) responder groups for the schedule.
	Responders *ScheduleOnCallRespondersDataRelationshipsResponders `json:"responders,omitempty"`
	// Defines the relationship to the schedule this on-call responders lookup was performed for.
	Schedule *ScheduleOnCallRespondersDataRelationshipsSchedule `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallRespondersDataRelationships instantiates a new ScheduleOnCallRespondersDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallRespondersDataRelationships() *ScheduleOnCallRespondersDataRelationships {
	this := ScheduleOnCallRespondersDataRelationships{}
	return &this
}

// NewScheduleOnCallRespondersDataRelationshipsWithDefaults instantiates a new ScheduleOnCallRespondersDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallRespondersDataRelationshipsWithDefaults() *ScheduleOnCallRespondersDataRelationships {
	this := ScheduleOnCallRespondersDataRelationships{}
	return &this
}

// GetResponders returns the Responders field value if set, zero value otherwise.
func (o *ScheduleOnCallRespondersDataRelationships) GetResponders() ScheduleOnCallRespondersDataRelationshipsResponders {
	if o == nil || o.Responders == nil {
		var ret ScheduleOnCallRespondersDataRelationshipsResponders
		return ret
	}
	return *o.Responders
}

// GetRespondersOk returns a tuple with the Responders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallRespondersDataRelationships) GetRespondersOk() (*ScheduleOnCallRespondersDataRelationshipsResponders, bool) {
	if o == nil || o.Responders == nil {
		return nil, false
	}
	return o.Responders, true
}

// HasResponders returns a boolean if a field has been set.
func (o *ScheduleOnCallRespondersDataRelationships) HasResponders() bool {
	return o != nil && o.Responders != nil
}

// SetResponders gets a reference to the given ScheduleOnCallRespondersDataRelationshipsResponders and assigns it to the Responders field.
func (o *ScheduleOnCallRespondersDataRelationships) SetResponders(v ScheduleOnCallRespondersDataRelationshipsResponders) {
	o.Responders = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ScheduleOnCallRespondersDataRelationships) GetSchedule() ScheduleOnCallRespondersDataRelationshipsSchedule {
	if o == nil || o.Schedule == nil {
		var ret ScheduleOnCallRespondersDataRelationshipsSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallRespondersDataRelationships) GetScheduleOk() (*ScheduleOnCallRespondersDataRelationshipsSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ScheduleOnCallRespondersDataRelationships) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given ScheduleOnCallRespondersDataRelationshipsSchedule and assigns it to the Schedule field.
func (o *ScheduleOnCallRespondersDataRelationships) SetSchedule(v ScheduleOnCallRespondersDataRelationshipsSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallRespondersDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Responders != nil {
		toSerialize["responders"] = o.Responders
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallRespondersDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Responders *ScheduleOnCallRespondersDataRelationshipsResponders `json:"responders,omitempty"`
		Schedule   *ScheduleOnCallRespondersDataRelationshipsSchedule   `json:"schedule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"responders", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Responders != nil && all.Responders.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responders = all.Responders
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
