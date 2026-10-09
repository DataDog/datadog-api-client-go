// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity Authenticate using an Azure managed identity.
type ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity struct {
	// The Azure credential kind. The value should always be `managed_identity`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind `json:"azure_credential_kind"`
	// The ID of the user-assigned managed identity. If omitted, the system-assigned managed identity is used.
	UserAssignedManagedIdentityId datadog.NullableString `json:"user_assigned_managed_identity_id,omitempty"`
	// The type of the user-assigned managed identity ID.
	UserAssignedManagedIdentityIdType *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType `json:"user_assigned_managed_identity_id_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity(azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind) *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity{}
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHMANAGEDIDENTITYKIND_MANAGED_IDENTITY
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind) {
	o.AzureCredentialKind = v
}

// GetUserAssignedManagedIdentityId returns the UserAssignedManagedIdentityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetUserAssignedManagedIdentityId() string {
	if o == nil || o.UserAssignedManagedIdentityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserAssignedManagedIdentityId.Get()
}

// GetUserAssignedManagedIdentityIdOk returns a tuple with the UserAssignedManagedIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetUserAssignedManagedIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAssignedManagedIdentityId.Get(), o.UserAssignedManagedIdentityId.IsSet()
}

// HasUserAssignedManagedIdentityId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) HasUserAssignedManagedIdentityId() bool {
	return o != nil && o.UserAssignedManagedIdentityId.IsSet()
}

// SetUserAssignedManagedIdentityId gets a reference to the given datadog.NullableString and assigns it to the UserAssignedManagedIdentityId field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) SetUserAssignedManagedIdentityId(v string) {
	o.UserAssignedManagedIdentityId.Set(&v)
}

// SetUserAssignedManagedIdentityIdNil sets the value for UserAssignedManagedIdentityId to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) SetUserAssignedManagedIdentityIdNil() {
	o.UserAssignedManagedIdentityId.Set(nil)
}

// UnsetUserAssignedManagedIdentityId ensures that no value is present for UserAssignedManagedIdentityId, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) UnsetUserAssignedManagedIdentityId() {
	o.UserAssignedManagedIdentityId.Unset()
}

// GetUserAssignedManagedIdentityIdType returns the UserAssignedManagedIdentityIdType field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetUserAssignedManagedIdentityIdType() ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType {
	if o == nil || o.UserAssignedManagedIdentityIdType == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType
		return ret
	}
	return *o.UserAssignedManagedIdentityIdType
}

// GetUserAssignedManagedIdentityIdTypeOk returns a tuple with the UserAssignedManagedIdentityIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) GetUserAssignedManagedIdentityIdTypeOk() (*ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType, bool) {
	if o == nil || o.UserAssignedManagedIdentityIdType == nil {
		return nil, false
	}
	return o.UserAssignedManagedIdentityIdType, true
}

// HasUserAssignedManagedIdentityIdType returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) HasUserAssignedManagedIdentityIdType() bool {
	return o != nil && o.UserAssignedManagedIdentityIdType != nil
}

// SetUserAssignedManagedIdentityIdType gets a reference to the given ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType and assigns it to the UserAssignedManagedIdentityIdType field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) SetUserAssignedManagedIdentityIdType(v ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType) {
	o.UserAssignedManagedIdentityIdType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind
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
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureCredentialKind               *ObservabilityPipelineAzureDataExplorerDestinationAuthManagedIdentityKind `json:"azure_credential_kind"`
		UserAssignedManagedIdentityId     datadog.NullableString                                                    `json:"user_assigned_managed_identity_id,omitempty"`
		UserAssignedManagedIdentityIdType *ObservabilityPipelineAzureDataExplorerDestinationManagedIdentityIdType   `json:"user_assigned_managed_identity_id_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_credential_kind", "user_assigned_managed_identity_id", "user_assigned_managed_identity_id_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}
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
