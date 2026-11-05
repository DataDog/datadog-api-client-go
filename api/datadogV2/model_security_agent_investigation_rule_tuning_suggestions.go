// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationRuleTuningSuggestions Suggestions to tune the detection rule that generated the signal.
type SecurityAgentInvestigationRuleTuningSuggestions struct {
	// Length of the window, in days.
	OccurrenceWindowDays int64 `json:"occurrence_window_days"`
	// Why tuning the rule is suggested.
	Reasoning string `json:"reasoning"`
	// ID of the detection rule.
	RuleId string `json:"rule_id"`
	// Name of the detection rule.
	RuleName string `json:"rule_name"`
	// Number of signals the rule generated in the window.
	RuleSignalTotal int64 `json:"rule_signal_total"`
	// The tuning suggestions.
	Suggestions []SecurityAgentInvestigationRuleTuningSuggestion `json:"suggestions"`
	// Summary of the suggestions.
	Summary string `json:"summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationRuleTuningSuggestions instantiates a new SecurityAgentInvestigationRuleTuningSuggestions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationRuleTuningSuggestions(occurrenceWindowDays int64, reasoning string, ruleId string, ruleName string, ruleSignalTotal int64, suggestions []SecurityAgentInvestigationRuleTuningSuggestion, summary string) *SecurityAgentInvestigationRuleTuningSuggestions {
	this := SecurityAgentInvestigationRuleTuningSuggestions{}
	this.OccurrenceWindowDays = occurrenceWindowDays
	this.Reasoning = reasoning
	this.RuleId = ruleId
	this.RuleName = ruleName
	this.RuleSignalTotal = ruleSignalTotal
	this.Suggestions = suggestions
	this.Summary = summary
	return &this
}

// NewSecurityAgentInvestigationRuleTuningSuggestionsWithDefaults instantiates a new SecurityAgentInvestigationRuleTuningSuggestions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationRuleTuningSuggestionsWithDefaults() *SecurityAgentInvestigationRuleTuningSuggestions {
	this := SecurityAgentInvestigationRuleTuningSuggestions{}
	return &this
}

// GetOccurrenceWindowDays returns the OccurrenceWindowDays field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetOccurrenceWindowDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OccurrenceWindowDays
}

// GetOccurrenceWindowDaysOk returns a tuple with the OccurrenceWindowDays field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetOccurrenceWindowDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccurrenceWindowDays, true
}

// SetOccurrenceWindowDays sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetOccurrenceWindowDays(v int64) {
	o.OccurrenceWindowDays = v
}

// GetReasoning returns the Reasoning field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetReasoning() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetReasoningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reasoning, true
}

// SetReasoning sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetReasoning(v string) {
	o.Reasoning = v
}

// GetRuleId returns the RuleId field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetRuleId(v string) {
	o.RuleId = v
}

// GetRuleName returns the RuleName field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleName, true
}

// SetRuleName sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetRuleName(v string) {
	o.RuleName = v
}

// GetRuleSignalTotal returns the RuleSignalTotal field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleSignalTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RuleSignalTotal
}

// GetRuleSignalTotalOk returns a tuple with the RuleSignalTotal field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetRuleSignalTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleSignalTotal, true
}

// SetRuleSignalTotal sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetRuleSignalTotal(v int64) {
	o.RuleSignalTotal = v
}

// GetSuggestions returns the Suggestions field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetSuggestions() []SecurityAgentInvestigationRuleTuningSuggestion {
	if o == nil {
		var ret []SecurityAgentInvestigationRuleTuningSuggestion
		return ret
	}
	return o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetSuggestionsOk() (*[]SecurityAgentInvestigationRuleTuningSuggestion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestions, true
}

// SetSuggestions sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetSuggestions(v []SecurityAgentInvestigationRuleTuningSuggestion) {
	o.Suggestions = v
}

// GetSummary returns the Summary field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) SetSummary(v string) {
	o.Summary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationRuleTuningSuggestions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["occurrence_window_days"] = o.OccurrenceWindowDays
	toSerialize["reasoning"] = o.Reasoning
	toSerialize["rule_id"] = o.RuleId
	toSerialize["rule_name"] = o.RuleName
	toSerialize["rule_signal_total"] = o.RuleSignalTotal
	toSerialize["suggestions"] = o.Suggestions
	toSerialize["summary"] = o.Summary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationRuleTuningSuggestions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OccurrenceWindowDays *int64                                            `json:"occurrence_window_days"`
		Reasoning            *string                                           `json:"reasoning"`
		RuleId               *string                                           `json:"rule_id"`
		RuleName             *string                                           `json:"rule_name"`
		RuleSignalTotal      *int64                                            `json:"rule_signal_total"`
		Suggestions          *[]SecurityAgentInvestigationRuleTuningSuggestion `json:"suggestions"`
		Summary              *string                                           `json:"summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OccurrenceWindowDays == nil {
		return fmt.Errorf("required field occurrence_window_days missing")
	}
	if all.Reasoning == nil {
		return fmt.Errorf("required field reasoning missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.RuleName == nil {
		return fmt.Errorf("required field rule_name missing")
	}
	if all.RuleSignalTotal == nil {
		return fmt.Errorf("required field rule_signal_total missing")
	}
	if all.Suggestions == nil {
		return fmt.Errorf("required field suggestions missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"occurrence_window_days", "reasoning", "rule_id", "rule_name", "rule_signal_total", "suggestions", "summary"})
	} else {
		return err
	}
	o.OccurrenceWindowDays = *all.OccurrenceWindowDays
	o.Reasoning = *all.Reasoning
	o.RuleId = *all.RuleId
	o.RuleName = *all.RuleName
	o.RuleSignalTotal = *all.RuleSignalTotal
	o.Suggestions = *all.Suggestions
	o.Summary = *all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
