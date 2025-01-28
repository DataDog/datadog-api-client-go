// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRule The definition of an automation pipeline rule scope.
// A rule can act on specific issue types, security rule types, security rule IDs, rule severities, or a query.
// The query can be used to filter resources on tags and attributes.
// The issue type and rule types fields are required.
type AutomationRule struct {
	// The type of issues on which the rule applies
	IssueType IssueType `json:"issue_type"`
	// The query is composed of one or several key:value pairs, which can be used to filter resources on tags and attributes.
	Query *string `json:"query,omitempty"`
	// Security rule ids
	RuleIds []string `json:"rule_ids,omitempty"`
	// Security rule types
	RuleTypes []SecurityRuleTypesItems `json:"rule_types"`
	// The security rules severities to consider
	Severities []SecurityRuleSeverity `json:"severities,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRule instantiates a new AutomationRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRule(issueType IssueType, ruleTypes []SecurityRuleTypesItems) *AutomationRule {
	this := AutomationRule{}
	this.IssueType = issueType
	this.RuleTypes = ruleTypes
	return &this
}

// NewAutomationRuleWithDefaults instantiates a new AutomationRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleWithDefaults() *AutomationRule {
	this := AutomationRule{}
	return &this
}

// GetIssueType returns the IssueType field value.
func (o *AutomationRule) GetIssueType() IssueType {
	if o == nil {
		var ret IssueType
		return ret
	}
	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *AutomationRule) GetIssueTypeOk() (*IssueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value.
func (o *AutomationRule) SetIssueType(v IssueType) {
	o.IssueType = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *AutomationRule) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRule) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *AutomationRule) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *AutomationRule) SetQuery(v string) {
	o.Query = &v
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *AutomationRule) GetRuleIds() []string {
	if o == nil || o.RuleIds == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRule) GetRuleIdsOk() (*[]string, bool) {
	if o == nil || o.RuleIds == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *AutomationRule) HasRuleIds() bool {
	return o != nil && o.RuleIds != nil
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *AutomationRule) SetRuleIds(v []string) {
	o.RuleIds = v
}

// GetRuleTypes returns the RuleTypes field value.
func (o *AutomationRule) GetRuleTypes() []SecurityRuleTypesItems {
	if o == nil {
		var ret []SecurityRuleTypesItems
		return ret
	}
	return o.RuleTypes
}

// GetRuleTypesOk returns a tuple with the RuleTypes field value
// and a boolean to check if the value has been set.
func (o *AutomationRule) GetRuleTypesOk() (*[]SecurityRuleTypesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleTypes, true
}

// SetRuleTypes sets field value.
func (o *AutomationRule) SetRuleTypes(v []SecurityRuleTypesItems) {
	o.RuleTypes = v
}

// GetSeverities returns the Severities field value if set, zero value otherwise.
func (o *AutomationRule) GetSeverities() []SecurityRuleSeverity {
	if o == nil || o.Severities == nil {
		var ret []SecurityRuleSeverity
		return ret
	}
	return o.Severities
}

// GetSeveritiesOk returns a tuple with the Severities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRule) GetSeveritiesOk() (*[]SecurityRuleSeverity, bool) {
	if o == nil || o.Severities == nil {
		return nil, false
	}
	return &o.Severities, true
}

// HasSeverities returns a boolean if a field has been set.
func (o *AutomationRule) HasSeverities() bool {
	return o != nil && o.Severities != nil
}

// SetSeverities gets a reference to the given []SecurityRuleSeverity and assigns it to the Severities field.
func (o *AutomationRule) SetSeverities(v []SecurityRuleSeverity) {
	o.Severities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["issue_type"] = o.IssueType
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleIds != nil {
		toSerialize["rule_ids"] = o.RuleIds
	}
	toSerialize["rule_types"] = o.RuleTypes
	if o.Severities != nil {
		toSerialize["severities"] = o.Severities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IssueType  *IssueType                `json:"issue_type"`
		Query      *string                   `json:"query,omitempty"`
		RuleIds    []string                  `json:"rule_ids,omitempty"`
		RuleTypes  *[]SecurityRuleTypesItems `json:"rule_types"`
		Severities []SecurityRuleSeverity    `json:"severities,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IssueType == nil {
		return fmt.Errorf("required field issue_type missing")
	}
	if all.RuleTypes == nil {
		return fmt.Errorf("required field rule_types missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issue_type", "query", "rule_ids", "rule_types", "severities"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.IssueType.IsValid() {
		hasInvalidField = true
	} else {
		o.IssueType = *all.IssueType
	}
	o.Query = all.Query
	o.RuleIds = all.RuleIds
	o.RuleTypes = *all.RuleTypes
	o.Severities = all.Severities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
