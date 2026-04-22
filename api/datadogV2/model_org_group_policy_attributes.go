// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyAttributes Attributes of an org group policy.
type OrgGroupPolicyAttributes struct {
	// The policy content as key-value pairs.
	Content map[string]interface{} `json:"content,omitempty"`
	// Timestamp when the policy was enforced.
	EnforcedAt time.Time `json:"enforced_at"`
	// Timestamp when the policy was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the policy.
	PolicyName string `json:"policy_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyAttributes instantiates a new OrgGroupPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyAttributes(enforcedAt time.Time, modifiedAt time.Time, policyName string) *OrgGroupPolicyAttributes {
	this := OrgGroupPolicyAttributes{}
	this.EnforcedAt = enforcedAt
	this.ModifiedAt = modifiedAt
	this.PolicyName = policyName
	return &this
}

// NewOrgGroupPolicyAttributesWithDefaults instantiates a new OrgGroupPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyAttributesWithDefaults() *OrgGroupPolicyAttributes {
	this := OrgGroupPolicyAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OrgGroupPolicyAttributes) GetContent() map[string]interface{} {
	if o == nil || o.Content == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OrgGroupPolicyAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *OrgGroupPolicyAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetEnforcedAt returns the EnforcedAt field value.
func (o *OrgGroupPolicyAttributes) GetEnforcedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EnforcedAt
}

// GetEnforcedAtOk returns a tuple with the EnforcedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyAttributes) GetEnforcedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforcedAt, true
}

// SetEnforcedAt sets field value.
func (o *OrgGroupPolicyAttributes) SetEnforcedAt(v time.Time) {
	o.EnforcedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *OrgGroupPolicyAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *OrgGroupPolicyAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetPolicyName returns the PolicyName field value.
func (o *OrgGroupPolicyAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *OrgGroupPolicyAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.EnforcedAt.Nanosecond() == 0 {
		toSerialize["enforced_at"] = o.EnforcedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["enforced_at"] = o.EnforcedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["policy_name"] = o.PolicyName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content    map[string]interface{} `json:"content,omitempty"`
		EnforcedAt *time.Time             `json:"enforced_at"`
		ModifiedAt *time.Time             `json:"modified_at"`
		PolicyName *string                `json:"policy_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnforcedAt == nil {
		return fmt.Errorf("required field enforced_at missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "enforced_at", "modified_at", "policy_name"})
	} else {
		return err
	}
	o.Content = all.Content
	o.EnforcedAt = *all.EnforcedAt
	o.ModifiedAt = *all.ModifiedAt
	o.PolicyName = *all.PolicyName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
