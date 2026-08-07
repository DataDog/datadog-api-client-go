// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceUpdateDataAttributes Attributes of a maintenance update resource.
type MaintenanceUpdateDataAttributes struct {
	// Components affected at the time of the update.
	ComponentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// The date and time the update was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The message body of the update.
	Description *string `json:"description,omitempty"`
	// Whether the update was applied manually by a user (true) or automatically by the system (false).
	ManualTransition *bool `json:"manual_transition,omitempty"`
	// The date and time the update was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The date and time the update started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// The status of the maintenance update.
	Status *MaintenanceUpdateDataAttributesStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceUpdateDataAttributes instantiates a new MaintenanceUpdateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceUpdateDataAttributes() *MaintenanceUpdateDataAttributes {
	this := MaintenanceUpdateDataAttributes{}
	return &this
}

// NewMaintenanceUpdateDataAttributesWithDefaults instantiates a new MaintenanceUpdateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceUpdateDataAttributesWithDefaults() *MaintenanceUpdateDataAttributes {
	this := MaintenanceUpdateDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetComponentsAffected() []CreateMaintenanceRequestDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []CreateMaintenanceRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetComponentsAffectedOk() (*[]CreateMaintenanceRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []CreateMaintenanceRequestDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *MaintenanceUpdateDataAttributes) SetComponentsAffected(v []CreateMaintenanceRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MaintenanceUpdateDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MaintenanceUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetManualTransition returns the ManualTransition field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetManualTransition() bool {
	if o == nil || o.ManualTransition == nil {
		var ret bool
		return ret
	}
	return *o.ManualTransition
}

// GetManualTransitionOk returns a tuple with the ManualTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetManualTransitionOk() (*bool, bool) {
	if o == nil || o.ManualTransition == nil {
		return nil, false
	}
	return o.ManualTransition, true
}

// HasManualTransition returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasManualTransition() bool {
	return o != nil && o.ManualTransition != nil
}

// SetManualTransition gets a reference to the given bool and assigns it to the ManualTransition field.
func (o *MaintenanceUpdateDataAttributes) SetManualTransition(v bool) {
	o.ManualTransition = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *MaintenanceUpdateDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *MaintenanceUpdateDataAttributes) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MaintenanceUpdateDataAttributes) GetStatus() MaintenanceUpdateDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret MaintenanceUpdateDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceUpdateDataAttributes) GetStatusOk() (*MaintenanceUpdateDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MaintenanceUpdateDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given MaintenanceUpdateDataAttributesStatus and assigns it to the Status field.
func (o *MaintenanceUpdateDataAttributes) SetStatus(v MaintenanceUpdateDataAttributesStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ManualTransition != nil {
		toSerialize["manual_transition"] = o.ManualTransition
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["started_at"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceUpdateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []CreateMaintenanceRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		CreatedAt          *time.Time                                                      `json:"created_at,omitempty"`
		Description        *string                                                         `json:"description,omitempty"`
		ManualTransition   *bool                                                           `json:"manual_transition,omitempty"`
		ModifiedAt         *time.Time                                                      `json:"modified_at,omitempty"`
		StartedAt          *time.Time                                                      `json:"started_at,omitempty"`
		Status             *MaintenanceUpdateDataAttributesStatus                          `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "created_at", "description", "manual_transition", "modified_at", "started_at", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentsAffected = all.ComponentsAffected
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.ManualTransition = all.ManualTransition
	o.ModifiedAt = all.ModifiedAt
	o.StartedAt = all.StartedAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
