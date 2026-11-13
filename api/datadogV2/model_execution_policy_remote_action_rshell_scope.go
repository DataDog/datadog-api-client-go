// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyRemoteActionRshellScope Restricts the policy to specific remote shell paths and commands.
type ExecutionPolicyRemoteActionRshellScope struct {
	// The non-elevated remote shell commands this policy contributes. The command list
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
	AllowedCommands *ExecutionPolicyRemoteActionRshellAllowedCommands `json:"allowed_commands,omitempty"`
	// The remote shell scope rules.
	Rules []ExecutionPolicyRemoteActionRshellScopeRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyRemoteActionRshellScope instantiates a new ExecutionPolicyRemoteActionRshellScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyRemoteActionRshellScope(rules []ExecutionPolicyRemoteActionRshellScopeRule) *ExecutionPolicyRemoteActionRshellScope {
	this := ExecutionPolicyRemoteActionRshellScope{}
	this.Rules = rules
	return &this
}

// NewExecutionPolicyRemoteActionRshellScopeWithDefaults instantiates a new ExecutionPolicyRemoteActionRshellScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyRemoteActionRshellScopeWithDefaults() *ExecutionPolicyRemoteActionRshellScope {
	this := ExecutionPolicyRemoteActionRshellScope{}
	return &this
}

// GetAllowedCommands returns the AllowedCommands field value if set, zero value otherwise.
func (o *ExecutionPolicyRemoteActionRshellScope) GetAllowedCommands() ExecutionPolicyRemoteActionRshellAllowedCommands {
	if o == nil || o.AllowedCommands == nil {
		var ret ExecutionPolicyRemoteActionRshellAllowedCommands
		return ret
	}
	return *o.AllowedCommands
}

// GetAllowedCommandsOk returns a tuple with the AllowedCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellScope) GetAllowedCommandsOk() (*ExecutionPolicyRemoteActionRshellAllowedCommands, bool) {
	if o == nil || o.AllowedCommands == nil {
		return nil, false
	}
	return o.AllowedCommands, true
}

// HasAllowedCommands returns a boolean if a field has been set.
func (o *ExecutionPolicyRemoteActionRshellScope) HasAllowedCommands() bool {
	return o != nil && o.AllowedCommands != nil
}

// SetAllowedCommands gets a reference to the given ExecutionPolicyRemoteActionRshellAllowedCommands and assigns it to the AllowedCommands field.
func (o *ExecutionPolicyRemoteActionRshellScope) SetAllowedCommands(v ExecutionPolicyRemoteActionRshellAllowedCommands) {
	o.AllowedCommands = &v
}

// GetRules returns the Rules field value.
func (o *ExecutionPolicyRemoteActionRshellScope) GetRules() []ExecutionPolicyRemoteActionRshellScopeRule {
	if o == nil {
		var ret []ExecutionPolicyRemoteActionRshellScopeRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellScope) GetRulesOk() (*[]ExecutionPolicyRemoteActionRshellScopeRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ExecutionPolicyRemoteActionRshellScope) SetRules(v []ExecutionPolicyRemoteActionRshellScopeRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyRemoteActionRshellScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedCommands != nil {
		toSerialize["allowed_commands"] = o.AllowedCommands
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyRemoteActionRshellScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedCommands *ExecutionPolicyRemoteActionRshellAllowedCommands `json:"allowed_commands,omitempty"`
		Rules           *[]ExecutionPolicyRemoteActionRshellScopeRule     `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_commands", "rules"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AllowedCommands != nil && all.AllowedCommands.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AllowedCommands = all.AllowedCommands
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
