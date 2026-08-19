// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyWriteAttributes Attributes used to create or update an execution policy.
type ExecutionPolicyWriteAttributes struct {
	// The set of actions this policy applies to.
	ActionPattern ExecutionPolicyActionPattern `json:"action_pattern"`
	// Whether the policy allows or denies matching actions.
	Effect ExecutionPolicyEffect `json:"effect"`
	// The name of the execution policy.
	Name string `json:"name"`
	// Restricts where the policy applies. At most one of `kubernetes`, `scripts`,
	// or `remote_action_rshell` can be set. An empty object means the policy has
	// no scope restriction.
	Scope *ExecutionPolicyScope `json:"scope,omitempty"`
	// The targets this policy applies to.
	Targets []ExecutionPolicyTarget `json:"targets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyWriteAttributes instantiates a new ExecutionPolicyWriteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyWriteAttributes(actionPattern ExecutionPolicyActionPattern, effect ExecutionPolicyEffect, name string) *ExecutionPolicyWriteAttributes {
	this := ExecutionPolicyWriteAttributes{}
	this.ActionPattern = actionPattern
	this.Effect = effect
	this.Name = name
	return &this
}

// NewExecutionPolicyWriteAttributesWithDefaults instantiates a new ExecutionPolicyWriteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyWriteAttributesWithDefaults() *ExecutionPolicyWriteAttributes {
	this := ExecutionPolicyWriteAttributes{}
	return &this
}

// GetActionPattern returns the ActionPattern field value.
func (o *ExecutionPolicyWriteAttributes) GetActionPattern() ExecutionPolicyActionPattern {
	if o == nil {
		var ret ExecutionPolicyActionPattern
		return ret
	}
	return o.ActionPattern
}

// GetActionPatternOk returns a tuple with the ActionPattern field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyWriteAttributes) GetActionPatternOk() (*ExecutionPolicyActionPattern, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPattern, true
}

// SetActionPattern sets field value.
func (o *ExecutionPolicyWriteAttributes) SetActionPattern(v ExecutionPolicyActionPattern) {
	o.ActionPattern = v
}

// GetEffect returns the Effect field value.
func (o *ExecutionPolicyWriteAttributes) GetEffect() ExecutionPolicyEffect {
	if o == nil {
		var ret ExecutionPolicyEffect
		return ret
	}
	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyWriteAttributes) GetEffectOk() (*ExecutionPolicyEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value.
func (o *ExecutionPolicyWriteAttributes) SetEffect(v ExecutionPolicyEffect) {
	o.Effect = v
}

// GetName returns the Name field value.
func (o *ExecutionPolicyWriteAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyWriteAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExecutionPolicyWriteAttributes) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ExecutionPolicyWriteAttributes) GetScope() ExecutionPolicyScope {
	if o == nil || o.Scope == nil {
		var ret ExecutionPolicyScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyWriteAttributes) GetScopeOk() (*ExecutionPolicyScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ExecutionPolicyWriteAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given ExecutionPolicyScope and assigns it to the Scope field.
func (o *ExecutionPolicyWriteAttributes) SetScope(v ExecutionPolicyScope) {
	o.Scope = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *ExecutionPolicyWriteAttributes) GetTargets() []ExecutionPolicyTarget {
	if o == nil || o.Targets == nil {
		var ret []ExecutionPolicyTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyWriteAttributes) GetTargetsOk() (*[]ExecutionPolicyTarget, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *ExecutionPolicyWriteAttributes) HasTargets() bool {
	return o != nil && o.Targets != nil
}

// SetTargets gets a reference to the given []ExecutionPolicyTarget and assigns it to the Targets field.
func (o *ExecutionPolicyWriteAttributes) SetTargets(v []ExecutionPolicyTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyWriteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_pattern"] = o.ActionPattern
	toSerialize["effect"] = o.Effect
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyWriteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionPattern *ExecutionPolicyActionPattern `json:"action_pattern"`
		Effect        *ExecutionPolicyEffect        `json:"effect"`
		Name          *string                       `json:"name"`
		Scope         *ExecutionPolicyScope         `json:"scope,omitempty"`
		Targets       []ExecutionPolicyTarget       `json:"targets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionPattern == nil {
		return fmt.Errorf("required field action_pattern missing")
	}
	if all.Effect == nil {
		return fmt.Errorf("required field effect missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_pattern", "effect", "name", "scope", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ActionPattern.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActionPattern = *all.ActionPattern
	if !all.Effect.IsValid() {
		hasInvalidField = true
	} else {
		o.Effect = *all.Effect
	}
	o.Name = *all.Name
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.Targets = all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
