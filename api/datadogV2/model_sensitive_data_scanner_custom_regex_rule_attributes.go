// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerCustomRegexRuleAttributes Attributes of a custom rule, which matches the regular expression in `pattern`.
type SensitiveDataScannerCustomRegexRuleAttributes struct {
	// Description of the rule.
	Description *string `json:"description,omitempty"`
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	ExcludedNamespaces []string `json:"excluded_namespaces,omitempty"`
	// Object defining a set of keywords and a number of characters that help reduce noise.
	// You can provide a list of keywords you would like to check within a defined proximity of the matching pattern.
	// If any of the keywords are found within the proximity check, the match is kept.
	// If none are found, the match is discarded.
	IncludedKeywordConfiguration *SensitiveDataScannerIncludedKeywordConfiguration `json:"included_keyword_configuration,omitempty"`
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned.
	// If both are missing the whole event is scanned.
	Namespaces []string `json:"namespaces,omitempty"`
	// Regular expression of the rule.
	Pattern string `json:"pattern"`
	// Severity of the issues found by the rule.
	Priority *SensitiveDataScannerRulePriority `json:"priority,omitempty"`
	// Type of a custom rule.
	RuleType SensitiveDataScannerCustomRegexRuleType `json:"rule_type"`
	// Object describing suppressions for a rule, or the recommended suppression lists published by a standard pattern.
	// There are three types of suppressions, `starts_with`, `ends_with`, and `exact_match`.
	// Rules based on the Datadog library can use recommended or custom suppressions. Custom rules can only use custom suppressions.
	// Suppressed matches are not obfuscated, counted in metrics, or displayed in the Findings page.
	Suppressions *SensitiveDataScannerSuppressions `json:"suppressions,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	// Object describing how the scanned event will be replaced.
	TextReplacement *SensitiveDataScannerTextReplacement `json:"text_replacement,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerCustomRegexRuleAttributes instantiates a new SensitiveDataScannerCustomRegexRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerCustomRegexRuleAttributes(pattern string, ruleType SensitiveDataScannerCustomRegexRuleType) *SensitiveDataScannerCustomRegexRuleAttributes {
	this := SensitiveDataScannerCustomRegexRuleAttributes{}
	this.Pattern = pattern
	this.RuleType = ruleType
	return &this
}

// NewSensitiveDataScannerCustomRegexRuleAttributesWithDefaults instantiates a new SensitiveDataScannerCustomRegexRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerCustomRegexRuleAttributesWithDefaults() *SensitiveDataScannerCustomRegexRuleAttributes {
	this := SensitiveDataScannerCustomRegexRuleAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExcludedNamespaces returns the ExcludedNamespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetExcludedNamespaces() []string {
	if o == nil || o.ExcludedNamespaces == nil {
		var ret []string
		return ret
	}
	return o.ExcludedNamespaces
}

// GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetExcludedNamespacesOk() (*[]string, bool) {
	if o == nil || o.ExcludedNamespaces == nil {
		return nil, false
	}
	return &o.ExcludedNamespaces, true
}

// HasExcludedNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasExcludedNamespaces() bool {
	return o != nil && o.ExcludedNamespaces != nil
}

// SetExcludedNamespaces gets a reference to the given []string and assigns it to the ExcludedNamespaces field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetExcludedNamespaces(v []string) {
	o.ExcludedNamespaces = v
}

// GetIncludedKeywordConfiguration returns the IncludedKeywordConfiguration field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetIncludedKeywordConfiguration() SensitiveDataScannerIncludedKeywordConfiguration {
	if o == nil || o.IncludedKeywordConfiguration == nil {
		var ret SensitiveDataScannerIncludedKeywordConfiguration
		return ret
	}
	return *o.IncludedKeywordConfiguration
}

// GetIncludedKeywordConfigurationOk returns a tuple with the IncludedKeywordConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetIncludedKeywordConfigurationOk() (*SensitiveDataScannerIncludedKeywordConfiguration, bool) {
	if o == nil || o.IncludedKeywordConfiguration == nil {
		return nil, false
	}
	return o.IncludedKeywordConfiguration, true
}

// HasIncludedKeywordConfiguration returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasIncludedKeywordConfiguration() bool {
	return o != nil && o.IncludedKeywordConfiguration != nil
}

