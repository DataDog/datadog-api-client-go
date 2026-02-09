// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowHeadlessExecutionConfig
type WorkflowHeadlessExecutionConfig struct {
	// List of connections to use for the workflow execution
	Connections []WorkflowHeadlessExecutionConnection `json:"connections"`
	// Input parameters for the workflow execution
	Inputs map[string]interface{} `json:"inputs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowHeadlessExecutionConfig instantiates a new WorkflowHeadlessExecutionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowHeadlessExecutionConfig(connections []WorkflowHeadlessExecutionConnection, inputs map[string]interface{}) *WorkflowHeadlessExecutionConfig {
	this := WorkflowHeadlessExecutionConfig{}
	this.Connections = connections
	this.Inputs = inputs
	return &this
}

// NewWorkflowHeadlessExecutionConfigWithDefaults instantiates a new WorkflowHeadlessExecutionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowHeadlessExecutionConfigWithDefaults() *WorkflowHeadlessExecutionConfig {
	this := WorkflowHeadlessExecutionConfig{}
	return &this
}

// GetConnections returns the Connections field value.
func (o *WorkflowHeadlessExecutionConfig) GetConnections() []WorkflowHeadlessExecutionConnection {
	if o == nil {
		var ret []WorkflowHeadlessExecutionConnection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionConfig) GetConnectionsOk() (*[]WorkflowHeadlessExecutionConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value.
func (o *WorkflowHeadlessExecutionConfig) SetConnections(v []WorkflowHeadlessExecutionConnection) {
	o.Connections = v
}

// GetInputs returns the Inputs field value.
func (o *WorkflowHeadlessExecutionConfig) GetInputs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *WorkflowHeadlessExecutionConfig) GetInputsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *WorkflowHeadlessExecutionConfig) SetInputs(v map[string]interface{}) {
	o.Inputs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowHeadlessExecutionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connections"] = o.Connections
	toSerialize["inputs"] = o.Inputs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowHeadlessExecutionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connections *[]WorkflowHeadlessExecutionConnection `json:"connections"`
		Inputs      *map[string]interface{}                `json:"inputs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Connections == nil {
		return fmt.Errorf("required field connections missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connections", "inputs"})
	} else {
		return err
	}
	o.Connections = *all.Connections
	o.Inputs = *all.Inputs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
