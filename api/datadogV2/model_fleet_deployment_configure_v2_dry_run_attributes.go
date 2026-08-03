// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentConfigureV2DryRunAttributes Attributes of a configuration deployment dry-run response.
type FleetDeploymentConfigureV2DryRunAttributes struct {
	// Validation result of a configuration deployment dry run.
	DryRun *FleetDeploymentConfigureV2DryRunResult `json:"dry_run,omitempty"`
	// Query used to filter and select target hosts for the deployment.
	Query *string `json:"query,omitempty"`
	// Total number of hosts targeted by the dry run.
	TotalHosts *int64 `json:"total_hosts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentConfigureV2DryRunAttributes instantiates a new FleetDeploymentConfigureV2DryRunAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentConfigureV2DryRunAttributes() *FleetDeploymentConfigureV2DryRunAttributes {
	this := FleetDeploymentConfigureV2DryRunAttributes{}
	return &this
}

// NewFleetDeploymentConfigureV2DryRunAttributesWithDefaults instantiates a new FleetDeploymentConfigureV2DryRunAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentConfigureV2DryRunAttributesWithDefaults() *FleetDeploymentConfigureV2DryRunAttributes {
	this := FleetDeploymentConfigureV2DryRunAttributes{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetDryRun() FleetDeploymentConfigureV2DryRunResult {
	if o == nil || o.DryRun == nil {
		var ret FleetDeploymentConfigureV2DryRunResult
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetDryRunOk() (*FleetDeploymentConfigureV2DryRunResult, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given FleetDeploymentConfigureV2DryRunResult and assigns it to the DryRun field.
func (o *FleetDeploymentConfigureV2DryRunAttributes) SetDryRun(v FleetDeploymentConfigureV2DryRunResult) {
	o.DryRun = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetDeploymentConfigureV2DryRunAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetTotalHosts() int64 {
	if o == nil || o.TotalHosts == nil {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) GetTotalHostsOk() (*int64, bool) {
	if o == nil || o.TotalHosts == nil {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2DryRunAttributes) HasTotalHosts() bool {
	return o != nil && o.TotalHosts != nil
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *FleetDeploymentConfigureV2DryRunAttributes) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentConfigureV2DryRunAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.TotalHosts != nil {
		toSerialize["total_hosts"] = o.TotalHosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentConfigureV2DryRunAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun     *FleetDeploymentConfigureV2DryRunResult `json:"dry_run,omitempty"`
		Query      *string                                 `json:"query,omitempty"`
		TotalHosts *int64                                  `json:"total_hosts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "query", "total_hosts"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DryRun != nil && all.DryRun.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DryRun = all.DryRun
	o.Query = all.Query
	o.TotalHosts = all.TotalHosts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
