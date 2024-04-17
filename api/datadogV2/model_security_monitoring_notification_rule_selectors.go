// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringNotificationRuleSelectors Selectors describing the notification rule.
type SecurityMonitoringNotificationRuleSelectors struct {
	// Set of rule and signal tags for which a notification will be triggered.
	Attributes []string `json:"attributes"`
	// Whether vulnerability_management rules are matched by default when the selector for rule type is empty.
	ImplicitVmRuleMatch bool `json:"implicit_vm_rule_match"`
	// True if the notification_rule was created with signal tags/attributes and migrated to an event query and false if it was created with an event query
	Migrated *bool `json:"migrated,omitempty"`
	// Query for matching the notification_rule against signal content
	Query *string `json:"query,omitempty"`
	// Set of rule tags for which a notification will be triggered.
	RuleTags []string `json:"rule_tags"`
	// Set of signal types (rule types) for which a notification will be triggered.
	RuleTypes []SecurityMonitoringRuleTypes `json:"rule_types"`
	// Set of signal severities (rule case statuses) for which a notification will be triggered.
	Severities []SecurityMonitoringRuleSeverity `json:"severities"`
	// Set of signal tags for which a notification will be triggered.
	SignalTags []string `json:"signal_tags"`
	// Specifies the evaluation result (Signal or Finding).
	TriggerSource *SecurityMonitoringNotificationRuleTriggerSource `json:"trigger_source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringNotificationRuleSelectors instantiates a new SecurityMonitoringNotificationRuleSelectors object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringNotificationRuleSelectors(attributes []string, implicitVmRuleMatch bool, ruleTags []string, ruleTypes []SecurityMonitoringRuleTypes, severities []SecurityMonitoringRuleSeverity, signalTags []string) *SecurityMonitoringNotificationRuleSelectors {
	this := SecurityMonitoringNotificationRuleSelectors{}
	this.Attributes = attributes
	this.ImplicitVmRuleMatch = implicitVmRuleMatch
	this.RuleTags = ruleTags
	this.RuleTypes = ruleTypes
	this.Severities = severities
	this.SignalTags = signalTags
	return &this
}

// NewSecurityMonitoringNotificationRuleSelectorsWithDefaults instantiates a new SecurityMonitoringNotificationRuleSelectors object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringNotificationRuleSelectorsWithDefaults() *SecurityMonitoringNotificationRuleSelectors {
	this := SecurityMonitoringNotificationRuleSelectors{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetAttributesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetAttributes(v []string) {
	o.Attributes = v
}

// GetImplicitVmRuleMatch returns the ImplicitVmRuleMatch field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetImplicitVmRuleMatch() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ImplicitVmRuleMatch
}

// GetImplicitVmRuleMatchOk returns a tuple with the ImplicitVmRuleMatch field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetImplicitVmRuleMatchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImplicitVmRuleMatch, true
}

// SetImplicitVmRuleMatch sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetImplicitVmRuleMatch(v bool) {
	o.ImplicitVmRuleMatch = v
}

// GetMigrated returns the Migrated field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleSelectors) GetMigrated() bool {
	if o == nil || o.Migrated == nil {
		var ret bool
		return ret
	}
	return *o.Migrated
}

// GetMigratedOk returns a tuple with the Migrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetMigratedOk() (*bool, bool) {
	if o == nil || o.Migrated == nil {
		return nil, false
	}
	return o.Migrated, true
}

// HasMigrated returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) HasMigrated() bool {
	return o != nil && o.Migrated != nil
}

// SetMigrated gets a reference to the given bool and assigns it to the Migrated field.
func (o *SecurityMonitoringNotificationRuleSelectors) SetMigrated(v bool) {
	o.Migrated = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleSelectors) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringNotificationRuleSelectors) SetQuery(v string) {
	o.Query = &v
}

// GetRuleTags returns the RuleTags field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetRuleTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RuleTags
}

// GetRuleTagsOk returns a tuple with the RuleTags field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetRuleTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleTags, true
}

// SetRuleTags sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetRuleTags(v []string) {
	o.RuleTags = v
}

// GetRuleTypes returns the RuleTypes field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetRuleTypes() []SecurityMonitoringRuleTypes {
	if o == nil {
		var ret []SecurityMonitoringRuleTypes
		return ret
	}
	return o.RuleTypes
}

// GetRuleTypesOk returns a tuple with the RuleTypes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetRuleTypesOk() (*[]SecurityMonitoringRuleTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleTypes, true
}

// SetRuleTypes sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetRuleTypes(v []SecurityMonitoringRuleTypes) {
	o.RuleTypes = v
}

// GetSeverities returns the Severities field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetSeverities() []SecurityMonitoringRuleSeverity {
	if o == nil {
		var ret []SecurityMonitoringRuleSeverity
		return ret
	}
	return o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetSeveritiesOk() (*[]SecurityMonitoringRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severities, true
}

// SetSeverities sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetSeverities(v []SecurityMonitoringRuleSeverity) {
	o.Severities = v
}

// GetSignalTags returns the SignalTags field value.
func (o *SecurityMonitoringNotificationRuleSelectors) GetSignalTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SignalTags
}

// GetSignalTagsOk returns a tuple with the SignalTags field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetSignalTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalTags, true
}

// SetSignalTags sets field value.
func (o *SecurityMonitoringNotificationRuleSelectors) SetSignalTags(v []string) {
	o.SignalTags = v
}

// GetTriggerSource returns the TriggerSource field value if set, zero value otherwise.
func (o *SecurityMonitoringNotificationRuleSelectors) GetTriggerSource() SecurityMonitoringNotificationRuleTriggerSource {
	if o == nil || o.TriggerSource == nil {
		var ret SecurityMonitoringNotificationRuleTriggerSource
		return ret
	}
	return *o.TriggerSource
}

// GetTriggerSourceOk returns a tuple with the TriggerSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) GetTriggerSourceOk() (*SecurityMonitoringNotificationRuleTriggerSource, bool) {
	if o == nil || o.TriggerSource == nil {
		return nil, false
	}
	return o.TriggerSource, true
}

// HasTriggerSource returns a boolean if a field has been set.
func (o *SecurityMonitoringNotificationRuleSelectors) HasTriggerSource() bool {
	return o != nil && o.TriggerSource != nil
}

// SetTriggerSource gets a reference to the given SecurityMonitoringNotificationRuleTriggerSource and assigns it to the TriggerSource field.
func (o *SecurityMonitoringNotificationRuleSelectors) SetTriggerSource(v SecurityMonitoringNotificationRuleTriggerSource) {
	o.TriggerSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringNotificationRuleSelectors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["implicit_vm_rule_match"] = o.ImplicitVmRuleMatch
	if o.Migrated != nil {
		toSerialize["migrated"] = o.Migrated
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["rule_tags"] = o.RuleTags
	toSerialize["rule_types"] = o.RuleTypes
	toSerialize["severities"] = o.Severities
	toSerialize["signal_tags"] = o.SignalTags
	if o.TriggerSource != nil {
		toSerialize["trigger_source"] = o.TriggerSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringNotificationRuleSelectors) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes          *[]string                                        `json:"attributes"`
		ImplicitVmRuleMatch *bool                                            `json:"implicit_vm_rule_match"`
		Migrated            *bool                                            `json:"migrated,omitempty"`
		Query               *string                                          `json:"query,omitempty"`
		RuleTags            *[]string                                        `json:"rule_tags"`
		RuleTypes           *[]SecurityMonitoringRuleTypes                   `json:"rule_types"`
		Severities          *[]SecurityMonitoringRuleSeverity                `json:"severities"`
		SignalTags          *[]string                                        `json:"signal_tags"`
		TriggerSource       *SecurityMonitoringNotificationRuleTriggerSource `json:"trigger_source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.ImplicitVmRuleMatch == nil {
		return fmt.Errorf("required field implicit_vm_rule_match missing")
	}
	if all.RuleTags == nil {
		return fmt.Errorf("required field rule_tags missing")
	}
	if all.RuleTypes == nil {
		return fmt.Errorf("required field rule_types missing")
	}
	if all.Severities == nil {
		return fmt.Errorf("required field severities missing")
	}
	if all.SignalTags == nil {
		return fmt.Errorf("required field signal_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "implicit_vm_rule_match", "migrated", "query", "rule_tags", "rule_types", "severities", "signal_tags", "trigger_source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attributes = *all.Attributes
	o.ImplicitVmRuleMatch = *all.ImplicitVmRuleMatch
	o.Migrated = all.Migrated
	o.Query = all.Query
	o.RuleTags = *all.RuleTags
	o.RuleTypes = *all.RuleTypes
	o.Severities = *all.Severities
	o.SignalTags = *all.SignalTags
	if all.TriggerSource != nil && !all.TriggerSource.IsValid() {
		hasInvalidField = true
	} else {
		o.TriggerSource = all.TriggerSource
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
