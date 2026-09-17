// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicyTarget A target this policy is scoped to, expressed as a set of Agent tags.
type ExecutionPolicyTarget struct {
	// The Agent tags identifying the target.
	AgentTags []string `json:"agent_tags"`
	// A human-readable name for the target.
	Name datadog.NullableString `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicyTarget instantiates a new ExecutionPolicyTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicyTarget(agentTags []string) *ExecutionPolicyTarget {
	this := ExecutionPolicyTarget{}
	this.AgentTags = agentTags
	return &this
}

// NewExecutionPolicyTargetWithDefaults instantiates a new ExecutionPolicyTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicyTargetWithDefaults() *ExecutionPolicyTarget {
	this := ExecutionPolicyTarget{}
	return &this
}

// GetAgentTags returns the AgentTags field value.
func (o *ExecutionPolicyTarget) GetAgentTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AgentTags
}

// GetAgentTagsOk returns a tuple with the AgentTags field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyTarget) GetAgentTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentTags, true
}

// SetAgentTags sets field value.
func (o *ExecutionPolicyTarget) SetAgentTags(v []string) {
	o.AgentTags = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionPolicyTarget) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExecutionPolicyTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExecutionPolicyTarget) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ExecutionPolicyTarget) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ExecutionPolicyTarget) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ExecutionPolicyTarget) UnsetName() {
	o.Name.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicyTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["agent_tags"] = o.AgentTags
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicyTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentTags *[]string              `json:"agent_tags"`
		Name      datadog.NullableString `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AgentTags == nil {
		return fmt.Errorf("required field agent_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_tags", "name"})
	} else {
		return err
	}
	o.AgentTags = *all.AgentTags
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
