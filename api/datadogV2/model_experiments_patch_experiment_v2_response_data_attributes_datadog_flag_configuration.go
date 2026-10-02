// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration Feature flag, environment, and targeting configuration for the experiment.
type ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration struct {
	// Key of the feature flag allocation linked to the experiment.
	AllocationKey *string `json:"allocation_key,omitempty"`
	// Datadog measure and filters used to select analyzed subjects.
	EntryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint `json:"entry_point"`
	// ID of the feature flag environment used by the experiment.
	EnvironmentId *string `json:"environment_id,omitempty"`
	// ID of the feature flag linked to the experiment.
	FeatureFlagId *string `json:"feature_flag_id,omitempty"`
	// Rules that select subjects for the experiment.
	TargetingRules []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration(entryPoint NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint) *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration{}
	this.EntryPoint = entryPoint
	return &this
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationWithDefaults instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationWithDefaults() *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration{}
	return &this
}

// GetAllocationKey returns the AllocationKey field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetAllocationKey() string {
	if o == nil || o.AllocationKey == nil {
		var ret string
		return ret
	}
	return *o.AllocationKey
}

// GetAllocationKeyOk returns a tuple with the AllocationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetAllocationKeyOk() (*string, bool) {
	if o == nil || o.AllocationKey == nil {
		return nil, false
	}
	return o.AllocationKey, true
}

// HasAllocationKey returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) HasAllocationKey() bool {
	return o != nil && o.AllocationKey != nil
}

// SetAllocationKey gets a reference to the given string and assigns it to the AllocationKey field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) SetAllocationKey(v string) {
	o.AllocationKey = &v
}

// GetEntryPoint returns the EntryPoint field value.
// If the value is explicit nil, the zero value for ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetEntryPoint() ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint {
	if o == nil || o.EntryPoint.Get() == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint
		return ret
	}
	return *o.EntryPoint.Get()
}

// GetEntryPointOk returns a tuple with the EntryPoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetEntryPointOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryPoint.Get(), o.EntryPoint.IsSet()
}

// SetEntryPoint sets field value.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) SetEntryPoint(v ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint) {
	o.EntryPoint.Set(&v)
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetFeatureFlagId() string {
	if o == nil || o.FeatureFlagId == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetFeatureFlagIdOk() (*string, bool) {
	if o == nil || o.FeatureFlagId == nil {
		return nil, false
	}
	return o.FeatureFlagId, true
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId != nil
}

// SetFeatureFlagId gets a reference to the given string and assigns it to the FeatureFlagId field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) SetFeatureFlagId(v string) {
	o.FeatureFlagId = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetTargetingRules() []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems {
	if o == nil || o.TargetingRules == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) GetTargetingRulesOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems and assigns it to the TargetingRules field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) SetTargetingRules(v []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllocationKey != nil {
		toSerialize["allocation_key"] = o.AllocationKey
	}
	toSerialize["entry_point"] = o.EntryPoint.Get()
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if o.FeatureFlagId != nil {
		toSerialize["feature_flag_id"] = o.FeatureFlagId
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
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationKey  *string                                                                                         `json:"allocation_key,omitempty"`
		EntryPoint     NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfigurationEntryPoint    `json:"entry_point"`
		EnvironmentId  *string                                                                                         `json:"environment_id,omitempty"`
		FeatureFlagId  *string                                                                                         `json:"feature_flag_id,omitempty"`
		TargetingRules []ExperimentsCreateExperimentV2RequestDataAttributesDatadogFlagConfigurationTargetingRulesItems `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.EntryPoint.IsSet() {
		return fmt.Errorf("required field entry_point missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocation_key", "entry_point", "environment_id", "feature_flag_id", "targeting_rules"})
	} else {
		return err
	}
	o.AllocationKey = all.AllocationKey
	o.EntryPoint = all.EntryPoint
	o.EnvironmentId = all.EnvironmentId
	o.FeatureFlagId = all.FeatureFlagId
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration handles when a null is used for ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration.
type NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration struct {
	value *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) Get() *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) Set(val *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration initializes the struct as if Set has been called.
func NewNullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration(val *ExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) *NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration {
	return &NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsPatchExperimentV2ResponseDataAttributesDatadogFlagConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
