// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentConfigureV2DryRunResult Validation result of a configuration deployment dry run.
type FleetDeploymentConfigureV2DryRunResult struct {
	// Whether the configuration passed schema validation.
	ConfigValidated *bool `json:"config_validated,omitempty"`
	// Breakdown of ineligible host counts by reason. Only includes reasons with a
	// non-zero count. Absent from the response when no targeted host is ineligible.
	NonUpgradableByReason map[string]int64 `json:"non_upgradable_by_reason,omitempty"`
	// Number of targeted hosts that are not eligible to receive this configuration.
	NonUpgradableHosts *int64 `json:"non_upgradable_hosts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentConfigureV2DryRunResult instantiates a new FleetDeploymentConfigureV2DryRunResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentConfigureV2DryRunResult() *FleetDeploymentConfigureV2DryRunResult {
	this := FleetDeploymentConfigureV2DryRunResult{}
	return &this
}

// NewFleetDeploymentConfigureV2DryRunResultWithDefaults instantiates a new FleetDeploymentConfigureV2DryRunResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentConfigureV2DryRunResultWithDefaults() *FleetDeploymentConfigureV2DryRunResult {
	this := FleetDeploymentConfigureV2DryRunResult{}
	return &this
}

// GetConfigValidated returns the ConfigValidated field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunResult) GetConfigValidated() bool {
	if o == nil || o.ConfigValidated == nil {
		var ret bool
		return ret
	}
	return *o.ConfigValidated
}

// GetConfigValidatedOk returns a tuple with the ConfigValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) GetConfigValidatedOk() (*bool, bool) {
	if o == nil || o.ConfigValidated == nil {
		return nil, false
	}
	return o.ConfigValidated, true
}

// HasConfigValidated returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) HasConfigValidated() bool {
	return o != nil && o.ConfigValidated != nil
}

// SetConfigValidated gets a reference to the given bool and assigns it to the ConfigValidated field.
func (o *FleetDeploymentConfigureV2DryRunResult) SetConfigValidated(v bool) {
	o.ConfigValidated = &v
}

// GetNonUpgradableByReason returns the NonUpgradableByReason field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunResult) GetNonUpgradableByReason() map[string]int64 {
	if o == nil || o.NonUpgradableByReason == nil {
		var ret map[string]int64
		return ret
	}
	return o.NonUpgradableByReason
}

// GetNonUpgradableByReasonOk returns a tuple with the NonUpgradableByReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) GetNonUpgradableByReasonOk() (*map[string]int64, bool) {
	if o == nil || o.NonUpgradableByReason == nil {
		return nil, false
	}
	return &o.NonUpgradableByReason, true
}

// HasNonUpgradableByReason returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) HasNonUpgradableByReason() bool {
	return o != nil && o.NonUpgradableByReason != nil
}

// SetNonUpgradableByReason gets a reference to the given map[string]int64 and assigns it to the NonUpgradableByReason field.
func (o *FleetDeploymentConfigureV2DryRunResult) SetNonUpgradableByReason(v map[string]int64) {
	o.NonUpgradableByReason = v
}

// GetNonUpgradableHosts returns the NonUpgradableHosts field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunResult) GetNonUpgradableHosts() int64 {
	if o == nil || o.NonUpgradableHosts == nil {
		var ret int64
		return ret
	}
	return *o.NonUpgradableHosts
}

// GetNonUpgradableHostsOk returns a tuple with the NonUpgradableHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) GetNonUpgradableHostsOk() (*int64, bool) {
	if o == nil || o.NonUpgradableHosts == nil {
		return nil, false
	}
	return o.NonUpgradableHosts, true
}

// HasNonUpgradableHosts returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunResult) HasNonUpgradableHosts() bool {
	return o != nil && o.NonUpgradableHosts != nil
}

// SetNonUpgradableHosts gets a reference to the given int64 and assigns it to the NonUpgradableHosts field.
func (o *FleetDeploymentConfigureV2DryRunResult) SetNonUpgradableHosts(v int64) {
	o.NonUpgradableHosts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentConfigureV2DryRunResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfigValidated != nil {
		toSerialize["config_validated"] = o.ConfigValidated
	}
	if o.NonUpgradableByReason != nil {
		toSerialize["non_upgradable_by_reason"] = o.NonUpgradableByReason
	}
	if o.NonUpgradableHosts != nil {
		toSerialize["non_upgradable_hosts"] = o.NonUpgradableHosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentConfigureV2DryRunResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigValidated       *bool            `json:"config_validated,omitempty"`
		NonUpgradableByReason map[string]int64 `json:"non_upgradable_by_reason,omitempty"`
		NonUpgradableHosts    *int64           `json:"non_upgradable_hosts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_validated", "non_upgradable_by_reason", "non_upgradable_hosts"})
	} else {
		return err
	}
	o.ConfigValidated = all.ConfigValidated
	o.NonUpgradableByReason = all.NonUpgradableByReason
	o.NonUpgradableHosts = all.NonUpgradableHosts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
