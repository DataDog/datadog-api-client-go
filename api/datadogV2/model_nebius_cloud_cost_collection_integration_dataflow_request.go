// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudCostCollectionIntegrationDataflowRequest Cloud cost data read from a Nebius Object Storage bucket export and used by [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Enabling it requires a bucket holding your cost export and a Nebius access key pair that can read it; a request that enables this dataflow without them is rejected with a `422` response.
type NebiusCloudCostCollectionIntegrationDataflowRequest struct {
	// Whether Datadog collects this data. Defaults to `false`; set to `true` to start collection.
	Enabled *bool `json:"enabled,omitempty"`
	// Settings of the Cloud Cost Management dataflow. Only the fields provided are changed.
	Settings *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudCostCollectionIntegrationDataflowRequest instantiates a new NebiusCloudCostCollectionIntegrationDataflowRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudCostCollectionIntegrationDataflowRequest() *NebiusCloudCostCollectionIntegrationDataflowRequest {
	this := NebiusCloudCostCollectionIntegrationDataflowRequest{}
	return &this
}

// NewNebiusCloudCostCollectionIntegrationDataflowRequestWithDefaults instantiates a new NebiusCloudCostCollectionIntegrationDataflowRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudCostCollectionIntegrationDataflowRequestWithDefaults() *NebiusCloudCostCollectionIntegrationDataflowRequest {
	this := NebiusCloudCostCollectionIntegrationDataflowRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) GetSettings() NebiusCloudCostCollectionIntegrationDataflowSettingsRequest {
	if o == nil || o.Settings == nil {
		var ret NebiusCloudCostCollectionIntegrationDataflowSettingsRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) GetSettingsOk() (*NebiusCloudCostCollectionIntegrationDataflowSettingsRequest, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given NebiusCloudCostCollectionIntegrationDataflowSettingsRequest and assigns it to the Settings field.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) SetSettings(v NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudCostCollectionIntegrationDataflowRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudCostCollectionIntegrationDataflowRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled  *bool                                                        `json:"enabled,omitempty"`
		Settings *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
