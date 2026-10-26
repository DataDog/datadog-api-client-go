// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchProtocolExposureScheduleInput Partial traffic schedule. Omitted fields remain unchanged. Null clears the whole schedule. Null clears nullable fields and rollout_steps; null sets strategy to an empty string and `autostart` to false.
type ExperimentsPatchProtocolExposureScheduleInput struct {
	// Whether the exposure schedule starts automatically.
	Autostart datadog.NullableBool `json:"autostart,omitempty"`
	// Action taken when a guardrail triggers during the exposure schedule.
	GuardrailTriggeredAction datadog.NullableString `json:"guardrail_triggered_action,omitempty"`
	// Rollout steps, replaced when supplied.
	RolloutSteps []ExperimentsProtocolExposureStepInput `json:"rollout_steps,omitempty"`
	// Interval between traffic selections, in milliseconds.
	SelectionIntervalMs datadog.NullableInt64 `json:"selection_interval_ms,omitempty"`
	// Method used to increase traffic exposure over the schedule.
	Strategy datadog.NullableString `json:"strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchProtocolExposureScheduleInput instantiates a new ExperimentsPatchProtocolExposureScheduleInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchProtocolExposureScheduleInput() *ExperimentsPatchProtocolExposureScheduleInput {
	this := ExperimentsPatchProtocolExposureScheduleInput{}
	return &this
}

// NewExperimentsPatchProtocolExposureScheduleInputWithDefaults instantiates a new ExperimentsPatchProtocolExposureScheduleInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchProtocolExposureScheduleInputWithDefaults() *ExperimentsPatchProtocolExposureScheduleInput {
	this := ExperimentsPatchProtocolExposureScheduleInput{}
	return &this
}

// GetAutostart returns the Autostart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetAutostart() bool {
	if o == nil || o.Autostart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Autostart.Get()
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetAutostartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autostart.Get(), o.Autostart.IsSet()
}

// HasAutostart returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolExposureScheduleInput) HasAutostart() bool {
	return o != nil && o.Autostart.IsSet()
}

// SetAutostart gets a reference to the given datadog.NullableBool and assigns it to the Autostart field.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetAutostart(v bool) {
	o.Autostart.Set(&v)
}

// SetAutostartNil sets the value for Autostart to be an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetAutostartNil() {
	o.Autostart.Set(nil)
}

// UnsetAutostart ensures that no value is present for Autostart, not even an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) UnsetAutostart() {
	o.Autostart.Unset()
}

// GetGuardrailTriggeredAction returns the GuardrailTriggeredAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetGuardrailTriggeredAction() string {
	if o == nil || o.GuardrailTriggeredAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.GuardrailTriggeredAction.Get()
}

// GetGuardrailTriggeredActionOk returns a tuple with the GuardrailTriggeredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetGuardrailTriggeredActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuardrailTriggeredAction.Get(), o.GuardrailTriggeredAction.IsSet()
}

// HasGuardrailTriggeredAction returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolExposureScheduleInput) HasGuardrailTriggeredAction() bool {
	return o != nil && o.GuardrailTriggeredAction.IsSet()
}

// SetGuardrailTriggeredAction gets a reference to the given datadog.NullableString and assigns it to the GuardrailTriggeredAction field.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetGuardrailTriggeredAction(v string) {
	o.GuardrailTriggeredAction.Set(&v)
}

// SetGuardrailTriggeredActionNil sets the value for GuardrailTriggeredAction to be an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetGuardrailTriggeredActionNil() {
	o.GuardrailTriggeredAction.Set(nil)
}

// UnsetGuardrailTriggeredAction ensures that no value is present for GuardrailTriggeredAction, not even an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) UnsetGuardrailTriggeredAction() {
	o.GuardrailTriggeredAction.Unset()
}

// GetRolloutSteps returns the RolloutSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetRolloutSteps() []ExperimentsProtocolExposureStepInput {
	if o == nil {
		var ret []ExperimentsProtocolExposureStepInput
		return ret
	}
	return o.RolloutSteps
}

// GetRolloutStepsOk returns a tuple with the RolloutSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetRolloutStepsOk() (*[]ExperimentsProtocolExposureStepInput, bool) {
	if o == nil || o.RolloutSteps == nil {
		return nil, false
	}
	return &o.RolloutSteps, true
}

// HasRolloutSteps returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolExposureScheduleInput) HasRolloutSteps() bool {
	return o != nil && o.RolloutSteps != nil
}

// SetRolloutSteps gets a reference to the given []ExperimentsProtocolExposureStepInput and assigns it to the RolloutSteps field.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetRolloutSteps(v []ExperimentsProtocolExposureStepInput) {
	o.RolloutSteps = v
}

// GetSelectionIntervalMs returns the SelectionIntervalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetSelectionIntervalMs() int64 {
	if o == nil || o.SelectionIntervalMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SelectionIntervalMs.Get()
}

// GetSelectionIntervalMsOk returns a tuple with the SelectionIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetSelectionIntervalMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectionIntervalMs.Get(), o.SelectionIntervalMs.IsSet()
}

// HasSelectionIntervalMs returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolExposureScheduleInput) HasSelectionIntervalMs() bool {
	return o != nil && o.SelectionIntervalMs.IsSet()
}

// SetSelectionIntervalMs gets a reference to the given datadog.NullableInt64 and assigns it to the SelectionIntervalMs field.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetSelectionIntervalMs(v int64) {
	o.SelectionIntervalMs.Set(&v)
}

// SetSelectionIntervalMsNil sets the value for SelectionIntervalMs to be an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetSelectionIntervalMsNil() {
	o.SelectionIntervalMs.Set(nil)
}

// UnsetSelectionIntervalMs ensures that no value is present for SelectionIntervalMs, not even an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) UnsetSelectionIntervalMs() {
	o.SelectionIntervalMs.Unset()
}

// GetStrategy returns the Strategy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetStrategy() string {
	if o == nil || o.Strategy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Strategy.Get()
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolExposureScheduleInput) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategy.Get(), o.Strategy.IsSet()
}

// HasStrategy returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolExposureScheduleInput) HasStrategy() bool {
	return o != nil && o.Strategy.IsSet()
}

// SetStrategy gets a reference to the given datadog.NullableString and assigns it to the Strategy field.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetStrategy(v string) {
	o.Strategy.Set(&v)
}

// SetStrategyNil sets the value for Strategy to be an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) SetStrategyNil() {
	o.Strategy.Set(nil)
}

// UnsetStrategy ensures that no value is present for Strategy, not even an explicit nil.
func (o *ExperimentsPatchProtocolExposureScheduleInput) UnsetStrategy() {
	o.Strategy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchProtocolExposureScheduleInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Autostart.IsSet() {
		toSerialize["autostart"] = o.Autostart.Get()
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
	if o.Strategy.IsSet() {
		toSerialize["strategy"] = o.Strategy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchProtocolExposureScheduleInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Autostart                datadog.NullableBool                   `json:"autostart,omitempty"`
		GuardrailTriggeredAction datadog.NullableString                 `json:"guardrail_triggered_action,omitempty"`
		RolloutSteps             []ExperimentsProtocolExposureStepInput `json:"rollout_steps,omitempty"`
		SelectionIntervalMs      datadog.NullableInt64                  `json:"selection_interval_ms,omitempty"`
		Strategy                 datadog.NullableString                 `json:"strategy,omitempty"`
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

// NullableExperimentsPatchProtocolExposureScheduleInput handles when a null is used for ExperimentsPatchProtocolExposureScheduleInput.
type NullableExperimentsPatchProtocolExposureScheduleInput struct {
	value *ExperimentsPatchProtocolExposureScheduleInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsPatchProtocolExposureScheduleInput) Get() *ExperimentsPatchProtocolExposureScheduleInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsPatchProtocolExposureScheduleInput) Set(val *ExperimentsPatchProtocolExposureScheduleInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsPatchProtocolExposureScheduleInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsPatchProtocolExposureScheduleInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsPatchProtocolExposureScheduleInput initializes the struct as if Set has been called.
func NewNullableExperimentsPatchProtocolExposureScheduleInput(val *ExperimentsPatchProtocolExposureScheduleInput) *NullableExperimentsPatchProtocolExposureScheduleInput {
	return &NullableExperimentsPatchProtocolExposureScheduleInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsPatchProtocolExposureScheduleInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsPatchProtocolExposureScheduleInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
