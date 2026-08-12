// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentTriggerWrapper Schema for an agent-based trigger.
type AgentTriggerWrapper struct {
	// Trigger a workflow from an agent via the MCP execute tool. Workflow can be executed from Bits Chat, Bits Agent Builder, Claude Code, Codex, Cursor, and any other coding agent using the Datadog MCP.
	AgentTrigger AgentTrigger `json:"agentTrigger"`
	// Names of existing workflow steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentTriggerWrapper instantiates a new AgentTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentTriggerWrapper(agentTrigger AgentTrigger) *AgentTriggerWrapper {
	this := AgentTriggerWrapper{}
	this.AgentTrigger = agentTrigger
	return &this
}

// NewAgentTriggerWrapperWithDefaults instantiates a new AgentTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentTriggerWrapperWithDefaults() *AgentTriggerWrapper {
	this := AgentTriggerWrapper{}
	return &this
}

// GetAgentTrigger returns the AgentTrigger field value.
func (o *AgentTriggerWrapper) GetAgentTrigger() AgentTrigger {
	if o == nil {
		var ret AgentTrigger
		return ret
	}
	return o.AgentTrigger
}

// GetAgentTriggerOk returns a tuple with the AgentTrigger field value
// and a boolean to check if the value has been set.
func (o *AgentTriggerWrapper) GetAgentTriggerOk() (*AgentTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentTrigger, true
}

// SetAgentTrigger sets field value.
func (o *AgentTriggerWrapper) SetAgentTrigger(v AgentTrigger) {
	o.AgentTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *AgentTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *AgentTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *AgentTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["agentTrigger"] = o.AgentTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentTrigger   *AgentTrigger `json:"agentTrigger"`
		StartStepNames []string      `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AgentTrigger == nil {
		return fmt.Errorf("required field agentTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agentTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AgentTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AgentTrigger = *all.AgentTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
