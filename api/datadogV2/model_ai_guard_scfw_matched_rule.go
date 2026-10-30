// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardScfwMatchedRule One rule the Supply Chain Firewall matched on a package, with the outcome the rule reached.
type AIGuardScfwMatchedRule struct {
	// The Supply Chain Firewall outcome for a package, a matched rule, or the whole request.
	Outcome AIGuardScfwOutcome `json:"outcome"`
	// The rule text the firewall reported, or the policy type when no text was reported.
	Rule string `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardScfwMatchedRule instantiates a new AIGuardScfwMatchedRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardScfwMatchedRule(outcome AIGuardScfwOutcome, rule string) *AIGuardScfwMatchedRule {
	this := AIGuardScfwMatchedRule{}
	this.Outcome = outcome
	this.Rule = rule
	return &this
}

// NewAIGuardScfwMatchedRuleWithDefaults instantiates a new AIGuardScfwMatchedRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardScfwMatchedRuleWithDefaults() *AIGuardScfwMatchedRule {
	this := AIGuardScfwMatchedRule{}
	return &this
}

// GetOutcome returns the Outcome field value.
func (o *AIGuardScfwMatchedRule) GetOutcome() AIGuardScfwOutcome {
	if o == nil {
		var ret AIGuardScfwOutcome
		return ret
	}
	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwMatchedRule) GetOutcomeOk() (*AIGuardScfwOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value.
func (o *AIGuardScfwMatchedRule) SetOutcome(v AIGuardScfwOutcome) {
	o.Outcome = v
}

// GetRule returns the Rule field value.
func (o *AIGuardScfwMatchedRule) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *AIGuardScfwMatchedRule) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *AIGuardScfwMatchedRule) SetRule(v string) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardScfwMatchedRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["outcome"] = o.Outcome
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardScfwMatchedRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Outcome *AIGuardScfwOutcome `json:"outcome"`
		Rule    *string             `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Outcome == nil {
		return fmt.Errorf("required field outcome missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"outcome", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Outcome.IsValid() {
		hasInvalidField = true
	} else {
		o.Outcome = *all.Outcome
	}
	o.Rule = *all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
