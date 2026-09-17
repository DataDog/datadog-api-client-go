// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsInitiator Run the workflow as the user who initiates the execution.
type WorkflowRunAsInitiator struct {
	// The initiator run-as type.
	Type WorkflowRunAsInitiatorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWorkflowRunAsInitiator instantiates a new WorkflowRunAsInitiator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowRunAsInitiator(typeVar WorkflowRunAsInitiatorType) *WorkflowRunAsInitiator {
	this := WorkflowRunAsInitiator{}
	this.Type = typeVar
	return &this
}

// NewWorkflowRunAsInitiatorWithDefaults instantiates a new WorkflowRunAsInitiator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowRunAsInitiatorWithDefaults() *WorkflowRunAsInitiator {
	this := WorkflowRunAsInitiator{}
	return &this
}

// GetType returns the Type field value.
func (o *WorkflowRunAsInitiator) GetType() WorkflowRunAsInitiatorType {
	if o == nil {
		var ret WorkflowRunAsInitiatorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunAsInitiator) GetTypeOk() (*WorkflowRunAsInitiatorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WorkflowRunAsInitiator) SetType(v WorkflowRunAsInitiatorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowRunAsInitiator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowRunAsInitiator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *WorkflowRunAsInitiatorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
