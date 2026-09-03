// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse Settings of the Databricks cloud cost metrics dataflow.
type DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse struct {
	// Whether cost data is collected for every workspace in the Databricks account rather than this workspace only.
	CcmCollectAllWorkspaces *bool `json:"ccm_collect_all_workspaces,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksCloudCostMetricsIntegrationDataflowSettingsResponse instantiates a new DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksCloudCostMetricsIntegrationDataflowSettingsResponse() *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse {
	this := DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse{}
	return &this
}

// NewDatabricksCloudCostMetricsIntegrationDataflowSettingsResponseWithDefaults instantiates a new DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksCloudCostMetricsIntegrationDataflowSettingsResponseWithDefaults() *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse {
	this := DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse{}
	return &this
}

// GetCcmCollectAllWorkspaces returns the CcmCollectAllWorkspaces field value if set, zero value otherwise.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) GetCcmCollectAllWorkspaces() bool {
	if o == nil || o.CcmCollectAllWorkspaces == nil {
		var ret bool
		return ret
	}
	return *o.CcmCollectAllWorkspaces
}

// GetCcmCollectAllWorkspacesOk returns a tuple with the CcmCollectAllWorkspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) GetCcmCollectAllWorkspacesOk() (*bool, bool) {
	if o == nil || o.CcmCollectAllWorkspaces == nil {
		return nil, false
	}
	return o.CcmCollectAllWorkspaces, true
}

// HasCcmCollectAllWorkspaces returns a boolean if a field has been set.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) HasCcmCollectAllWorkspaces() bool {
	return o != nil && o.CcmCollectAllWorkspaces != nil
}

// SetCcmCollectAllWorkspaces gets a reference to the given bool and assigns it to the CcmCollectAllWorkspaces field.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) SetCcmCollectAllWorkspaces(v bool) {
	o.CcmCollectAllWorkspaces = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CcmCollectAllWorkspaces != nil {
		toSerialize["ccm_collect_all_workspaces"] = o.CcmCollectAllWorkspaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksCloudCostMetricsIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CcmCollectAllWorkspaces *bool `json:"ccm_collect_all_workspaces,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ccm_collect_all_workspaces"})
	} else {
		return err
	}
	o.CcmCollectAllWorkspaces = all.CcmCollectAllWorkspaces

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
