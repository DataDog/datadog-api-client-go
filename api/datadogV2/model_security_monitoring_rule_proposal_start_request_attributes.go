// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalStartRequestAttributes Attributes of the rule proposal to start.
type SecurityMonitoringRuleProposalStartRequestAttributes struct {
	// The natural-language description or rule to generate the proposed rule from, in the
	// format given by `inputType`.
	Input string `json:"input"`
	// The type of input to generate the proposed rule from: `natural_language` for a
	// plain-English description of the behavior to detect, `spl` for a Splunk SPL search or
	// saved search definition, or `kql` for a Microsoft KQL query. SPL inputs must be
	// self-contained, with search macros expanded.
	InputType SecurityMonitoringRuleProposalInputType `json:"inputType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleProposalStartRequestAttributes instantiates a new SecurityMonitoringRuleProposalStartRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleProposalStartRequestAttributes(input string, inputType SecurityMonitoringRuleProposalInputType) *SecurityMonitoringRuleProposalStartRequestAttributes {
	this := SecurityMonitoringRuleProposalStartRequestAttributes{}
	this.Input = input
	this.InputType = inputType
	return &this
}

// NewSecurityMonitoringRuleProposalStartRequestAttributesWithDefaults instantiates a new SecurityMonitoringRuleProposalStartRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleProposalStartRequestAttributesWithDefaults() *SecurityMonitoringRuleProposalStartRequestAttributes {
	this := SecurityMonitoringRuleProposalStartRequestAttributes{}
	return &this
}

// GetInput returns the Input field value.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) SetInput(v string) {
	o.Input = v
}

// GetInputType returns the InputType field value.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) GetInputType() SecurityMonitoringRuleProposalInputType {
	if o == nil {
		var ret SecurityMonitoringRuleProposalInputType
		return ret
	}
	return o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) GetInputTypeOk() (*SecurityMonitoringRuleProposalInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputType, true
}

// SetInputType sets field value.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) SetInputType(v SecurityMonitoringRuleProposalInputType) {
	o.InputType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleProposalStartRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["input"] = o.Input
	toSerialize["inputType"] = o.InputType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleProposalStartRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Input     *string                                  `json:"input"`
		InputType *SecurityMonitoringRuleProposalInputType `json:"inputType"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Input == nil {
		return fmt.Errorf("required field input missing")
	}
	if all.InputType == nil {
		return fmt.Errorf("required field inputType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"input", "inputType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Input = *all.Input
	if !all.InputType.IsValid() {
		hasInvalidField = true
	} else {
		o.InputType = *all.InputType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
