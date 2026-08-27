// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsMonitor Monitor options for deployment rules.
type DeploymentRuleOptionsMonitor struct {
	// Seconds the monitor needs to stay in OK status for the rule to pass.
	Duration *int64 `json:"duration,omitempty"`
	// Whether the rule should fail if a matching monitor group is in a NO DATA state.
	FailOnNoData *bool `json:"fail_on_no_data,omitempty"`
	// Whether the rule should fail if no monitor groups are found for the query.
	FailOnNoGroupsFound *bool `json:"fail_on_no_groups_found,omitempty"`
	// Monitors that match this query are evaluated.
	Query string `json:"query"`
	// Seconds to wait after a deployment starts before evaluating the monitor's status.
	Warmup *int64 `json:"warmup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsMonitor instantiates a new DeploymentRuleOptionsMonitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsMonitor(query string) *DeploymentRuleOptionsMonitor {
	this := DeploymentRuleOptionsMonitor{}
	var failOnNoData bool = true
	this.FailOnNoData = &failOnNoData
	var failOnNoGroupsFound bool = false
	this.FailOnNoGroupsFound = &failOnNoGroupsFound
	this.Query = query
	var warmup int64 = 0
	this.Warmup = &warmup
	return &this
}

// NewDeploymentRuleOptionsMonitorWithDefaults instantiates a new DeploymentRuleOptionsMonitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsMonitorWithDefaults() *DeploymentRuleOptionsMonitor {
	this := DeploymentRuleOptionsMonitor{}
	var failOnNoData bool = true
	this.FailOnNoData = &failOnNoData
	var failOnNoGroupsFound bool = false
	this.FailOnNoGroupsFound = &failOnNoGroupsFound
	var warmup int64 = 0
	this.Warmup = &warmup
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitor) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitor) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRuleOptionsMonitor) SetDuration(v int64) {
	o.Duration = &v
}

// GetFailOnNoData returns the FailOnNoData field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitor) GetFailOnNoData() bool {
	if o == nil || o.FailOnNoData == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoData
}

// GetFailOnNoDataOk returns a tuple with the FailOnNoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetFailOnNoDataOk() (*bool, bool) {
	if o == nil || o.FailOnNoData == nil {
		return nil, false
	}
	return o.FailOnNoData, true
}

// HasFailOnNoData returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitor) HasFailOnNoData() bool {
	return o != nil && o.FailOnNoData != nil
}

// SetFailOnNoData gets a reference to the given bool and assigns it to the FailOnNoData field.
func (o *DeploymentRuleOptionsMonitor) SetFailOnNoData(v bool) {
	o.FailOnNoData = &v
}

// GetFailOnNoGroupsFound returns the FailOnNoGroupsFound field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitor) GetFailOnNoGroupsFound() bool {
	if o == nil || o.FailOnNoGroupsFound == nil {
		var ret bool
		return ret
	}
	return *o.FailOnNoGroupsFound
}

// GetFailOnNoGroupsFoundOk returns a tuple with the FailOnNoGroupsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetFailOnNoGroupsFoundOk() (*bool, bool) {
	if o == nil || o.FailOnNoGroupsFound == nil {
		return nil, false
	}
	return o.FailOnNoGroupsFound, true
}

// HasFailOnNoGroupsFound returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitor) HasFailOnNoGroupsFound() bool {
	return o != nil && o.FailOnNoGroupsFound != nil
}

// SetFailOnNoGroupsFound gets a reference to the given bool and assigns it to the FailOnNoGroupsFound field.
func (o *DeploymentRuleOptionsMonitor) SetFailOnNoGroupsFound(v bool) {
	o.FailOnNoGroupsFound = &v
}

// GetQuery returns the Query field value.
func (o *DeploymentRuleOptionsMonitor) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DeploymentRuleOptionsMonitor) SetQuery(v string) {
	o.Query = v
}

// GetWarmup returns the Warmup field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsMonitor) GetWarmup() int64 {
	if o == nil || o.Warmup == nil {
		var ret int64
		return ret
	}
	return *o.Warmup
}

// GetWarmupOk returns a tuple with the Warmup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsMonitor) GetWarmupOk() (*int64, bool) {
	if o == nil || o.Warmup == nil {
		return nil, false
	}
	return o.Warmup, true
}

// HasWarmup returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsMonitor) HasWarmup() bool {
	return o != nil && o.Warmup != nil
}

// SetWarmup gets a reference to the given int64 and assigns it to the Warmup field.
func (o *DeploymentRuleOptionsMonitor) SetWarmup(v int64) {
	o.Warmup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsMonitor) MarshalJSON() ([]byte, error) {
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
	toSerialize["query"] = o.Query
	if o.Warmup != nil {
		toSerialize["warmup"] = o.Warmup
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsMonitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration            *int64  `json:"duration,omitempty"`
		FailOnNoData        *bool   `json:"fail_on_no_data,omitempty"`
		FailOnNoGroupsFound *bool   `json:"fail_on_no_groups_found,omitempty"`
		Query               *string `json:"query"`
		Warmup              *int64  `json:"warmup,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.Duration = all.Duration
	o.FailOnNoData = all.FailOnNoData
	o.FailOnNoGroupsFound = all.FailOnNoGroupsFound
	o.Query = *all.Query
	o.Warmup = all.Warmup

	return nil
}
