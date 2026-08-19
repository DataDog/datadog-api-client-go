// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyScope Restricts where the policy applies. At most one of `kubernetes`, `scripts`,
// or `remote_action_rshell` can be set. An empty object means the policy has
// no scope restriction.
type ExecutionPolicyScope struct {
	// Restricts the policy to specific Kubernetes namespaces.
	Kubernetes *ExecutionPolicyKubernetesScope `json:"kubernetes,omitempty"`
	// Restricts the policy to specific remote shell paths.
	RemoteActionRshell *ExecutionPolicyRemoteActionRshellScope `json:"remote_action_rshell,omitempty"`
	// Restricts the policy to specific scripts.
	Scripts *ExecutionPolicyScriptScope `json:"scripts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyScope instantiates a new ExecutionPolicyScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyScope() *ExecutionPolicyScope {
	this := ExecutionPolicyScope{}
	return &this
}

// NewExecutionPolicyScopeWithDefaults instantiates a new ExecutionPolicyScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyScopeWithDefaults() *ExecutionPolicyScope {
	this := ExecutionPolicyScope{}
	return &this
}

// GetKubernetes returns the Kubernetes field value if set, zero value otherwise.
func (o *ExecutionPolicyScope) GetKubernetes() ExecutionPolicyKubernetesScope {
	if o == nil || o.Kubernetes == nil {
		var ret ExecutionPolicyKubernetesScope
		return ret
	}
	return *o.Kubernetes
}

// GetKubernetesOk returns a tuple with the Kubernetes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyScope) GetKubernetesOk() (*ExecutionPolicyKubernetesScope, bool) {
	if o == nil || o.Kubernetes == nil {
		return nil, false
	}
	return o.Kubernetes, true
}

// HasKubernetes returns a boolean if a field has been set.
func (o *ExecutionPolicyScope) HasKubernetes() bool {
	return o != nil && o.Kubernetes != nil
}

// SetKubernetes gets a reference to the given ExecutionPolicyKubernetesScope and assigns it to the Kubernetes field.
func (o *ExecutionPolicyScope) SetKubernetes(v ExecutionPolicyKubernetesScope) {
	o.Kubernetes = &v
}

// GetRemoteActionRshell returns the RemoteActionRshell field value if set, zero value otherwise.
func (o *ExecutionPolicyScope) GetRemoteActionRshell() ExecutionPolicyRemoteActionRshellScope {
	if o == nil || o.RemoteActionRshell == nil {
		var ret ExecutionPolicyRemoteActionRshellScope
		return ret
	}
	return *o.RemoteActionRshell
}

// GetRemoteActionRshellOk returns a tuple with the RemoteActionRshell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyScope) GetRemoteActionRshellOk() (*ExecutionPolicyRemoteActionRshellScope, bool) {
	if o == nil || o.RemoteActionRshell == nil {
		return nil, false
	}
	return o.RemoteActionRshell, true
}

// HasRemoteActionRshell returns a boolean if a field has been set.
func (o *ExecutionPolicyScope) HasRemoteActionRshell() bool {
	return o != nil && o.RemoteActionRshell != nil
}

// SetRemoteActionRshell gets a reference to the given ExecutionPolicyRemoteActionRshellScope and assigns it to the RemoteActionRshell field.
func (o *ExecutionPolicyScope) SetRemoteActionRshell(v ExecutionPolicyRemoteActionRshellScope) {
	o.RemoteActionRshell = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *ExecutionPolicyScope) GetScripts() ExecutionPolicyScriptScope {
	if o == nil || o.Scripts == nil {
		var ret ExecutionPolicyScriptScope
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyScope) GetScriptsOk() (*ExecutionPolicyScriptScope, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *ExecutionPolicyScope) HasScripts() bool {
	return o != nil && o.Scripts != nil
}

// SetScripts gets a reference to the given ExecutionPolicyScriptScope and assigns it to the Scripts field.
func (o *ExecutionPolicyScope) SetScripts(v ExecutionPolicyScriptScope) {
	o.Scripts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Kubernetes != nil {
		toSerialize["kubernetes"] = o.Kubernetes
	}
	if o.RemoteActionRshell != nil {
		toSerialize["remote_action_rshell"] = o.RemoteActionRshell
	}
	if o.Scripts != nil {
		toSerialize["scripts"] = o.Scripts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kubernetes         *ExecutionPolicyKubernetesScope         `json:"kubernetes,omitempty"`
		RemoteActionRshell *ExecutionPolicyRemoteActionRshellScope `json:"remote_action_rshell,omitempty"`
		Scripts            *ExecutionPolicyScriptScope             `json:"scripts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"kubernetes", "remote_action_rshell", "scripts"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Kubernetes != nil && all.Kubernetes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Kubernetes = all.Kubernetes
	if all.RemoteActionRshell != nil && all.RemoteActionRshell.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RemoteActionRshell = all.RemoteActionRshell
	if all.Scripts != nil && all.Scripts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scripts = all.Scripts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
