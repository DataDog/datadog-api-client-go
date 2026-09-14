// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsMonitorIds Specific monitor options for deployment rules.
type DeploymentRuleOptionsMonitorIds struct {
	// Seconds the monitors need to stay in OK status for the rule to pass.
	Duration *int64 `json:"duration,omitempty"`
	// Whether the rule should fail if a selected monitor group is in a NO DATA state.
	FailOnNoData *bool `json:"fail_on_no_data,omitempty"`
	// Whether the rule should fail if no monitor groups are found for the selected monitors.
	FailOnNoGroupsFound *bool `json:"fail_on_no_groups_found,omitempty"`
	// A non-empty list of specific monitors to evaluate.
	MonitorIds []DeploymentRuleOptionsMonitorId `json:"monitor_ids"`
	// Seconds to wait after a deployment starts before evaluating the monitors' statuses.
	Warmup *int64 `json:"warmup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsMonitorIds instantiates a new DeploymentRuleOptionsMonitorIds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsMonitorIds(monitorIds []DeploymentRuleOptionsMonitorId) *DeploymentRuleOptionsMonitorIds {
	this := DeploymentRuleOptionsMonitorIds{}
	var failOnNoData bool = true
	this.FailOnNoData = &failOnNoData
	var failOnNoGroupsFound bool = false
	this.FailOnNoGroupsFound = &failOnNoGroupsFound
	this.MonitorIds = monitorIds
	var warmup int64 = 0
	this.Warmup = &warmup
	return &this
}

// NewDeploymentRuleOptionsMonitorIdsWithDefaults instantiates a new DeploymentRuleOptionsMonitorIds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsMonitorIdsWithDefaults() *DeploymentRuleOptionsMonitorIds {
	this := DeploymentRuleOptionsMonitorIds{}
	var failOnNoData bool = true
	this.FailOnNoData = &failOnNoData
	var failOnNoGroupsFound bool = false
	this.FailOnNoGroupsFound = &failOnNoGroupsFound
	var warmup int64 = 0
	this.Warmup = &warmup
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitorIds) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorIds) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitorIds) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRuleOptionsMonitorIds) SetDuration(v int64) {
	o.Duration = &v
}

// GetFailOnNoData returns the FailOnNoData field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitorIds) GetFailOnNoData() bool {
	if o == nil || o.FailOnNoData == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoData
}

// GetFailOnNoDataOk returns a tuple with the FailOnNoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorIds) GetFailOnNoDataOk() (*bool, bool) {
	if o == nil || o.FailOnNoData == nil {
		return nil, false
	}
	return o.FailOnNoData, true
}

// HasFailOnNoData returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitorIds) HasFailOnNoData() bool {
	return o != nil && o.FailOnNoData != nil
}

// SetFailOnNoData gets a reference to the given bool and assigns it to the FailOnNoData field.
func (o *DeploymentRuleOptionsMonitorIds) SetFailOnNoData(v bool) {
	o.FailOnNoData = &v
}

// GetFailOnNoGroupsFound returns the FailOnNoGroupsFound field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitorIds) GetFailOnNoGroupsFound() bool {
	if o == nil || o.FailOnNoGroupsFound == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoGroupsFound
}

// GetFailOnNoGroupsFoundOk returns a tuple with the FailOnNoGroupsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorIds) GetFailOnNoGroupsFoundOk() (*bool, bool) {
	if o == nil || o.FailOnNoGroupsFound == nil {
		return nil, false
	}
	return o.FailOnNoGroupsFound, true
}

// HasFailOnNoGroupsFound returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitorIds) HasFailOnNoGroupsFound() bool {
	return o != nil && o.FailOnNoGroupsFound != nil
}

// SetFailOnNoGroupsFound gets a reference to the given bool and assigns it to the FailOnNoGroupsFound field.
func (o *DeploymentRuleOptionsMonitorIds) SetFailOnNoGroupsFound(v bool) {
	o.FailOnNoGroupsFound = &v
}

// GetMonitorIds returns the MonitorIds field value.
func (o *DeploymentRuleOptionsMonitorIds) GetMonitorIds() []DeploymentRuleOptionsMonitorId {
	if o == nil {
		var ret []DeploymentRuleOptionsMonitorId
		return ret
	}
	return o.MonitorIds
}

// GetMonitorIdsOk returns a tuple with the MonitorIds field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorIds) GetMonitorIdsOk() (*[]DeploymentRuleOptionsMonitorId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorIds, true
}

// SetMonitorIds sets field value.
func (o *DeploymentRuleOptionsMonitorIds) SetMonitorIds(v []DeploymentRuleOptionsMonitorId) {
	o.MonitorIds = v
}

// GetWarmup returns the Warmup field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitorIds) GetWarmup() int64 {
	if o == nil || o.Warmup == nil {
		var ret int64
		return ret
	}
	return *o.Warmup
}

// GetWarmupOk returns a tuple with the Warmup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitorIds) GetWarmupOk() (*int64, bool) {
	if o == nil || o.Warmup == nil {
		return nil, false
	}
	return o.Warmup, true
}

// HasWarmup returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitorIds) HasWarmup() bool {
	return o != nil && o.Warmup != nil
}

// SetWarmup gets a reference to the given int64 and assigns it to the Warmup field.
func (o *DeploymentRuleOptionsMonitorIds) SetWarmup(v int64) {
	o.Warmup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsMonitorIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.FailOnNoData != nil {
		toSerialize["fail_on_no_data"] = o.FailOnNoData
	}
	if o.FailOnNoGroupsFound != nil {
		toSerialize["fail_on_no_groups_found"] = o.FailOnNoGroupsFound
	}
	toSerialize["monitor_ids"] = o.MonitorIds
	if o.Warmup != nil {
		toSerialize["warmup"] = o.Warmup
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsMonitorIds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration            *int64                            `json:"duration,omitempty"`
		FailOnNoData        *bool                             `json:"fail_on_no_data,omitempty"`
		FailOnNoGroupsFound *bool                             `json:"fail_on_no_groups_found,omitempty"`
		MonitorIds          *[]DeploymentRuleOptionsMonitorId `json:"monitor_ids"`
		Warmup              *int64                            `json:"warmup,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorIds == nil {
		return fmt.Errorf("required field monitor_ids missing")
	}
	o.Duration = all.Duration
	o.FailOnNoData = all.FailOnNoData
	o.FailOnNoGroupsFound = all.FailOnNoGroupsFound
	o.MonitorIds = *all.MonitorIds
	o.Warmup = all.Warmup

	return nil
}
