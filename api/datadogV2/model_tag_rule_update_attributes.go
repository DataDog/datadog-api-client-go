// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleUpdateAttributes Mutable attributes of a tag rule. Each field is optional; omitting a field leaves its
// current value unchanged. The `source` of a rule cannot be changed.
type TagRuleUpdateAttributes struct {
	// Whether the rule is currently enforced.
	Enabled *bool `json:"enabled,omitempty"`
	// Human-readable name for the tag rule.
	Name *string `json:"name,omitempty"`
	// When `true`, the rule matches tag values that do NOT match any of the supplied patterns.
	Negated *bool `json:"negated,omitempty"`
	// When `true`, telemetry without this tag is treated as a violation.
	Required *bool `json:"required,omitempty"`
	// How the rule is enforced. `blocking` rejects telemetry that violates the rule.
	// `surfacing` only highlights non-compliant telemetry without blocking it.
	RuleType *TagRuleType `json:"rule_type,omitempty"`
	// The scope the rule applies within.
	Scope *string `json:"scope,omitempty"`
	// The tag key that the rule governs.
	TagKey *string `json:"tag_key,omitempty"`
	// One or more patterns that valid values for the tag key must match.
	TagValuePatterns []string `json:"tag_value_patterns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagRuleUpdateAttributes instantiates a new TagRuleUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagRuleUpdateAttributes() *TagRuleUpdateAttributes {
	this := TagRuleUpdateAttributes{}
	return &this
}

// NewTagRuleUpdateAttributesWithDefaults instantiates a new TagRuleUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagRuleUpdateAttributesWithDefaults() *TagRuleUpdateAttributes {
	this := TagRuleUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TagRuleUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagRuleUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetNegated() bool {
	if o == nil || o.Negated == nil {
		var ret bool
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetNegatedOk() (*bool, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasNegated() bool {
	return o != nil && o.Negated != nil
}

// SetNegated gets a reference to the given bool and assigns it to the Negated field.
func (o *TagRuleUpdateAttributes) SetNegated(v bool) {
	o.Negated = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TagRuleUpdateAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetRuleType() TagRuleType {
	if o == nil || o.RuleType == nil {
		var ret TagRuleType
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetRuleTypeOk() (*TagRuleType, bool) {
	if o == nil || o.RuleType == nil {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasRuleType() bool {
	return o != nil && o.RuleType != nil
}

// SetRuleType gets a reference to the given TagRuleType and assigns it to the RuleType field.
func (o *TagRuleUpdateAttributes) SetRuleType(v TagRuleType) {
	o.RuleType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *TagRuleUpdateAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetTagKey() string {
	if o == nil || o.TagKey == nil {
		var ret string
		return ret
	}
	return *o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil || o.TagKey == nil {
		return nil, false
	}
	return o.TagKey, true
}

// HasTagKey returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasTagKey() bool {
	return o != nil && o.TagKey != nil
}

// SetTagKey gets a reference to the given string and assigns it to the TagKey field.
func (o *TagRuleUpdateAttributes) SetTagKey(v string) {
	o.TagKey = &v
}

// GetTagValuePatterns returns the TagValuePatterns field value if set, zero value otherwise.
func (o *TagRuleUpdateAttributes) GetTagValuePatterns() []string {
	if o == nil || o.TagValuePatterns == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagRuleUpdateAttributes) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil || o.TagValuePatterns == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// HasTagValuePatterns returns a boolean if a field has been set.
func (o *TagRuleUpdateAttributes) HasTagValuePatterns() bool {
	return o != nil && o.TagValuePatterns != nil
}

// SetTagValuePatterns gets a reference to the given []string and assigns it to the TagValuePatterns field.
func (o *TagRuleUpdateAttributes) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagRuleUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Negated != nil {
		toSerialize["negated"] = o.Negated
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.RuleType != nil {
		toSerialize["rule_type"] = o.RuleType
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TagKey != nil {
		toSerialize["tag_key"] = o.TagKey
	}
	if o.TagValuePatterns != nil {
		toSerialize["tag_value_patterns"] = o.TagValuePatterns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagRuleUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool        `json:"enabled,omitempty"`
		Name             *string      `json:"name,omitempty"`
		Negated          *bool        `json:"negated,omitempty"`
		Required         *bool        `json:"required,omitempty"`
		RuleType         *TagRuleType `json:"rule_type,omitempty"`
		Scope            *string      `json:"scope,omitempty"`
		TagKey           *string      `json:"tag_key,omitempty"`
		TagValuePatterns []string     `json:"tag_value_patterns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "name", "negated", "required", "rule_type", "scope", "tag_key", "tag_value_patterns"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Name = all.Name
	o.Negated = all.Negated
	o.Required = all.Required
	if all.RuleType != nil && !all.RuleType.IsValid() {
		hasInvalidField = true
	} else {
		o.RuleType = all.RuleType
	}
	o.Scope = all.Scope
	o.TagKey = all.TagKey
	o.TagValuePatterns = all.TagValuePatterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
