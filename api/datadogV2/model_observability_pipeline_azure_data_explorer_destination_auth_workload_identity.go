// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity Authenticate using Azure Workload Identity (for example, on Kubernetes).
type ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity struct {
	// The Azure credential kind. The value should always be `workload_identity`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind `json:"azure_credential_kind"`
	// The client ID of the Microsoft Entra application. If omitted, it is read from the environment.
	ClientId datadog.NullableString `json:"client_id,omitempty"`
	// The Microsoft Entra tenant ID. If omitted, it is read from the environment.
	TenantId datadog.NullableString `json:"tenant_id,omitempty"`
	// Path to the federated token file. If omitted, it is read from the environment.
	TokenFilePath datadog.NullableString `json:"token_file_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity(azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity{}
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHWORKLOADIDENTITYKIND_WORKLOAD_IDENTITY
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind) {
	o.AzureCredentialKind = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetClientId() string {
	if o == nil || o.ClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) HasClientId() bool {
	return o != nil && o.ClientId.IsSet()
}

// SetClientId gets a reference to the given datadog.NullableString and assigns it to the ClientId field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) UnsetClientId() {
	o.ClientId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) HasTenantId() bool {
	return o != nil && o.TenantId.IsSet()
}

// SetTenantId gets a reference to the given datadog.NullableString and assigns it to the TenantId field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// SetTenantIdNil sets the value for TenantId to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTokenFilePath returns the TokenFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetTokenFilePath() string {
	if o == nil || o.TokenFilePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.TokenFilePath.Get()
}

// GetTokenFilePathOk returns a tuple with the TokenFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) GetTokenFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenFilePath.Get(), o.TokenFilePath.IsSet()
}

// HasTokenFilePath returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) HasTokenFilePath() bool {
	return o != nil && o.TokenFilePath.IsSet()
}

// SetTokenFilePath gets a reference to the given datadog.NullableString and assigns it to the TokenFilePath field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetTokenFilePath(v string) {
	o.TokenFilePath.Set(&v)
}

// SetTokenFilePathNil sets the value for TokenFilePath to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) SetTokenFilePathNil() {
	o.TokenFilePath.Set(nil)
}

// UnsetTokenFilePath ensures that no value is present for TokenFilePath, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) UnsetTokenFilePath() {
	o.TokenFilePath.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind
	if o.ClientId.IsSet() {
		toSerialize["client_id"] = o.ClientId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenant_id"] = o.TenantId.Get()
	}
	if o.TokenFilePath.IsSet() {
		toSerialize["token_file_path"] = o.TokenFilePath.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureCredentialKind *ObservabilityPipelineAzureDataExplorerDestinationAuthWorkloadIdentityKind `json:"azure_credential_kind"`
		ClientId            datadog.NullableString                                                     `json:"client_id,omitempty"`
		TenantId            datadog.NullableString                                                     `json:"tenant_id,omitempty"`
		TokenFilePath       datadog.NullableString                                                     `json:"token_file_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_credential_kind", "client_id", "tenant_id", "token_file_path"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}
	o.ClientId = all.ClientId
	o.TenantId = all.TenantId
	o.TokenFilePath = all.TokenFilePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
