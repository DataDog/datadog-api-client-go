// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationRuleTuningCondition A condition of a rule tuning suggestion.
type SecurityAgentInvestigationRuleTuningCondition struct {
	// Attribute the condition checks.
	Field string `json:"field"`
	// Value the attribute is compared with.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationRuleTuningCondition instantiates a new SecurityAgentInvestigationRuleTuningCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationRuleTuningCondition(field string, value string) *SecurityAgentInvestigationRuleTuningCondition {
	this := SecurityAgentInvestigationRuleTuningCondition{}
	this.Field = field
	this.Value = value
	return &this
}

// NewSecurityAgentInvestigationRuleTuningConditionWithDefaults instantiates a new SecurityAgentInvestigationRuleTuningCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationRuleTuningConditionWithDefaults() *SecurityAgentInvestigationRuleTuningCondition {
	this := SecurityAgentInvestigationRuleTuningCondition{}
	return &this
}

// GetField returns the Field field value.
func (o *SecurityAgentInvestigationRuleTuningCondition) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningCondition) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *SecurityAgentInvestigationRuleTuningCondition) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value.
func (o *SecurityAgentInvestigationRuleTuningCondition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationRuleTuningCondition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *SecurityAgentInvestigationRuleTuningCondition) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationRuleTuningCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationRuleTuningCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field *string `json:"field"`
		Value *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "value"})
	} else {
		return err
	}
	o.Field = *all.Field
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
