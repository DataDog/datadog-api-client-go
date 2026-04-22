// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyCreateAttributes Attributes for creating an org group policy.
type OrgGroupPolicyCreateAttributes struct {
	// The policy content as key-value pairs.
	Content map[string]interface{} `json:"content"`
	// The name of the policy.
	PolicyName string `json:"policy_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyCreateAttributes instantiates a new OrgGroupPolicyCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyCreateAttributes(content map[string]interface{}, policyName string) *OrgGroupPolicyCreateAttributes {
	this := OrgGroupPolicyCreateAttributes{}
	this.Content = content
	this.PolicyName = policyName
	return &this
}

// NewOrgGroupPolicyCreateAttributesWithDefaults instantiates a new OrgGroupPolicyCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyCreateAttributesWithDefaults() *OrgGroupPolicyCreateAttributes {
	this := OrgGroupPolicyCreateAttributes{}
	return &this
}

// GetContent returns the Content field value.
func (o *OrgGroupPolicyCreateAttributes) GetContent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *OrgGroupPolicyCreateAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetPolicyName returns the PolicyName field value.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyCreateAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *OrgGroupPolicyCreateAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content
	toSerialize["policy_name"] = o.PolicyName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content    *map[string]interface{} `json:"content"`
		PolicyName *string                 `json:"policy_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "policy_name"})
	} else {
		return err
	}
	o.Content = *all.Content
	o.PolicyName = *all.PolicyName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
