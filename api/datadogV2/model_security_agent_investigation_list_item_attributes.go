// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationListItemAttributes Attributes of an investigation in a list. To get an investigation's full content, use the get investigation
// endpoint.
type SecurityAgentInvestigationListItemAttributes struct {
	// When the investigation started.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Whether this is the latest investigation of the signal. Always true when `filter[signal_id]` isn't set.
	IsLatest bool `json:"is_latest"`
	// ID of the detection rule that generated the signal. It can be empty while the investigation runs.
	RuleId string `json:"rule_id"`
	// Summary of the signal the investigation is about.
	Signal *SecurityAgentInvestigationSignal `json:"signal,omitempty"`
	// ID of the signal, in its V1 form.
	SignalId string `json:"signal_id"`
	// State of the investigation: `investigating` while it runs, `completed` after it has a result, or `failed`
	// when it produced no result before its timeout.
	State SecurityAgentInvestigationState `json:"state"`
	// When the investigation last changed.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Verdict of the investigation. Only completed investigations have one.
	Verdict *SecurityAgentInvestigationVerdict `json:"verdict,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationListItemAttributes instantiates a new SecurityAgentInvestigationListItemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationListItemAttributes(isLatest bool, ruleId string, signalId string, state SecurityAgentInvestigationState) *SecurityAgentInvestigationListItemAttributes {
	this := SecurityAgentInvestigationListItemAttributes{}
	this.IsLatest = isLatest
	this.RuleId = ruleId
	this.SignalId = signalId
	this.State = state
	return &this
}

// NewSecurityAgentInvestigationListItemAttributesWithDefaults instantiates a new SecurityAgentInvestigationListItemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationListItemAttributesWithDefaults() *SecurityAgentInvestigationListItemAttributes {
	this := SecurityAgentInvestigationListItemAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationListItemAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationListItemAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SecurityAgentInvestigationListItemAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetIsLatest returns the IsLatest field value.
func (o *SecurityAgentInvestigationListItemAttributes) GetIsLatest() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsLatest
}

// GetIsLatestOk returns a tuple with the IsLatest field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetIsLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLatest, true
}

// SetIsLatest sets field value.
func (o *SecurityAgentInvestigationListItemAttributes) SetIsLatest(v bool) {
	o.IsLatest = v
}

// GetRuleId returns the RuleId field value.
func (o *SecurityAgentInvestigationListItemAttributes) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value.
func (o *SecurityAgentInvestigationListItemAttributes) SetRuleId(v string) {
	o.RuleId = v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationListItemAttributes) GetSignal() SecurityAgentInvestigationSignal {
	if o == nil || o.Signal == nil {
		var ret SecurityAgentInvestigationSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetSignalOk() (*SecurityAgentInvestigationSignal, bool) {
	if o == nil || o.Signal == nil {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationListItemAttributes) HasSignal() bool {
	return o != nil && o.Signal != nil
}

// SetSignal gets a reference to the given SecurityAgentInvestigationSignal and assigns it to the Signal field.
func (o *SecurityAgentInvestigationListItemAttributes) SetSignal(v SecurityAgentInvestigationSignal) {
	o.Signal = &v
}

// GetSignalId returns the SignalId field value.
func (o *SecurityAgentInvestigationListItemAttributes) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value.
func (o *SecurityAgentInvestigationListItemAttributes) SetSignalId(v string) {
	o.SignalId = v
}

// GetState returns the State field value.
func (o *SecurityAgentInvestigationListItemAttributes) GetState() SecurityAgentInvestigationState {
	if o == nil {
		var ret SecurityAgentInvestigationState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetStateOk() (*SecurityAgentInvestigationState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *SecurityAgentInvestigationListItemAttributes) SetState(v SecurityAgentInvestigationState) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationListItemAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationListItemAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SecurityAgentInvestigationListItemAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerdict returns the Verdict field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationListItemAttributes) GetVerdict() SecurityAgentInvestigationVerdict {
	if o == nil || o.Verdict == nil {
		var ret SecurityAgentInvestigationVerdict
		return ret
	}
	return *o.Verdict
}

// GetVerdictOk returns a tuple with the Verdict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationListItemAttributes) GetVerdictOk() (*SecurityAgentInvestigationVerdict, bool) {
	if o == nil || o.Verdict == nil {
		return nil, false
	}
	return o.Verdict, true
}

// HasVerdict returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationListItemAttributes) HasVerdict() bool {
	return o != nil && o.Verdict != nil
}

// SetVerdict gets a reference to the given SecurityAgentInvestigationVerdict and assigns it to the Verdict field.
func (o *SecurityAgentInvestigationListItemAttributes) SetVerdict(v SecurityAgentInvestigationVerdict) {
	o.Verdict = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationListItemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["is_latest"] = o.IsLatest
	toSerialize["rule_id"] = o.RuleId
	if o.Signal != nil {
		toSerialize["signal"] = o.Signal
	}
	toSerialize["signal_id"] = o.SignalId
	toSerialize["state"] = o.State
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
func (o *SecurityAgentInvestigationListItemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                         `json:"created_at,omitempty"`
		IsLatest  *bool                              `json:"is_latest"`
		RuleId    *string                            `json:"rule_id"`
		Signal    *SecurityAgentInvestigationSignal  `json:"signal,omitempty"`
		SignalId  *string                            `json:"signal_id"`
		State     *SecurityAgentInvestigationState   `json:"state"`
		UpdatedAt *time.Time                         `json:"updated_at,omitempty"`
		Verdict   *SecurityAgentInvestigationVerdict `json:"verdict,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsLatest == nil {
		return fmt.Errorf("required field is_latest missing")
	}
	if all.RuleId == nil {
		return fmt.Errorf("required field rule_id missing")
	}
	if all.SignalId == nil {
		return fmt.Errorf("required field signal_id missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "is_latest", "rule_id", "signal", "signal_id", "state", "updated_at", "verdict"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.IsLatest = *all.IsLatest
	o.RuleId = *all.RuleId
	if all.Signal != nil && all.Signal.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Signal = all.Signal
	o.SignalId = *all.SignalId
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
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
