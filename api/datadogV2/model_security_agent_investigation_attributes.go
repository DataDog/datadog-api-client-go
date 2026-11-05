// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationAttributes Attributes of an investigation. When `fields[investigations]` is set, only listed attributes are returned.
// Attributes the investigation doesn't have yet are absent.
type SecurityAgentInvestigationAttributes struct {
	// Summary of the investigation's findings.
	Conclusion *string `json:"conclusion,omitempty"`
	// When the investigation started.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The entities the investigation found, by role.
	KeyEntities *SecurityAgentInvestigationKeyEntities `json:"key_entities,omitempty"`
	// The evidence the verdict rests on.
	KeyEvidence []SecurityAgentInvestigationKeyEvidence `json:"key_evidence,omitempty"`
	// Suggested responses to the signal.
	ResponseSuggestions *SecurityAgentInvestigationResponseSuggestions `json:"response_suggestions,omitempty"`
	// ID of the detection rule that generated the signal.
	RuleId *string `json:"rule_id,omitempty"`
	// Summary of the signal the investigation is about.
	Signal *SecurityAgentInvestigationSignal `json:"signal,omitempty"`
	// ID of the signal, in its V1 form.
	SignalId *string `json:"signal_id,omitempty"`
	// State of the investigation: `investigating` while it runs, `completed` after it has a result, or `failed`
	// when it produced no result before its timeout.
	State *SecurityAgentInvestigationState `json:"state,omitempty"`
	// The steps of the investigation, in order.
	Steps []SecurityAgentInvestigationStep `json:"steps,omitempty"`
	// When the investigation last changed.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Verdict of the investigation. Only completed investigations have one.
	Verdict *SecurityAgentInvestigationVerdict `json:"verdict,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationAttributes instantiates a new SecurityAgentInvestigationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationAttributes() *SecurityAgentInvestigationAttributes {
	this := SecurityAgentInvestigationAttributes{}
	return &this
}

// NewSecurityAgentInvestigationAttributesWithDefaults instantiates a new SecurityAgentInvestigationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationAttributesWithDefaults() *SecurityAgentInvestigationAttributes {
	this := SecurityAgentInvestigationAttributes{}
	return &this
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetConclusion() string {
	if o == nil || o.Conclusion == nil {
		var ret string
		return ret
	}
	return *o.Conclusion
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetConclusionOk() (*string, bool) {
	if o == nil || o.Conclusion == nil {
		return nil, false
	}
	return o.Conclusion, true
}

// HasConclusion returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasConclusion() bool {
	return o != nil && o.Conclusion != nil
}

// SetConclusion gets a reference to the given string and assigns it to the Conclusion field.
func (o *SecurityAgentInvestigationAttributes) SetConclusion(v string) {
	o.Conclusion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SecurityAgentInvestigationAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetKeyEntities returns the KeyEntities field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetKeyEntities() SecurityAgentInvestigationKeyEntities {
	if o == nil || o.KeyEntities == nil {
		var ret SecurityAgentInvestigationKeyEntities
		return ret
	}
	return *o.KeyEntities
}

// GetKeyEntitiesOk returns a tuple with the KeyEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetKeyEntitiesOk() (*SecurityAgentInvestigationKeyEntities, bool) {
	if o == nil || o.KeyEntities == nil {
		return nil, false
	}
	return o.KeyEntities, true
}

// HasKeyEntities returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasKeyEntities() bool {
	return o != nil && o.KeyEntities != nil
}

// SetKeyEntities gets a reference to the given SecurityAgentInvestigationKeyEntities and assigns it to the KeyEntities field.
func (o *SecurityAgentInvestigationAttributes) SetKeyEntities(v SecurityAgentInvestigationKeyEntities) {
	o.KeyEntities = &v
}

// GetKeyEvidence returns the KeyEvidence field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetKeyEvidence() []SecurityAgentInvestigationKeyEvidence {
	if o == nil || o.KeyEvidence == nil {
		var ret []SecurityAgentInvestigationKeyEvidence
		return ret
	}
	return o.KeyEvidence
}

// GetKeyEvidenceOk returns a tuple with the KeyEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetKeyEvidenceOk() (*[]SecurityAgentInvestigationKeyEvidence, bool) {
	if o == nil || o.KeyEvidence == nil {
		return nil, false
	}
	return &o.KeyEvidence, true
}

// HasKeyEvidence returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasKeyEvidence() bool {
	return o != nil && o.KeyEvidence != nil
}

// SetKeyEvidence gets a reference to the given []SecurityAgentInvestigationKeyEvidence and assigns it to the KeyEvidence field.
func (o *SecurityAgentInvestigationAttributes) SetKeyEvidence(v []SecurityAgentInvestigationKeyEvidence) {
	o.KeyEvidence = v
}

// GetResponseSuggestions returns the ResponseSuggestions field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetResponseSuggestions() SecurityAgentInvestigationResponseSuggestions {
	if o == nil || o.ResponseSuggestions == nil {
		var ret SecurityAgentInvestigationResponseSuggestions
		return ret
	}
	return *o.ResponseSuggestions
}

// GetResponseSuggestionsOk returns a tuple with the ResponseSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetResponseSuggestionsOk() (*SecurityAgentInvestigationResponseSuggestions, bool) {
	if o == nil || o.ResponseSuggestions == nil {
		return nil, false
	}
	return o.ResponseSuggestions, true
}

// HasResponseSuggestions returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasResponseSuggestions() bool {
	return o != nil && o.ResponseSuggestions != nil
}

// SetResponseSuggestions gets a reference to the given SecurityAgentInvestigationResponseSuggestions and assigns it to the ResponseSuggestions field.
func (o *SecurityAgentInvestigationAttributes) SetResponseSuggestions(v SecurityAgentInvestigationResponseSuggestions) {
	o.ResponseSuggestions = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasRuleId() bool {
	return o != nil && o.RuleId != nil
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *SecurityAgentInvestigationAttributes) SetRuleId(v string) {
	o.RuleId = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetSignal() SecurityAgentInvestigationSignal {
	if o == nil || o.Signal == nil {
		var ret SecurityAgentInvestigationSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetSignalOk() (*SecurityAgentInvestigationSignal, bool) {
	if o == nil || o.Signal == nil {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasSignal() bool {
	return o != nil && o.Signal != nil
}

// SetSignal gets a reference to the given SecurityAgentInvestigationSignal and assigns it to the Signal field.
func (o *SecurityAgentInvestigationAttributes) SetSignal(v SecurityAgentInvestigationSignal) {
	o.Signal = &v
}

// GetSignalId returns the SignalId field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetSignalId() string {
	if o == nil || o.SignalId == nil {
		var ret string
		return ret
	}
	return *o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil || o.SignalId == nil {
		return nil, false
	}
	return o.SignalId, true
}

// HasSignalId returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasSignalId() bool {
	return o != nil && o.SignalId != nil
}

// SetSignalId gets a reference to the given string and assigns it to the SignalId field.
func (o *SecurityAgentInvestigationAttributes) SetSignalId(v string) {
	o.SignalId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetState() SecurityAgentInvestigationState {
	if o == nil || o.State == nil {
		var ret SecurityAgentInvestigationState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetStateOk() (*SecurityAgentInvestigationState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given SecurityAgentInvestigationState and assigns it to the State field.
func (o *SecurityAgentInvestigationAttributes) SetState(v SecurityAgentInvestigationState) {
	o.State = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetSteps() []SecurityAgentInvestigationStep {
	if o == nil || o.Steps == nil {
		var ret []SecurityAgentInvestigationStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetStepsOk() (*[]SecurityAgentInvestigationStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SecurityAgentInvestigationStep and assigns it to the Steps field.
func (o *SecurityAgentInvestigationAttributes) SetSteps(v []SecurityAgentInvestigationStep) {
	o.Steps = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SecurityAgentInvestigationAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerdict returns the Verdict field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationAttributes) GetVerdict() SecurityAgentInvestigationVerdict {
	if o == nil || o.Verdict == nil {
		var ret SecurityAgentInvestigationVerdict
		return ret
	}
	return *o.Verdict
}

// GetVerdictOk returns a tuple with the Verdict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationAttributes) GetVerdictOk() (*SecurityAgentInvestigationVerdict, bool) {
	if o == nil || o.Verdict == nil {
		return nil, false
	}
	return o.Verdict, true
}

// HasVerdict returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationAttributes) HasVerdict() bool {
	return o != nil && o.Verdict != nil
}

// SetVerdict gets a reference to the given SecurityAgentInvestigationVerdict and assigns it to the Verdict field.
func (o *SecurityAgentInvestigationAttributes) SetVerdict(v SecurityAgentInvestigationVerdict) {
	o.Verdict = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Conclusion != nil {
		toSerialize["conclusion"] = o.Conclusion
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.KeyEntities != nil {
		toSerialize["key_entities"] = o.KeyEntities
	}
	if o.KeyEvidence != nil {
		toSerialize["key_evidence"] = o.KeyEvidence
	}
	if o.ResponseSuggestions != nil {
		toSerialize["response_suggestions"] = o.ResponseSuggestions
	}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	if o.Signal != nil {
		toSerialize["signal"] = o.Signal
	}
	if o.SignalId != nil {
		toSerialize["signal_id"] = o.SignalId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Verdict != nil {
		toSerialize["verdict"] = o.Verdict
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conclusion          *string                                        `json:"conclusion,omitempty"`
		CreatedAt           *time.Time                                     `json:"created_at,omitempty"`
		KeyEntities         *SecurityAgentInvestigationKeyEntities         `json:"key_entities,omitempty"`
		KeyEvidence         []SecurityAgentInvestigationKeyEvidence        `json:"key_evidence,omitempty"`
		ResponseSuggestions *SecurityAgentInvestigationResponseSuggestions `json:"response_suggestions,omitempty"`
		RuleId              *string                                        `json:"rule_id,omitempty"`
		Signal              *SecurityAgentInvestigationSignal              `json:"signal,omitempty"`
		SignalId            *string                                        `json:"signal_id,omitempty"`
		State               *SecurityAgentInvestigationState               `json:"state,omitempty"`
		Steps               []SecurityAgentInvestigationStep               `json:"steps,omitempty"`
		UpdatedAt           *time.Time                                     `json:"updated_at,omitempty"`
		Verdict             *SecurityAgentInvestigationVerdict             `json:"verdict,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conclusion", "created_at", "key_entities", "key_evidence", "response_suggestions", "rule_id", "signal", "signal_id", "state", "steps", "updated_at", "verdict"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Conclusion = all.Conclusion
	o.CreatedAt = all.CreatedAt
	if all.KeyEntities != nil && all.KeyEntities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.KeyEntities = all.KeyEntities
	o.KeyEvidence = all.KeyEvidence
	if all.ResponseSuggestions != nil && all.ResponseSuggestions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ResponseSuggestions = all.ResponseSuggestions
	o.RuleId = all.RuleId
	if all.Signal != nil && all.Signal.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Signal = all.Signal
	o.SignalId = all.SignalId
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	o.Steps = all.Steps
	o.UpdatedAt = all.UpdatedAt
	if all.Verdict != nil && !all.Verdict.IsValid() {
		hasInvalidField = true
	} else {
		o.Verdict = all.Verdict
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
