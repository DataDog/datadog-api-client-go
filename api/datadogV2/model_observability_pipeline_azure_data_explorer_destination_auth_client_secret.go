// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret Authenticate using a Microsoft Entra application client secret.
type ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret struct {
	// The Microsoft Entra application (client) ID.
	AzureClientId string `json:"azure_client_id"`
	// Name of the environment variable or secret that holds the Microsoft Entra application client secret.
	AzureClientSecretKey string `json:"azure_client_secret_key"`
	// The Azure credential kind. The value should always be `client_secret_credential`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind `json:"azure_credential_kind"`
	// The Microsoft Entra tenant ID.
	AzureTenantId string `json:"azure_tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret(azureClientId string, azureClientSecretKey string, azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind, azureTenantId string) *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret{}
	this.AzureClientId = azureClientId
	this.AzureClientSecretKey = azureClientSecretKey
	this.AzureCredentialKind = azureCredentialKind
	this.AzureTenantId = azureTenantId
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTSECRETKIND_CLIENT_SECRET_CREDENTIAL
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureClientId returns the AzureClientId field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureClientId, true
}

// SetAzureClientId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) SetAzureClientId(v string) {
	o.AzureClientId = v
}

// GetAzureClientSecretKey returns the AzureClientSecretKey field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureClientSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureClientSecretKey
}

// GetAzureClientSecretKeyOk returns a tuple with the AzureClientSecretKey field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureClientSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureClientSecretKey, true
}

// SetAzureClientSecretKey sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) SetAzureClientSecretKey(v string) {
	o.AzureClientSecretKey = v
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind) {
	o.AzureCredentialKind = v
}

// GetAzureTenantId returns the AzureTenantId field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureTenantId, true
}

// SetAzureTenantId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) SetAzureTenantId(v string) {
	o.AzureTenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_client_id"] = o.AzureClientId
	toSerialize["azure_client_secret_key"] = o.AzureClientSecretKey
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind
	toSerialize["azure_tenant_id"] = o.AzureTenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecret) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureClientId        *string                                                                `json:"azure_client_id"`
		AzureClientSecretKey *string                                                                `json:"azure_client_secret_key"`
		AzureCredentialKind  *ObservabilityPipelineAzureDataExplorerDestinationAuthClientSecretKind `json:"azure_credential_kind"`
		AzureTenantId        *string                                                                `json:"azure_tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureClientId == nil {
		return fmt.Errorf("required field azure_client_id missing")
	}
	if all.AzureClientSecretKey == nil {
		return fmt.Errorf("required field azure_client_secret_key missing")
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	if all.AzureTenantId == nil {
		return fmt.Errorf("required field azure_tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_client_id", "azure_client_secret_key", "azure_credential_kind", "azure_tenant_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AzureClientId = *all.AzureClientId
	o.AzureClientSecretKey = *all.AzureClientSecretKey
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}
	o.AzureTenantId = *all.AzureTenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
