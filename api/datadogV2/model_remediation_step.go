// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationStep A single execution-oriented step in a remediation plan.
type RemediationStep struct {
	// Fully qualified action type, for example ecs.update_service_memory.
	ActionFqn *string `json:"action_fqn,omitempty"`
	// Approval state for a remediation step.
	ApprovalState *RemediationStepApprovalState `json:"approval_state,omitempty"`
	// Human-readable description of the step.
	Description *string `json:"description,omitempty"`
	// Whether the step can be reversed after execution.
	Reversible *bool `json:"reversible,omitempty"`
	// Estimated risk of a remediation step or proposed fix.
	Risk *RemediationRiskLevel `json:"risk,omitempty"`
	// Unique ID for the step within the plan.
	StepId *string `json:"step_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationStep instantiates a new RemediationStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationStep() *RemediationStep {
	this := RemediationStep{}
	return &this
}

// NewRemediationStepWithDefaults instantiates a new RemediationStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationStepWithDefaults() *RemediationStep {
	this := RemediationStep{}
	return &this
}

// GetActionFqn returns the ActionFqn field value if set, zero value otherwise.
func (o *RemediationStep) GetActionFqn() string {
	if o == nil || o.ActionFqn == nil {
		var ret string
		return ret
	}
	return *o.ActionFqn
}

// GetActionFqnOk returns a tuple with the ActionFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetActionFqnOk() (*string, bool) {
	if o == nil || o.ActionFqn == nil {
		return nil, false
	}
	return o.ActionFqn, true
}

// HasActionFqn returns a boolean if a field has been set.
func (o *RemediationStep) HasActionFqn() bool {
	return o != nil && o.ActionFqn != nil
}

// SetActionFqn gets a reference to the given string and assigns it to the ActionFqn field.
func (o *RemediationStep) SetActionFqn(v string) {
	o.ActionFqn = &v
}

// GetApprovalState returns the ApprovalState field value if set, zero value otherwise.
func (o *RemediationStep) GetApprovalState() RemediationStepApprovalState {
	if o == nil || o.ApprovalState == nil {
		var ret RemediationStepApprovalState
		return ret
	}
	return *o.ApprovalState
}

// GetApprovalStateOk returns a tuple with the ApprovalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetApprovalStateOk() (*RemediationStepApprovalState, bool) {
	if o == nil || o.ApprovalState == nil {
		return nil, false
	}
	return o.ApprovalState, true
}

// HasApprovalState returns a boolean if a field has been set.
func (o *RemediationStep) HasApprovalState() bool {
	return o != nil && o.ApprovalState != nil
}

// SetApprovalState gets a reference to the given RemediationStepApprovalState and assigns it to the ApprovalState field.
func (o *RemediationStep) SetApprovalState(v RemediationStepApprovalState) {
	o.ApprovalState = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RemediationStep) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RemediationStep) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RemediationStep) SetDescription(v string) {
	o.Description = &v
}

// GetReversible returns the Reversible field value if set, zero value otherwise.
func (o *RemediationStep) GetReversible() bool {
	if o == nil || o.Reversible == nil {
		var ret bool
		return ret
	}
	return *o.Reversible
}

// GetReversibleOk returns a tuple with the Reversible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetReversibleOk() (*bool, bool) {
	if o == nil || o.Reversible == nil {
		return nil, false
	}
	return o.Reversible, true
}

// HasReversible returns a boolean if a field has been set.
func (o *RemediationStep) HasReversible() bool {
	return o != nil && o.Reversible != nil
}

// SetReversible gets a reference to the given bool and assigns it to the Reversible field.
func (o *RemediationStep) SetReversible(v bool) {
	o.Reversible = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *RemediationStep) GetRisk() RemediationRiskLevel {
	if o == nil || o.Risk == nil {
		var ret RemediationRiskLevel
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetRiskOk() (*RemediationRiskLevel, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *RemediationStep) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given RemediationRiskLevel and assigns it to the Risk field.
func (o *RemediationStep) SetRisk(v RemediationRiskLevel) {
	o.Risk = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *RemediationStep) GetStepId() string {
	if o == nil || o.StepId == nil {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationStep) GetStepIdOk() (*string, bool) {
	if o == nil || o.StepId == nil {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *RemediationStep) HasStepId() bool {
	return o != nil && o.StepId != nil
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *RemediationStep) SetStepId(v string) {
	o.StepId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ActionFqn != nil {
		toSerialize["action_fqn"] = o.ActionFqn
	}
	if o.ApprovalState != nil {
		toSerialize["approval_state"] = o.ApprovalState
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Reversible != nil {
		toSerialize["reversible"] = o.Reversible
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.StepId != nil {
		toSerialize["step_id"] = o.StepId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionFqn     *string                       `json:"action_fqn,omitempty"`
		ApprovalState *RemediationStepApprovalState `json:"approval_state,omitempty"`
		Description   *string                       `json:"description,omitempty"`
		Reversible    *bool                         `json:"reversible,omitempty"`
		Risk          *RemediationRiskLevel         `json:"risk,omitempty"`
		StepId        *string                       `json:"step_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_fqn", "approval_state", "description", "reversible", "risk", "step_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionFqn = all.ActionFqn
	if all.ApprovalState != nil && !all.ApprovalState.IsValid() {
		hasInvalidField = true
	} else {
		o.ApprovalState = all.ApprovalState
	}
	o.Description = all.Description
	o.Reversible = all.Reversible
	if all.Risk != nil && !all.Risk.IsValid() {
		hasInvalidField = true
	} else {
		o.Risk = all.Risk
	}
	o.StepId = all.StepId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
