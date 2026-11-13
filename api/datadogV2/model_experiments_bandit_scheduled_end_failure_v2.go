// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditScheduledEndFailureV2 Why the scheduled end couldn't run. It isn't retried until the end or fallback changes. Absent when there's no failure.
type ExperimentsBanditScheduledEndFailureV2 struct {
	// When the scheduled end failed.
	FailedAt time.Time `json:"failed_at"`
	// Why the scheduled end failed.
	Reason string `json:"reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditScheduledEndFailureV2 instantiates a new ExperimentsBanditScheduledEndFailureV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditScheduledEndFailureV2(failedAt time.Time, reason string) *ExperimentsBanditScheduledEndFailureV2 {
	this := ExperimentsBanditScheduledEndFailureV2{}
	this.FailedAt = failedAt
	this.Reason = reason
	return &this
}

// NewExperimentsBanditScheduledEndFailureV2WithDefaults instantiates a new ExperimentsBanditScheduledEndFailureV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditScheduledEndFailureV2WithDefaults() *ExperimentsBanditScheduledEndFailureV2 {
	this := ExperimentsBanditScheduledEndFailureV2{}
	return &this
}

// GetFailedAt returns the FailedAt field value.
func (o *ExperimentsBanditScheduledEndFailureV2) GetFailedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.FailedAt
}

// GetFailedAtOk returns a tuple with the FailedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditScheduledEndFailureV2) GetFailedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedAt, true
}

// SetFailedAt sets field value.
func (o *ExperimentsBanditScheduledEndFailureV2) SetFailedAt(v time.Time) {
	o.FailedAt = v
}

// GetReason returns the Reason field value.
func (o *ExperimentsBanditScheduledEndFailureV2) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditScheduledEndFailureV2) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *ExperimentsBanditScheduledEndFailureV2) SetReason(v string) {
	o.Reason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditScheduledEndFailureV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FailedAt.Nanosecond() == 0 {
		toSerialize["failed_at"] = o.FailedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["failed_at"] = o.FailedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditScheduledEndFailureV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailedAt *time.Time `json:"failed_at"`
		Reason   *string    `json:"reason"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FailedAt == nil {
		return fmt.Errorf("required field failed_at missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failed_at", "reason"})
	} else {
		return err
	}
	o.FailedAt = *all.FailedAt
	o.Reason = *all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
