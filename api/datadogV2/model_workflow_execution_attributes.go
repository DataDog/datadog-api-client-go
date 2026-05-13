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

// WorkflowExecutionAttributes Attributes of a workflow execution.
type WorkflowExecutionAttributes struct {
	// The ID of the parent AI workflow.
	AiWorkflowId uuid.UUID `json:"ai_workflow_id"`
	// Timestamp when the execution was created.
	CreatedAt time.Time `json:"created_at"`
	// The list of entities processed by this execution.
	Entities []Entity `json:"entities"`
	// The Datadog Workflow Automation instance ID for this execution.
	InstanceId uuid.UUID `json:"instance_id"`
	// Timestamp when the execution was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowExecutionAttributes instantiates a new WorkflowExecutionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowExecutionAttributes(aiWorkflowId uuid.UUID, createdAt time.Time, entities []Entity, instanceId uuid.UUID, updatedAt time.Time) *WorkflowExecutionAttributes {
	this := WorkflowExecutionAttributes{}
	this.AiWorkflowId = aiWorkflowId
	this.CreatedAt = createdAt
	this.Entities = entities
	this.InstanceId = instanceId
	this.UpdatedAt = updatedAt
	return &this
}

// NewWorkflowExecutionAttributesWithDefaults instantiates a new WorkflowExecutionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowExecutionAttributesWithDefaults() *WorkflowExecutionAttributes {
	this := WorkflowExecutionAttributes{}
	return &this
}

// GetAiWorkflowId returns the AiWorkflowId field value.
func (o *WorkflowExecutionAttributes) GetAiWorkflowId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.AiWorkflowId
}

// GetAiWorkflowIdOk returns a tuple with the AiWorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionAttributes) GetAiWorkflowIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AiWorkflowId, true
}

// SetAiWorkflowId sets field value.
func (o *WorkflowExecutionAttributes) SetAiWorkflowId(v uuid.UUID) {
	o.AiWorkflowId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *WorkflowExecutionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *WorkflowExecutionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEntities returns the Entities field value.
func (o *WorkflowExecutionAttributes) GetEntities() []Entity {
	if o == nil {
		var ret []Entity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionAttributes) GetEntitiesOk() (*[]Entity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value.
func (o *WorkflowExecutionAttributes) SetEntities(v []Entity) {
	o.Entities = v
}

// GetInstanceId returns the InstanceId field value.
func (o *WorkflowExecutionAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *WorkflowExecutionAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *WorkflowExecutionAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *WorkflowExecutionAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowExecutionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ai_workflow_id"] = o.AiWorkflowId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["entities"] = o.Entities
	toSerialize["instance_id"] = o.InstanceId
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowExecutionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AiWorkflowId *uuid.UUID `json:"ai_workflow_id"`
		CreatedAt    *time.Time `json:"created_at"`
		Entities     *[]Entity  `json:"entities"`
		InstanceId   *uuid.UUID `json:"instance_id"`
		UpdatedAt    *time.Time `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AiWorkflowId == nil {
		return fmt.Errorf("required field ai_workflow_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Entities == nil {
		return fmt.Errorf("required field entities missing")
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ai_workflow_id", "created_at", "entities", "instance_id", "updated_at"})
	} else {
		return err
	}
	o.AiWorkflowId = *all.AiWorkflowId
	o.CreatedAt = *all.CreatedAt
	o.Entities = *all.Entities
	o.InstanceId = *all.InstanceId
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
