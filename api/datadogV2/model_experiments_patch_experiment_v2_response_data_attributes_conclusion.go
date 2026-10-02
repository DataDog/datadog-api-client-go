// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2ResponseDataAttributesConclusion Outcome and supporting text recorded when the experiment is concluded.
type ExperimentsPatchExperimentV2ResponseDataAttributesConclusion struct {
	// Reason for the recorded decision.
	DecisionReason datadog.NullableString `json:"decision_reason,omitempty"`
	// Recorded experiment outcome.
	Outcome NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome `json:"outcome,omitempty"`
	// Summary of the experiment conclusion.
	Summary datadog.NullableString `json:"summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesConclusion instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesConclusion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2ResponseDataAttributesConclusion() *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesConclusion{}
	return &this
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesConclusionWithDefaults instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesConclusion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2ResponseDataAttributesConclusionWithDefaults() *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesConclusion{}
	return &this
}

// GetDecisionReason returns the DecisionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetDecisionReason() string {
	if o == nil || o.DecisionReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.DecisionReason.Get()
}

// GetDecisionReasonOk returns a tuple with the DecisionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetDecisionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionReason.Get(), o.DecisionReason.IsSet()
}

// HasDecisionReason returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) HasDecisionReason() bool {
	return o != nil && o.DecisionReason.IsSet()
}

// SetDecisionReason gets a reference to the given datadog.NullableString and assigns it to the DecisionReason field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetDecisionReason(v string) {
	o.DecisionReason.Set(&v)
}

// SetDecisionReasonNil sets the value for DecisionReason to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetDecisionReasonNil() {
	o.DecisionReason.Set(nil)
}

// UnsetDecisionReason ensures that no value is present for DecisionReason, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) UnsetDecisionReason() {
	o.DecisionReason.Unset()
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetOutcome() ExperimentsExperimentV2DTODataAttributesConclusionOutcome {
	if o == nil || o.Outcome.Get() == nil {
		var ret ExperimentsExperimentV2DTODataAttributesConclusionOutcome
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetOutcomeOk() (*ExperimentsExperimentV2DTODataAttributesConclusionOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) HasOutcome() bool {
	return o != nil && o.Outcome.IsSet()
}

// SetOutcome gets a reference to the given NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome and assigns it to the Outcome field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetOutcome(v ExperimentsExperimentV2DTODataAttributesConclusionOutcome) {
	o.Outcome.Set(&v)
}

// SetOutcomeNil sets the value for Outcome to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetSummary() string {
	if o == nil || o.Summary.Get() == nil {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) HasSummary() bool {
	return o != nil && o.Summary.IsSet()
}

// SetSummary gets a reference to the given datadog.NullableString and assigns it to the Summary field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetSummary(v string) {
	o.Summary.Set(&v)
}

// SetSummaryNil sets the value for Summary to be an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) UnsetSummary() {
	o.Summary.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DecisionReason.IsSet() {
		toSerialize["decision_reason"] = o.DecisionReason.Get()
	}
	if o.Outcome.IsSet() {
		toSerialize["outcome"] = o.Outcome.Get()
	}
	if o.Summary.IsSet() {
		toSerialize["summary"] = o.Summary.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesConclusion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DecisionReason datadog.NullableString                                            `json:"decision_reason,omitempty"`
		Outcome        NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome `json:"outcome,omitempty"`
		Summary        datadog.NullableString                                            `json:"summary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"decision_reason", "outcome", "summary"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DecisionReason = all.DecisionReason
	if all.Outcome.Get() != nil && !all.Outcome.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Outcome = all.Outcome
	}
	o.Summary = all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
