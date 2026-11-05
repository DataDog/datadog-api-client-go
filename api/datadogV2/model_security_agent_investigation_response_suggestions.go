// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationResponseSuggestions Suggested responses to the signal.
type SecurityAgentInvestigationResponseSuggestions struct {
	// Suggested mitigations, in the order to apply them.
	MitigationSuggestions []SecurityAgentInvestigationMitigationSuggestion `json:"mitigation_suggestions,omitempty"`
	// Suggestions to tune the detection rule that generated the signal.
	RuleTuningSuggestions *SecurityAgentInvestigationRuleTuningSuggestions `json:"rule_tuning_suggestions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationResponseSuggestions instantiates a new SecurityAgentInvestigationResponseSuggestions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationResponseSuggestions() *SecurityAgentInvestigationResponseSuggestions {
	this := SecurityAgentInvestigationResponseSuggestions{}
	return &this
}

// NewSecurityAgentInvestigationResponseSuggestionsWithDefaults instantiates a new SecurityAgentInvestigationResponseSuggestions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationResponseSuggestionsWithDefaults() *SecurityAgentInvestigationResponseSuggestions {
	this := SecurityAgentInvestigationResponseSuggestions{}
	return &this
}

// GetMitigationSuggestions returns the MitigationSuggestions field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationResponseSuggestions) GetMitigationSuggestions() []SecurityAgentInvestigationMitigationSuggestion {
	if o == nil || o.MitigationSuggestions == nil {
		var ret []SecurityAgentInvestigationMitigationSuggestion
		return ret
	}
	return o.MitigationSuggestions
}

// GetMitigationSuggestionsOk returns a tuple with the MitigationSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationResponseSuggestions) GetMitigationSuggestionsOk() (*[]SecurityAgentInvestigationMitigationSuggestion, bool) {
	if o == nil || o.MitigationSuggestions == nil {
		return nil, false
	}
	return &o.MitigationSuggestions, true
}

// HasMitigationSuggestions returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationResponseSuggestions) HasMitigationSuggestions() bool {
	return o != nil && o.MitigationSuggestions != nil
}

// SetMitigationSuggestions gets a reference to the given []SecurityAgentInvestigationMitigationSuggestion and assigns it to the MitigationSuggestions field.
func (o *SecurityAgentInvestigationResponseSuggestions) SetMitigationSuggestions(v []SecurityAgentInvestigationMitigationSuggestion) {
	o.MitigationSuggestions = v
}

// GetRuleTuningSuggestions returns the RuleTuningSuggestions field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationResponseSuggestions) GetRuleTuningSuggestions() SecurityAgentInvestigationRuleTuningSuggestions {
	if o == nil || o.RuleTuningSuggestions == nil {
		var ret SecurityAgentInvestigationRuleTuningSuggestions
		return ret
	}
	return *o.RuleTuningSuggestions
}

// GetRuleTuningSuggestionsOk returns a tuple with the RuleTuningSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationResponseSuggestions) GetRuleTuningSuggestionsOk() (*SecurityAgentInvestigationRuleTuningSuggestions, bool) {
	if o == nil || o.RuleTuningSuggestions == nil {
		return nil, false
	}
	return o.RuleTuningSuggestions, true
}

// HasRuleTuningSuggestions returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationResponseSuggestions) HasRuleTuningSuggestions() bool {
	return o != nil && o.RuleTuningSuggestions != nil
}

// SetRuleTuningSuggestions gets a reference to the given SecurityAgentInvestigationRuleTuningSuggestions and assigns it to the RuleTuningSuggestions field.
func (o *SecurityAgentInvestigationResponseSuggestions) SetRuleTuningSuggestions(v SecurityAgentInvestigationRuleTuningSuggestions) {
	o.RuleTuningSuggestions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationResponseSuggestions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MitigationSuggestions != nil {
		toSerialize["mitigation_suggestions"] = o.MitigationSuggestions
	}
	if o.RuleTuningSuggestions != nil {
		toSerialize["rule_tuning_suggestions"] = o.RuleTuningSuggestions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationResponseSuggestions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MitigationSuggestions []SecurityAgentInvestigationMitigationSuggestion `json:"mitigation_suggestions,omitempty"`
		RuleTuningSuggestions *SecurityAgentInvestigationRuleTuningSuggestions `json:"rule_tuning_suggestions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mitigation_suggestions", "rule_tuning_suggestions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MitigationSuggestions = all.MitigationSuggestions
	if all.RuleTuningSuggestions != nil && all.RuleTuningSuggestions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RuleTuningSuggestions = all.RuleTuningSuggestions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
