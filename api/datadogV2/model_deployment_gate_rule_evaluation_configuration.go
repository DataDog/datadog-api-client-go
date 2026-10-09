// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateRuleEvaluationConfiguration Evaluated rule configuration. Fields depend on rule type and unset fields are omitted.
// Monitor rules can include `duration`, `query`, `monitor_ids`, `warmup`, `fail_on_no_groups_found`, and `fail_on_no_data`.
// Faulty deployment detection rules can include `duration`, `allowed_resources`, and `excluded_resources`.
type DeploymentGateRuleEvaluationConfiguration struct {
	// APM resources explicitly allowed by faulty deployment detection.
	AllowedResources []string `json:"allowed_resources,omitempty"`
	// Evaluation duration configured for this rule.
	Duration *int64 `json:"duration,omitempty"`
	// APM resources excluded from faulty deployment detection.
	ExcludedResources []string `json:"excluded_resources,omitempty"`
	// Whether a monitor rule fails when no data is found.
	FailOnNoData *bool `json:"fail_on_no_data,omitempty"`
	// Whether a monitor rule fails when no groups are found.
	FailOnNoGroupsFound *bool `json:"fail_on_no_groups_found,omitempty"`
	// Monitor IDs evaluated by a monitor rule.
	MonitorIds []string `json:"monitor_ids,omitempty"`
	// Monitor query used by a monitor rule.
	Query *string `json:"query,omitempty"`
	// Warm-up duration in seconds for a monitor rule. Omitted when zero.
	Warmup *int64 `json:"warmup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentGateRuleEvaluationConfiguration instantiates a new DeploymentGateRuleEvaluationConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateRuleEvaluationConfiguration() *DeploymentGateRuleEvaluationConfiguration {
	this := DeploymentGateRuleEvaluationConfiguration{}
	return &this
}

// NewDeploymentGateRuleEvaluationConfigurationWithDefaults instantiates a new DeploymentGateRuleEvaluationConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateRuleEvaluationConfigurationWithDefaults() *DeploymentGateRuleEvaluationConfiguration {
	this := DeploymentGateRuleEvaluationConfiguration{}
	return &this
}

// GetAllowedResources returns the AllowedResources field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetAllowedResources() []string {
	if o == nil || o.AllowedResources == nil {
		var ret []string
		return ret
	}
	return o.AllowedResources
}

// GetAllowedResourcesOk returns a tuple with the AllowedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetAllowedResourcesOk() (*[]string, bool) {
	if o == nil || o.AllowedResources == nil {
		return nil, false
	}
	return &o.AllowedResources, true
}

// HasAllowedResources returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasAllowedResources() bool {
	return o != nil && o.AllowedResources != nil
}

// SetAllowedResources gets a reference to the given []string and assigns it to the AllowedResources field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetAllowedResources(v []string) {
	o.AllowedResources = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetDuration(v int64) {
	o.Duration = &v
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetExcludedResources() []string {
	if o == nil || o.ExcludedResources == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetExcludedResourcesOk() (*[]string, bool) {
	if o == nil || o.ExcludedResources == nil {
		return nil, false
	}
	return &o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasExcludedResources() bool {
	return o != nil && o.ExcludedResources != nil
}

// SetExcludedResources gets a reference to the given []string and assigns it to the ExcludedResources field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetExcludedResources(v []string) {
	o.ExcludedResources = v
}

// GetFailOnNoData returns the FailOnNoData field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetFailOnNoData() bool {
	if o == nil || o.FailOnNoData == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoData
}

// GetFailOnNoDataOk returns a tuple with the FailOnNoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetFailOnNoDataOk() (*bool, bool) {
	if o == nil || o.FailOnNoData == nil {
		return nil, false
	}
	return o.FailOnNoData, true
}

// HasFailOnNoData returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasFailOnNoData() bool {
	return o != nil && o.FailOnNoData != nil
}

// SetFailOnNoData gets a reference to the given bool and assigns it to the FailOnNoData field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetFailOnNoData(v bool) {
	o.FailOnNoData = &v
}

// GetFailOnNoGroupsFound returns the FailOnNoGroupsFound field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetFailOnNoGroupsFound() bool {
	if o == nil || o.FailOnNoGroupsFound == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoGroupsFound
}

// GetFailOnNoGroupsFoundOk returns a tuple with the FailOnNoGroupsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetFailOnNoGroupsFoundOk() (*bool, bool) {
	if o == nil || o.FailOnNoGroupsFound == nil {
		return nil, false
	}
	return o.FailOnNoGroupsFound, true
}

// HasFailOnNoGroupsFound returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasFailOnNoGroupsFound() bool {
	return o != nil && o.FailOnNoGroupsFound != nil
}

// SetFailOnNoGroupsFound gets a reference to the given bool and assigns it to the FailOnNoGroupsFound field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetFailOnNoGroupsFound(v bool) {
	o.FailOnNoGroupsFound = &v
}

// GetMonitorIds returns the MonitorIds field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetMonitorIds() []string {
	if o == nil || o.MonitorIds == nil {
		var ret []string
		return ret
	}
	return o.MonitorIds
}

// GetMonitorIdsOk returns a tuple with the MonitorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetMonitorIdsOk() (*[]string, bool) {
	if o == nil || o.MonitorIds == nil {
		return nil, false
	}
	return &o.MonitorIds, true
}

// HasMonitorIds returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasMonitorIds() bool {
	return o != nil && o.MonitorIds != nil
}

// SetMonitorIds gets a reference to the given []string and assigns it to the MonitorIds field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetMonitorIds(v []string) {
	o.MonitorIds = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetQuery(v string) {
	o.Query = &v
}

// GetWarmup returns the Warmup field value if set, zero value otherwise.
func (o *DeploymentGateRuleEvaluationConfiguration) GetWarmup() int64 {
	if o == nil || o.Warmup == nil {
		var ret int64
		return ret
	}
	return *o.Warmup
}

// GetWarmupOk returns a tuple with the Warmup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) GetWarmupOk() (*int64, bool) {
	if o == nil || o.Warmup == nil {
		return nil, false
	}
	return o.Warmup, true
}

// HasWarmup returns a boolean if a field has been set.
func (o *DeploymentGateRuleEvaluationConfiguration) HasWarmup() bool {
	return o != nil && o.Warmup != nil
}

// SetWarmup gets a reference to the given int64 and assigns it to the Warmup field.
func (o *DeploymentGateRuleEvaluationConfiguration) SetWarmup(v int64) {
	o.Warmup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateRuleEvaluationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedResources != nil {
		toSerialize["allowed_resources"] = o.AllowedResources
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ExcludedResources != nil {
		toSerialize["excluded_resources"] = o.ExcludedResources
	}
	if o.FailOnNoData != nil {
		toSerialize["fail_on_no_data"] = o.FailOnNoData
	}
	if o.FailOnNoGroupsFound != nil {
		toSerialize["fail_on_no_groups_found"] = o.FailOnNoGroupsFound
	}
	if o.MonitorIds != nil {
		toSerialize["monitor_ids"] = o.MonitorIds
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Warmup != nil {
		toSerialize["warmup"] = o.Warmup
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateRuleEvaluationConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedResources    []string `json:"allowed_resources,omitempty"`
		Duration            *int64   `json:"duration,omitempty"`
		ExcludedResources   []string `json:"excluded_resources,omitempty"`
		FailOnNoData        *bool    `json:"fail_on_no_data,omitempty"`
		FailOnNoGroupsFound *bool    `json:"fail_on_no_groups_found,omitempty"`
		MonitorIds          []string `json:"monitor_ids,omitempty"`
		Query               *string  `json:"query,omitempty"`
		Warmup              *int64   `json:"warmup,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.AllowedResources = all.AllowedResources
	o.Duration = all.Duration
	o.ExcludedResources = all.ExcludedResources
	o.FailOnNoData = all.FailOnNoData
	o.FailOnNoGroupsFound = all.FailOnNoGroupsFound
	o.MonitorIds = all.MonitorIds
	o.Query = all.Query
	o.Warmup = all.Warmup

	return nil
}
