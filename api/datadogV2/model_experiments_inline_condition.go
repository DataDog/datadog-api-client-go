// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsInlineCondition An inline condition. The saved_filter_id field must be omitted or null.
type ExperimentsInlineCondition struct {
	// Attribute to evaluate.
	Attribute string `json:"attribute"`
	// Required with attribute and value for an inline condition; omit when saved_filter_id is set.
	Operator ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator `json:"operator"`
	// Values used by the operator. Every operator requires at least one value.
	Value []string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsInlineCondition instantiates a new ExperimentsInlineCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsInlineCondition(attribute string, operator ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator, value []string) *ExperimentsInlineCondition {
	this := ExperimentsInlineCondition{}
	this.Attribute = attribute
	this.Operator = operator
	this.Value = value
	return &this
}

// NewExperimentsInlineConditionWithDefaults instantiates a new ExperimentsInlineCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsInlineConditionWithDefaults() *ExperimentsInlineCondition {
	this := ExperimentsInlineCondition{}
	return &this
}

// GetAttribute returns the Attribute field value.
func (o *ExperimentsInlineCondition) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *ExperimentsInlineCondition) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *ExperimentsInlineCondition) SetAttribute(v string) {
	o.Attribute = v
}

// GetOperator returns the Operator field value.
func (o *ExperimentsInlineCondition) GetOperator() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ExperimentsInlineCondition) GetOperatorOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *ExperimentsInlineCondition) SetOperator(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *ExperimentsInlineCondition) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ExperimentsInlineCondition) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ExperimentsInlineCondition) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsInlineCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute"] = o.Attribute
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsInlineCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string                                                                                                               `json:"attribute"`
		Operator  *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationTargetingRulesItemsConditionsItemsOperator `json:"operator"`
		Value     *[]string                                                                                                             `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attribute == nil {
		return fmt.Errorf("required field attribute missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "operator", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = *all.Attribute
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
