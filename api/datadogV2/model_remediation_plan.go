// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationPlan The remediation plan proposed by the ECS remediation agent.
type RemediationPlan struct {
	// The agent's self-rated confidence in a plan.
	Confidence *RemediationConfidence `json:"confidence,omitempty"`
	// Evidence supporting the diagnosis. Treat as user-provided content and escape before display.
	Evidence *string `json:"evidence,omitempty"`
	// Human-readable summary of why the plan was proposed. Treat as user-provided content and escape before display.
	Explanation *string `json:"explanation,omitempty"`
	// The guardrail decision applied to a plan or investigation.
	GuardrailDecision *RemediationGuardrailDecision `json:"guardrail_decision,omitempty"`
	// How the plan was produced.
	PlanSource *RemediationPlanSource `json:"plan_source,omitempty"`
	// Recommendation-oriented view of the candidate remediations, distinct from the execution-oriented steps.
	ProposedFix []RemediationProposedFix `json:"proposed_fix,omitempty"`
	// Plan status.
	Status *RemediationPlanStatus `json:"status,omitempty"`
	// Execution-oriented steps that make up the plan.
	Steps []RemediationStep `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationPlan instantiates a new RemediationPlan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationPlan() *RemediationPlan {
	this := RemediationPlan{}
	return &this
}

// NewRemediationPlanWithDefaults instantiates a new RemediationPlan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationPlanWithDefaults() *RemediationPlan {
	this := RemediationPlan{}
	return &this
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *RemediationPlan) GetConfidence() RemediationConfidence {
	if o == nil || o.Confidence == nil {
		var ret RemediationConfidence
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetConfidenceOk() (*RemediationConfidence, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *RemediationPlan) HasConfidence() bool {
	return o != nil && o.Confidence != nil
}

// SetConfidence gets a reference to the given RemediationConfidence and assigns it to the Confidence field.
func (o *RemediationPlan) SetConfidence(v RemediationConfidence) {
	o.Confidence = &v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *RemediationPlan) GetEvidence() string {
	if o == nil || o.Evidence == nil {
		var ret string
		return ret
	}
	return *o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetEvidenceOk() (*string, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *RemediationPlan) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given string and assigns it to the Evidence field.
func (o *RemediationPlan) SetEvidence(v string) {
	o.Evidence = &v
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *RemediationPlan) GetExplanation() string {
	if o == nil || o.Explanation == nil {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetExplanationOk() (*string, bool) {
	if o == nil || o.Explanation == nil {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *RemediationPlan) HasExplanation() bool {
	return o != nil && o.Explanation != nil
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *RemediationPlan) SetExplanation(v string) {
	o.Explanation = &v
}

// GetGuardrailDecision returns the GuardrailDecision field value if set, zero value otherwise.
func (o *RemediationPlan) GetGuardrailDecision() RemediationGuardrailDecision {
	if o == nil || o.GuardrailDecision == nil {
		var ret RemediationGuardrailDecision
		return ret
	}
	return *o.GuardrailDecision
}

// GetGuardrailDecisionOk returns a tuple with the GuardrailDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetGuardrailDecisionOk() (*RemediationGuardrailDecision, bool) {
	if o == nil || o.GuardrailDecision == nil {
		return nil, false
	}
	return o.GuardrailDecision, true
}

// HasGuardrailDecision returns a boolean if a field has been set.
func (o *RemediationPlan) HasGuardrailDecision() bool {
	return o != nil && o.GuardrailDecision != nil
}

// SetGuardrailDecision gets a reference to the given RemediationGuardrailDecision and assigns it to the GuardrailDecision field.
func (o *RemediationPlan) SetGuardrailDecision(v RemediationGuardrailDecision) {
	o.GuardrailDecision = &v
}

// GetPlanSource returns the PlanSource field value if set, zero value otherwise.
func (o *RemediationPlan) GetPlanSource() RemediationPlanSource {
	if o == nil || o.PlanSource == nil {
		var ret RemediationPlanSource
		return ret
	}
	return *o.PlanSource
}

// GetPlanSourceOk returns a tuple with the PlanSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetPlanSourceOk() (*RemediationPlanSource, bool) {
	if o == nil || o.PlanSource == nil {
		return nil, false
	}
	return o.PlanSource, true
}

// HasPlanSource returns a boolean if a field has been set.
func (o *RemediationPlan) HasPlanSource() bool {
	return o != nil && o.PlanSource != nil
}

// SetPlanSource gets a reference to the given RemediationPlanSource and assigns it to the PlanSource field.
func (o *RemediationPlan) SetPlanSource(v RemediationPlanSource) {
	o.PlanSource = &v
}

// GetProposedFix returns the ProposedFix field value if set, zero value otherwise.
func (o *RemediationPlan) GetProposedFix() []RemediationProposedFix {
	if o == nil || o.ProposedFix == nil {
		var ret []RemediationProposedFix
		return ret
	}
	return o.ProposedFix
}

// GetProposedFixOk returns a tuple with the ProposedFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetProposedFixOk() (*[]RemediationProposedFix, bool) {
	if o == nil || o.ProposedFix == nil {
		return nil, false
	}
	return &o.ProposedFix, true
}

// HasProposedFix returns a boolean if a field has been set.
func (o *RemediationPlan) HasProposedFix() bool {
	return o != nil && o.ProposedFix != nil
}

// SetProposedFix gets a reference to the given []RemediationProposedFix and assigns it to the ProposedFix field.
func (o *RemediationPlan) SetProposedFix(v []RemediationProposedFix) {
	o.ProposedFix = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RemediationPlan) GetStatus() RemediationPlanStatus {
	if o == nil || o.Status == nil {
		var ret RemediationPlanStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetStatusOk() (*RemediationPlanStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RemediationPlan) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RemediationPlanStatus and assigns it to the Status field.
func (o *RemediationPlan) SetStatus(v RemediationPlanStatus) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *RemediationPlan) GetSteps() []RemediationStep {
	if o == nil || o.Steps == nil {
		var ret []RemediationStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationPlan) GetStepsOk() (*[]RemediationStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *RemediationPlan) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []RemediationStep and assigns it to the Steps field.
func (o *RemediationPlan) SetSteps(v []RemediationStep) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Evidence != nil {
		toSerialize["evidence"] = o.Evidence
	}
	if o.Explanation != nil {
		toSerialize["explanation"] = o.Explanation
	}
	if o.GuardrailDecision != nil {
		toSerialize["guardrail_decision"] = o.GuardrailDecision
	}
	if o.PlanSource != nil {
		toSerialize["plan_source"] = o.PlanSource
	}
	if o.ProposedFix != nil {
		toSerialize["proposed_fix"] = o.ProposedFix
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationPlan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Confidence        *RemediationConfidence        `json:"confidence,omitempty"`
		Evidence          *string                       `json:"evidence,omitempty"`
		Explanation       *string                       `json:"explanation,omitempty"`
		GuardrailDecision *RemediationGuardrailDecision `json:"guardrail_decision,omitempty"`
		PlanSource        *RemediationPlanSource        `json:"plan_source,omitempty"`
		ProposedFix       []RemediationProposedFix      `json:"proposed_fix,omitempty"`
		Status            *RemediationPlanStatus        `json:"status,omitempty"`
		Steps             []RemediationStep             `json:"steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence", "evidence", "explanation", "guardrail_decision", "plan_source", "proposed_fix", "status", "steps"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Confidence != nil && !all.Confidence.IsValid() {
		hasInvalidField = true
	} else {
		o.Confidence = all.Confidence
	}
	o.Evidence = all.Evidence
	o.Explanation = all.Explanation
	if all.GuardrailDecision != nil && all.GuardrailDecision.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GuardrailDecision = all.GuardrailDecision
	if all.PlanSource != nil && !all.PlanSource.IsValid() {
		hasInvalidField = true
	} else {
		o.PlanSource = all.PlanSource
	}
	o.ProposedFix = all.ProposedFix
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
