// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyUpdateAttributes Attributes for updating an org group policy. `policy_name`, `content`, and `enforcement_tier` may be omitted individually to leave them unchanged.
type OrgGroupPolicyUpdateAttributes struct {
	// The policy content as key-value pairs. For `org_config` policies, an arbitrary key-value map (for example, `{"value": "UTC"}`). For `role` policies, a `permissions` key containing an array of permission UUIDs (for example, `{"permissions": ["<uuid>", ...]}`).
	Content map[string]interface{} `json:"content,omitempty"`
	// The enforcement tier of the policy. `OVERRIDE_ALLOWED` means the policy is set but member orgs may mutate it. `GROUP_MANAGED` means the policy is strictly controlled and mutations are blocked for affected orgs. `DELEGATE` means each member org controls its own value. `role` policies only support `GROUP_MANAGED` and `DELEGATE` — `OVERRIDE_ALLOWED` is rejected for this policy type. Transitioning a `role` policy to `DELEGATE` (disabling it) is one-way — the policy cannot be transitioned back to `GROUP_MANAGED` afterward.
	EnforcementTier *OrgGroupPolicyEnforcementTier `json:"enforcement_tier,omitempty"`
	// The name of the policy. This becomes the name of the resource created across orgs in the group (for example, for `role` policies, the name of the created role). Omit to leave unchanged.
	PolicyName *string `json:"policy_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyUpdateAttributes instantiates a new OrgGroupPolicyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyUpdateAttributes() *OrgGroupPolicyUpdateAttributes {
	this := OrgGroupPolicyUpdateAttributes{}
	return &this
}

// NewOrgGroupPolicyUpdateAttributesWithDefaults instantiates a new OrgGroupPolicyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyUpdateAttributesWithDefaults() *OrgGroupPolicyUpdateAttributes {
	this := OrgGroupPolicyUpdateAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OrgGroupPolicyUpdateAttributes) GetContent() map[string]interface{} {
	if o == nil || o.Content == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyUpdateAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OrgGroupPolicyUpdateAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *OrgGroupPolicyUpdateAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetEnforcementTier returns the EnforcementTier field value if set, zero value otherwise.
func (o *OrgGroupPolicyUpdateAttributes) GetEnforcementTier() OrgGroupPolicyEnforcementTier {
	if o == nil || o.EnforcementTier == nil {
		var ret OrgGroupPolicyEnforcementTier
		return ret
	}
	return *o.EnforcementTier
}

// GetEnforcementTierOk returns a tuple with the EnforcementTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyUpdateAttributes) GetEnforcementTierOk() (*OrgGroupPolicyEnforcementTier, bool) {
	if o == nil || o.EnforcementTier == nil {
		return nil, false
	}
	return o.EnforcementTier, true
}

// HasEnforcementTier returns a boolean if a field has been set.
func (o *OrgGroupPolicyUpdateAttributes) HasEnforcementTier() bool {
	return o != nil && o.EnforcementTier != nil
}

// SetEnforcementTier gets a reference to the given OrgGroupPolicyEnforcementTier and assigns it to the EnforcementTier field.
func (o *OrgGroupPolicyUpdateAttributes) SetEnforcementTier(v OrgGroupPolicyEnforcementTier) {
	o.EnforcementTier = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *OrgGroupPolicyUpdateAttributes) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyUpdateAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *OrgGroupPolicyUpdateAttributes) HasPolicyName() bool {
	return o != nil && o.PolicyName != nil
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *OrgGroupPolicyUpdateAttributes) SetPolicyName(v string) {
	o.PolicyName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.EnforcementTier != nil {
		toSerialize["enforcement_tier"] = o.EnforcementTier
	}
	if o.PolicyName != nil {
		toSerialize["policy_name"] = o.PolicyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content         map[string]interface{}         `json:"content,omitempty"`
		EnforcementTier *OrgGroupPolicyEnforcementTier `json:"enforcement_tier,omitempty"`
		PolicyName      *string                        `json:"policy_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "enforcement_tier", "policy_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Content = all.Content
	if all.EnforcementTier != nil && !all.EnforcementTier.IsValid() {
		hasInvalidField = true
	} else {
		o.EnforcementTier = all.EnforcementTier
	}
	o.PolicyName = all.PolicyName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
