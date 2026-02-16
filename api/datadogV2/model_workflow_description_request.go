// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowDescriptionRequest
type WorkflowDescriptionRequest struct {
	// The name of the workflow.
	Name string `json:"name"`
	// The workflow specification as a JSON object.
	Spec interface{} `json:"spec"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowDescriptionRequest instantiates a new WorkflowDescriptionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowDescriptionRequest(name string, spec interface{}) *WorkflowDescriptionRequest {
	this := WorkflowDescriptionRequest{}
	this.Name = name
	this.Spec = spec
	return &this
}

// NewWorkflowDescriptionRequestWithDefaults instantiates a new WorkflowDescriptionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowDescriptionRequestWithDefaults() *WorkflowDescriptionRequest {
	this := WorkflowDescriptionRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *WorkflowDescriptionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowDescriptionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WorkflowDescriptionRequest) SetName(v string) {
	o.Name = v
}

// GetSpec returns the Spec field value.
func (o *WorkflowDescriptionRequest) GetSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *WorkflowDescriptionRequest) GetSpecOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value.
func (o *WorkflowDescriptionRequest) SetSpec(v interface{}) {
	o.Spec = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowDescriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["spec"] = o.Spec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowDescriptionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string      `json:"name"`
		Spec *interface{} `json:"spec"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Spec == nil {
		return fmt.Errorf("required field spec missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "spec"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Spec = *all.Spec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
