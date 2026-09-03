// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataJobMonitoringIntegrationDataflowRequest The Databricks Data Jobs Monitoring dataflow.
type DatabricksDataJobMonitoringIntegrationDataflowRequest struct {
	// Whether the Databricks dataflow is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Settings of the Databricks Data Jobs Monitoring dataflow. Only the fields provided are changed.
	Settings *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksDataJobMonitoringIntegrationDataflowRequest instantiates a new DatabricksDataJobMonitoringIntegrationDataflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataJobMonitoringIntegrationDataflowRequest() *DatabricksDataJobMonitoringIntegrationDataflowRequest {
	this := DatabricksDataJobMonitoringIntegrationDataflowRequest{}
	return &this
}

// NewDatabricksDataJobMonitoringIntegrationDataflowRequestWithDefaults instantiates a new DatabricksDataJobMonitoringIntegrationDataflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataJobMonitoringIntegrationDataflowRequestWithDefaults() *DatabricksDataJobMonitoringIntegrationDataflowRequest {
	this := DatabricksDataJobMonitoringIntegrationDataflowRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) GetSettings() DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest {
	if o == nil || o.Settings == nil {
		var ret DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) GetSettingsOk() (*DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest and assigns it to the Settings field.
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) SetSettings(v DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataJobMonitoringIntegrationDataflowRequest) MarshalJSON() ([]byte, error) {
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
func (o *DatabricksDataJobMonitoringIntegrationDataflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled  *bool                                                          `json:"enabled,omitempty"`
		Settings *DatabricksDataJobMonitoringIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
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
