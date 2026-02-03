// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationStepOutput Output from an investigation step.
type SecurityMonitoringSignalInvestigationStepOutput struct {
	// A one-line summary of the step analysis.
	CurrentStepAnalysisOneliner *string `json:"currentStepAnalysisOneliner,omitempty"`
	// A detailed summary of the step analysis.
	CurrentStepAnalysisSummary string `json:"currentStepAnalysisSummary"`
	// The name of the investigation step.
	Name string `json:"name"`
	// The verdict from the investigation step.
	Verdict SecurityMonitoringSignalInvestigationStepOutputVerdict `json:"verdict"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationStepOutput instantiates a new SecurityMonitoringSignalInvestigationStepOutput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationStepOutput(currentStepAnalysisSummary string, name string, verdict SecurityMonitoringSignalInvestigationStepOutputVerdict) *SecurityMonitoringSignalInvestigationStepOutput {
	this := SecurityMonitoringSignalInvestigationStepOutput{}
	this.CurrentStepAnalysisSummary = currentStepAnalysisSummary
	this.Name = name
	this.Verdict = verdict
	return &this
}

// NewSecurityMonitoringSignalInvestigationStepOutputWithDefaults instantiates a new SecurityMonitoringSignalInvestigationStepOutput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationStepOutputWithDefaults() *SecurityMonitoringSignalInvestigationStepOutput {
	this := SecurityMonitoringSignalInvestigationStepOutput{}
	return &this
}

// GetCurrentStepAnalysisOneliner returns the CurrentStepAnalysisOneliner field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetCurrentStepAnalysisOneliner() string {
	if o == nil || o.CurrentStepAnalysisOneliner == nil {
		var ret string
		return ret
	}
	return *o.CurrentStepAnalysisOneliner
}

// GetCurrentStepAnalysisOnelinerOk returns a tuple with the CurrentStepAnalysisOneliner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetCurrentStepAnalysisOnelinerOk() (*string, bool) {
	if o == nil || o.CurrentStepAnalysisOneliner == nil {
		return nil, false
	}
	return o.CurrentStepAnalysisOneliner, true
}

// HasCurrentStepAnalysisOneliner returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationStepOutput) HasCurrentStepAnalysisOneliner() bool {
	return o != nil && o.CurrentStepAnalysisOneliner != nil
}

// SetCurrentStepAnalysisOneliner gets a reference to the given string and assigns it to the CurrentStepAnalysisOneliner field.
func (o *SecurityMonitoringSignalInvestigationStepOutput) SetCurrentStepAnalysisOneliner(v string) {
	o.CurrentStepAnalysisOneliner = &v
}

// GetCurrentStepAnalysisSummary returns the CurrentStepAnalysisSummary field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetCurrentStepAnalysisSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CurrentStepAnalysisSummary
}

// GetCurrentStepAnalysisSummaryOk returns a tuple with the CurrentStepAnalysisSummary field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetCurrentStepAnalysisSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStepAnalysisSummary, true
}

// SetCurrentStepAnalysisSummary sets field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) SetCurrentStepAnalysisSummary(v string) {
	o.CurrentStepAnalysisSummary = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) SetName(v string) {
	o.Name = v
}

// GetVerdict returns the Verdict field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetVerdict() SecurityMonitoringSignalInvestigationStepOutputVerdict {
	if o == nil {
		var ret SecurityMonitoringSignalInvestigationStepOutputVerdict
		return ret
	}
	return o.Verdict
}

// GetVerdictOk returns a tuple with the Verdict field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationStepOutput) GetVerdictOk() (*SecurityMonitoringSignalInvestigationStepOutputVerdict, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verdict, true
}

// SetVerdict sets field value.
func (o *SecurityMonitoringSignalInvestigationStepOutput) SetVerdict(v SecurityMonitoringSignalInvestigationStepOutputVerdict) {
	o.Verdict = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationStepOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentStepAnalysisOneliner != nil {
		toSerialize["currentStepAnalysisOneliner"] = o.CurrentStepAnalysisOneliner
	}
	toSerialize["currentStepAnalysisSummary"] = o.CurrentStepAnalysisSummary
	toSerialize["name"] = o.Name
	toSerialize["verdict"] = o.Verdict

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationStepOutput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentStepAnalysisOneliner *string                                                 `json:"currentStepAnalysisOneliner,omitempty"`
		CurrentStepAnalysisSummary  *string                                                 `json:"currentStepAnalysisSummary"`
		Name                        *string                                                 `json:"name"`
		Verdict                     *SecurityMonitoringSignalInvestigationStepOutputVerdict `json:"verdict"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CurrentStepAnalysisSummary == nil {
		return fmt.Errorf("required field currentStepAnalysisSummary missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Verdict == nil {
		return fmt.Errorf("required field verdict missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"currentStepAnalysisOneliner", "currentStepAnalysisSummary", "name", "verdict"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CurrentStepAnalysisOneliner = all.CurrentStepAnalysisOneliner
	o.CurrentStepAnalysisSummary = *all.CurrentStepAnalysisSummary
	o.Name = *all.Name
	if !all.Verdict.IsValid() {
		hasInvalidField = true
	} else {
		o.Verdict = *all.Verdict
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
