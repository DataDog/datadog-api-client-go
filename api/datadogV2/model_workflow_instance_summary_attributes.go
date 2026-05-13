// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowInstanceSummaryAttributes Attributes of a workflow instance summary.
type WorkflowInstanceSummaryAttributes struct {
	// Timestamp when the workflow instance was created.
	CreatedAt time.Time `json:"created_at"`
	// The entities being processed by this workflow instance.
	Entities []Entity `json:"entities"`
	// The current status of the step.
	Status WorkflowInstanceSummaryAttributesStatus `json:"status"`
	// A human-readable display name for the current status.
	StatusDisplay string `json:"status_display"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowInstanceSummaryAttributes instantiates a new WorkflowInstanceSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowInstanceSummaryAttributes(createdAt time.Time, entities []Entity, status WorkflowInstanceSummaryAttributesStatus, statusDisplay string) *WorkflowInstanceSummaryAttributes {
	this := WorkflowInstanceSummaryAttributes{}
	this.CreatedAt = createdAt
	this.Entities = entities
	this.Status = status
	this.StatusDisplay = statusDisplay
	return &this
}

// NewWorkflowInstanceSummaryAttributesWithDefaults instantiates a new WorkflowInstanceSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowInstanceSummaryAttributesWithDefaults() *WorkflowInstanceSummaryAttributes {
	this := WorkflowInstanceSummaryAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *WorkflowInstanceSummaryAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceSummaryAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *WorkflowInstanceSummaryAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEntities returns the Entities field value.
func (o *WorkflowInstanceSummaryAttributes) GetEntities() []Entity {
	if o == nil {
		var ret []Entity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceSummaryAttributes) GetEntitiesOk() (*[]Entity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value.
func (o *WorkflowInstanceSummaryAttributes) SetEntities(v []Entity) {
	o.Entities = v
}

// GetStatus returns the Status field value.
func (o *WorkflowInstanceSummaryAttributes) GetStatus() WorkflowInstanceSummaryAttributesStatus {
	if o == nil {
		var ret WorkflowInstanceSummaryAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceSummaryAttributes) GetStatusOk() (*WorkflowInstanceSummaryAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *WorkflowInstanceSummaryAttributes) SetStatus(v WorkflowInstanceSummaryAttributesStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value.
func (o *WorkflowInstanceSummaryAttributes) GetStatusDisplay() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceSummaryAttributes) GetStatusDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDisplay, true
}

// SetStatusDisplay sets field value.
func (o *WorkflowInstanceSummaryAttributes) SetStatusDisplay(v string) {
	o.StatusDisplay = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowInstanceSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["entities"] = o.Entities
	toSerialize["status"] = o.Status
	toSerialize["status_display"] = o.StatusDisplay

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowInstanceSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time                               `json:"created_at"`
		Entities      *[]Entity                                `json:"entities"`
		Status        *WorkflowInstanceSummaryAttributesStatus `json:"status"`
		StatusDisplay *string                                  `json:"status_display"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Entities == nil {
		return fmt.Errorf("required field entities missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StatusDisplay == nil {
		return fmt.Errorf("required field status_display missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "entities", "status", "status_display"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.Entities = *all.Entities
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.StatusDisplay = *all.StatusDisplay

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
