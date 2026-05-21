// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentBatchUpdateRuleExecutionStatesDataAttributes Attributes for batch updating rule execution states.
type IncidentBatchUpdateRuleExecutionStatesDataAttributes struct {
	// List of rules to update execution states for.
	Rules []IncidentRuleExecutionStateRule `json:"rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentBatchUpdateRuleExecutionStatesDataAttributes instantiates a new IncidentBatchUpdateRuleExecutionStatesDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentBatchUpdateRuleExecutionStatesDataAttributes(rules []IncidentRuleExecutionStateRule) *IncidentBatchUpdateRuleExecutionStatesDataAttributes {
	this := IncidentBatchUpdateRuleExecutionStatesDataAttributes{}
	this.Rules = rules
	return &this
}

// NewIncidentBatchUpdateRuleExecutionStatesDataAttributesWithDefaults instantiates a new IncidentBatchUpdateRuleExecutionStatesDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentBatchUpdateRuleExecutionStatesDataAttributesWithDefaults() *IncidentBatchUpdateRuleExecutionStatesDataAttributes {
	this := IncidentBatchUpdateRuleExecutionStatesDataAttributes{}
	return &this
}

// GetRules returns the Rules field value.
func (o *IncidentBatchUpdateRuleExecutionStatesDataAttributes) GetRules() []IncidentRuleExecutionStateRule {
	if o == nil {
		var ret []IncidentRuleExecutionStateRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *IncidentBatchUpdateRuleExecutionStatesDataAttributes) GetRulesOk() (*[]IncidentRuleExecutionStateRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *IncidentBatchUpdateRuleExecutionStatesDataAttributes) SetRules(v []IncidentRuleExecutionStateRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentBatchUpdateRuleExecutionStatesDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rules"] = o.Rules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentBatchUpdateRuleExecutionStatesDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rules *[]IncidentRuleExecutionStateRule `json:"rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rules"})
	} else {
		return err
	}
	o.Rules = *all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
