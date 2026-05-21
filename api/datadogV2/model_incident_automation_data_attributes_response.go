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

// IncidentAutomationDataAttributesResponse Attributes of incident automation data.
type IncidentAutomationDataAttributesResponse struct {
	// Timestamp when the data was created.
	CreatedAt time.Time `json:"created_at"`
	// The incident identifier.
	IncidentId uuid.UUID `json:"incident_id"`
	// The automation data key.
	Key string `json:"key"`
	// Timestamp when the data was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The automation data value.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentAutomationDataAttributesResponse instantiates a new IncidentAutomationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAutomationDataAttributesResponse(createdAt time.Time, incidentId uuid.UUID, key string, updatedAt time.Time, value string) *IncidentAutomationDataAttributesResponse {
	this := IncidentAutomationDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.IncidentId = incidentId
	this.Key = key
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewIncidentAutomationDataAttributesResponseWithDefaults instantiates a new IncidentAutomationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAutomationDataAttributesResponseWithDefaults() *IncidentAutomationDataAttributesResponse {
	this := IncidentAutomationDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentAutomationDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentAutomationDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentAutomationDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentAutomationDataAttributesResponse) GetIncidentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentAutomationDataAttributesResponse) GetIncidentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentAutomationDataAttributesResponse) SetIncidentId(v uuid.UUID) {
	o.IncidentId = v
}

// GetKey returns the Key field value.
func (o *IncidentAutomationDataAttributesResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IncidentAutomationDataAttributesResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *IncidentAutomationDataAttributesResponse) SetKey(v string) {
	o.Key = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *IncidentAutomationDataAttributesResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentAutomationDataAttributesResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *IncidentAutomationDataAttributesResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value.
func (o *IncidentAutomationDataAttributesResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IncidentAutomationDataAttributesResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *IncidentAutomationDataAttributesResponse) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAutomationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["incident_id"] = o.IncidentId
	toSerialize["key"] = o.Key
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAutomationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time `json:"created_at"`
		IncidentId *uuid.UUID `json:"incident_id"`
		Key        *string    `json:"key"`
		UpdatedAt  *time.Time `json:"updated_at"`
		Value      *string    `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "incident_id", "key", "updated_at", "value"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.IncidentId = *all.IncidentId
	o.Key = *all.Key
	o.UpdatedAt = *all.UpdatedAt
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
