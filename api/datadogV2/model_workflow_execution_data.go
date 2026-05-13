// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowExecutionData A single workflow execution resource.
type WorkflowExecutionData struct {
	// Attributes of a workflow execution.
	Attributes WorkflowExecutionAttributes `json:"attributes"`
	// The unique identifier of the workflow execution.
	Id uuid.UUID `json:"id"`
	// The resource type for workflow executions.
	Type WorkflowExecutionDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowExecutionData instantiates a new WorkflowExecutionData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowExecutionData(attributes WorkflowExecutionAttributes, id uuid.UUID, typeVar WorkflowExecutionDataType) *WorkflowExecutionData {
	this := WorkflowExecutionData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewWorkflowExecutionDataWithDefaults instantiates a new WorkflowExecutionData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowExecutionDataWithDefaults() *WorkflowExecutionData {
	this := WorkflowExecutionData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *WorkflowExecutionData) GetAttributes() WorkflowExecutionAttributes {
	if o == nil {
		var ret WorkflowExecutionAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionData) GetAttributesOk() (*WorkflowExecutionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *WorkflowExecutionData) SetAttributes(v WorkflowExecutionAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *WorkflowExecutionData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *WorkflowExecutionData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *WorkflowExecutionData) GetType() WorkflowExecutionDataType {
	if o == nil {
		var ret WorkflowExecutionDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionData) GetTypeOk() (*WorkflowExecutionDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WorkflowExecutionData) SetType(v WorkflowExecutionDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowExecutionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowExecutionData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *WorkflowExecutionAttributes `json:"attributes"`
		Id         *uuid.UUID                   `json:"id"`
		Type       *WorkflowExecutionDataType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
