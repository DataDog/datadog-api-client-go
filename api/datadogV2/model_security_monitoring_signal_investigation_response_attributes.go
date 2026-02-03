// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationResponseAttributes Attributes of a signal investigation response.
type SecurityMonitoringSignalInvestigationResponseAttributes struct {
	// The unique ID of the investigation.
	InvestigationId string `json:"investigation_id"`
	// The ID of the rule that triggered the signal.
	RuleId string `json:"rule_id"`
	// The unique ID of the security signal.
	SignalId string `json:"signal_id"`
	// The state of the investigation.
	State *SecurityMonitoringSignalInvestigationState `json:"state,omitempty"`
	// Information about an investigation step.
	Step *SecurityMonitoringSignalInvestigationStep `json:"step,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalInvestigationResponseAttributes instantiates a new SecurityMonitoringSignalInvestigationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalInvestigationResponseAttributes(investigationId string, ruleId string, signalId string) *SecurityMonitoringSignalInvestigationResponseAttributes {
	this := SecurityMonitoringSignalInvestigationResponseAttributes{}
	this.InvestigationId = investigationId
	this.RuleId = ruleId
	this.SignalId = signalId
	return &this
}

// NewSecurityMonitoringSignalInvestigationResponseAttributesWithDefaults instantiates a new SecurityMonitoringSignalInvestigationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalInvestigationResponseAttributesWithDefaults() *SecurityMonitoringSignalInvestigationResponseAttributes {
	this := SecurityMonitoringSignalInvestigationResponseAttributes{}
	return &this
}

// GetInvestigationId returns the InvestigationId field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetInvestigationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InvestigationId
}

// GetInvestigationIdOk returns a tuple with the InvestigationId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetInvestigationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestigationId, true
}

// SetInvestigationId sets field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) SetInvestigationId(v string) {
	o.InvestigationId = v
}

// GetRuleId returns the RuleId field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) SetRuleId(v string) {
	o.RuleId = v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetState() SecurityMonitoringSignalInvestigationState {
	if o == nil || o.State == nil {
		var ret SecurityMonitoringSignalInvestigationState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetStateOk() (*SecurityMonitoringSignalInvestigationState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given SecurityMonitoringSignalInvestigationState and assigns it to the State field.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) SetState(v SecurityMonitoringSignalInvestigationState) {
	o.State = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetStep() SecurityMonitoringSignalInvestigationStep {
	if o == nil || o.Step == nil {
		var ret SecurityMonitoringSignalInvestigationStep
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) GetStepOk() (*SecurityMonitoringSignalInvestigationStep, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) HasStep() bool {
	return o != nil && o.Step != nil
}

// SetStep gets a reference to the given SecurityMonitoringSignalInvestigationStep and assigns it to the Step field.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) SetStep(v SecurityMonitoringSignalInvestigationStep) {
	o.Step = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalInvestigationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["investigation_id"] = o.InvestigationId
	toSerialize["rule_id"] = o.RuleId
	toSerialize["signal_id"] = o.SignalId
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Step != nil {
		toSerialize["step"] = o.Step
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalInvestigationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InvestigationId *string                                     `json:"investigation_id"`
		RuleId          *string                                     `json:"rule_id"`
		SignalId        *string                                     `json:"signal_id"`
		State           *SecurityMonitoringSignalInvestigationState `json:"state,omitempty"`
		Step            *SecurityMonitoringSignalInvestigationStep  `json:"step,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InvestigationId == nil {
		return fmt.Errorf("required field investigation_id missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"investigation_id", "rule_id", "signal_id", "state", "step"})
	} else {
		return err
	}

	hasInvalidField := false
	o.InvestigationId = *all.InvestigationId
	o.RuleId = *all.RuleId
	o.SignalId = *all.SignalId
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	if all.Step != nil && all.Step.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Step = all.Step

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
