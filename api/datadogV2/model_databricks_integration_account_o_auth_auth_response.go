// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountOAuthAuthResponse The Databricks OAuth authentication method and service principal configured on the account.
type DatabricksIntegrationAccountOAuthAuthResponse struct {
	// The authentication method type.
	AuthType DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
	// Microsoft Entra ID tenant of the service principal, for Azure Databricks workspaces.
	AzureTenantId *string `json:"azure_tenant_id,omitempty"`
	// Client ID of the Databricks service principal.
	ClientId string `json:"client_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountOAuthAuthResponse instantiates a new DatabricksIntegrationAccountOAuthAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountOAuthAuthResponse(authType DatabricksIntegrationAccountOAuthAuthType, clientId string) *DatabricksIntegrationAccountOAuthAuthResponse {
	this := DatabricksIntegrationAccountOAuthAuthResponse{}
	this.AuthType = authType
	this.ClientId = clientId
	return &this
}

// NewDatabricksIntegrationAccountOAuthAuthResponseWithDefaults instantiates a new DatabricksIntegrationAccountOAuthAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountOAuthAuthResponseWithDefaults() *DatabricksIntegrationAccountOAuthAuthResponse {
	this := DatabricksIntegrationAccountOAuthAuthResponse{}
	var authType DatabricksIntegrationAccountOAuthAuthType = DATABRICKSINTEGRATIONACCOUNTOAUTHAUTHTYPE_DATABRICKS_OAUTH
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetAuthType() DatabricksIntegrationAccountOAuthAuthType {
	if o == nil {
		var ret DatabricksIntegrationAccountOAuthAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetAuthTypeOk() (*DatabricksIntegrationAccountOAuthAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) SetAuthType(v DatabricksIntegrationAccountOAuthAuthType) {
	o.AuthType = v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) HasAzureTenantId() bool {
	return o != nil && o.AzureTenantId != nil
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetClientId returns the ClientId field value.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) SetClientId(v string) {
	o.ClientId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountOAuthAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.AzureTenantId != nil {
		toSerialize["azure_tenant_id"] = o.AzureTenantId
	}
	toSerialize["client_id"] = o.ClientId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountOAuthAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType      *DatabricksIntegrationAccountOAuthAuthType `json:"auth_type"`
		AzureTenantId *string                                    `json:"azure_tenant_id,omitempty"`
		ClientId      *string                                    `json:"client_id"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "azure_tenant_id", "client_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.AzureTenantId = all.AzureTenantId
	o.ClientId = *all.ClientId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
