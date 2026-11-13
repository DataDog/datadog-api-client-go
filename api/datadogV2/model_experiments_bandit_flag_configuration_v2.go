// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditFlagConfigurationV2 The flag rule the bandit serves through. Absent until a draft is put on a flag, and when the caller lacks feature flag and environment read permissions (`meta.warnings` then says so).
type ExperimentsBanditFlagConfigurationV2 struct {
	// The key of the bandit's rule on the flag. Exposures log it.
	AllocationKey string `json:"allocation_key"`
	// The environment the bandit serves in.
	EnvironmentId uuid.UUID `json:"environment_id"`
	// The flag the bandit serves through.
	FeatureFlagId uuid.UUID `json:"feature_flag_id"`
	// Who the bandit serves. Empty serves everyone in the environment.
	TargetingRules []ExperimentsTargetingRule `json:"targeting_rules"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditFlagConfigurationV2 instantiates a new ExperimentsBanditFlagConfigurationV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditFlagConfigurationV2(allocationKey string, environmentId uuid.UUID, featureFlagId uuid.UUID, targetingRules []ExperimentsTargetingRule) *ExperimentsBanditFlagConfigurationV2 {
	this := ExperimentsBanditFlagConfigurationV2{}
	this.AllocationKey = allocationKey
	this.EnvironmentId = environmentId
	this.FeatureFlagId = featureFlagId
	this.TargetingRules = targetingRules
	return &this
}

// NewExperimentsBanditFlagConfigurationV2WithDefaults instantiates a new ExperimentsBanditFlagConfigurationV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditFlagConfigurationV2WithDefaults() *ExperimentsBanditFlagConfigurationV2 {
	this := ExperimentsBanditFlagConfigurationV2{}
	return &this
}

// GetAllocationKey returns the AllocationKey field value.
func (o *ExperimentsBanditFlagConfigurationV2) GetAllocationKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AllocationKey
}

// GetAllocationKeyOk returns a tuple with the AllocationKey field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationV2) GetAllocationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocationKey, true
}

// SetAllocationKey sets field value.
func (o *ExperimentsBanditFlagConfigurationV2) SetAllocationKey(v string) {
	o.AllocationKey = v
}

// GetEnvironmentId returns the EnvironmentId field value.
func (o *ExperimentsBanditFlagConfigurationV2) GetEnvironmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationV2) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *ExperimentsBanditFlagConfigurationV2) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = v
}

// GetFeatureFlagId returns the FeatureFlagId field value.
func (o *ExperimentsBanditFlagConfigurationV2) GetFeatureFlagId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationV2) GetFeatureFlagIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagId, true
}

// SetFeatureFlagId sets field value.
func (o *ExperimentsBanditFlagConfigurationV2) SetFeatureFlagId(v uuid.UUID) {
	o.FeatureFlagId = v
}

// GetTargetingRules returns the TargetingRules field value.
func (o *ExperimentsBanditFlagConfigurationV2) GetTargetingRules() []ExperimentsTargetingRule {
	if o == nil {
		var ret []ExperimentsTargetingRule
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditFlagConfigurationV2) GetTargetingRulesOk() (*[]ExperimentsTargetingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingRules, true
}

// SetTargetingRules sets field value.
func (o *ExperimentsBanditFlagConfigurationV2) SetTargetingRules(v []ExperimentsTargetingRule) {
	o.TargetingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditFlagConfigurationV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allocation_key"] = o.AllocationKey
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["feature_flag_id"] = o.FeatureFlagId
	toSerialize["targeting_rules"] = o.TargetingRules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditFlagConfigurationV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationKey  *string                     `json:"allocation_key"`
		EnvironmentId  *uuid.UUID                  `json:"environment_id"`
		FeatureFlagId  *uuid.UUID                  `json:"feature_flag_id"`
		TargetingRules *[]ExperimentsTargetingRule `json:"targeting_rules"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllocationKey == nil {
		return fmt.Errorf("required field allocation_key missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"allocation_key", "environment_id", "feature_flag_id", "targeting_rules"})
	} else {
		return err
	}
	o.AllocationKey = *all.AllocationKey
	o.EnvironmentId = *all.EnvironmentId
	o.FeatureFlagId = *all.FeatureFlagId
	o.TargetingRules = *all.TargetingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
