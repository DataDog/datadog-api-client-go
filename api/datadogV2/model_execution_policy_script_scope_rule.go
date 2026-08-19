// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyScriptScopeRule A rule restricting a script scope to specific script names.
type ExecutionPolicyScriptScopeRule struct {
	// The script names this rule applies to.
	TargetScriptNames []string `json:"target_script_names"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyScriptScopeRule instantiates a new ExecutionPolicyScriptScopeRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyScriptScopeRule(targetScriptNames []string) *ExecutionPolicyScriptScopeRule {
	this := ExecutionPolicyScriptScopeRule{}
	this.TargetScriptNames = targetScriptNames
	return &this
}

// NewExecutionPolicyScriptScopeRuleWithDefaults instantiates a new ExecutionPolicyScriptScopeRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyScriptScopeRuleWithDefaults() *ExecutionPolicyScriptScopeRule {
	this := ExecutionPolicyScriptScopeRule{}
	return &this
}

// GetTargetScriptNames returns the TargetScriptNames field value.
func (o *ExecutionPolicyScriptScopeRule) GetTargetScriptNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetScriptNames
}

// GetTargetScriptNamesOk returns a tuple with the TargetScriptNames field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyScriptScopeRule) GetTargetScriptNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetScriptNames, true
}

// SetTargetScriptNames sets field value.
func (o *ExecutionPolicyScriptScopeRule) SetTargetScriptNames(v []string) {
	o.TargetScriptNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyScriptScopeRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target_script_names"] = o.TargetScriptNames

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyScriptScopeRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TargetScriptNames *[]string `json:"target_script_names"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TargetScriptNames == nil {
		return fmt.Errorf("required field target_script_names missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"target_script_names"})
	} else {
		return err
	}
	o.TargetScriptNames = *all.TargetScriptNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
