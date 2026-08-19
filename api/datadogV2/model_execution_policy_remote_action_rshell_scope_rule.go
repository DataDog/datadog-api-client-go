// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyRemoteActionRshellScopeRule A rule restricting remote shell access to specific paths.
type ExecutionPolicyRemoteActionRshellScopeRule struct {
	// The level of remote shell access granted for the target paths.
	Access ExecutionPolicyRemoteActionRshellAccess `json:"access"`
	// The file system paths this rule applies to.
	TargetPaths []string `json:"target_paths"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyRemoteActionRshellScopeRule instantiates a new ExecutionPolicyRemoteActionRshellScopeRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyRemoteActionRshellScopeRule(access ExecutionPolicyRemoteActionRshellAccess, targetPaths []string) *ExecutionPolicyRemoteActionRshellScopeRule {
	this := ExecutionPolicyRemoteActionRshellScopeRule{}
	this.Access = access
	this.TargetPaths = targetPaths
	return &this
}

// NewExecutionPolicyRemoteActionRshellScopeRuleWithDefaults instantiates a new ExecutionPolicyRemoteActionRshellScopeRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyRemoteActionRshellScopeRuleWithDefaults() *ExecutionPolicyRemoteActionRshellScopeRule {
	this := ExecutionPolicyRemoteActionRshellScopeRule{}
	return &this
}

// GetAccess returns the Access field value.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) GetAccess() ExecutionPolicyRemoteActionRshellAccess {
	if o == nil {
		var ret ExecutionPolicyRemoteActionRshellAccess
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) GetAccessOk() (*ExecutionPolicyRemoteActionRshellAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) SetAccess(v ExecutionPolicyRemoteActionRshellAccess) {
	o.Access = v
}

// GetTargetPaths returns the TargetPaths field value.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) GetTargetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetPaths
}

// GetTargetPathsOk returns a tuple with the TargetPaths field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) GetTargetPathsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPaths, true
}

// SetTargetPaths sets field value.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) SetTargetPaths(v []string) {
	o.TargetPaths = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyRemoteActionRshellScopeRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access"] = o.Access
	toSerialize["target_paths"] = o.TargetPaths

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyRemoteActionRshellScopeRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Access      *ExecutionPolicyRemoteActionRshellAccess `json:"access"`
		TargetPaths *[]string                                `json:"target_paths"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Access == nil {
		return fmt.Errorf("required field access missing")
	}
	if all.TargetPaths == nil {
		return fmt.Errorf("required field target_paths missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access", "target_paths"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Access.IsValid() {
		hasInvalidField = true
	} else {
		o.Access = *all.Access
	}
	o.TargetPaths = *all.TargetPaths

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
