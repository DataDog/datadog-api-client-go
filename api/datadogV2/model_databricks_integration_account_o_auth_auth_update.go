// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountOAuthAuthUpdate Databricks OAuth machine-to-machine authentication using a service principal. Omit `client_secret` to keep the stored one; `client_id` must be provided on every submission. Omitting `azure_tenant_id` clears it.
type DatabricksIntegrationAccountOAuthAuthUpdate struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
	// Microsoft Entra ID tenant of the service principal, for Azure Databricks workspaces.
	AzureTenantId *string `json:"azure_tenant_id,omitempty"`
	// Client ID of the Databricks service principal.
	ClientId string `json:"client_id"`
	// Secret of the Databricks service principal.
	ClientSecret *string `json:"client_secret,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountOAuthAuthUpdate instantiates a new DatabricksIntegrationAccountOAuthAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountOAuthAuthUpdate(authType DatabricksIntegrationAccountOAuthAuthType, clientId string) *DatabricksIntegrationAccountOAuthAuthUpdate {
	this := DatabricksIntegrationAccountOAuthAuthUpdate{}
	this.AuthType = authType
	this.ClientId = clientId
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

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) HasAzureTenantId() bool {
	return o != nil && o.AzureTenantId != nil
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetClientId returns the ClientId field value.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) SetClientId(v string) {
	o.ClientId = v
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
	if o.AzureTenantId != nil {
		toSerialize["azure_tenant_id"] = o.AzureTenantId
	}
	toSerialize["client_id"] = o.ClientId
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountOAuthAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType      *DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
		AzureTenantId *string                                    `json:"azure_tenant_id,omitempty"`
		ClientId      *string                                    `json:"client_id"`
		ClientSecret  *string                                    `json:"client_secret,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.AzureTenantId = all.AzureTenantId
	o.ClientId = *all.ClientId
	o.ClientSecret = all.ClientSecret

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
