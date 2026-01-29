// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDependenciesAttributesRequest Attributes for dataset dependencies request.
type SecurityMonitoringDatasetDependenciesAttributesRequest struct {
	// Array of dataset IDs to check dependencies for (minimum 1, maximum 100).
	DatasetIds []string `json:"datasetIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDependenciesAttributesRequest instantiates a new SecurityMonitoringDatasetDependenciesAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDependenciesAttributesRequest(datasetIds []string) *SecurityMonitoringDatasetDependenciesAttributesRequest {
	this := SecurityMonitoringDatasetDependenciesAttributesRequest{}
	this.DatasetIds = datasetIds
	return &this
}

// NewSecurityMonitoringDatasetDependenciesAttributesRequestWithDefaults instantiates a new SecurityMonitoringDatasetDependenciesAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDependenciesAttributesRequestWithDefaults() *SecurityMonitoringDatasetDependenciesAttributesRequest {
	this := SecurityMonitoringDatasetDependenciesAttributesRequest{}
	return &this
}

// GetDatasetIds returns the DatasetIds field value.
func (o *SecurityMonitoringDatasetDependenciesAttributesRequest) GetDatasetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DatasetIds
}

// GetDatasetIdsOk returns a tuple with the DatasetIds field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDependenciesAttributesRequest) GetDatasetIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetIds, true
}

// SetDatasetIds sets field value.
func (o *SecurityMonitoringDatasetDependenciesAttributesRequest) SetDatasetIds(v []string) {
	o.DatasetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDependenciesAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["datasetIds"] = o.DatasetIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetDependenciesAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetIds *[]string `json:"datasetIds"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetIds == nil {
		return fmt.Errorf("required field datasetIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datasetIds"})
	} else {
		return err
	}
	o.DatasetIds = *all.DatasetIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
