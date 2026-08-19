// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyScriptScope Restricts the policy to specific scripts.
type ExecutionPolicyScriptScope struct {
	// The script scope rules.
	Rules []ExecutionPolicyScriptScopeRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyScriptScope instantiates a new ExecutionPolicyScriptScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyScriptScope(rules []ExecutionPolicyScriptScopeRule) *ExecutionPolicyScriptScope {
	this := ExecutionPolicyScriptScope{}
	this.Rules = rules
	return &this
}

// NewExecutionPolicyScriptScopeWithDefaults instantiates a new ExecutionPolicyScriptScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyScriptScopeWithDefaults() *ExecutionPolicyScriptScope {
	this := ExecutionPolicyScriptScope{}
	return &this
}

// GetRules returns the Rules field value.
func (o *ExecutionPolicyScriptScope) GetRules() []ExecutionPolicyScriptScopeRule {
	if o == nil {
		var ret []ExecutionPolicyScriptScopeRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyScriptScope) GetRulesOk() (*[]ExecutionPolicyScriptScopeRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *ExecutionPolicyScriptScope) SetRules(v []ExecutionPolicyScriptScopeRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyScriptScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyScriptScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rules *[]ExecutionPolicyScriptScopeRule `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rules"})
	} else {
		return err
	}
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
