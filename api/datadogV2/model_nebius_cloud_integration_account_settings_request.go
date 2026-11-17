// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountSettingsRequest Settings for creating the Nebius Cloud integration account.
type NebiusCloudIntegrationAccountSettingsRequest struct {
	// Tags applied to all metrics, logs, and resources collected from this Nebius Cloud account. Defaults to no tags.
	AccountTags []string `json:"account_tags,omitempty"`
	// Configuration for metric collection from this Nebius Cloud account.
	MetricsConfig *NebiusCloudIntegrationAccountMetricsConfigRequest `json:"metrics_config,omitempty"`
	// ID of the Nebius tenant to monitor, the organizational workspace identifier. Found in the Nebius console under Administration > Organization.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountSettingsRequest instantiates a new NebiusCloudIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountSettingsRequest(tenantId string) *NebiusCloudIntegrationAccountSettingsRequest {
	this := NebiusCloudIntegrationAccountSettingsRequest{}
	this.TenantId = tenantId
	return &this
}

// NewNebiusCloudIntegrationAccountSettingsRequestWithDefaults instantiates a new NebiusCloudIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountSettingsRequestWithDefaults() *NebiusCloudIntegrationAccountSettingsRequest {
	this := NebiusCloudIntegrationAccountSettingsRequest{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountSettingsRequest) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *NebiusCloudIntegrationAccountSettingsRequest) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetMetricsConfig() NebiusCloudIntegrationAccountMetricsConfigRequest {
	if o == nil || o.MetricsConfig == nil {
		var ret NebiusCloudIntegrationAccountMetricsConfigRequest
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetMetricsConfigOk() (*NebiusCloudIntegrationAccountMetricsConfigRequest, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountSettingsRequest) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given NebiusCloudIntegrationAccountMetricsConfigRequest and assigns it to the MetricsConfig field.
func (o *NebiusCloudIntegrationAccountSettingsRequest) SetMetricsConfig(v NebiusCloudIntegrationAccountMetricsConfigRequest) {
	o.MetricsConfig = &v
}

// GetTenantId returns the TenantId field value.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *NebiusCloudIntegrationAccountSettingsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags != nil {
		toSerialize["account_tags"] = o.AccountTags
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	toSerialize["tenant_id"] = o.TenantId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags   []string                                           `json:"account_tags,omitempty"`
		MetricsConfig *NebiusCloudIntegrationAccountMetricsConfigRequest `json:"metrics_config,omitempty"`
		TenantId      *string                                            `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}

	hasInvalidField := false
	o.AccountTags = all.AccountTags
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	o.TenantId = *all.TenantId

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
