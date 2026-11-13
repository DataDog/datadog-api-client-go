// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshV2 The most recently started refresh. Null until one starts.
type ExperimentsBanditRefreshV2 struct {
	// When the refresh was requested.
	CreatedAt time.Time `json:"created_at"`
	// When the refresh finished. Null while it runs.
	EndedAt datadog.NullableTime `json:"ended_at"`
	// The refresh ID.
	Id uuid.UUID `json:"id"`
	// When the refresh started running. Null while it is pending.
	StartedAt datadog.NullableTime `json:"started_at"`
	// The status of the refresh.
	Status ExperimentsBanditRefreshStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditRefreshV2 instantiates a new ExperimentsBanditRefreshV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditRefreshV2(createdAt time.Time, endedAt datadog.NullableTime, id uuid.UUID, startedAt datadog.NullableTime, status ExperimentsBanditRefreshStatus) *ExperimentsBanditRefreshV2 {
	this := ExperimentsBanditRefreshV2{}
	this.CreatedAt = createdAt
	this.EndedAt = endedAt
	this.Id = id
	this.StartedAt = startedAt
	this.Status = status
	return &this
}

// NewExperimentsBanditRefreshV2WithDefaults instantiates a new ExperimentsBanditRefreshV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditRefreshV2WithDefaults() *ExperimentsBanditRefreshV2 {
	this := ExperimentsBanditRefreshV2{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ExperimentsBanditRefreshV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ExperimentsBanditRefreshV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEndedAt returns the EndedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditRefreshV2) GetEndedAt() time.Time {
	if o == nil || o.EndedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditRefreshV2) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// SetEndedAt sets field value.
func (o *ExperimentsBanditRefreshV2) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}

// GetId returns the Id field value.
func (o *ExperimentsBanditRefreshV2) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ExperimentsBanditRefreshV2) SetId(v uuid.UUID) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *ExperimentsBanditRefreshV2) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsBanditRefreshV2) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value.
func (o *ExperimentsBanditRefreshV2) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetStatus returns the Status field value.
func (o *ExperimentsBanditRefreshV2) GetStatus() ExperimentsBanditRefreshStatus {
	if o == nil {
		var ret ExperimentsBanditRefreshStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2) GetStatusOk() (*ExperimentsBanditRefreshStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ExperimentsBanditRefreshV2) SetStatus(v ExperimentsBanditRefreshStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditRefreshV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["ended_at"] = o.EndedAt.Get()
	toSerialize["id"] = o.Id
	toSerialize["started_at"] = o.StartedAt.Get()
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditRefreshV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                      `json:"created_at"`
		EndedAt   datadog.NullableTime            `json:"ended_at"`
		Id        *uuid.UUID                      `json:"id"`
		StartedAt datadog.NullableTime            `json:"started_at"`
		Status    *ExperimentsBanditRefreshStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.EndedAt.IsSet() {
		return fmt.Errorf("required field ended_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if !all.StartedAt.IsSet() {
		return fmt.Errorf("required field started_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "ended_at", "id", "started_at", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.EndedAt = all.EndedAt
	o.Id = *all.Id
	o.StartedAt = all.StartedAt
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableExperimentsBanditRefreshV2 handles when a null is used for ExperimentsBanditRefreshV2.
type NullableExperimentsBanditRefreshV2 struct {
	value *ExperimentsBanditRefreshV2
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditRefreshV2) Get() *ExperimentsBanditRefreshV2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditRefreshV2) Set(val *ExperimentsBanditRefreshV2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditRefreshV2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsBanditRefreshV2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditRefreshV2 initializes the struct as if Set has been called.
func NewNullableExperimentsBanditRefreshV2(val *ExperimentsBanditRefreshV2) *NullableExperimentsBanditRefreshV2 {
	return &NullableExperimentsBanditRefreshV2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditRefreshV2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditRefreshV2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
