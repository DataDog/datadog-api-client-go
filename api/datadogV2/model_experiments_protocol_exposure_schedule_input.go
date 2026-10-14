// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsProtocolExposureScheduleInput Traffic schedule. Publishing and replacement require `NO_ROLLOUT` with exactly one step, or `UNIFORM_INTERVALS` with a positive selection interval and at least one step. Requires `environment_id`.
type ExperimentsProtocolExposureScheduleInput struct {
	// Whether the exposure schedule starts automatically.
	Autostart *bool `json:"autostart,omitempty"`
	// Action taken when a guardrail triggers during the exposure schedule.
	GuardrailTriggeredAction datadog.NullableString `json:"guardrail_triggered_action,omitempty"`
	// Rollout steps, replaced when supplied.
	RolloutSteps []ExperimentsProtocolExposureStepInput `json:"rollout_steps,omitempty"`
	// Interval between traffic selections, in milliseconds.
	SelectionIntervalMs datadog.NullableInt64 `json:"selection_interval_ms,omitempty"`
	// Method used to increase traffic exposure over the schedule.
	Strategy *string `json:"strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsProtocolExposureScheduleInput instantiates a new ExperimentsProtocolExposureScheduleInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsProtocolExposureScheduleInput() *ExperimentsProtocolExposureScheduleInput {
	this := ExperimentsProtocolExposureScheduleInput{}
	return &this
}

// NewExperimentsProtocolExposureScheduleInputWithDefaults instantiates a new ExperimentsProtocolExposureScheduleInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsProtocolExposureScheduleInputWithDefaults() *ExperimentsProtocolExposureScheduleInput {
	this := ExperimentsProtocolExposureScheduleInput{}
	return &this
}

// GetAutostart returns the Autostart field value if set, zero value otherwise.
func (o *ExperimentsProtocolExposureScheduleInput) GetAutostart() bool {
	if o == nil || o.Autostart == nil {
		var ret bool
		return ret
	}
	return *o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolExposureScheduleInput) GetAutostartOk() (*bool, bool) {
	if o == nil || o.Autostart == nil {
		return nil, false
	}
	return o.Autostart, true
}

// HasAutostart returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureScheduleInput) HasAutostart() bool {
	return o != nil && o.Autostart != nil
}

// SetAutostart gets a reference to the given bool and assigns it to the Autostart field.
func (o *ExperimentsProtocolExposureScheduleInput) SetAutostart(v bool) {
	o.Autostart = &v
}

// GetGuardrailTriggeredAction returns the GuardrailTriggeredAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolExposureScheduleInput) GetGuardrailTriggeredAction() string {
	if o == nil || o.GuardrailTriggeredAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.GuardrailTriggeredAction.Get()
}

// GetGuardrailTriggeredActionOk returns a tuple with the GuardrailTriggeredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolExposureScheduleInput) GetGuardrailTriggeredActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailTriggeredAction.Get(), o.GuardrailTriggeredAction.IsSet()
}

// HasGuardrailTriggeredAction returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureScheduleInput) HasGuardrailTriggeredAction() bool {
	return o != nil && o.GuardrailTriggeredAction.IsSet()
}

// SetGuardrailTriggeredAction gets a reference to the given datadog.NullableString and assigns it to the GuardrailTriggeredAction field.
func (o *ExperimentsProtocolExposureScheduleInput) SetGuardrailTriggeredAction(v string) {
	o.GuardrailTriggeredAction.Set(&v)
}

// SetGuardrailTriggeredActionNil sets the value for GuardrailTriggeredAction to be an explicit nil.
func (o *ExperimentsProtocolExposureScheduleInput) SetGuardrailTriggeredActionNil() {
	o.GuardrailTriggeredAction.Set(nil)
}

// UnsetGuardrailTriggeredAction ensures that no value is present for GuardrailTriggeredAction, not even an explicit nil.
func (o *ExperimentsProtocolExposureScheduleInput) UnsetGuardrailTriggeredAction() {
	o.GuardrailTriggeredAction.Unset()
}

// GetRolloutSteps returns the RolloutSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolExposureScheduleInput) GetRolloutSteps() []ExperimentsProtocolExposureStepInput {
	if o == nil {
		var ret []ExperimentsProtocolExposureStepInput
		return ret
	}
	return o.RolloutSteps
}

// GetRolloutStepsOk returns a tuple with the RolloutSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolExposureScheduleInput) GetRolloutStepsOk() (*[]ExperimentsProtocolExposureStepInput, bool) {
	if o == nil || o.RolloutSteps == nil {
		return nil, false
	}
	return &o.RolloutSteps, true
}

// HasRolloutSteps returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureScheduleInput) HasRolloutSteps() bool {
	return o != nil && o.RolloutSteps != nil
}

// SetRolloutSteps gets a reference to the given []ExperimentsProtocolExposureStepInput and assigns it to the RolloutSteps field.
func (o *ExperimentsProtocolExposureScheduleInput) SetRolloutSteps(v []ExperimentsProtocolExposureStepInput) {
	o.RolloutSteps = v
}

// GetSelectionIntervalMs returns the SelectionIntervalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolExposureScheduleInput) GetSelectionIntervalMs() int64 {
	if o == nil || o.SelectionIntervalMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SelectionIntervalMs.Get()
}

// GetSelectionIntervalMsOk returns a tuple with the SelectionIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolExposureScheduleInput) GetSelectionIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectionIntervalMs.Get(), o.SelectionIntervalMs.IsSet()
}

// HasSelectionIntervalMs returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureScheduleInput) HasSelectionIntervalMs() bool {
	return o != nil && o.SelectionIntervalMs.IsSet()
}

// SetSelectionIntervalMs gets a reference to the given datadog.NullableInt64 and assigns it to the SelectionIntervalMs field.
func (o *ExperimentsProtocolExposureScheduleInput) SetSelectionIntervalMs(v int64) {
	o.SelectionIntervalMs.Set(&v)
}

// SetSelectionIntervalMsNil sets the value for SelectionIntervalMs to be an explicit nil.
func (o *ExperimentsProtocolExposureScheduleInput) SetSelectionIntervalMsNil() {
	o.SelectionIntervalMs.Set(nil)
}

// UnsetSelectionIntervalMs ensures that no value is present for SelectionIntervalMs, not even an explicit nil.
func (o *ExperimentsProtocolExposureScheduleInput) UnsetSelectionIntervalMs() {
	o.SelectionIntervalMs.Unset()
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ExperimentsProtocolExposureScheduleInput) GetStrategy() string {
	if o == nil || o.Strategy == nil {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolExposureScheduleInput) GetStrategyOk() (*string, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ExperimentsProtocolExposureScheduleInput) HasStrategy() bool {
	return o != nil && o.Strategy != nil
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *ExperimentsProtocolExposureScheduleInput) SetStrategy(v string) {
	o.Strategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsProtocolExposureScheduleInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Autostart != nil {
		toSerialize["autostart"] = o.Autostart
	}
	if o.GuardrailTriggeredAction.IsSet() {
		toSerialize["guardrail_triggered_action"] = o.GuardrailTriggeredAction.Get()
	}
	if o.RolloutSteps != nil {
		toSerialize["rollout_steps"] = o.RolloutSteps
	}
	if o.SelectionIntervalMs.IsSet() {
		toSerialize["selection_interval_ms"] = o.SelectionIntervalMs.Get()
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
func (o *ExperimentsProtocolExposureScheduleInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autostart                *bool                                  `json:"autostart,omitempty"`
		GuardrailTriggeredAction datadog.NullableString                 `json:"guardrail_triggered_action,omitempty"`
		RolloutSteps             []ExperimentsProtocolExposureStepInput `json:"rollout_steps,omitempty"`
		SelectionIntervalMs      datadog.NullableInt64                  `json:"selection_interval_ms,omitempty"`
		Strategy                 *string                                `json:"strategy,omitempty"`
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

// NullableExperimentsProtocolExposureScheduleInput handles when a null is used for ExperimentsProtocolExposureScheduleInput.
type NullableExperimentsProtocolExposureScheduleInput struct {
	value *ExperimentsProtocolExposureScheduleInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsProtocolExposureScheduleInput) Get() *ExperimentsProtocolExposureScheduleInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsProtocolExposureScheduleInput) Set(val *ExperimentsProtocolExposureScheduleInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsProtocolExposureScheduleInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsProtocolExposureScheduleInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsProtocolExposureScheduleInput initializes the struct as if Set has been called.
func NewNullableExperimentsProtocolExposureScheduleInput(val *ExperimentsProtocolExposureScheduleInput) *NullableExperimentsProtocolExposureScheduleInput {
	return &NullableExperimentsProtocolExposureScheduleInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsProtocolExposureScheduleInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsProtocolExposureScheduleInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
