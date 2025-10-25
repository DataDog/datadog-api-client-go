// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleOptionsFaultyDeploymentDetection Faulty deployment detection options for deployment rules.
type DeploymentRuleOptionsFaultyDeploymentDetection struct {
	// Resources to exclude from faulty deployment detection.
	ExcludedResources []string `json:"excluded_resources,omitempty"`
	// The operation name for faulty deployment detection.
	OperationName *string `json:"operation_name,omitempty"`
	// The wait time for faulty deployment detection.
	WaitTime *int64 `json:"wait_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDeploymentRuleOptionsFaultyDeploymentDetection instantiates a new DeploymentRuleOptionsFaultyDeploymentDetection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleOptionsFaultyDeploymentDetection() *DeploymentRuleOptionsFaultyDeploymentDetection {
	this := DeploymentRuleOptionsFaultyDeploymentDetection{}
	return &this
}

// NewDeploymentRuleOptionsFaultyDeploymentDetectionWithDefaults instantiates a new DeploymentRuleOptionsFaultyDeploymentDetection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleOptionsFaultyDeploymentDetectionWithDefaults() *DeploymentRuleOptionsFaultyDeploymentDetection {
	this := DeploymentRuleOptionsFaultyDeploymentDetection{}
	return &this
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetExcludedResources() []string {
	if o == nil || o.ExcludedResources == nil {
		var ret []string
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetExcludedResourcesOk() (*[]string, bool) {
	if o == nil || o.ExcludedResources == nil {
		return nil, false
	}
	return &o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) HasExcludedResources() bool {
	return o != nil && o.ExcludedResources != nil
}

// SetExcludedResources gets a reference to the given []string and assigns it to the ExcludedResources field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) SetExcludedResources(v []string) {
	o.ExcludedResources = v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) SetOperationName(v string) {
	o.OperationName = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetWaitTime() int64 {
	if o == nil || o.WaitTime == nil {
		var ret int64
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) GetWaitTimeOk() (*int64, bool) {
	if o == nil || o.WaitTime == nil {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) HasWaitTime() bool {
	return o != nil && o.WaitTime != nil
}

// SetWaitTime gets a reference to the given int64 and assigns it to the WaitTime field.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) SetWaitTime(v int64) {
	o.WaitTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleOptionsFaultyDeploymentDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludedResources != nil {
		toSerialize["excluded_resources"] = o.ExcludedResources
	}
	if o.OperationName != nil {
		toSerialize["operation_name"] = o.OperationName
	}
	if o.WaitTime != nil {
		toSerialize["wait_time"] = o.WaitTime
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleOptionsFaultyDeploymentDetection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludedResources []string `json:"excluded_resources,omitempty"`
		OperationName     *string  `json:"operation_name,omitempty"`
		WaitTime          *int64   `json:"wait_time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ExcludedResources = all.ExcludedResources
	o.OperationName = all.OperationName
	o.WaitTime = all.WaitTime

	return nil
}
