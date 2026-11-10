// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudLogsIntegrationDataflowSettingsRequest Settings of the logs dataflow. Only the fields provided are changed.
type NebiusCloudLogsIntegrationDataflowSettingsRequest struct {
	// Nebius log integrations to collect. One of `nebius-cloud-compute`, `nebius-cloud-kubernetes`, `nebius-cloud-mlflow`, `nebius-cloud-postgresql`, `nebius-cloud-applications`, or `default` for the streams shared by every service. A non-empty list is the sole source of truth and replaces the stored list: those services are collected and any others stop. An empty list stops log collection entirely. Omit it to keep the stored list.
	EnabledServices []string `json:"enabled_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudLogsIntegrationDataflowSettingsRequest instantiates a new NebiusCloudLogsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudLogsIntegrationDataflowSettingsRequest() *NebiusCloudLogsIntegrationDataflowSettingsRequest {
	this := NebiusCloudLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewNebiusCloudLogsIntegrationDataflowSettingsRequestWithDefaults instantiates a new NebiusCloudLogsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudLogsIntegrationDataflowSettingsRequestWithDefaults() *NebiusCloudLogsIntegrationDataflowSettingsRequest {
	this := NebiusCloudLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetEnabledServices returns the EnabledServices field value if set, zero value otherwise.
func (o *NebiusCloudLogsIntegrationDataflowSettingsRequest) GetEnabledServices() []string {
	if o == nil || o.EnabledServices == nil {
		var ret []string
		return ret
	}
	return o.EnabledServices
}

// GetEnabledServicesOk returns a tuple with the EnabledServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudLogsIntegrationDataflowSettingsRequest) GetEnabledServicesOk() (*[]string, bool) {
	if o == nil || o.EnabledServices == nil {
		return nil, false
	}
	return &o.EnabledServices, true
}

// HasEnabledServices returns a boolean if a field has been set.
func (o *NebiusCloudLogsIntegrationDataflowSettingsRequest) HasEnabledServices() bool {
	return o != nil && o.EnabledServices != nil
}

// SetEnabledServices gets a reference to the given []string and assigns it to the EnabledServices field.
func (o *NebiusCloudLogsIntegrationDataflowSettingsRequest) SetEnabledServices(v []string) {
	o.EnabledServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudLogsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnabledServices != nil {
		toSerialize["enabled_services"] = o.EnabledServices
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudLogsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnabledServices []string `json:"enabled_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.EnabledServices = all.EnabledServices

	return nil
}
