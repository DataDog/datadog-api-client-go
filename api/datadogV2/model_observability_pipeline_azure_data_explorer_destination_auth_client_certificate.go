// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate Authenticate using a Microsoft Entra application client certificate.
type ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate struct {
	// The Microsoft Entra application (client) ID.
	AzureClientId string `json:"azure_client_id"`
	// The Azure credential kind. The value should always be `client_certificate_credential`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind `json:"azure_credential_kind"`
	// The Microsoft Entra tenant ID.
	AzureTenantId string `json:"azure_tenant_id"`
	// Name of the environment variable or secret that holds the password for the client certificate.
	CertificatePasswordKey datadog.NullableString `json:"certificate_password_key,omitempty"`
	// Path to the `.pfx` client certificate file on the Worker.
	CertificatePath string `json:"certificate_path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate(azureClientId string, azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind, azureTenantId string, certificatePath string) *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate{}
	this.AzureClientId = azureClientId
	this.AzureCredentialKind = azureCredentialKind
	this.AzureTenantId = azureTenantId
	this.CertificatePath = certificatePath
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHCLIENTCERTIFICATEKIND_CLIENT_CERTIFICATE_CREDENTIAL
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureClientId returns the AzureClientId field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureClientId, true
}

// SetAzureClientId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetAzureClientId(v string) {
	o.AzureClientId = v
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind) {
	o.AzureCredentialKind = v
}

// GetAzureTenantId returns the AzureTenantId field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetAzureTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureTenantId, true
}

// SetAzureTenantId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetAzureTenantId(v string) {
	o.AzureTenantId = v
}

// GetCertificatePasswordKey returns the CertificatePasswordKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetCertificatePasswordKey() string {
	if o == nil || o.CertificatePasswordKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.CertificatePasswordKey.Get()
}

// GetCertificatePasswordKeyOk returns a tuple with the CertificatePasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetCertificatePasswordKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificatePasswordKey.Get(), o.CertificatePasswordKey.IsSet()
}

// HasCertificatePasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) HasCertificatePasswordKey() bool {
	return o != nil && o.CertificatePasswordKey.IsSet()
}

// SetCertificatePasswordKey gets a reference to the given datadog.NullableString and assigns it to the CertificatePasswordKey field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetCertificatePasswordKey(v string) {
	o.CertificatePasswordKey.Set(&v)
}

// SetCertificatePasswordKeyNil sets the value for CertificatePasswordKey to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetCertificatePasswordKeyNil() {
	o.CertificatePasswordKey.Set(nil)
}

// UnsetCertificatePasswordKey ensures that no value is present for CertificatePasswordKey, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) UnsetCertificatePasswordKey() {
	o.CertificatePasswordKey.Unset()
}

// GetCertificatePath returns the CertificatePath field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetCertificatePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CertificatePath
}

// GetCertificatePathOk returns a tuple with the CertificatePath field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) GetCertificatePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificatePath, true
}

// SetCertificatePath sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) SetCertificatePath(v string) {
	o.CertificatePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_client_id"] = o.AzureClientId
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind
	toSerialize["azure_tenant_id"] = o.AzureTenantId
	if o.CertificatePasswordKey.IsSet() {
		toSerialize["certificate_password_key"] = o.CertificatePasswordKey.Get()
	}
	toSerialize["certificate_path"] = o.CertificatePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureClientId          *string                                                                     `json:"azure_client_id"`
		AzureCredentialKind    *ObservabilityPipelineAzureDataExplorerDestinationAuthClientCertificateKind `json:"azure_credential_kind"`
		AzureTenantId          *string                                                                     `json:"azure_tenant_id"`
		CertificatePasswordKey datadog.NullableString                                                      `json:"certificate_password_key,omitempty"`
		CertificatePath        *string                                                                     `json:"certificate_path"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureClientId == nil {
		return fmt.Errorf("required field azure_client_id missing")
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	if all.AzureTenantId == nil {
		return fmt.Errorf("required field azure_tenant_id missing")
	}
	if all.CertificatePath == nil {
		return fmt.Errorf("required field certificate_path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_client_id", "azure_credential_kind", "azure_tenant_id", "certificate_password_key", "certificate_path"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AzureClientId = *all.AzureClientId
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}
	o.AzureTenantId = *all.AzureTenantId
	o.CertificatePasswordKey = all.CertificatePasswordKey
	o.CertificatePath = *all.CertificatePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
