// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleCreateAttributes Attributes that can be supplied when creating a tag rule.
type TagRuleCreateAttributes struct {
	// Whether the rule is currently enforced. Defaults to `true` for newly created rules.
	Enabled *bool `json:"enabled,omitempty"`
	// When `true`, the rule matches tag values that do NOT match any of the supplied patterns. Defaults to `false`.
	Negated *bool `json:"negated,omitempty"`
	// Human-readable name for the tag rule.
	PolicyName string `json:"policy_name"`
	// The rule type allowed when creating a tag rule. Only `surfacing` is accepted at
	// creation time.
	PolicyType TagRuleCreateType `json:"policy_type"`
	// When `true`, telemetry without this tag is treated as a violation. Defaults to `false`.
	Required *bool `json:"required,omitempty"`
	// The scope the rule applies within. Typically an environment, team, or
	// organization-level identifier used to limit where the rule is enforced.
	Scope string `json:"scope"`
	// The telemetry source that a tag rule applies to.
	Source TagRuleSource `json:"source"`
	// The tag key that the rule governs (for example, `service`).
	TagKey string `json:"tag_key"`
	// One or more patterns that valid values for the tag key must match. At least one
	// pattern is required.
	TagValuePatterns []string `json:"tag_value_patterns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagRuleCreateAttributes instantiates a new TagRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagRuleCreateAttributes(policyName string, policyType TagRuleCreateType, scope string, source TagRuleSource, tagKey string, tagValuePatterns []string) *TagRuleCreateAttributes {
	this := TagRuleCreateAttributes{}
	this.PolicyName = policyName
	this.PolicyType = policyType
	this.Scope = scope
	this.Source = source
	this.TagKey = tagKey
	this.TagValuePatterns = tagValuePatterns
	return &this
}

// NewTagRuleCreateAttributesWithDefaults instantiates a new TagRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagRuleCreateAttributesWithDefaults() *TagRuleCreateAttributes {
	this := TagRuleCreateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TagRuleCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TagRuleCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TagRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *TagRuleCreateAttributes) GetNegated() bool {
	if o == nil || o.Negated == nil {
		var ret bool
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetNegatedOk() (*bool, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *TagRuleCreateAttributes) HasNegated() bool {
	return o != nil && o.Negated != nil
}

// SetNegated gets a reference to the given bool and assigns it to the Negated field.
func (o *TagRuleCreateAttributes) SetNegated(v bool) {
	o.Negated = &v
}

// GetPolicyName returns the PolicyName field value.
func (o *TagRuleCreateAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *TagRuleCreateAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetPolicyType returns the PolicyType field value.
func (o *TagRuleCreateAttributes) GetPolicyType() TagRuleCreateType {
	if o == nil {
		var ret TagRuleCreateType
		return ret
	}
	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetPolicyTypeOk() (*TagRuleCreateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value.
func (o *TagRuleCreateAttributes) SetPolicyType(v TagRuleCreateType) {
	o.PolicyType = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TagRuleCreateAttributes) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TagRuleCreateAttributes) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TagRuleCreateAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetScope returns the Scope field value.
func (o *TagRuleCreateAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *TagRuleCreateAttributes) SetScope(v string) {
	o.Scope = v
}

// GetSource returns the Source field value.
func (o *TagRuleCreateAttributes) GetSource() TagRuleSource {
	if o == nil {
		var ret TagRuleSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetSourceOk() (*TagRuleSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TagRuleCreateAttributes) SetSource(v TagRuleSource) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *TagRuleCreateAttributes) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *TagRuleCreateAttributes) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValuePatterns returns the TagValuePatterns field value.
func (o *TagRuleCreateAttributes) GetTagValuePatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value
// and a boolean to check if the value has been set.
func (o *TagRuleCreateAttributes) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// SetTagValuePatterns sets field value.
func (o *TagRuleCreateAttributes) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Negated != nil {
		toSerialize["negated"] = o.Negated
	}
	toSerialize["policy_name"] = o.PolicyName
	toSerialize["policy_type"] = o.PolicyType
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
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
func (o *TagRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool              `json:"enabled,omitempty"`
		Negated          *bool              `json:"negated,omitempty"`
		PolicyName       *string            `json:"policy_name"`
		PolicyType       *TagRuleCreateType `json:"policy_type"`
		Required         *bool              `json:"required,omitempty"`
		Scope            *string            `json:"scope"`
		Source           *TagRuleSource     `json:"source"`
		TagKey           *string            `json:"tag_key"`
		TagValuePatterns *[]string          `json:"tag_value_patterns"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	if all.PolicyType == nil {
		return fmt.Errorf("required field policy_type missing")
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "negated", "policy_name", "policy_type", "required", "scope", "source", "tag_key", "tag_value_patterns"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Negated = all.Negated
	o.PolicyName = *all.PolicyName
	if !all.PolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.PolicyType = *all.PolicyType
	}
	o.Required = all.Required
	o.Scope = *all.Scope
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.TagKey = *all.TagKey
	o.TagValuePatterns = *all.TagValuePatterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
