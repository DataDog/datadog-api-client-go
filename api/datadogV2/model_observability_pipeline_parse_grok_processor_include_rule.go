// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorIncludeRule A Grok parsing rule selected using the `include` query. Each rule defines how to extract structured fields
// from logs matching a Datadog search query.
type ObservabilityPipelineParseGrokProcessorIncludeRule struct {
	// A Datadog search query used to determine which logs this Grok rule targets.
	Include string `json:"include"`
	// A list of Grok parsing rules that define how to extract fields from matching logs.
	// Each rule must contain a name and a valid Grok pattern.
	MatchRules []ObservabilityPipelineParseGrokProcessorRuleMatchRule `json:"match_rules"`
	// A list of Grok helper rules that can be referenced by the parsing rules.
	SupportRules []ObservabilityPipelineParseGrokProcessorRuleSupportRule `json:"support_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseGrokProcessorIncludeRule instantiates a new ObservabilityPipelineParseGrokProcessorIncludeRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseGrokProcessorIncludeRule(include string, matchRules []ObservabilityPipelineParseGrokProcessorRuleMatchRule) *ObservabilityPipelineParseGrokProcessorIncludeRule {
	this := ObservabilityPipelineParseGrokProcessorIncludeRule{}
	this.Include = include
	this.MatchRules = matchRules
	return &this
}

// NewObservabilityPipelineParseGrokProcessorIncludeRuleWithDefaults instantiates a new ObservabilityPipelineParseGrokProcessorIncludeRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseGrokProcessorIncludeRuleWithDefaults() *ObservabilityPipelineParseGrokProcessorIncludeRule {
	this := ObservabilityPipelineParseGrokProcessorIncludeRule{}
	return &this
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) SetInclude(v string) {
	o.Include = v
}

// GetMatchRules returns the MatchRules field value.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetMatchRules() []ObservabilityPipelineParseGrokProcessorRuleMatchRule {
	if o == nil {
		var ret []ObservabilityPipelineParseGrokProcessorRuleMatchRule
		return ret
	}
	return o.MatchRules
}

// GetMatchRulesOk returns a tuple with the MatchRules field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetMatchRulesOk() (*[]ObservabilityPipelineParseGrokProcessorRuleMatchRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchRules, true
}

// SetMatchRules sets field value.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) SetMatchRules(v []ObservabilityPipelineParseGrokProcessorRuleMatchRule) {
	o.MatchRules = v
}

// GetSupportRules returns the SupportRules field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetSupportRules() []ObservabilityPipelineParseGrokProcessorRuleSupportRule {
	if o == nil || o.SupportRules == nil {
		var ret []ObservabilityPipelineParseGrokProcessorRuleSupportRule
		return ret
	}
	return o.SupportRules
}

// GetSupportRulesOk returns a tuple with the SupportRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) GetSupportRulesOk() (*[]ObservabilityPipelineParseGrokProcessorRuleSupportRule, bool) {
	if o == nil || o.SupportRules == nil {
		return nil, false
	}
	return &o.SupportRules, true
}

// HasSupportRules returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) HasSupportRules() bool {
	return o != nil && o.SupportRules != nil
}

// SetSupportRules gets a reference to the given []ObservabilityPipelineParseGrokProcessorRuleSupportRule and assigns it to the SupportRules field.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) SetSupportRules(v []ObservabilityPipelineParseGrokProcessorRuleSupportRule) {
	o.SupportRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseGrokProcessorIncludeRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include"] = o.Include
	toSerialize["match_rules"] = o.MatchRules
	if o.SupportRules != nil {
		toSerialize["support_rules"] = o.SupportRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineParseGrokProcessorIncludeRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Include      *string                                                  `json:"include"`
		MatchRules   *[]ObservabilityPipelineParseGrokProcessorRuleMatchRule  `json:"match_rules"`
		SupportRules []ObservabilityPipelineParseGrokProcessorRuleSupportRule `json:"support_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.MatchRules == nil {
		return fmt.Errorf("required field match_rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include", "match_rules", "support_rules"})
	} else {
		return err
	}
	o.Include = *all.Include
	o.MatchRules = *all.MatchRules
	o.SupportRules = all.SupportRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
