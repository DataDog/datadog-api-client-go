// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerMachineLearningRuleAttributes Attributes of a rule based on a standard machine learning rule. It has no pattern or keywords.
type SensitiveDataScannerMachineLearningRuleAttributes struct {
	// Description of the rule.
	Description *string `json:"description,omitempty"`
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	ExcludedNamespaces []string `json:"excluded_namespaces,omitempty"`
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned.
	// If both are missing the whole event is scanned.
	Namespaces []string `json:"namespaces,omitempty"`
	// Severity of the issues found by the rule.
	Priority *SensitiveDataScannerRulePriority `json:"priority,omitempty"`
	// Type of a rule based on a standard machine learning rule.
	RuleType SensitiveDataScannerMachineLearningRuleType `json:"rule_type"`
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

// NewSensitiveDataScannerMachineLearningRuleAttributes instantiates a new SensitiveDataScannerMachineLearningRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerMachineLearningRuleAttributes(ruleType SensitiveDataScannerMachineLearningRuleType) *SensitiveDataScannerMachineLearningRuleAttributes {
	this := SensitiveDataScannerMachineLearningRuleAttributes{}
	this.RuleType = ruleType
	return &this
}

// NewSensitiveDataScannerMachineLearningRuleAttributesWithDefaults instantiates a new SensitiveDataScannerMachineLearningRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerMachineLearningRuleAttributesWithDefaults() *SensitiveDataScannerMachineLearningRuleAttributes {
	this := SensitiveDataScannerMachineLearningRuleAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExcludedNamespaces returns the ExcludedNamespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetExcludedNamespaces() []string {
	if o == nil || o.ExcludedNamespaces == nil {
		var ret []string
		return ret
	}
	return o.ExcludedNamespaces
}

// GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetExcludedNamespacesOk() (*[]string, bool) {
	if o == nil || o.ExcludedNamespaces == nil {
		return nil, false
	}
	return &o.ExcludedNamespaces, true
}

// HasExcludedNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasExcludedNamespaces() bool {
	return o != nil && o.ExcludedNamespaces != nil
}

// SetExcludedNamespaces gets a reference to the given []string and assigns it to the ExcludedNamespaces field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetExcludedNamespaces(v []string) {
	o.ExcludedNamespaces = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetPriority() SensitiveDataScannerRulePriority {
	if o == nil || o.Priority == nil {
		var ret SensitiveDataScannerRulePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetPriorityOk() (*SensitiveDataScannerRulePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given SensitiveDataScannerRulePriority and assigns it to the Priority field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetPriority(v SensitiveDataScannerRulePriority) {
	o.Priority = &v
}

// GetRuleType returns the RuleType field value.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetRuleType() SensitiveDataScannerMachineLearningRuleType {
	if o == nil {
		var ret SensitiveDataScannerMachineLearningRuleType
		return ret
	}
	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetRuleTypeOk() (*SensitiveDataScannerMachineLearningRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetRuleType(v SensitiveDataScannerMachineLearningRuleType) {
	o.RuleType = v
}

// GetSuppressions returns the Suppressions field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetSuppressions() SensitiveDataScannerSuppressions {
	if o == nil || o.Suppressions == nil {
		var ret SensitiveDataScannerSuppressions
		return ret
	}
	return *o.Suppressions
}

// GetSuppressionsOk returns a tuple with the Suppressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetSuppressionsOk() (*SensitiveDataScannerSuppressions, bool) {
	if o == nil || o.Suppressions == nil {
		return nil, false
	}
	return o.Suppressions, true
}

// HasSuppressions returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasSuppressions() bool {
	return o != nil && o.Suppressions != nil
}

// SetSuppressions gets a reference to the given SensitiveDataScannerSuppressions and assigns it to the Suppressions field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetSuppressions(v SensitiveDataScannerSuppressions) {
	o.Suppressions = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTextReplacement returns the TextReplacement field value if set, zero value otherwise.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetTextReplacement() SensitiveDataScannerTextReplacement {
	if o == nil || o.TextReplacement == nil {
		var ret SensitiveDataScannerTextReplacement
		return ret
	}
	return *o.TextReplacement
}

// GetTextReplacementOk returns a tuple with the TextReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) GetTextReplacementOk() (*SensitiveDataScannerTextReplacement, bool) {
	if o == nil || o.TextReplacement == nil {
		return nil, false
	}
	return o.TextReplacement, true
}

// HasTextReplacement returns a boolean if a field has been set.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) HasTextReplacement() bool {
	return o != nil && o.TextReplacement != nil
}

// SetTextReplacement gets a reference to the given SensitiveDataScannerTextReplacement and assigns it to the TextReplacement field.
func (o *SensitiveDataScannerMachineLearningRuleAttributes) SetTextReplacement(v SensitiveDataScannerTextReplacement) {
	o.TextReplacement = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerMachineLearningRuleAttributes) MarshalJSON() ([]byte, error) {
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
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
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
func (o *SensitiveDataScannerMachineLearningRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description        *string                                      `json:"description,omitempty"`
		ExcludedNamespaces []string                                     `json:"excluded_namespaces,omitempty"`
		IsEnabled          *bool                                        `json:"is_enabled,omitempty"`
		Name               *string                                      `json:"name,omitempty"`
		Namespaces         []string                                     `json:"namespaces,omitempty"`
		Priority           *SensitiveDataScannerRulePriority            `json:"priority,omitempty"`
		RuleType           *SensitiveDataScannerMachineLearningRuleType `json:"rule_type"`
		Suppressions       *SensitiveDataScannerSuppressions            `json:"suppressions,omitempty"`
		Tags               []string                                     `json:"tags,omitempty"`
		TextReplacement    *SensitiveDataScannerTextReplacement         `json:"text_replacement,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RuleType == nil {
		return fmt.Errorf("required field rule_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "excluded_namespaces", "is_enabled", "name", "namespaces", "priority", "rule_type", "suppressions", "tags", "text_replacement"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.ExcludedNamespaces = all.ExcludedNamespaces
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Namespaces = all.Namespaces
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
