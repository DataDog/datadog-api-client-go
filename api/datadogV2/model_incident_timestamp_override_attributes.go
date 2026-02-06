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

// IncidentTimestampOverrideAttributes Attributes of an incident timestamp override.
type IncidentTimestampOverrideAttributes struct {
	// Timestamp of when the override was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp of when the override was deleted (soft delete).
	DeletedAt datadog.NullableTime `json:"deleted_at"`
	// The UUID of the incident.
	IncidentId uuid.UUID `json:"incident_id"`
	// Timestamp of when the override was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The type of timestamp being overridden.
	TimestampType TimestampType `json:"timestamp_type"`
	// The timestamp value for the override.
	TimestampValue time.Time `json:"timestamp_value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTimestampOverrideAttributes instantiates a new IncidentTimestampOverrideAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTimestampOverrideAttributes(createdAt time.Time, deletedAt datadog.NullableTime, incidentId uuid.UUID, modifiedAt time.Time, timestampType TimestampType, timestampValue time.Time) *IncidentTimestampOverrideAttributes {
	this := IncidentTimestampOverrideAttributes{}
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.IncidentId = incidentId
	this.ModifiedAt = modifiedAt
	this.TimestampType = timestampType
	this.TimestampValue = timestampValue
	return &this
}

// NewIncidentTimestampOverrideAttributesWithDefaults instantiates a new IncidentTimestampOverrideAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTimestampOverrideAttributesWithDefaults() *IncidentTimestampOverrideAttributes {
	this := IncidentTimestampOverrideAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentTimestampOverrideAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentTimestampOverrideAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *IncidentTimestampOverrideAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentTimestampOverrideAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// SetDeletedAt sets field value.
func (o *IncidentTimestampOverrideAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentTimestampOverrideAttributes) GetIncidentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideAttributes) GetIncidentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentTimestampOverrideAttributes) SetIncidentId(v uuid.UUID) {
	o.IncidentId = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentTimestampOverrideAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentTimestampOverrideAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetTimestampType returns the TimestampType field value.
func (o *IncidentTimestampOverrideAttributes) GetTimestampType() TimestampType {
	if o == nil {
		var ret TimestampType
		return ret
	}
	return o.TimestampType
}

// GetTimestampTypeOk returns a tuple with the TimestampType field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideAttributes) GetTimestampTypeOk() (*TimestampType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampType, true
}

// SetTimestampType sets field value.
func (o *IncidentTimestampOverrideAttributes) SetTimestampType(v TimestampType) {
	o.TimestampType = v
}

// GetTimestampValue returns the TimestampValue field value.
func (o *IncidentTimestampOverrideAttributes) GetTimestampValue() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.TimestampValue
}

// GetTimestampValueOk returns a tuple with the TimestampValue field value
// and a boolean to check if the value has been set.
func (o *IncidentTimestampOverrideAttributes) GetTimestampValueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampValue, true
}

// SetTimestampValue sets field value.
func (o *IncidentTimestampOverrideAttributes) SetTimestampValue(v time.Time) {
	o.TimestampValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTimestampOverrideAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["deleted_at"] = o.DeletedAt.Get()
	toSerialize["incident_id"] = o.IncidentId
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["timestamp_type"] = o.TimestampType
	if o.TimestampValue.Nanosecond() == 0 {
		toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp_value"] = o.TimestampValue.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTimestampOverrideAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time           `json:"created_at"`
		DeletedAt      datadog.NullableTime `json:"deleted_at"`
		IncidentId     *uuid.UUID           `json:"incident_id"`
		ModifiedAt     *time.Time           `json:"modified_at"`
		TimestampType  *TimestampType       `json:"timestamp_type"`
		TimestampValue *time.Time           `json:"timestamp_value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.DeletedAt.IsSet() {
		return fmt.Errorf("required field deleted_at missing")
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.TimestampType == nil {
		return fmt.Errorf("required field timestamp_type missing")
	}
	if all.TimestampValue == nil {
		return fmt.Errorf("required field timestamp_value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "deleted_at", "incident_id", "modified_at", "timestamp_type", "timestamp_value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.IncidentId = *all.IncidentId
	o.ModifiedAt = *all.ModifiedAt
	if !all.TimestampType.IsValid() {
		hasInvalidField = true
	} else {
		o.TimestampType = *all.TimestampType
	}
	o.TimestampValue = *all.TimestampValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
