// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems Use an empty array when no targeting rules apply.
type ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems struct {
	// Conditions that must all match for this rule. Each condition must use exactly one shape: saved_filter_id alone or operator plus attribute plus value.
	Conditions []ExperimentsTargetingRuleCondition `json:"conditions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems(conditions []ExperimentsTargetingRuleCondition) *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems{}
	this.Conditions = conditions
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) GetConditions() []ExperimentsTargetingRuleCondition {
	if o == nil {
		var ret []ExperimentsTargetingRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) GetConditionsOk() (*[]ExperimentsTargetingRuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) SetConditions(v []ExperimentsTargetingRuleCondition) {
	o.Conditions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]ExperimentsTargetingRuleCondition `json:"conditions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions"})
	} else {
		return err
	}
	o.Conditions = *all.Conditions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
