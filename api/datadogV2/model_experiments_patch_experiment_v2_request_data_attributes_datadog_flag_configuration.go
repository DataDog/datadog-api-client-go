// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration Feature flag, environment, and targeting configuration for the experiment.
type ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration struct {
	// Datadog measure and filters used to select analyzed subjects.
	EntryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint `json:"entry_point,omitempty"`
	// ID of the feature flag environment used by the experiment.
	EnvironmentId *string `json:"environment_id,omitempty"`
	// ID of the feature flag linked to the experiment.
	FeatureFlagId *string `json:"feature_flag_id,omitempty"`
	// Set true to replace an existing draft allocation when feature_flag_id changes. The replacement resets all flag-bound randomization state.
	ResetOnFeatureFlagChange *bool `json:"reset_on_feature_flag_change,omitempty"`
	// Omit to keep the stored rules. Use an empty array to remove targeting rules.
	TargetingRules []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration instantiates a new ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration() *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	this := ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration{}
	return &this
}

// NewExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfigurationWithDefaults instantiates a new ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfigurationWithDefaults() *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	this := ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration{}
	return &this
}

// GetEntryPoint returns the EntryPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEntryPoint() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint {
	if o == nil || o.EntryPoint.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint
		return ret
	}
	return *o.EntryPoint.Get()
}

// GetEntryPointOk returns a tuple with the EntryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEntryPointOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryPoint.Get(), o.EntryPoint.IsSet()
}

// HasEntryPoint returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasEntryPoint() bool {
	return o != nil && o.EntryPoint.IsSet()
}

// SetEntryPoint gets a reference to the given NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint and assigns it to the EntryPoint field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetEntryPoint(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint) {
	o.EntryPoint.Set(&v)
}

// SetEntryPointNil sets the value for EntryPoint to be an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetEntryPointNil() {
	o.EntryPoint.Set(nil)
}

// UnsetEntryPoint ensures that no value is present for EntryPoint, not even an explicit nil.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) UnsetEntryPoint() {
	o.EntryPoint.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetFeatureFlagId() string {
	if o == nil || o.FeatureFlagId == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetFeatureFlagIdOk() (*string, bool) {
	if o == nil || o.FeatureFlagId == nil {
		return nil, false
	}
	return o.FeatureFlagId, true
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId != nil
}

// SetFeatureFlagId gets a reference to the given string and assigns it to the FeatureFlagId field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetFeatureFlagId(v string) {
	o.FeatureFlagId = &v
}

// GetResetOnFeatureFlagChange returns the ResetOnFeatureFlagChange field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetResetOnFeatureFlagChange() bool {
	if o == nil || o.ResetOnFeatureFlagChange == nil {
		var ret bool
		return ret
	}
	return *o.ResetOnFeatureFlagChange
}

// GetResetOnFeatureFlagChangeOk returns a tuple with the ResetOnFeatureFlagChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetResetOnFeatureFlagChangeOk() (*bool, bool) {
	if o == nil || o.ResetOnFeatureFlagChange == nil {
		return nil, false
	}
	return o.ResetOnFeatureFlagChange, true
}

// HasResetOnFeatureFlagChange returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasResetOnFeatureFlagChange() bool {
	return o != nil && o.ResetOnFeatureFlagChange != nil
}

// SetResetOnFeatureFlagChange gets a reference to the given bool and assigns it to the ResetOnFeatureFlagChange field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetResetOnFeatureFlagChange(v bool) {
	o.ResetOnFeatureFlagChange = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetTargetingRules() []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems {
	if o == nil || o.TargetingRules == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) GetTargetingRulesOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems and assigns it to the TargetingRules field.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) SetTargetingRules(v []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EntryPoint.IsSet() {
		toSerialize["entry_point"] = o.EntryPoint.Get()
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.FeatureFlagId != nil {
		toSerialize["feature_flag_id"] = o.FeatureFlagId
	}
	if o.ResetOnFeatureFlagChange != nil {
		toSerialize["reset_on_feature_flag_change"] = o.ResetOnFeatureFlagChange
	}
	if o.TargetingRules != nil {
		toSerialize["targeting_rules"] = o.TargetingRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntryPoint               NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint    `json:"entry_point,omitempty"`
		EnvironmentId            *string                                                                                         `json:"environment_id,omitempty"`
		FeatureFlagId            *string                                                                                         `json:"feature_flag_id,omitempty"`
		ResetOnFeatureFlagChange *bool                                                                                           `json:"reset_on_feature_flag_change,omitempty"`
		TargetingRules           []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entry_point", "environment_id", "feature_flag_id", "reset_on_feature_flag_change", "targeting_rules"})
	} else {
		return err
	}
	o.EntryPoint = all.EntryPoint
	o.EnvironmentId = all.EnvironmentId
	o.FeatureFlagId = all.FeatureFlagId
	o.ResetOnFeatureFlagChange = all.ResetOnFeatureFlagChange
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration handles when a null is used for ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration.
type NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration struct {
	value *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) Get() *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) Set(val *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration initializes the struct as if Set has been called.
func NewNullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration(val *ExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) *NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration {
	return &NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsPatchExperimentV2RequestDataAttributesDatadogFlagConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
