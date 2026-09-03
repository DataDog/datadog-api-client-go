// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest Settings of the Databricks Data Jobs Monitoring dataflow. Only the fields provided are changed.
type DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest struct {
	// ID of the Datadog API key the global init script uses to submit data. Setting or changing it requires `dd_api_key_secret` in the same request.
	DdApiKeyId *string `json:"dd_api_key_id,omitempty"`
	// Secret value of the Datadog API key identified by `dd_api_key_id`.
	DdApiKeySecret *string `json:"dd_api_key_secret,omitempty"`
	// Whether Datadog manages the global init script that installs the Agent on your Databricks clusters.
	DjmGlobalInitScriptEnabled *bool `json:"djm_global_init_script_enabled,omitempty"`
	// Whether GPU metrics are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`.
	ScriptGpumEnabled *bool `json:"script_gpum_enabled,omitempty"`
	// Whether logs are collected from your Databricks clusters. The Agent installed by the global init script performs the collection, so this requires the dataflow to be enabled with `djm_global_init_script_enabled` set to `true`.
	ScriptLogsEnabled *bool `json:"script_logs_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksDataJobMonitoringIntegrationDataflowSettingsRequest instantiates a new DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataJobMonitoringIntegrationDataflowSettingsRequest() *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest {
	this := DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest{}
	return &this
}

// NewDatabricksDataJobMonitoringIntegrationDataflowSettingsRequestWithDefaults instantiates a new DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataJobMonitoringIntegrationDataflowSettingsRequestWithDefaults() *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest {
	this := DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest{}
	return &this
}

// GetDdApiKeyId returns the DdApiKeyId field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeyId() string {
	if o == nil || o.DdApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.DdApiKeyId
}

// GetDdApiKeyIdOk returns a tuple with the DdApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeyIdOk() (*string, bool) {
	if o == nil || o.DdApiKeyId == nil {
		return nil, false
	}
	return o.DdApiKeyId, true
}

// HasDdApiKeyId returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) HasDdApiKeyId() bool {
	return o != nil && o.DdApiKeyId != nil
}

// SetDdApiKeyId gets a reference to the given string and assigns it to the DdApiKeyId field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) SetDdApiKeyId(v string) {
	o.DdApiKeyId = &v
}

// GetDdApiKeySecret returns the DdApiKeySecret field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeySecret() string {
	if o == nil || o.DdApiKeySecret == nil {
		var ret string
		return ret
	}
	return *o.DdApiKeySecret
}

// GetDdApiKeySecretOk returns a tuple with the DdApiKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDdApiKeySecretOk() (*string, bool) {
	if o == nil || o.DdApiKeySecret == nil {
		return nil, false
	}
	return o.DdApiKeySecret, true
}

// HasDdApiKeySecret returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) HasDdApiKeySecret() bool {
	return o != nil && o.DdApiKeySecret != nil
}

// SetDdApiKeySecret gets a reference to the given string and assigns it to the DdApiKeySecret field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) SetDdApiKeySecret(v string) {
	o.DdApiKeySecret = &v
}

// GetDjmGlobalInitScriptEnabled returns the DjmGlobalInitScriptEnabled field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDjmGlobalInitScriptEnabled() bool {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DjmGlobalInitScriptEnabled
}

// GetDjmGlobalInitScriptEnabledOk returns a tuple with the DjmGlobalInitScriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetDjmGlobalInitScriptEnabledOk() (*bool, bool) {
	if o == nil || o.DjmGlobalInitScriptEnabled == nil {
		return nil, false
	}
	return o.DjmGlobalInitScriptEnabled, true
}

// HasDjmGlobalInitScriptEnabled returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) HasDjmGlobalInitScriptEnabled() bool {
	return o != nil && o.DjmGlobalInitScriptEnabled != nil
}

// SetDjmGlobalInitScriptEnabled gets a reference to the given bool and assigns it to the DjmGlobalInitScriptEnabled field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) SetDjmGlobalInitScriptEnabled(v bool) {
	o.DjmGlobalInitScriptEnabled = &v
}

// GetScriptGpumEnabled returns the ScriptGpumEnabled field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetScriptGpumEnabled() bool {
	if o == nil || o.ScriptGpumEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptGpumEnabled
}

// GetScriptGpumEnabledOk returns a tuple with the ScriptGpumEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetScriptGpumEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptGpumEnabled == nil {
		return nil, false
	}
	return o.ScriptGpumEnabled, true
}

// HasScriptGpumEnabled returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) HasScriptGpumEnabled() bool {
	return o != nil && o.ScriptGpumEnabled != nil
}

// SetScriptGpumEnabled gets a reference to the given bool and assigns it to the ScriptGpumEnabled field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) SetScriptGpumEnabled(v bool) {
	o.ScriptGpumEnabled = &v
}

// GetScriptLogsEnabled returns the ScriptLogsEnabled field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetScriptLogsEnabled() bool {
	if o == nil || o.ScriptLogsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ScriptLogsEnabled
}

// GetScriptLogsEnabledOk returns a tuple with the ScriptLogsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) GetScriptLogsEnabledOk() (*bool, bool) {
	if o == nil || o.ScriptLogsEnabled == nil {
		return nil, false
	}
	return o.ScriptLogsEnabled, true
}

// HasScriptLogsEnabled returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) HasScriptLogsEnabled() bool {
	return o != nil && o.ScriptLogsEnabled != nil
}

// SetScriptLogsEnabled gets a reference to the given bool and assigns it to the ScriptLogsEnabled field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) SetScriptLogsEnabled(v bool) {
	o.ScriptLogsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdApiKeyId                 *string `json:"dd_api_key_id,omitempty"`
		DdApiKeySecret             *string `json:"dd_api_key_secret,omitempty"`
		DjmGlobalInitScriptEnabled *bool   `json:"djm_global_init_script_enabled,omitempty"`
		ScriptGpumEnabled          *bool   `json:"script_gpum_enabled,omitempty"`
		ScriptLogsEnabled          *bool   `json:"script_logs_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.DdApiKeyId = all.DdApiKeyId
	o.DdApiKeySecret = all.DdApiKeySecret
	o.DjmGlobalInitScriptEnabled = all.DjmGlobalInitScriptEnabled
	o.ScriptGpumEnabled = all.ScriptGpumEnabled
	o.ScriptLogsEnabled = all.ScriptLogsEnabled

	return nil
}
