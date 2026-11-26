// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunWorkflowExecutionInput Test-only generated schema for the dashboard versioning proof of concept.
type RunWorkflowExecutionInput struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunWorkflowExecutionInput instantiates a new RunWorkflowExecutionInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunWorkflowExecutionInput(name string, value string) *RunWorkflowExecutionInput {
	this := RunWorkflowExecutionInput{}
	this.Name = name
	this.Value = value
	return &this
}

// NewRunWorkflowExecutionInputWithDefaults instantiates a new RunWorkflowExecutionInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunWorkflowExecutionInputWithDefaults() *RunWorkflowExecutionInput {
	this := RunWorkflowExecutionInput{}
	return &this
}

// GetName returns the Name field value.
func (o *RunWorkflowExecutionInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RunWorkflowExecutionInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RunWorkflowExecutionInput) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *RunWorkflowExecutionInput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RunWorkflowExecutionInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RunWorkflowExecutionInput) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunWorkflowExecutionInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunWorkflowExecutionInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name"`
		Value *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	o.Name = *all.Name
	o.Value = *all.Value

	return nil
}
