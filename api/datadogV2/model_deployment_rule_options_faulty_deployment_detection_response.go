// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsFaultyDeploymentDetectionResponse Faulty deployment detection options returned in deployment rule responses. The `excluded_resources` field is always present, which allows disambiguating this type from monitor options when both share a `duration` field.
type DeploymentRuleOptionsFaultyDeploymentDetectionResponse struct {
	// The duration for faulty deployment detection.
	Duration *int64 `json:"duration,omitempty"`
	// Resources to exclude from faulty deployment detection.
	ExcludedResources []string `json:"excluded_resources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsFaultyDeploymentDetectionResponse instantiates a new DeploymentRuleOptionsFaultyDeploymentDetectionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsFaultyDeploymentDetectionResponse(excludedResources []string) *DeploymentRuleOptionsFaultyDeploymentDetectionResponse {
	this := DeploymentRuleOptionsFaultyDeploymentDetectionResponse{}
	this.ExcludedResources = excludedResources
	return &this
}

// NewDeploymentRuleOptionsFaultyDeploymentDetectionResponseWithDefaults instantiates a new DeploymentRuleOptionsFaultyDeploymentDetectionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsFaultyDeploymentDetectionResponseWithDefaults() *DeploymentRuleOptionsFaultyDeploymentDetectionResponse {
	this := DeploymentRuleOptionsFaultyDeploymentDetectionResponse{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) SetDuration(v int64) {
	o.Duration = &v
}

// GetExcludedResources returns the ExcludedResources field value.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) GetExcludedResources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) GetExcludedResourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludedResources, true
}

// SetExcludedResources sets field value.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) SetExcludedResources(v []string) {
	o.ExcludedResources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsFaultyDeploymentDetectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["excluded_resources"] = o.ExcludedResources
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsFaultyDeploymentDetectionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration          *int64    `json:"duration,omitempty"`
		ExcludedResources *[]string `json:"excluded_resources"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExcludedResources == nil {
		return fmt.Errorf("required field excluded_resources missing")
	}
	o.Duration = all.Duration
	o.ExcludedResources = *all.ExcludedResources

	return nil
}
