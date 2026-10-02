// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration Feature flag, environment, and targeting configuration for a Datadog experiment.
type ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration struct {
	// Datadog measure and filters used to select analyzed subjects.
	EntryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint `json:"entry_point"`
	// Identifier of the feature flag environment.
	EnvironmentId uuid.UUID `json:"environment_id"`
	// Identifier of the Datadog feature flag used by the experiment.
	FeatureFlagId uuid.UUID `json:"feature_flag_id"`
	// Accepted on create but has no effect.
	ResetOnFeatureFlagChange *bool `json:"reset_on_feature_flag_change,omitempty"`
	// Use an empty array when no targeting rules apply.
	TargetingRules []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration(entryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint, environmentId uuid.UUID, featureFlagId uuid.UUID, targetingRules []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration{}
	this.EntryPoint = entryPoint
	this.EnvironmentId = environmentId
	this.FeatureFlagId = featureFlagId
	this.TargetingRules = targetingRules
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration{}
	return &this
}

// GetEntryPoint returns the EntryPoint field value.
// If the value is explicit nil, the zero value for ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEntryPoint() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint {
	if o == nil || o.EntryPoint.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint
		return ret
	}
	return *o.EntryPoint.Get()
}

// GetEntryPointOk returns a tuple with the EntryPoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEntryPointOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryPoint.Get(), o.EntryPoint.IsSet()
}

// SetEntryPoint sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetEntryPoint(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint) {
	o.EntryPoint.Set(&v)
}

// GetEnvironmentId returns the EnvironmentId field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEnvironmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = v
}

// GetFeatureFlagId returns the FeatureFlagId field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetFeatureFlagId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetFeatureFlagIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagId, true
}

// SetFeatureFlagId sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetFeatureFlagId(v uuid.UUID) {
	o.FeatureFlagId = v
}

// GetResetOnFeatureFlagChange returns the ResetOnFeatureFlagChange field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetResetOnFeatureFlagChange() bool {
	if o == nil || o.ResetOnFeatureFlagChange == nil {
		var ret bool
		return ret
	}
	return *o.ResetOnFeatureFlagChange
}

// GetResetOnFeatureFlagChangeOk returns a tuple with the ResetOnFeatureFlagChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetResetOnFeatureFlagChangeOk() (*bool, bool) {
	if o == nil || o.ResetOnFeatureFlagChange == nil {
		return nil, false
	}
	return o.ResetOnFeatureFlagChange, true
}

// HasResetOnFeatureFlagChange returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasResetOnFeatureFlagChange() bool {
	return o != nil && o.ResetOnFeatureFlagChange != nil
}

// SetResetOnFeatureFlagChange gets a reference to the given bool and assigns it to the ResetOnFeatureFlagChange field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetResetOnFeatureFlagChange(v bool) {
	o.ResetOnFeatureFlagChange = &v
}

// GetTargetingRules returns the TargetingRules field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetTargetingRules() []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems {
	if o == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetTargetingRulesOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// SetTargetingRules sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetTargetingRules(v []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entry_point"] = o.EntryPoint.Get()
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["feature_flag_id"] = o.FeatureFlagId
	if o.ResetOnFeatureFlagChange != nil {
		toSerialize["reset_on_feature_flag_change"] = o.ResetOnFeatureFlagChange
	}
	toSerialize["targeting_rules"] = o.TargetingRules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntryPoint               NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint     `json:"entry_point"`
		EnvironmentId            *uuid.UUID                                                                                       `json:"environment_id"`
		FeatureFlagId            *uuid.UUID                                                                                       `json:"feature_flag_id"`
		ResetOnFeatureFlagChange *bool                                                                                            `json:"reset_on_feature_flag_change,omitempty"`
		TargetingRules           *[]ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.EntryPoint.IsSet() {
		return fmt.Errorf("required field entry_point missing")
	}
	if all.EnvironmentId == nil {
		return fmt.Errorf("required field environment_id missing")
	}
	if all.FeatureFlagId == nil {
		return fmt.Errorf("required field feature_flag_id missing")
	}
	if all.TargetingRules == nil {
		return fmt.Errorf("required field targeting_rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entry_point", "environment_id", "feature_flag_id", "reset_on_feature_flag_change", "targeting_rules"})
	} else {
		return err
	}
	o.EntryPoint = all.EntryPoint
	o.EnvironmentId = *all.EnvironmentId
	o.FeatureFlagId = *all.FeatureFlagId
	o.ResetOnFeatureFlagChange = all.ResetOnFeatureFlagChange
	o.TargetingRules = *all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration handles when a null is used for ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration.
type NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration struct {
	value *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) Get() *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) Set(val *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration initializes the struct as if Set has been called.
func NewNullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration(val *ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) *NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	return &NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
