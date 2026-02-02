// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyAttributesRequest Attributes for creating or updating a tag policy.
type TagPolicyAttributesRequest struct {
	// Whether the policy is enabled.
	Enabled bool `json:"enabled"`
	// Whether the policy is negated.
	Negated bool `json:"negated"`
	// The name of the tag policy.
	PolicyName string `json:"policy_name"`
	// Whether the tag is required.
	Required bool `json:"required"`
	// The scope of the tag policy.
	Scope string `json:"scope"`
	// The data source for the tag policy (e.g., logs, metrics).
	Source string `json:"source"`
	// The tag key that the policy applies to.
	TagKey string `json:"tag_key"`
	// List of patterns that tag values must match.
	TagValuePatterns []string `json:"tag_value_patterns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyAttributesRequest instantiates a new TagPolicyAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyAttributesRequest(enabled bool, negated bool, policyName string, required bool, scope string, source string, tagKey string, tagValuePatterns []string) *TagPolicyAttributesRequest {
	this := TagPolicyAttributesRequest{}
	this.Enabled = enabled
	this.Negated = negated
	this.PolicyName = policyName
	this.Required = required
	this.Scope = scope
	this.Source = source
	this.TagKey = tagKey
	this.TagValuePatterns = tagValuePatterns
	return &this
}

// NewTagPolicyAttributesRequestWithDefaults instantiates a new TagPolicyAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyAttributesRequestWithDefaults() *TagPolicyAttributesRequest {
	this := TagPolicyAttributesRequest{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *TagPolicyAttributesRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *TagPolicyAttributesRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNegated returns the Negated field value.
func (o *TagPolicyAttributesRequest) GetNegated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetNegatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negated, true
}

// SetNegated sets field value.
func (o *TagPolicyAttributesRequest) SetNegated(v bool) {
	o.Negated = v
}

// GetPolicyName returns the PolicyName field value.
func (o *TagPolicyAttributesRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *TagPolicyAttributesRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetRequired returns the Required field value.
func (o *TagPolicyAttributesRequest) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *TagPolicyAttributesRequest) SetRequired(v bool) {
	o.Required = v
}

// GetScope returns the Scope field value.
func (o *TagPolicyAttributesRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *TagPolicyAttributesRequest) SetScope(v string) {
	o.Scope = v
}

// GetSource returns the Source field value.
func (o *TagPolicyAttributesRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TagPolicyAttributesRequest) SetSource(v string) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *TagPolicyAttributesRequest) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *TagPolicyAttributesRequest) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValuePatterns returns the TagValuePatterns field value.
func (o *TagPolicyAttributesRequest) GetTagValuePatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesRequest) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// SetTagValuePatterns sets field value.
func (o *TagPolicyAttributesRequest) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["negated"] = o.Negated
	toSerialize["policy_name"] = o.PolicyName
	toSerialize["required"] = o.Required
	toSerialize["scope"] = o.Scope
	toSerialize["source"] = o.Source
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_value_patterns"] = o.TagValuePatterns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool     `json:"enabled"`
		Negated          *bool     `json:"negated"`
		PolicyName       *string   `json:"policy_name"`
		Required         *bool     `json:"required"`
		Scope            *string   `json:"scope"`
		Source           *string   `json:"source"`
		TagKey           *string   `json:"tag_key"`
		TagValuePatterns *[]string `json:"tag_value_patterns"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Negated == nil {
		return fmt.Errorf("required field negated missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagValuePatterns == nil {
		return fmt.Errorf("required field tag_value_patterns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "negated", "policy_name", "required", "scope", "source", "tag_key", "tag_value_patterns"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Negated = *all.Negated
	o.PolicyName = *all.PolicyName
	o.Required = *all.Required
	o.Scope = *all.Scope
	o.Source = *all.Source
	o.TagKey = *all.TagKey
	o.TagValuePatterns = *all.TagValuePatterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
