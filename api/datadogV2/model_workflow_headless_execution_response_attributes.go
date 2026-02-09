// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowHeadlessExecutionResponseAttributes
type WorkflowHeadlessExecutionResponseAttributes struct {
	// The ID of the workflow instance that was created
	InstanceId uuid.UUID `json:"instance_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowHeadlessExecutionResponseAttributes instantiates a new WorkflowHeadlessExecutionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowHeadlessExecutionResponseAttributes(instanceId uuid.UUID) *WorkflowHeadlessExecutionResponseAttributes {
	this := WorkflowHeadlessExecutionResponseAttributes{}
	this.InstanceId = instanceId
	return &this
}

// NewWorkflowHeadlessExecutionResponseAttributesWithDefaults instantiates a new WorkflowHeadlessExecutionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowHeadlessExecutionResponseAttributesWithDefaults() *WorkflowHeadlessExecutionResponseAttributes {
	this := WorkflowHeadlessExecutionResponseAttributes{}
	return &this
}

// GetInstanceId returns the InstanceId field value.
func (o *WorkflowHeadlessExecutionResponseAttributes) GetInstanceId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionResponseAttributes) GetInstanceIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value.
func (o *WorkflowHeadlessExecutionResponseAttributes) SetInstanceId(v uuid.UUID) {
	o.InstanceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowHeadlessExecutionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance_id"] = o.InstanceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowHeadlessExecutionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceId *uuid.UUID `json:"instance_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InstanceId == nil {
		return fmt.Errorf("required field instance_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_id"})
	} else {
		return err
	}
	o.InstanceId = *all.InstanceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
