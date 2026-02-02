// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyAttributesUpdateRequest Attributes for updating a tag policy. All fields are optional.
type TagPolicyAttributesUpdateRequest struct {
	// Whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether the policy is negated.
	Negated *bool `json:"negated,omitempty"`
	// The name of the tag policy.
	PolicyName *string `json:"policy_name,omitempty"`
	// Whether the tag is required.
	Required *bool `json:"required,omitempty"`
	// The scope of the tag policy.
	Scope *string `json:"scope,omitempty"`
	// The data source for the tag policy (e.g., logs, metrics).
	Source *string `json:"source,omitempty"`
	// The tag key that the policy applies to.
	TagKey *string `json:"tag_key,omitempty"`
	// List of patterns that tag values must match.
	TagValuePatterns []string `json:"tag_value_patterns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyAttributesUpdateRequest instantiates a new TagPolicyAttributesUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyAttributesUpdateRequest() *TagPolicyAttributesUpdateRequest {
	this := TagPolicyAttributesUpdateRequest{}
	return &this
}

// NewTagPolicyAttributesUpdateRequestWithDefaults instantiates a new TagPolicyAttributesUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyAttributesUpdateRequestWithDefaults() *TagPolicyAttributesUpdateRequest {
	this := TagPolicyAttributesUpdateRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TagPolicyAttributesUpdateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetNegated() bool {
	if o == nil || o.Negated == nil {
		var ret bool
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetNegatedOk() (*bool, bool) {
	if o == nil || o.Negated == nil {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasNegated() bool {
	return o != nil && o.Negated != nil
}

// SetNegated gets a reference to the given bool and assigns it to the Negated field.
func (o *TagPolicyAttributesUpdateRequest) SetNegated(v bool) {
	o.Negated = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasPolicyName() bool {
	return o != nil && o.PolicyName != nil
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *TagPolicyAttributesUpdateRequest) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TagPolicyAttributesUpdateRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *TagPolicyAttributesUpdateRequest) SetScope(v string) {
	o.Scope = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TagPolicyAttributesUpdateRequest) SetSource(v string) {
	o.Source = &v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetTagKey() string {
	if o == nil || o.TagKey == nil {
		var ret string
		return ret
	}
	return *o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetTagKeyOk() (*string, bool) {
	if o == nil || o.TagKey == nil {
		return nil, false
	}
	return o.TagKey, true
}

// HasTagKey returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasTagKey() bool {
	return o != nil && o.TagKey != nil
}

// SetTagKey gets a reference to the given string and assigns it to the TagKey field.
func (o *TagPolicyAttributesUpdateRequest) SetTagKey(v string) {
	o.TagKey = &v
}

// GetTagValuePatterns returns the TagValuePatterns field value if set, zero value otherwise.
func (o *TagPolicyAttributesUpdateRequest) GetTagValuePatterns() []string {
	if o == nil || o.TagValuePatterns == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesUpdateRequest) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil || o.TagValuePatterns == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// HasTagValuePatterns returns a boolean if a field has been set.
func (o *TagPolicyAttributesUpdateRequest) HasTagValuePatterns() bool {
	return o != nil && o.TagValuePatterns != nil
}

// SetTagValuePatterns gets a reference to the given []string and assigns it to the TagValuePatterns field.
func (o *TagPolicyAttributesUpdateRequest) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyAttributesUpdateRequest) MarshalJSON() ([]byte, error) {
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
	if o.PolicyName != nil {
		toSerialize["policy_name"] = o.PolicyName
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
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
func (o *TagPolicyAttributesUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled          *bool    `json:"enabled,omitempty"`
		Negated          *bool    `json:"negated,omitempty"`
		PolicyName       *string  `json:"policy_name,omitempty"`
		Required         *bool    `json:"required,omitempty"`
		Scope            *string  `json:"scope,omitempty"`
		Source           *string  `json:"source,omitempty"`
		TagKey           *string  `json:"tag_key,omitempty"`
		TagValuePatterns []string `json:"tag_value_patterns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "negated", "policy_name", "required", "scope", "source", "tag_key", "tag_value_patterns"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Negated = all.Negated
	o.PolicyName = all.PolicyName
	o.Required = all.Required
	o.Scope = all.Scope
	o.Source = all.Source
	o.TagKey = all.TagKey
	o.TagValuePatterns = all.TagValuePatterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
