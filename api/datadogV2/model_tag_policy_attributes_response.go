// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyAttributesResponse Attributes of a tag policy response.
type TagPolicyAttributesResponse struct {
	// Timestamp when the policy was created.
	CreatedAt time.Time `json:"created_at"`
	// User who created the policy.
	CreatedBy string `json:"created_by"`
	// Whether the policy is enabled.
	Enabled bool `json:"enabled"`
	// Timestamp when the policy was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// User who last modified the policy.
	ModifiedBy string `json:"modified_by"`
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
	// The version of the tag policy.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyAttributesResponse instantiates a new TagPolicyAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyAttributesResponse(createdAt time.Time, createdBy string, enabled bool, modifiedAt time.Time, modifiedBy string, negated bool, policyName string, required bool, scope string, source string, tagKey string, tagValuePatterns []string, version int64) *TagPolicyAttributesResponse {
	this := TagPolicyAttributesResponse{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Negated = negated
	this.PolicyName = policyName
	this.Required = required
	this.Scope = scope
	this.Source = source
	this.TagKey = tagKey
	this.TagValuePatterns = tagValuePatterns
	this.Version = version
	return &this
}

// NewTagPolicyAttributesResponseWithDefaults instantiates a new TagPolicyAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyAttributesResponseWithDefaults() *TagPolicyAttributesResponse {
	this := TagPolicyAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TagPolicyAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TagPolicyAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *TagPolicyAttributesResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *TagPolicyAttributesResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *TagPolicyAttributesResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *TagPolicyAttributesResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *TagPolicyAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *TagPolicyAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *TagPolicyAttributesResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *TagPolicyAttributesResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetNegated returns the Negated field value.
func (o *TagPolicyAttributesResponse) GetNegated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetNegatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negated, true
}

// SetNegated sets field value.
func (o *TagPolicyAttributesResponse) SetNegated(v bool) {
	o.Negated = v
}

// GetPolicyName returns the PolicyName field value.
func (o *TagPolicyAttributesResponse) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *TagPolicyAttributesResponse) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetRequired returns the Required field value.
func (o *TagPolicyAttributesResponse) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *TagPolicyAttributesResponse) SetRequired(v bool) {
	o.Required = v
}

// GetScope returns the Scope field value.
func (o *TagPolicyAttributesResponse) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *TagPolicyAttributesResponse) SetScope(v string) {
	o.Scope = v
}

// GetSource returns the Source field value.
func (o *TagPolicyAttributesResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TagPolicyAttributesResponse) SetSource(v string) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *TagPolicyAttributesResponse) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *TagPolicyAttributesResponse) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValuePatterns returns the TagValuePatterns field value.
func (o *TagPolicyAttributesResponse) GetTagValuePatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// SetTagValuePatterns sets field value.
func (o *TagPolicyAttributesResponse) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// GetVersion returns the Version field value.
func (o *TagPolicyAttributesResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagPolicyAttributesResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *TagPolicyAttributesResponse) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["negated"] = o.Negated
	toSerialize["policy_name"] = o.PolicyName
	toSerialize["required"] = o.Required
	toSerialize["scope"] = o.Scope
	toSerialize["source"] = o.Source
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_value_patterns"] = o.TagValuePatterns
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time `json:"created_at"`
		CreatedBy        *string    `json:"created_by"`
		Enabled          *bool      `json:"enabled"`
		ModifiedAt       *time.Time `json:"modified_at"`
		ModifiedBy       *string    `json:"modified_by"`
		Negated          *bool      `json:"negated"`
		PolicyName       *string    `json:"policy_name"`
		Required         *bool      `json:"required"`
		Scope            *string    `json:"scope"`
		Source           *string    `json:"source"`
		TagKey           *string    `json:"tag_key"`
		TagValuePatterns *[]string  `json:"tag_value_patterns"`
		Version          *int64     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
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
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "enabled", "modified_at", "modified_by", "negated", "policy_name", "required", "scope", "source", "tag_key", "tag_value_patterns", "version"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Negated = *all.Negated
	o.PolicyName = *all.PolicyName
	o.Required = *all.Required
	o.Scope = *all.Scope
	o.Source = *all.Source
	o.TagKey = *all.TagKey
	o.TagValuePatterns = *all.TagValuePatterns
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
