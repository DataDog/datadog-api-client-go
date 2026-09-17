// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsOwner Run the workflow as its owner.
type WorkflowRunAsOwner struct {
	// The owner run-as type.
	Type WorkflowRunAsOwnerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWorkflowRunAsOwner instantiates a new WorkflowRunAsOwner object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowRunAsOwner(typeVar WorkflowRunAsOwnerType) *WorkflowRunAsOwner {
	this := WorkflowRunAsOwner{}
	this.Type = typeVar
	return &this
}

// NewWorkflowRunAsOwnerWithDefaults instantiates a new WorkflowRunAsOwner object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowRunAsOwnerWithDefaults() *WorkflowRunAsOwner {
	this := WorkflowRunAsOwner{}
	return &this
}

// GetType returns the Type field value.
func (o *WorkflowRunAsOwner) GetType() WorkflowRunAsOwnerType {
	if o == nil {
		var ret WorkflowRunAsOwnerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunAsOwner) GetTypeOk() (*WorkflowRunAsOwnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WorkflowRunAsOwner) SetType(v WorkflowRunAsOwnerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowRunAsOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowRunAsOwner) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *WorkflowRunAsOwnerType `json:"type"`
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
