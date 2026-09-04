// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest Settings of the Databricks cloud cost metrics dataflow. Only the fields provided are changed.
type DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest struct {
	// Whether cost data is collected for every workspace in the Databricks account rather than this workspace only.
	CcmCollectAllWorkspaces *bool `json:"ccm_collect_all_workspaces,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksCloudCostMetricsIntegrationDataflowSettingsRequest instantiates a new DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksCloudCostMetricsIntegrationDataflowSettingsRequest() *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest {
	this := DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewDatabricksCloudCostMetricsIntegrationDataflowSettingsRequestWithDefaults instantiates a new DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksCloudCostMetricsIntegrationDataflowSettingsRequestWithDefaults() *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest {
	this := DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetCcmCollectAllWorkspaces returns the CcmCollectAllWorkspaces field value if set, zero value otherwise.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) GetCcmCollectAllWorkspaces() bool {
	if o == nil || o.CcmCollectAllWorkspaces == nil {
		var ret bool
		return ret
	}
	return *o.CcmCollectAllWorkspaces
}

// GetCcmCollectAllWorkspacesOk returns a tuple with the CcmCollectAllWorkspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) GetCcmCollectAllWorkspacesOk() (*bool, bool) {
	if o == nil || o.CcmCollectAllWorkspaces == nil {
		return nil, false
	}
	return o.CcmCollectAllWorkspaces, true
}

// HasCcmCollectAllWorkspaces returns a boolean if a field has been set.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) HasCcmCollectAllWorkspaces() bool {
	return o != nil && o.CcmCollectAllWorkspaces != nil
}

// SetCcmCollectAllWorkspaces gets a reference to the given bool and assigns it to the CcmCollectAllWorkspaces field.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) SetCcmCollectAllWorkspaces(v bool) {
	o.CcmCollectAllWorkspaces = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CcmCollectAllWorkspaces != nil {
		toSerialize["ccm_collect_all_workspaces"] = o.CcmCollectAllWorkspaces
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CcmCollectAllWorkspaces *bool `json:"ccm_collect_all_workspaces,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.CcmCollectAllWorkspaces = all.CcmCollectAllWorkspaces

	return nil
}
