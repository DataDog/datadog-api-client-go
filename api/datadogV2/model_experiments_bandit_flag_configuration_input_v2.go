// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditFlagConfigurationInputV2 The flag and environment the bandit serves on, and who it serves there. The server chooses the rule's `allocation_key`. On create, `feature_flag_id` and `environment_id` are required, and `null` is the same as omitting it. On update, omitted fields are unchanged, and `null` takes a draft off its flag, which also clears its variants and scheduled end. Moving a draft to another flag clears its variants unless the request sets them. Only a draft can change it.
type ExperimentsBanditFlagConfigurationInputV2 struct {
	// The environment the bandit serves in.
	EnvironmentId *uuid.UUID `json:"environment_id,omitempty"`
	// The flag the bandit serves through.
	FeatureFlagId *uuid.UUID `json:"feature_flag_id,omitempty"`
	// Who the bandit serves. When empty, the bandit serves everyone in the environment.
	TargetingRules []ExperimentsTargetingRule `json:"targeting_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditFlagConfigurationInputV2 instantiates a new ExperimentsBanditFlagConfigurationInputV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditFlagConfigurationInputV2() *ExperimentsBanditFlagConfigurationInputV2 {
	this := ExperimentsBanditFlagConfigurationInputV2{}
	return &this
}

// NewExperimentsBanditFlagConfigurationInputV2WithDefaults instantiates a new ExperimentsBanditFlagConfigurationInputV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditFlagConfigurationInputV2WithDefaults() *ExperimentsBanditFlagConfigurationInputV2 {
	this := ExperimentsBanditFlagConfigurationInputV2{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given uuid.UUID and assigns it to the EnvironmentId field.
func (o *ExperimentsBanditFlagConfigurationInputV2) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = &v
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetFeatureFlagId() uuid.UUID {
	if o == nil || o.FeatureFlagId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetFeatureFlagIdOk() (*uuid.UUID, bool) {
	if o == nil || o.FeatureFlagId == nil {
		return nil, false
	}
	return o.FeatureFlagId, true
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId != nil
}

// SetFeatureFlagId gets a reference to the given uuid.UUID and assigns it to the FeatureFlagId field.
func (o *ExperimentsBanditFlagConfigurationInputV2) SetFeatureFlagId(v uuid.UUID) {
	o.FeatureFlagId = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetTargetingRules() []ExperimentsTargetingRule {
	if o == nil || o.TargetingRules == nil {
		var ret []ExperimentsTargetingRule
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) GetTargetingRulesOk() (*[]ExperimentsTargetingRule, bool) {
	if o == nil || o.TargetingRules == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *ExperimentsBanditFlagConfigurationInputV2) HasTargetingRules() bool {
	return o != nil && o.TargetingRules != nil
}

// SetTargetingRules gets a reference to the given []ExperimentsTargetingRule and assigns it to the TargetingRules field.
func (o *ExperimentsBanditFlagConfigurationInputV2) SetTargetingRules(v []ExperimentsTargetingRule) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditFlagConfigurationInputV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
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
func (o *ExperimentsBanditFlagConfigurationInputV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentId  *uuid.UUID                 `json:"environment_id,omitempty"`
		FeatureFlagId  *uuid.UUID                 `json:"feature_flag_id,omitempty"`
		TargetingRules []ExperimentsTargetingRule `json:"targeting_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"environment_id", "feature_flag_id", "targeting_rules"})
	} else {
		return err
	}
	o.EnvironmentId = all.EnvironmentId
	o.FeatureFlagId = all.FeatureFlagId
	o.TargetingRules = all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsBanditFlagConfigurationInputV2 handles when a null is used for ExperimentsBanditFlagConfigurationInputV2.
type NullableExperimentsBanditFlagConfigurationInputV2 struct {
	value *ExperimentsBanditFlagConfigurationInputV2
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditFlagConfigurationInputV2) Get() *ExperimentsBanditFlagConfigurationInputV2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditFlagConfigurationInputV2) Set(val *ExperimentsBanditFlagConfigurationInputV2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditFlagConfigurationInputV2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsBanditFlagConfigurationInputV2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditFlagConfigurationInputV2 initializes the struct as if Set has been called.
func NewNullableExperimentsBanditFlagConfigurationInputV2(val *ExperimentsBanditFlagConfigurationInputV2) *NullableExperimentsBanditFlagConfigurationInputV2 {
	return &NullableExperimentsBanditFlagConfigurationInputV2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditFlagConfigurationInputV2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditFlagConfigurationInputV2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
