// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleAttributes The attributes of a tag rule resource.
type TagRuleAttributes struct {
	// The RFC 3339 timestamp at which the rule was created.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user who created the rule.
	CreatedBy string `json:"created_by"`
	// The RFC 3339 timestamp at which the rule was soft-deleted. `null` if the rule has not been deleted. Only present when `include_deleted=true` is requested.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// The identifier of the user who soft-deleted the rule. `null` if the rule has not been deleted.
	DeletedBy datadog.NullableString `json:"deleted_by,omitempty"`
	// Whether the rule is currently enforced.
	Enabled bool `json:"enabled"`
	// The RFC 3339 timestamp at which the rule was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The identifier of the user who last modified the rule.
	ModifiedBy string `json:"modified_by"`
	// Human-readable name for the tag rule.
	Name string `json:"name"`
	// When `true`, the rule matches tag values that do NOT match any of the supplied patterns.
	Negated bool `json:"negated"`
	// When `true`, telemetry without this tag is treated as a violation.
	Required bool `json:"required"`
	// How the rule is enforced. `blocking` rejects telemetry that violates the rule.
	// `surfacing` only highlights non-compliant telemetry without blocking it.
	RuleType TagRuleType `json:"rule_type"`
	// The scope the rule applies within.
	Scope string `json:"scope"`
	// The telemetry source that a tag rule applies to.
	Source TagRuleSource `json:"source"`
	// The tag key that the rule governs.
	TagKey string `json:"tag_key"`
	// The patterns that valid values for the tag key must match.
	TagValuePatterns []string `json:"tag_value_patterns"`
	// A monotonically increasing version counter that is incremented on each update.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagRuleAttributes instantiates a new TagRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagRuleAttributes(createdAt time.Time, createdBy string, enabled bool, modifiedAt time.Time, modifiedBy string, name string, negated bool, required bool, ruleType TagRuleType, scope string, source TagRuleSource, tagKey string, tagValuePatterns []string, version int64) *TagRuleAttributes {
	this := TagRuleAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.Negated = negated
	this.Required = required
	this.RuleType = ruleType
	this.Scope = scope
	this.Source = source
	this.TagKey = tagKey
	this.TagValuePatterns = tagValuePatterns
	this.Version = version
	return &this
}

// NewTagRuleAttributesWithDefaults instantiates a new TagRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagRuleAttributesWithDefaults() *TagRuleAttributes {
	this := TagRuleAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TagRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TagRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *TagRuleAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *TagRuleAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagRuleAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagRuleAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TagRuleAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *TagRuleAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *TagRuleAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *TagRuleAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagRuleAttributes) GetDeletedBy() string {
	if o == nil || o.DeletedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeletedBy.Get()
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagRuleAttributes) GetDeletedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedBy.Get(), o.DeletedBy.IsSet()
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *TagRuleAttributes) HasDeletedBy() bool {
	return o != nil && o.DeletedBy.IsSet()
}

// SetDeletedBy gets a reference to the given datadog.NullableString and assigns it to the DeletedBy field.
func (o *TagRuleAttributes) SetDeletedBy(v string) {
	o.DeletedBy.Set(&v)
}

// SetDeletedByNil sets the value for DeletedBy to be an explicit nil.
func (o *TagRuleAttributes) SetDeletedByNil() {
	o.DeletedBy.Set(nil)
}

// UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil.
func (o *TagRuleAttributes) UnsetDeletedBy() {
	o.DeletedBy.Unset()
}

// GetEnabled returns the Enabled field value.
func (o *TagRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *TagRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *TagRuleAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *TagRuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *TagRuleAttributes) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *TagRuleAttributes) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *TagRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TagRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetNegated returns the Negated field value.
func (o *TagRuleAttributes) GetNegated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetNegatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negated, true
}

// SetNegated sets field value.
func (o *TagRuleAttributes) SetNegated(v bool) {
	o.Negated = v
}

// GetRequired returns the Required field value.
func (o *TagRuleAttributes) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *TagRuleAttributes) SetRequired(v bool) {
	o.Required = v
}

// GetRuleType returns the RuleType field value.
func (o *TagRuleAttributes) GetRuleType() TagRuleType {
	if o == nil {
		var ret TagRuleType
		return ret
	}
	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetRuleTypeOk() (*TagRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value.
func (o *TagRuleAttributes) SetRuleType(v TagRuleType) {
	o.RuleType = v
}

// GetScope returns the Scope field value.
func (o *TagRuleAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *TagRuleAttributes) SetScope(v string) {
	o.Scope = v
}

// GetSource returns the Source field value.
func (o *TagRuleAttributes) GetSource() TagRuleSource {
	if o == nil {
		var ret TagRuleSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetSourceOk() (*TagRuleSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TagRuleAttributes) SetSource(v TagRuleSource) {
	o.Source = v
}

// GetTagKey returns the TagKey field value.
func (o *TagRuleAttributes) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *TagRuleAttributes) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValuePatterns returns the TagValuePatterns field value.
func (o *TagRuleAttributes) GetTagValuePatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValuePatterns
}

// GetTagValuePatternsOk returns a tuple with the TagValuePatterns field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetTagValuePatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValuePatterns, true
}

// SetTagValuePatterns sets field value.
func (o *TagRuleAttributes) SetTagValuePatterns(v []string) {
	o.TagValuePatterns = v
}

// GetVersion returns the Version field value.
func (o *TagRuleAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagRuleAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *TagRuleAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagRuleAttributes) MarshalJSON() ([]byte, error) {
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
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.DeletedBy.IsSet() {
		toSerialize["deleted_by"] = o.DeletedBy.Get()
	}
	toSerialize["enabled"] = o.Enabled
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["negated"] = o.Negated
	toSerialize["required"] = o.Required
	toSerialize["rule_type"] = o.RuleType
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
func (o *TagRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time             `json:"created_at"`
		CreatedBy        *string                `json:"created_by"`
		DeletedAt        datadog.NullableTime   `json:"deleted_at,omitempty"`
		DeletedBy        datadog.NullableString `json:"deleted_by,omitempty"`
		Enabled          *bool                  `json:"enabled"`
		ModifiedAt       *time.Time             `json:"modified_at"`
		ModifiedBy       *string                `json:"modified_by"`
		Name             *string                `json:"name"`
		Negated          *bool                  `json:"negated"`
		Required         *bool                  `json:"required"`
		RuleType         *TagRuleType           `json:"rule_type"`
		Scope            *string                `json:"scope"`
		Source           *TagRuleSource         `json:"source"`
		TagKey           *string                `json:"tag_key"`
		TagValuePatterns *[]string              `json:"tag_value_patterns"`
		Version          *int64                 `json:"version"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Negated == nil {
		return fmt.Errorf("required field negated missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.RuleType == nil {
		return fmt.Errorf("required field rule_type missing")
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "deleted_at", "deleted_by", "enabled", "modified_at", "modified_by", "name", "negated", "required", "rule_type", "scope", "source", "tag_key", "tag_value_patterns", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.DeletedAt = all.DeletedAt
	o.DeletedBy = all.DeletedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.Negated = *all.Negated
	o.Required = *all.Required
	if !all.RuleType.IsValid() {
		hasInvalidField = true
	} else {
		o.RuleType = *all.RuleType
	}
	o.Scope = *all.Scope
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.TagKey = *all.TagKey
	o.TagValuePatterns = *all.TagValuePatterns
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
