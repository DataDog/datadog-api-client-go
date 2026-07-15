// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersDataAttributes Attributes for a schedule's on-call responders lookup.
type ScheduleOnCallRespondersDataAttributes struct {
	// The timestamp the responders were resolved at.
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallRespondersDataAttributes instantiates a new ScheduleOnCallRespondersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallRespondersDataAttributes() *ScheduleOnCallRespondersDataAttributes {
	this := ScheduleOnCallRespondersDataAttributes{}
	return &this
}

// NewScheduleOnCallRespondersDataAttributesWithDefaults instantiates a new ScheduleOnCallRespondersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallRespondersDataAttributesWithDefaults() *ScheduleOnCallRespondersDataAttributes {
	this := ScheduleOnCallRespondersDataAttributes{}
	return &this
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *ScheduleOnCallRespondersDataAttributes) GetScheduledAt() time.Time {
	if o == nil || o.ScheduledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallRespondersDataAttributes) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledAt == nil {
		return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *ScheduleOnCallRespondersDataAttributes) HasScheduledAt() bool {
	return o != nil && o.ScheduledAt != nil
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *ScheduleOnCallRespondersDataAttributes) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallRespondersDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ScheduledAt != nil {
		if o.ScheduledAt.Nanosecond() == 0 {
			toSerialize["scheduled_at"] = o.ScheduledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["scheduled_at"] = o.ScheduledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallRespondersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"scheduled_at"})
	} else {
		return err
	}
	o.ScheduledAt = all.ScheduledAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
