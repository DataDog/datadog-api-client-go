// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyUpdateAttributes Update an existing Cloud Workload Security Agent policy
type CloudWorkloadSecurityAgentPolicyUpdateAttributes struct {
	// The description of the policy
	Description *string `json:"description,omitempty"`
	// Whether the policy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The host tags defining where this policy is deployed
	HostTags []string `json:"hostTags,omitempty"`
	// The name of the policy
	Name *string `json:"name,omitempty"`
	// The priority of the policy
	Priority *int64 `json:"priority,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentPolicyUpdateAttributes instantiates a new CloudWorkloadSecurityAgentPolicyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentPolicyUpdateAttributes() *CloudWorkloadSecurityAgentPolicyUpdateAttributes {
	this := CloudWorkloadSecurityAgentPolicyUpdateAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentPolicyUpdateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentPolicyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentPolicyUpdateAttributesWithDefaults() *CloudWorkloadSecurityAgentPolicyUpdateAttributes {
	this := CloudWorkloadSecurityAgentPolicyUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostTags returns the HostTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetHostTags() []string {
	if o == nil || o.HostTags == nil {
		var ret []string
		return ret
	}
	return o.HostTags
}

// GetHostTagsOk returns a tuple with the HostTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetHostTagsOk() (*[]string, bool) {
	if o == nil || o.HostTags == nil {
		return nil, false
	}
	return &o.HostTags, true
}

// HasHostTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) HasHostTags() bool {
	return o != nil && o.HostTags != nil
}

// SetHostTags gets a reference to the given []string and assigns it to the HostTags field.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) SetHostTags(v []string) {
	o.HostTags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) SetPriority(v int64) {
	o.Priority = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentPolicyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HostTags != nil {
		toSerialize["hostTags"] = o.HostTags
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentPolicyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description,omitempty"`
		Enabled     *bool    `json:"enabled,omitempty"`
		HostTags    []string `json:"hostTags,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Priority    *int64   `json:"priority,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "hostTags", "name", "priority"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.HostTags = all.HostTags
	o.Name = all.Name
	o.Priority = all.Priority

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
