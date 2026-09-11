// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest Settings of the Data Jobs Monitoring dataflow. Only the fields provided are changed.
type DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest struct {
	// ID of the Datadog API key the global init script uses to submit data. Setting or changing it requires `dd_api_key_secret` in the same request.
	DdApiKeyId *string `json:"dd_api_key_id,omitempty"`
	// Secret value of the Datadog API key identified by `dd_api_key_id`.
	DdApiKeySecret *string `json:"dd_api_key_secret,omitempty"`
	// Whether Datadog installs and manages the Agent on your Databricks clusters through a global init script. Installation can take up to 15 minutes and requires Databricks Workspace Admin permissions. The script does not apply to clusters in Standard access mode. Leave this `false` to install the Agent yourself. Defaults to `false`.
	DjmGlobalInitScriptEnabled *bool `json:"djm_global_init_script_enabled,omitempty"`
	// Whether GPU metrics are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`. Defaults to `false`.
	ScriptGpumEnabled *bool `json:"script_gpum_enabled,omitempty"`
	// Whether driver and worker logs are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`. Defaults to `false`.
	ScriptLogsEnabled *bool `json:"script_logs_enabled,omitempty"`
	// Whether health and cost data is collected for jobs running on Serverless or SQL Warehouse compute. This compute has no clusters for the global init script to target, so collection reads the Databricks system tables and requires `system_tables_sql_warehouse_id`. Defaults to `true`.
	ServerlessJobsEnabled *bool `json:"serverless_jobs_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest{}
	return &this
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequestWithDefaults instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequestWithDefaults() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest{}
	return &this
}

// GetDdApiKeyId returns the DdApiKeyId field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeyId() string {
	if o == nil || o.DdApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.DdApiKeyId
}

// GetDdApiKeyIdOk returns a tuple with the DdApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeyIdOk() (*string, bool) {
	if o == nil || o.DdApiKeyId == nil {
		return nil, false
	}
	return o.DdApiKeyId, true
}

// HasDdApiKeyId returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasDdApiKeyId() bool {
	return o != nil && o.DdApiKeyId != nil
}

// SetDdApiKeyId gets a reference to the given string and assigns it to the DdApiKeyId field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetDdApiKeyId(v string) {
	o.DdApiKeyId = &v
}

// GetDdApiKeySecret returns the DdApiKeySecret field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeySecret() string {
	if o == nil || o.DdApiKeySecret == nil {
		var ret string
		return ret
	}
	return *o.DdApiKeySecret
}

// GetDdApiKeySecretOk returns a tuple with the DdApiKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeySecretOk() (*string, bool) {
	if o == nil || o.DdApiKeySecret == nil {
		return nil, false
	}
	return o.DdApiKeySecret, true
}

// HasDdApiKeySecret returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasDdApiKeySecret() bool {
	return o != nil && o.DdApiKeySecret != nil
}

// SetDdApiKeySecret gets a reference to the given string and assigns it to the DdApiKeySecret field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetDdApiKeySecret(v string) {
	o.DdApiKeySecret = &v
}

// GetDjmGlobalInitScriptEnabled returns the DjmGlobalInitScriptEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDjmGlobalInitScriptEnabled() bool {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DjmGlobalInitScriptEnabled
}

// GetDjmGlobalInitScriptEnabledOk returns a tuple with the DjmGlobalInitScriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetDjmGlobalInitScriptEnabledOk() (*bool, bool) {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		return nil, false
	}
	return o.DjmGlobalInitScriptEnabled, true
}

// HasDjmGlobalInitScriptEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasDjmGlobalInitScriptEnabled() bool {
	return o != nil && o.DjmGlobalInitScriptEnabled != nil
}

// SetDjmGlobalInitScriptEnabled gets a reference to the given bool and assigns it to the DjmGlobalInitScriptEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetDjmGlobalInitScriptEnabled(v bool) {
	o.DjmGlobalInitScriptEnabled = &v
}

// GetScriptGpumEnabled returns the ScriptGpumEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetScriptGpumEnabled() bool {
	if o == nil || o.ScriptGpumEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptGpumEnabled
}

// GetScriptGpumEnabledOk returns a tuple with the ScriptGpumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetScriptGpumEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptGpumEnabled == nil {
		return nil, false
	}
	return o.ScriptGpumEnabled, true
}

// HasScriptGpumEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasScriptGpumEnabled() bool {
	return o != nil && o.ScriptGpumEnabled != nil
}

// SetScriptGpumEnabled gets a reference to the given bool and assigns it to the ScriptGpumEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetScriptGpumEnabled(v bool) {
	o.ScriptGpumEnabled = &v
}

// GetScriptLogsEnabled returns the ScriptLogsEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetScriptLogsEnabled() bool {
	if o == nil || o.ScriptLogsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptLogsEnabled
}

// GetScriptLogsEnabledOk returns a tuple with the ScriptLogsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetScriptLogsEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptLogsEnabled == nil {
		return nil, false
	}
	return o.ScriptLogsEnabled, true
}

// HasScriptLogsEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasScriptLogsEnabled() bool {
	return o != nil && o.ScriptLogsEnabled != nil
}

// SetScriptLogsEnabled gets a reference to the given bool and assigns it to the ScriptLogsEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetScriptLogsEnabled(v bool) {
	o.ScriptLogsEnabled = &v
}

// GetServerlessJobsEnabled returns the ServerlessJobsEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetServerlessJobsEnabled() bool {
	if o == nil || o.ServerlessJobsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServerlessJobsEnabled
}

// GetServerlessJobsEnabledOk returns a tuple with the ServerlessJobsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) GetServerlessJobsEnabledOk() (*bool, bool) {
	if o == nil || o.ServerlessJobsEnabled == nil {
		return nil, false
	}
	return o.ServerlessJobsEnabled, true
}

// HasServerlessJobsEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) HasServerlessJobsEnabled() bool {
	return o != nil && o.ServerlessJobsEnabled != nil
}

// SetServerlessJobsEnabled gets a reference to the given bool and assigns it to the ServerlessJobsEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) SetServerlessJobsEnabled(v bool) {
	o.ServerlessJobsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DdApiKeyId != nil {
		toSerialize["dd_api_key_id"] = o.DdApiKeyId
	}
	if o.DdApiKeySecret != nil {
		toSerialize["dd_api_key_secret"] = o.DdApiKeySecret
	}
	if o.DjmGlobalInitScriptEnabled != nil {
		toSerialize["djm_global_init_script_enabled"] = o.DjmGlobalInitScriptEnabled
	}
	if o.ScriptGpumEnabled != nil {
		toSerialize["script_gpum_enabled"] = o.ScriptGpumEnabled
	}
	if o.ScriptLogsEnabled != nil {
		toSerialize["script_logs_enabled"] = o.ScriptLogsEnabled
	}
	if o.ServerlessJobsEnabled != nil {
		toSerialize["serverless_jobs_enabled"] = o.ServerlessJobsEnabled
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdApiKeyId                 *string `json:"dd_api_key_id,omitempty"`
		DdApiKeySecret             *string `json:"dd_api_key_secret,omitempty"`
		DjmGlobalInitScriptEnabled *bool   `json:"djm_global_init_script_enabled,omitempty"`
		ScriptGpumEnabled          *bool   `json:"script_gpum_enabled,omitempty"`
		ScriptLogsEnabled          *bool   `json:"script_logs_enabled,omitempty"`
		ServerlessJobsEnabled      *bool   `json:"serverless_jobs_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.DdApiKeyId = all.DdApiKeyId
	o.DdApiKeySecret = all.DdApiKeySecret
	o.DjmGlobalInitScriptEnabled = all.DjmGlobalInitScriptEnabled
	o.ScriptGpumEnabled = all.ScriptGpumEnabled
	o.ScriptLogsEnabled = all.ScriptLogsEnabled
	o.ServerlessJobsEnabled = all.ServerlessJobsEnabled

	return nil
}