// SetIncludedKeywordConfiguration gets a reference to the given SensitiveDataScannerIncludedKeywordConfiguration and assigns it to the IncludedKeywordConfiguration field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetIncludedKeywordConfiguration(v SensitiveDataScannerIncludedKeywordConfiguration) {
	o.IncludedKeywordConfiguration = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetPattern returns the Pattern field value.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetPattern(v string) {
	o.Pattern = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetPriority() SensitiveDataScannerRulePriority {
	if o == nil || o.Priority == nil {
		var ret SensitiveDataScannerRulePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetPriorityOk() (*SensitiveDataScannerRulePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given SensitiveDataScannerRulePriority and assigns it to the Priority field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetPriority(v SensitiveDataScannerRulePriority) {
	o.Priority = &v
}

// GetRuleType returns the RuleType field value.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetRuleType() SensitiveDataScannerCustomRegexRuleType {
	if o == nil {
		var ret SensitiveDataScannerCustomRegexRuleType
		return ret
	}
	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetRuleTypeOk() (*SensitiveDataScannerCustomRegexRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetRuleType(v SensitiveDataScannerCustomRegexRuleType) {
	o.RuleType = v
}

// GetSuppressions returns the Suppressions field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetSuppressions() SensitiveDataScannerSuppressions {
	if o == nil || o.Suppressions == nil {
		var ret SensitiveDataScannerSuppressions
		return ret
	}
	return *o.Suppressions
}

// GetSuppressionsOk returns a tuple with the Suppressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetSuppressionsOk() (*SensitiveDataScannerSuppressions, bool) {
	if o == nil || o.Suppressions == nil {
		return nil, false
	}
	return o.Suppressions, true
}

// HasSuppressions returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasSuppressions() bool {
	return o != nil && o.Suppressions != nil
}

// SetSuppressions gets a reference to the given SensitiveDataScannerSuppressions and assigns it to the Suppressions field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetSuppressions(v SensitiveDataScannerSuppressions) {
	o.Suppressions = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTextReplacement returns the TextReplacement field value if set, zero value otherwise.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetTextReplacement() SensitiveDataScannerTextReplacement {
	if o == nil || o.TextReplacement == nil {
		var ret SensitiveDataScannerTextReplacement
		return ret
	}
	return *o.TextReplacement
}

// GetTextReplacementOk returns a tuple with the TextReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) GetTextReplacementOk() (*SensitiveDataScannerTextReplacement, bool) {
	if o == nil || o.TextReplacement == nil {
		return nil, false
	}
	return o.TextReplacement, true
}

// HasTextReplacement returns a boolean if a field has been set.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) HasTextReplacement() bool {
	return o != nil && o.TextReplacement != nil
}

// SetTextReplacement gets a reference to the given SensitiveDataScannerTextReplacement and assigns it to the TextReplacement field.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) SetTextReplacement(v SensitiveDataScannerTextReplacement) {
	o.TextReplacement = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerCustomRegexRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExcludedNamespaces != nil {
		toSerialize["excluded_namespaces"] = o.ExcludedNamespaces
	}
	if o.IncludedKeywordConfiguration != nil {
		toSerialize["included_keyword_configuration"] = o.IncludedKeywordConfiguration
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	toSerialize["pattern"] = o.Pattern
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["rule_type"] = o.RuleType
	if o.Suppressions != nil {
		toSerialize["suppressions"] = o.Suppressions
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TextReplacement != nil {
		toSerialize["text_replacement"] = o.TextReplacement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerCustomRegexRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description                  *string                                           `json:"description,omitempty"`
		ExcludedNamespaces           []string                                          `json:"excluded_namespaces,omitempty"`
		IncludedKeywordConfiguration *SensitiveDataScannerIncludedKeywordConfiguration `json:"included_keyword_configuration,omitempty"`
		IsEnabled                    *bool                                             `json:"is_enabled,omitempty"`
		Name                         *string                                           `json:"name,omitempty"`
		Namespaces                   []string                                          `json:"namespaces,omitempty"`
		Pattern                      *string                                           `json:"pattern"`
		Priority                     *SensitiveDataScannerRulePriority                 `json:"priority,omitempty"`
		RuleType                     *SensitiveDataScannerCustomRegexRuleType          `json:"rule_type"`
		Suppressions                 *SensitiveDataScannerSuppressions                 `json:"suppressions,omitempty"`
		Tags                         []string                                          `json:"tags,omitempty"`
		TextReplacement              *SensitiveDataScannerTextReplacement              `json:"text_replacement,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	if all.RuleType == nil {
		return fmt.Errorf("required field rule_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "excluded_namespaces", "included_keyword_configuration", "is_enabled", "name", "namespaces", "pattern", "priority", "rule_type", "suppressions", "tags", "text_replacement"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.ExcludedNamespaces = all.ExcludedNamespaces
	if all.IncludedKeywordConfiguration != nil && all.IncludedKeywordConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncludedKeywordConfiguration = all.IncludedKeywordConfiguration
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Namespaces = all.Namespaces
	o.Pattern = *all.Pattern
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	if !all.RuleType.IsValid() {
		hasInvalidField = true
	} else {
		o.RuleType = *all.RuleType
	}
	if all.Suppressions != nil && all.Suppressions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Suppressions = all.Suppressions
	o.Tags = all.Tags
	if all.TextReplacement != nil && all.TextReplacement.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TextReplacement = all.TextReplacement

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
