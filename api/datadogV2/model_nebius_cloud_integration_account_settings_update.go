// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountSettingsUpdate Settings for updating the Nebius Cloud integration account. Only the fields provided are changed.
type NebiusCloudIntegrationAccountSettingsUpdate struct {
	// Tags applied to all metrics, logs, and resources collected from this Nebius Cloud account. Sending it replaces the whole list; send an empty array to remove all tags, or omit it to keep the stored list.
	AccountTags []string `json:"account_tags,omitempty"`
	// ID of the Nebius tenant to monitor. An empty value is ignored and the stored tenant ID is kept.
	TenantId *string `json:"tenant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountSettingsUpdate instantiates a new NebiusCloudIntegrationAccountSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountSettingsUpdate() *NebiusCloudIntegrationAccountSettingsUpdate {
	this := NebiusCloudIntegrationAccountSettingsUpdate{}
	return &this
}

// NewNebiusCloudIntegrationAccountSettingsUpdateWithDefaults instantiates a new NebiusCloudIntegrationAccountSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountSettingsUpdateWithDefaults() *NebiusCloudIntegrationAccountSettingsUpdate {
	this := NebiusCloudIntegrationAccountSettingsUpdate{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) SetTenantId(v string) {
	o.TenantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags != nil {
		toSerialize["account_tags"] = o.AccountTags
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags []string `json:"account_tags,omitempty"`
		TenantId    *string  `json:"tenant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.AccountTags = all.AccountTags
	o.TenantId = all.TenantId

	return nil
}
