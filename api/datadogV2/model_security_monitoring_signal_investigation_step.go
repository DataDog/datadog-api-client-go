// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationStep Information about an investigation step.
type SecurityMonitoringSignalInvestigationStep struct {
	// Array of step outputs.
	StepOutputs []SecurityMonitoringSignalInvestigationStepOutput `json:"stepOutputs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationStep instantiates a new SecurityMonitoringSignalInvestigationStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationStep(stepOutputs []SecurityMonitoringSignalInvestigationStepOutput) *SecurityMonitoringSignalInvestigationStep {
	this := SecurityMonitoringSignalInvestigationStep{}
	this.StepOutputs = stepOutputs
	return &this
}

// NewSecurityMonitoringSignalInvestigationStepWithDefaults instantiates a new SecurityMonitoringSignalInvestigationStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationStepWithDefaults() *SecurityMonitoringSignalInvestigationStep {
	this := SecurityMonitoringSignalInvestigationStep{}
	return &this
}

// GetStepOutputs returns the StepOutputs field value.
func (o *SecurityMonitoringSignalInvestigationStep) GetStepOutputs() []SecurityMonitoringSignalInvestigationStepOutput {
	if o == nil {
		var ret []SecurityMonitoringSignalInvestigationStepOutput
		return ret
	}
	return o.StepOutputs
}

// GetStepOutputsOk returns a tuple with the StepOutputs field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationStep) GetStepOutputsOk() (*[]SecurityMonitoringSignalInvestigationStepOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepOutputs, true
}

// SetStepOutputs sets field value.
func (o *SecurityMonitoringSignalInvestigationStep) SetStepOutputs(v []SecurityMonitoringSignalInvestigationStepOutput) {
	o.StepOutputs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["stepOutputs"] = o.StepOutputs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StepOutputs *[]SecurityMonitoringSignalInvestigationStepOutput `json:"stepOutputs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.StepOutputs == nil {
		return fmt.Errorf("required field stepOutputs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"stepOutputs"})
	} else {
		return err
	}
	o.StepOutputs = *all.StepOutputs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
