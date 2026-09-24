// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountOAuthAuthUpdate Databricks OAuth machine-to-machine authentication using a service principal. Only the fields provided are changed; omit `client_secret` to keep the stored one.
type DatabricksIntegrationAccountOAuthAuthUpdate struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
	// Microsoft Entra ID tenant of the service principal, for Azure Databricks workspaces. Omit it to keep the stored tenant, send `null` or an empty string to remove it, or send a value to replace it.
	AzureTenantId datadog.NullableString `json:"azure_tenant_id,omitempty"`
	// Client ID of the Databricks service principal.
	ClientId *string `json:"client_id,omitempty"`
	// Secret of the Databricks service principal. Generate it under User management > Service principals > Credentials & secrets in Databricks.
	ClientSecret *string `json:"client_secret,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountOAuthAuthUpdate instantiates a new DatabricksIntegrationAccountOAuthAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountOAuthAuthUpdate(authType DatabricksIntegrationAccountOAuthAuthType) *DatabricksIntegrationAccountOAuthAuthUpdate {
	this := DatabricksIntegrationAccountOAuthAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewDatabricksIntegrationAccountOAuthAuthUpdateWithDefaults instantiates a new DatabricksIntegrationAccountOAuthAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountOAuthAuthUpdateWithDefaults() *DatabricksIntegrationAccountOAuthAuthUpdate {
	this := DatabricksIntegrationAccountOAuthAuthUpdate{}
	var authType DatabricksIntegrationAccountOAuthAuthType = DATABRICKSINTEGRATIONACCOUNTOAUTHAUTHTYPE_DATABRICKS_OAUTH
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAuthType() DatabricksIntegrationAccountOAuthAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountOAuthAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAuthTypeOk() (*DatabricksIntegrationAccountOAuthAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetAuthType(v DatabricksIntegrationAccountOAuthAuthType) {
	o.AuthType = v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId.Get()
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureTenantId.Get(), o.AzureTenantId.IsSet()
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) HasAzureTenantId() bool {
	return o != nil && o.AzureTenantId.IsSet()
}

// SetAzureTenantId gets a reference to the given datadog.NullableString and assigns it to the AzureTenantId field.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetAzureTenantId(v string) {
	o.AzureTenantId.Set(&v)
}

// SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetAzureTenantIdNil() {
	o.AzureTenantId.Set(nil)
}

// UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) UnsetAzureTenantId() {
	o.AzureTenantId.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountOAuthAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.AzureTenantId.IsSet() {
		toSerialize["azure_tenant_id"] = o.AzureTenantId.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType      *DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
		AzureTenantId datadog.NullableString                     `json:"azure_tenant_id,omitempty"`
		ClientId      *string                                    `json:"client_id,omitempty"`
		ClientSecret  *string                                    `json:"client_secret,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.AzureTenantId = all.AzureTenantId
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
