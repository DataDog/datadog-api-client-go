// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyRemoteActionRshellAllowedCommand A non-elevated remote shell command allowed by the policy.
type ExecutionPolicyRemoteActionRshellAllowedCommand struct {
	// The fully qualified command identifier. Use the exact `rshell:*` identifier
	// to allow all commands in the organization's Balto allowlist, or use a literal
	// identifier such as `rshell:cat`.
	Command string `json:"command"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyRemoteActionRshellAllowedCommand instantiates a new ExecutionPolicyRemoteActionRshellAllowedCommand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyRemoteActionRshellAllowedCommand(command string) *ExecutionPolicyRemoteActionRshellAllowedCommand {
	this := ExecutionPolicyRemoteActionRshellAllowedCommand{}
	this.Command = command
	return &this
}

// NewExecutionPolicyRemoteActionRshellAllowedCommandWithDefaults instantiates a new ExecutionPolicyRemoteActionRshellAllowedCommand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyRemoteActionRshellAllowedCommandWithDefaults() *ExecutionPolicyRemoteActionRshellAllowedCommand {
	this := ExecutionPolicyRemoteActionRshellAllowedCommand{}
	return &this
}

// GetCommand returns the Command field value.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommand) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommand) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommand) SetCommand(v string) {
	o.Command = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyRemoteActionRshellAllowedCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["command"] = o.Command

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Command *string `json:"command"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Command == nil {
		return fmt.Errorf("required field command missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"command"})
	} else {
		return err
	}
	o.Command = *all.Command

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
