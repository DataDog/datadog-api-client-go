// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLiveCallOver Information about a live call that has ended.
type OnCallPageLiveCallOver struct {
	// The outcome of a live call that has ended.
	Outcome *OnCallPageLiveCallOutcome `json:"outcome,omitempty"`
	// When support hours start, if the call was placed outside support hours.
	SupportHoursStartAt *time.Time `json:"support_hours_start_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageLiveCallOver instantiates a new OnCallPageLiveCallOver object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageLiveCallOver() *OnCallPageLiveCallOver {
	this := OnCallPageLiveCallOver{}
	return &this
}

// NewOnCallPageLiveCallOverWithDefaults instantiates a new OnCallPageLiveCallOver object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageLiveCallOverWithDefaults() *OnCallPageLiveCallOver {
	this := OnCallPageLiveCallOver{}
	return &this
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *OnCallPageLiveCallOver) GetOutcome() OnCallPageLiveCallOutcome {
	if o == nil || o.Outcome == nil {
		var ret OnCallPageLiveCallOutcome
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageLiveCallOver) GetOutcomeOk() (*OnCallPageLiveCallOutcome, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *OnCallPageLiveCallOver) HasOutcome() bool {
	return o != nil && o.Outcome != nil
}

// SetOutcome gets a reference to the given OnCallPageLiveCallOutcome and assigns it to the Outcome field.
func (o *OnCallPageLiveCallOver) SetOutcome(v OnCallPageLiveCallOutcome) {
	o.Outcome = &v
}

// GetSupportHoursStartAt returns the SupportHoursStartAt field value if set, zero value otherwise.
func (o *OnCallPageLiveCallOver) GetSupportHoursStartAt() time.Time {
	if o == nil || o.SupportHoursStartAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SupportHoursStartAt
}

// GetSupportHoursStartAtOk returns a tuple with the SupportHoursStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageLiveCallOver) GetSupportHoursStartAtOk() (*time.Time, bool) {
	if o == nil || o.SupportHoursStartAt == nil {
		return nil, false
	}
	return o.SupportHoursStartAt, true
}

// HasSupportHoursStartAt returns a boolean if a field has been set.
func (o *OnCallPageLiveCallOver) HasSupportHoursStartAt() bool {
	return o != nil && o.SupportHoursStartAt != nil
}

// SetSupportHoursStartAt gets a reference to the given time.Time and assigns it to the SupportHoursStartAt field.
func (o *OnCallPageLiveCallOver) SetSupportHoursStartAt(v time.Time) {
	o.SupportHoursStartAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageLiveCallOver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.SupportHoursStartAt != nil {
		if o.SupportHoursStartAt.Nanosecond() == 0 {
			toSerialize["support_hours_start_at"] = o.SupportHoursStartAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["support_hours_start_at"] = o.SupportHoursStartAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageLiveCallOver) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Outcome             *OnCallPageLiveCallOutcome `json:"outcome,omitempty"`
		SupportHoursStartAt *time.Time                 `json:"support_hours_start_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"outcome", "support_hours_start_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Outcome != nil && !all.Outcome.IsValid() {
		hasInvalidField = true
	} else {
		o.Outcome = all.Outcome
	}
	o.SupportHoursStartAt = all.SupportHoursStartAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
