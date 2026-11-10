// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountSettingsResponse Settings configured on the Nebius Cloud integration account.
type NebiusCloudIntegrationAccountSettingsResponse struct {
	// Tags applied to all data collected from this Nebius Cloud account.
	AccountTags []string `json:"account_tags,omitempty"`
	// ID of the Nebius tenant being monitored.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountSettingsResponse instantiates a new NebiusCloudIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountSettingsResponse(tenantId string) *NebiusCloudIntegrationAccountSettingsResponse {
	this := NebiusCloudIntegrationAccountSettingsResponse{}
	this.TenantId = tenantId
	return &this
}

// NewNebiusCloudIntegrationAccountSettingsResponseWithDefaults instantiates a new NebiusCloudIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountSettingsResponseWithDefaults() *NebiusCloudIntegrationAccountSettingsResponse {
	this := NebiusCloudIntegrationAccountSettingsResponse{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountSettingsResponse) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsResponse) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountSettingsResponse) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *NebiusCloudIntegrationAccountSettingsResponse) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetTenantId returns the TenantId field value.
func (o *NebiusCloudIntegrationAccountSettingsResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsResponse) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *NebiusCloudIntegrationAccountSettingsResponse) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags != nil {
		toSerialize["account_tags"] = o.AccountTags
	}
	toSerialize["tenant_id"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags []string `json:"account_tags,omitempty"`
		TenantId    *string  `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "tenant_id"})
	} else {
		return err
	}
	o.AccountTags = all.AccountTags
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
