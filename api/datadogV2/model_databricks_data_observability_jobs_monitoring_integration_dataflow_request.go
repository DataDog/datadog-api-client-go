// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest Data Jobs Monitoring, which collects performance, reliability, and cost data for your Databricks jobs.
type DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest struct {
	// Whether Datadog collects this data. Defaults to `true`; set to `false` to stop collection.
	Enabled *bool `json:"enabled,omitempty"`
	// Settings of the Data Jobs Monitoring dataflow. Only the fields provided are changed.
	Settings *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest{}
	return &this
}

// NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequestWithDefaults instantiates a new DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequestWithDefaults() *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest {
	this := DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) GetSettings() DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest {
	if o == nil || o.Settings == nil {
		var ret DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) GetSettingsOk() (*DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest and assigns it to the Settings field.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) SetSettings(v DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled  *bool                                                                        `json:"enabled,omitempty"`
		Settings *DatabricksDataObservabilityJobsMonitoringIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
