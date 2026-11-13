// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyRemoteActionRshellAllowedCommands The non-elevated remote shell commands this policy contributes. The command list
// is additive across matching policies and is capped by the organization's Balto
// allowlist. An explicit empty list (`commands: []`) contributes no commands.
// The exact wildcard `rshell:*` contributes the current Balto command ceiling.
// A literal list can contain values such as `rshell:cat` and `rshell:ls`.
// For compatibility with policies created before this field existed,
// `allowed_commands` may be omitted temporarily. Compatibility readers currently
// interpret that omission as inheriting the organization's Balto allowlist; it is
// not equivalent to an explicit empty list. After supported writers and write
// guards are deployed and legacy policies are backfilled, a later rollout will
// require this choice for applicable `rshell` ALLOW policies and reject it when
// missing. That future requirement does not apply to unrelated action families or
// DENY policies, and this transitional schema does not enforce it yet.
type ExecutionPolicyRemoteActionRshellAllowedCommands struct {
	// The command identifiers contributed by this policy. Set this to `[]` to
	// contribute no commands; do not omit `commands` when `allowed_commands` is
	// present.
	Commands []ExecutionPolicyRemoteActionRshellAllowedCommand `json:"commands"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyRemoteActionRshellAllowedCommands instantiates a new ExecutionPolicyRemoteActionRshellAllowedCommands object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyRemoteActionRshellAllowedCommands(commands []ExecutionPolicyRemoteActionRshellAllowedCommand) *ExecutionPolicyRemoteActionRshellAllowedCommands {
	this := ExecutionPolicyRemoteActionRshellAllowedCommands{}
	this.Commands = commands
	return &this
}

// NewExecutionPolicyRemoteActionRshellAllowedCommandsWithDefaults instantiates a new ExecutionPolicyRemoteActionRshellAllowedCommands object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyRemoteActionRshellAllowedCommandsWithDefaults() *ExecutionPolicyRemoteActionRshellAllowedCommands {
	this := ExecutionPolicyRemoteActionRshellAllowedCommands{}
	return &this
}

// GetCommands returns the Commands field value.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommands) GetCommands() []ExecutionPolicyRemoteActionRshellAllowedCommand {
	if o == nil {
		var ret []ExecutionPolicyRemoteActionRshellAllowedCommand
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommands) GetCommandsOk() (*[]ExecutionPolicyRemoteActionRshellAllowedCommand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commands, true
}

// SetCommands sets field value.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommands) SetCommands(v []ExecutionPolicyRemoteActionRshellAllowedCommand) {
	o.Commands = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyRemoteActionRshellAllowedCommands) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commands"] = o.Commands

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyRemoteActionRshellAllowedCommands) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Commands *[]ExecutionPolicyRemoteActionRshellAllowedCommand `json:"commands"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Commands == nil {
		return fmt.Errorf("required field commands missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commands"})
	} else {
		return err
	}
	o.Commands = *all.Commands

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
