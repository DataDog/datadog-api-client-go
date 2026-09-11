// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse Settings of the Data Jobs Monitoring dataflow.
type DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse struct {
	// ID of the Datadog API key the global init script uses to submit data.
	DdApiKeyId *string `json:"dd_api_key_id,omitempty"`
	// Whether Datadog installs and manages the Agent on your Databricks clusters through a global init script. The script does not apply to clusters in Standard access mode. When `false`, the Agent is installed manually.
	DjmGlobalInitScriptEnabled *bool `json:"djm_global_init_script_enabled,omitempty"`
	// Whether GPU metrics are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`.
	ScriptGpumEnabled *bool `json:"script_gpum_enabled,omitempty"`
	// Whether driver and worker logs are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`.
	ScriptLogsEnabled *bool `json:"script_logs_enabled,omitempty"`
	// Whether health and cost data is collected for jobs running on Serverless or SQL Warehouse compute. This compute has no clusters for the global init script to target, so collection reads the Databricks system tables and requires `system_tables_sql_warehouse_id`.
	ServerlessJobsEnabled *bool `json:"serverless_jobs_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponseWithDefaults instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponseWithDefaults() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// GetDdApiKeyId returns the DdApiKeyId field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetDdApiKeyId() string {
	if o == nil || o.DdApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.DdApiKeyId
}

// GetDdApiKeyIdOk returns a tuple with the DdApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetDdApiKeyIdOk() (*string, bool) {
	if o == nil || o.DdApiKeyId == nil {
		return nil, false
	}
	return o.DdApiKeyId, true
}

// HasDdApiKeyId returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) HasDdApiKeyId() bool {
	return o != nil && o.DdApiKeyId != nil
}

// SetDdApiKeyId gets a reference to the given string and assigns it to the DdApiKeyId field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) SetDdApiKeyId(v string) {
	o.DdApiKeyId = &v
}

// GetDjmGlobalInitScriptEnabled returns the DjmGlobalInitScriptEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetDjmGlobalInitScriptEnabled() bool {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DjmGlobalInitScriptEnabled
}

// GetDjmGlobalInitScriptEnabledOk returns a tuple with the DjmGlobalInitScriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetDjmGlobalInitScriptEnabledOk() (*bool, bool) {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		return nil, false
	}
	return o.DjmGlobalInitScriptEnabled, true
}

// HasDjmGlobalInitScriptEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) HasDjmGlobalInitScriptEnabled() bool {
	return o != nil && o.DjmGlobalInitScriptEnabled != nil
}

// SetDjmGlobalInitScriptEnabled gets a reference to the given bool and assigns it to the DjmGlobalInitScriptEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) SetDjmGlobalInitScriptEnabled(v bool) {
	o.DjmGlobalInitScriptEnabled = &v
}

// GetScriptGpumEnabled returns the ScriptGpumEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetScriptGpumEnabled() bool {
	if o == nil || o.ScriptGpumEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptGpumEnabled
}

// GetScriptGpumEnabledOk returns a tuple with the ScriptGpumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetScriptGpumEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptGpumEnabled == nil {
		return nil, false
	}
	return o.ScriptGpumEnabled, true
}

// HasScriptGpumEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) HasScriptGpumEnabled() bool {
	return o != nil && o.ScriptGpumEnabled != nil
}

// SetScriptGpumEnabled gets a reference to the given bool and assigns it to the ScriptGpumEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) SetScriptGpumEnabled(v bool) {
	o.ScriptGpumEnabled = &v
}

// GetScriptLogsEnabled returns the ScriptLogsEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetScriptLogsEnabled() bool {
	if o == nil || o.ScriptLogsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptLogsEnabled
}

// GetScriptLogsEnabledOk returns a tuple with the ScriptLogsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetScriptLogsEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptLogsEnabled == nil {
		return nil, false
	}
	return o.ScriptLogsEnabled, true
}

// HasScriptLogsEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) HasScriptLogsEnabled() bool {
	return o != nil && o.ScriptLogsEnabled != nil
}

// SetScriptLogsEnabled gets a reference to the given bool and assigns it to the ScriptLogsEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) SetScriptLogsEnabled(v bool) {
	o.ScriptLogsEnabled = &v
}

// GetServerlessJobsEnabled returns the ServerlessJobsEnabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetServerlessJobsEnabled() bool {
	if o == nil || o.ServerlessJobsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServerlessJobsEnabled
}

// GetServerlessJobsEnabledOk returns a tuple with the ServerlessJobsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) GetServerlessJobsEnabledOk() (*bool, bool) {
	if o == nil || o.ServerlessJobsEnabled == nil {
		return nil, false
	}
	return o.ServerlessJobsEnabled, true
}

// HasServerlessJobsEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) HasServerlessJobsEnabled() bool {
	return o != nil && o.ServerlessJobsEnabled != nil
}

// SetServerlessJobsEnabled gets a reference to the given bool and assigns it to the ServerlessJobsEnabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) SetServerlessJobsEnabled(v bool) {
	o.ServerlessJobsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DdApiKeyId != nil {
		toSerialize["dd_api_key_id"] = o.DdApiKeyId
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdApiKeyId                 *string `json:"dd_api_key_id,omitempty"`
		DjmGlobalInitScriptEnabled *bool   `json:"djm_global_init_script_enabled,omitempty"`
		ScriptGpumEnabled          *bool   `json:"script_gpum_enabled,omitempty"`
		ScriptLogsEnabled          *bool   `json:"script_logs_enabled,omitempty"`
		ServerlessJobsEnabled      *bool   `json:"serverless_jobs_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dd_api_key_id", "djm_global_init_script_enabled", "script_gpum_enabled", "script_logs_enabled", "serverless_jobs_enabled"})
	} else {
		return err
	}
	o.DdApiKeyId = all.DdApiKeyId
	o.DjmGlobalInitScriptEnabled = all.DjmGlobalInitScriptEnabled
	o.ScriptGpumEnabled = all.ScriptGpumEnabled
	o.ScriptLogsEnabled = all.ScriptLogsEnabled
	o.ServerlessJobsEnabled = all.ServerlessJobsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
