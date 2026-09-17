// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyKubernetesScopeRule A rule restricting a Kubernetes scope to specific namespaces.
type ExecutionPolicyKubernetesScopeRule struct {
	// The Kubernetes namespaces this rule applies to.
	TargetNamespaces []string `json:"target_namespaces"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyKubernetesScopeRule instantiates a new ExecutionPolicyKubernetesScopeRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyKubernetesScopeRule(targetNamespaces []string) *ExecutionPolicyKubernetesScopeRule {
	this := ExecutionPolicyKubernetesScopeRule{}
	this.TargetNamespaces = targetNamespaces
	return &this
}

// NewExecutionPolicyKubernetesScopeRuleWithDefaults instantiates a new ExecutionPolicyKubernetesScopeRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyKubernetesScopeRuleWithDefaults() *ExecutionPolicyKubernetesScopeRule {
	this := ExecutionPolicyKubernetesScopeRule{}
	return &this
}

// GetTargetNamespaces returns the TargetNamespaces field value.
func (o *ExecutionPolicyKubernetesScopeRule) GetTargetNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetNamespaces
}

// GetTargetNamespacesOk returns a tuple with the TargetNamespaces field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyKubernetesScopeRule) GetTargetNamespacesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetNamespaces, true
}

// SetTargetNamespaces sets field value.
func (o *ExecutionPolicyKubernetesScopeRule) SetTargetNamespaces(v []string) {
	o.TargetNamespaces = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyKubernetesScopeRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["target_namespaces"] = o.TargetNamespaces

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyKubernetesScopeRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TargetNamespaces *[]string `json:"target_namespaces"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TargetNamespaces == nil {
		return fmt.Errorf("required field target_namespaces missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"target_namespaces"})
	} else {
		return err
	}
	o.TargetNamespaces = *all.TargetNamespaces

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
