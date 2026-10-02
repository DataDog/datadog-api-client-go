// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesExposureSchedule Schedule that controls traffic exposure for experiments created from the protocol.
type ExperimentsPublicProtocolResponseDataAttributesExposureSchedule struct {
	// Whether the exposure schedule starts automatically.
	Autostart *bool `json:"autostart,omitempty"`
	// Action taken when a guardrail triggers during the exposure schedule.
	GuardrailTriggeredAction *string `json:"guardrail_triggered_action,omitempty"`
	// Ordered steps that define changes in traffic exposure.
	RolloutSteps []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems `json:"rollout_steps,omitempty"`
	// Interval between traffic selections, in milliseconds.
	SelectionIntervalMs *int64 `json:"selection_interval_ms,omitempty"`
	// Method used to increase traffic exposure over the schedule.
	Strategy *string `json:"strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesExposureSchedule instantiates a new ExperimentsPublicProtocolResponseDataAttributesExposureSchedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesExposureSchedule() *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule {
	this := ExperimentsPublicProtocolResponseDataAttributesExposureSchedule{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesExposureSchedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesExposureScheduleWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule {
	this := ExperimentsPublicProtocolResponseDataAttributesExposureSchedule{}
	return &this
}

// GetAutostart returns the Autostart field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetAutostart() bool {
	if o == nil || o.Autostart == nil {
		var ret bool
		return ret
	}
	return *o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetAutostartOk() (*bool, bool) {
	if o == nil || o.Autostart == nil {
		return nil, false
	}
	return o.Autostart, true
}

// HasAutostart returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) HasAutostart() bool {
	return o != nil && o.Autostart != nil
}

// SetAutostart gets a reference to the given bool and assigns it to the Autostart field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) SetAutostart(v bool) {
	o.Autostart = &v
}

// GetGuardrailTriggeredAction returns the GuardrailTriggeredAction field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetGuardrailTriggeredAction() string {
	if o == nil || o.GuardrailTriggeredAction == nil {
		var ret string
		return ret
	}
	return *o.GuardrailTriggeredAction
}

// GetGuardrailTriggeredActionOk returns a tuple with the GuardrailTriggeredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetGuardrailTriggeredActionOk() (*string, bool) {
	if o == nil || o.GuardrailTriggeredAction == nil {
		return nil, false
	}
	return o.GuardrailTriggeredAction, true
}

// HasGuardrailTriggeredAction returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) HasGuardrailTriggeredAction() bool {
	return o != nil && o.GuardrailTriggeredAction != nil
}

// SetGuardrailTriggeredAction gets a reference to the given string and assigns it to the GuardrailTriggeredAction field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) SetGuardrailTriggeredAction(v string) {
	o.GuardrailTriggeredAction = &v
}

// GetRolloutSteps returns the RolloutSteps field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetRolloutSteps() []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems {
	if o == nil || o.RolloutSteps == nil {
		var ret []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems
		return ret
	}
	return o.RolloutSteps
}

// GetRolloutStepsOk returns a tuple with the RolloutSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetRolloutStepsOk() (*[]ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems, bool) {
	if o == nil || o.RolloutSteps == nil {
		return nil, false
	}
	return &o.RolloutSteps, true
}

// HasRolloutSteps returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) HasRolloutSteps() bool {
	return o != nil && o.RolloutSteps != nil
}

// SetRolloutSteps gets a reference to the given []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems and assigns it to the RolloutSteps field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) SetRolloutSteps(v []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems) {
	o.RolloutSteps = v
}

// GetSelectionIntervalMs returns the SelectionIntervalMs field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetSelectionIntervalMs() int64 {
	if o == nil || o.SelectionIntervalMs == nil {
		var ret int64
		return ret
	}
	return *o.SelectionIntervalMs
}

// GetSelectionIntervalMsOk returns a tuple with the SelectionIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetSelectionIntervalMsOk() (*int64, bool) {
	if o == nil || o.SelectionIntervalMs == nil {
		return nil, false
	}
	return o.SelectionIntervalMs, true
}

// HasSelectionIntervalMs returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) HasSelectionIntervalMs() bool {
	return o != nil && o.SelectionIntervalMs != nil
}

// SetSelectionIntervalMs gets a reference to the given int64 and assigns it to the SelectionIntervalMs field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) SetSelectionIntervalMs(v int64) {
	o.SelectionIntervalMs = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetStrategy() string {
	if o == nil || o.Strategy == nil {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) GetStrategyOk() (*string, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) HasStrategy() bool {
	return o != nil && o.Strategy != nil
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) SetStrategy(v string) {
	o.Strategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Autostart != nil {
		toSerialize["autostart"] = o.Autostart
	}
	if o.GuardrailTriggeredAction != nil {
		toSerialize["guardrail_triggered_action"] = o.GuardrailTriggeredAction
	}
	if o.RolloutSteps != nil {
		toSerialize["rollout_steps"] = o.RolloutSteps
	}
	if o.SelectionIntervalMs != nil {
		toSerialize["selection_interval_ms"] = o.SelectionIntervalMs
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributesExposureSchedule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autostart                *bool                                                                              `json:"autostart,omitempty"`
		GuardrailTriggeredAction *string                                                                            `json:"guardrail_triggered_action,omitempty"`
		RolloutSteps             []ExperimentsPublicProtocolResponseDataAttributesExposureScheduleRolloutStepsItems `json:"rollout_steps,omitempty"`
		SelectionIntervalMs      *int64                                                                             `json:"selection_interval_ms,omitempty"`
		Strategy                 *string                                                                            `json:"strategy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"autostart", "guardrail_triggered_action", "rollout_steps", "selection_interval_ms", "strategy"})
	} else {
		return err
	}
	o.Autostart = all.Autostart
	o.GuardrailTriggeredAction = all.GuardrailTriggeredAction
	o.RolloutSteps = all.RolloutSteps
	o.SelectionIntervalMs = all.SelectionIntervalMs
	o.Strategy = all.Strategy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
