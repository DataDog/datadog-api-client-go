// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion Authenticate using a managed identity as a client assertion for a Microsoft Entra application.
type ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion struct {
	// The Azure credential kind. The value should always be `managed_identity_client_assertion`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind `json:"azure_credential_kind"`
	// The client ID of the Microsoft Entra application that trusts the managed identity.
	ClientAssertionClientId string `json:"client_assertion_client_id"`
	// The tenant ID of the Microsoft Entra application that trusts the managed identity.
	ClientAssertionTenantId string `json:"client_assertion_tenant_id"`
	// The ID of the user-assigned managed identity. If omitted, the system-assigned managed identity is used.
	UserAssignedManagedIdentityId datadog.NullableString `json:"user_assigned_managed_identity_id,omitempty"`
	// The type of the user-assigned managed identity ID.
	UserAssignedManagedIdentityIdType *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType `json:"user_assigned_managed_identity_id_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion instantiates a new ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion(azureCredentialKind ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind, clientAssertionClientId string, clientAssertionTenantId string) *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion {
	this := ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion{}
	this.AzureCredentialKind = azureCredentialKind
	this.ClientAssertionClientId = clientAssertionClientId
	this.ClientAssertionTenantId = clientAssertionTenantId
	return &this
}

// NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionWithDefaults() *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion {
	this := ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERMANAGEDIDENTITYCLIENTASSERTIONKIND_MANAGED_IDENTITY_CLIENT_ASSERTION
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind) {
	o.AzureCredentialKind = v
}

// GetClientAssertionClientId returns the ClientAssertionClientId field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetClientAssertionClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientAssertionClientId
}

// GetClientAssertionClientIdOk returns a tuple with the ClientAssertionClientId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetClientAssertionClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAssertionClientId, true
}

// SetClientAssertionClientId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetClientAssertionClientId(v string) {
	o.ClientAssertionClientId = v
}

// GetClientAssertionTenantId returns the ClientAssertionTenantId field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetClientAssertionTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientAssertionTenantId
}

// GetClientAssertionTenantIdOk returns a tuple with the ClientAssertionTenantId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetClientAssertionTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAssertionTenantId, true
}

// SetClientAssertionTenantId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetClientAssertionTenantId(v string) {
	o.ClientAssertionTenantId = v
}

// GetUserAssignedManagedIdentityId returns the UserAssignedManagedIdentityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetUserAssignedManagedIdentityId() string {
	if o == nil || o.UserAssignedManagedIdentityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserAssignedManagedIdentityId.Get()
}

// GetUserAssignedManagedIdentityIdOk returns a tuple with the UserAssignedManagedIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetUserAssignedManagedIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAssignedManagedIdentityId.Get(), o.UserAssignedManagedIdentityId.IsSet()
}

// HasUserAssignedManagedIdentityId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) HasUserAssignedManagedIdentityId() bool {
	return o != nil && o.UserAssignedManagedIdentityId.IsSet()
}

// SetUserAssignedManagedIdentityId gets a reference to the given datadog.NullableString and assigns it to the UserAssignedManagedIdentityId field.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetUserAssignedManagedIdentityId(v string) {
	o.UserAssignedManagedIdentityId.Set(&v)
}

// SetUserAssignedManagedIdentityIdNil sets the value for UserAssignedManagedIdentityId to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetUserAssignedManagedIdentityIdNil() {
	o.UserAssignedManagedIdentityId.Set(nil)
}

// UnsetUserAssignedManagedIdentityId ensures that no value is present for UserAssignedManagedIdentityId, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) UnsetUserAssignedManagedIdentityId() {
	o.UserAssignedManagedIdentityId.Unset()
}

// GetUserAssignedManagedIdentityIdType returns the UserAssignedManagedIdentityIdType field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetUserAssignedManagedIdentityIdType() ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType {
	if o == nil || o.UserAssignedManagedIdentityIdType == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType
		return ret
	}
	return *o.UserAssignedManagedIdentityIdType
}

// GetUserAssignedManagedIdentityIdTypeOk returns a tuple with the UserAssignedManagedIdentityIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) GetUserAssignedManagedIdentityIdTypeOk() (*ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType, bool) {
	if o == nil || o.UserAssignedManagedIdentityIdType == nil {
		return nil, false
	}
	return o.UserAssignedManagedIdentityIdType, true
}

// HasUserAssignedManagedIdentityIdType returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) HasUserAssignedManagedIdentityIdType() bool {
	return o != nil && o.UserAssignedManagedIdentityIdType != nil
}

// SetUserAssignedManagedIdentityIdType gets a reference to the given ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType and assigns it to the UserAssignedManagedIdentityIdType field.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) SetUserAssignedManagedIdentityIdType(v ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) {
	o.UserAssignedManagedIdentityIdType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind
	toSerialize["client_assertion_client_id"] = o.ClientAssertionClientId
	toSerialize["client_assertion_tenant_id"] = o.ClientAssertionTenantId
	if o.UserAssignedManagedIdentityId.IsSet() {
		toSerialize["user_assigned_managed_identity_id"] = o.UserAssignedManagedIdentityId.Get()
	}
	if o.UserAssignedManagedIdentityIdType != nil {
		toSerialize["user_assigned_managed_identity_id_type"] = o.UserAssignedManagedIdentityIdType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureCredentialKind               *ObservabilityPipelineAzureDataExplorerManagedIdentityClientAssertionKind `json:"azure_credential_kind"`
		ClientAssertionClientId           *string                                                                   `json:"client_assertion_client_id"`
		ClientAssertionTenantId           *string                                                                   `json:"client_assertion_tenant_id"`
		UserAssignedManagedIdentityId     datadog.NullableString                                                    `json:"user_assigned_managed_identity_id,omitempty"`
		UserAssignedManagedIdentityIdType *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType   `json:"user_assigned_managed_identity_id_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	if all.ClientAssertionClientId == nil {
		return fmt.Errorf("required field client_assertion_client_id missing")
	}
	if all.ClientAssertionTenantId == nil {
		return fmt.Errorf("required field client_assertion_tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_credential_kind", "client_assertion_client_id", "client_assertion_tenant_id", "user_assigned_managed_identity_id", "user_assigned_managed_identity_id_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}
	o.ClientAssertionClientId = *all.ClientAssertionClientId
	o.ClientAssertionTenantId = *all.ClientAssertionTenantId
	o.UserAssignedManagedIdentityId = all.UserAssignedManagedIdentityId
	if all.UserAssignedManagedIdentityIdType != nil && !all.UserAssignedManagedIdentityIdType.IsValid() {
		hasInvalidField = true
	} else {
		o.UserAssignedManagedIdentityIdType = all.UserAssignedManagedIdentityIdType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
