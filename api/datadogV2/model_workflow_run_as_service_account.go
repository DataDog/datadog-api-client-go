// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowRunAsServiceAccount Run the workflow as a service account.
type WorkflowRunAsServiceAccount struct {
	// The service account identifier.
	Id string `json:"id"`
	// The service account run-as type.
	Type WorkflowRunAsServiceAccountType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWorkflowRunAsServiceAccount instantiates a new WorkflowRunAsServiceAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowRunAsServiceAccount(id string, typeVar WorkflowRunAsServiceAccountType) *WorkflowRunAsServiceAccount {
	this := WorkflowRunAsServiceAccount{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewWorkflowRunAsServiceAccountWithDefaults instantiates a new WorkflowRunAsServiceAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowRunAsServiceAccountWithDefaults() *WorkflowRunAsServiceAccount {
	this := WorkflowRunAsServiceAccount{}
	return &this
}

// GetId returns the Id field value.
func (o *WorkflowRunAsServiceAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunAsServiceAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *WorkflowRunAsServiceAccount) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *WorkflowRunAsServiceAccount) GetType() WorkflowRunAsServiceAccountType {
	if o == nil {
		var ret WorkflowRunAsServiceAccountType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunAsServiceAccount) GetTypeOk() (*WorkflowRunAsServiceAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WorkflowRunAsServiceAccount) SetType(v WorkflowRunAsServiceAccountType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowRunAsServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowRunAsServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                          `json:"id"`
		Type *WorkflowRunAsServiceAccountType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Id = *all.Id
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
