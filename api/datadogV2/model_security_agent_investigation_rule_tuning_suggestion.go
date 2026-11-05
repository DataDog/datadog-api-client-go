// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationRuleTuningSuggestion A suggestion to tune the detection rule.
type SecurityAgentInvestigationRuleTuningSuggestion struct {
	// Kind of change the suggestion makes to the rule.
	ActionType string `json:"action_type"`
	// The conditions the change applies to.
	Conditions []SecurityAgentInvestigationRuleTuningCondition `json:"conditions"`
	// What the suggestion changes.
	Description string `json:"description"`
	// Estimated number of signals the change would have silenced in the window.
	EstimatedSilencedCount *int64 `json:"estimated_silenced_count,omitempty"`
	// ID of the suggestion.
	Id string `json:"id"`
	// How the conditions combine, such as AND or OR.
	MatchLogic string `json:"match_logic"`
	// Whether the suggestion is recommended.
	Recommended bool `json:"recommended"`
	// ID of the detection rule.
	RuleId string `json:"rule_id"`
	// Title of the suggestion.
	Title string `json:"title"`
	// What the change could hide. It can be empty.
	Warning string `json:"warning"`
	// Why the suggestion is made.
	Why string `json:"why"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationRuleTuningSuggestion instantiates a new SecurityAgentInvestigationRuleTuningSuggestion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationRuleTuningSuggestion(actionType string, conditions []SecurityAgentInvestigationRuleTuningCondition, description string, id string, matchLogic string, recommended bool, ruleId string, title string, warning string, why string) *SecurityAgentInvestigationRuleTuningSuggestion {
	this := SecurityAgentInvestigationRuleTuningSuggestion{}
	this.ActionType = actionType
	this.Conditions = conditions
	this.Description = description
	this.Id = id
	this.MatchLogic = matchLogic
	this.Recommended = recommended
	this.RuleId = ruleId
	this.Title = title
	this.Warning = warning
	this.Why = why
	return &this
}

// NewSecurityAgentInvestigationRuleTuningSuggestionWithDefaults instantiates a new SecurityAgentInvestigationRuleTuningSuggestion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationRuleTuningSuggestionWithDefaults() *SecurityAgentInvestigationRuleTuningSuggestion {
	this := SecurityAgentInvestigationRuleTuningSuggestion{}
	return &this
}

// GetActionType returns the ActionType field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetActionType(v string) {
	o.ActionType = v
}

// GetConditions returns the Conditions field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetConditions() []SecurityAgentInvestigationRuleTuningCondition {
	if o == nil {
		var ret []SecurityAgentInvestigationRuleTuningCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetConditionsOk() (*[]SecurityAgentInvestigationRuleTuningCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetConditions(v []SecurityAgentInvestigationRuleTuningCondition) {
	o.Conditions = v
}

// GetDescription returns the Description field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetDescription(v string) {
	o.Description = v
}

// GetEstimatedSilencedCount returns the EstimatedSilencedCount field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetEstimatedSilencedCount() int64 {
	if o == nil || o.EstimatedSilencedCount == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedSilencedCount
}

// GetEstimatedSilencedCountOk returns a tuple with the EstimatedSilencedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetEstimatedSilencedCountOk() (*int64, bool) {
	if o == nil || o.EstimatedSilencedCount == nil {
		return nil, false
	}
	return o.EstimatedSilencedCount, true
}

// HasEstimatedSilencedCount returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) HasEstimatedSilencedCount() bool {
	return o != nil && o.EstimatedSilencedCount != nil
}

// SetEstimatedSilencedCount gets a reference to the given int64 and assigns it to the EstimatedSilencedCount field.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetEstimatedSilencedCount(v int64) {
	o.EstimatedSilencedCount = &v
}

// GetId returns the Id field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetId(v string) {
	o.Id = v
}

// GetMatchLogic returns the MatchLogic field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetMatchLogic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MatchLogic
}

// GetMatchLogicOk returns a tuple with the MatchLogic field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetMatchLogicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchLogic, true
}

// SetMatchLogic sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetMatchLogic(v string) {
	o.MatchLogic = v
}

// GetRecommended returns the Recommended field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetRecommended() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetRecommendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetRecommended(v bool) {
	o.Recommended = v
}

// GetRuleId returns the RuleId field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetRuleId(v string) {
	o.RuleId = v
}

// GetTitle returns the Title field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetTitle(v string) {
	o.Title = v
}

// GetWarning returns the Warning field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetWarning() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetWarning(v string) {
	o.Warning = v
}

// GetWhy returns the Why field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetWhy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Why
}

// GetWhyOk returns a tuple with the Why field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) GetWhyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Why, true
}

// SetWhy sets field value.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) SetWhy(v string) {
	o.Why = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationRuleTuningSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action_type"] = o.ActionType
	toSerialize["conditions"] = o.Conditions
	toSerialize["description"] = o.Description
	if o.EstimatedSilencedCount != nil {
		toSerialize["estimated_silenced_count"] = o.EstimatedSilencedCount
	}
	toSerialize["id"] = o.Id
	toSerialize["match_logic"] = o.MatchLogic
	toSerialize["recommended"] = o.Recommended
	toSerialize["rule_id"] = o.RuleId
	toSerialize["title"] = o.Title
	toSerialize["warning"] = o.Warning
	toSerialize["why"] = o.Why

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationRuleTuningSuggestion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionType             *string                                          `json:"action_type"`
		Conditions             *[]SecurityAgentInvestigationRuleTuningCondition `json:"conditions"`
		Description            *string                                          `json:"description"`
		EstimatedSilencedCount *int64                                           `json:"estimated_silenced_count,omitempty"`
		Id                     *string                                          `json:"id"`
		MatchLogic             *string                                          `json:"match_logic"`
		Recommended            *bool                                            `json:"recommended"`
		RuleId                 *string                                          `json:"rule_id"`
		Title                  *string                                          `json:"title"`
		Warning                *string                                          `json:"warning"`
		Why                    *string                                          `json:"why"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionType == nil {
		return fmt.Errorf("required field action_type missing")
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.MatchLogic == nil {
		return fmt.Errorf("required field match_logic missing")
	}
	if all.Recommended == nil {
		return fmt.Errorf("required field recommended missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Warning == nil {
		return fmt.Errorf("required field warning missing")
	}
	if all.Why == nil {
		return fmt.Errorf("required field why missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_type", "conditions", "description", "estimated_silenced_count", "id", "match_logic", "recommended", "rule_id", "title", "warning", "why"})
	} else {
		return err
	}
	o.ActionType = *all.ActionType
	o.Conditions = *all.Conditions
	o.Description = *all.Description
	o.EstimatedSilencedCount = all.EstimatedSilencedCount
	o.Id = *all.Id
	o.MatchLogic = *all.MatchLogic
	o.Recommended = *all.Recommended
	o.RuleId = *all.RuleId
	o.Title = *all.Title
	o.Warning = *all.Warning
	o.Why = *all.Why

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
