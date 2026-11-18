// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountServiceAccountAuthRequest Service-account key pair authentication, the only method Nebius Cloud integration accounts support. Create a service account in the Nebius console, add it to the admin group, upload an authorized key pair to it, and issue a static API key with the OBSERVABILITY scope through the Nebius CLI.
type NebiusCloudIntegrationAccountServiceAccountAuthRequest struct {
	// The authentication method type.
	AuthType NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
	// The private key of the authorized key pair, in PEM format. Required when log, resource, or service-quota collection is enabled; a request that enables one of them without it is rejected with a `422` response.
	PrivateKey *string `json:"private_key,omitempty"`
	// ID of the public key half of the authorized key pair, such as `publickey-xxxxxxxx`. Required when log, resource, or service-quota collection is enabled; a request that enables one of them without it is rejected with a `422` response.
	PublicKeyId *string `json:"public_key_id,omitempty"`
	// Non-secret ID of the Nebius service account the account authenticates as.
	ServiceAccountId string `json:"service_account_id"`
	// Nebius static API key with the OBSERVABILITY scope, used to query logs. Required when log, resource, or service-quota collection is enabled; a request that enables one of them without it is rejected with a `422` response.
	StaticApiKey *string `json:"static_api_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthRequest instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountServiceAccountAuthRequest(authType NebiusCloudIntegrationAccountServiceAccountAuthType, serviceAccountId string) *NebiusCloudIntegrationAccountServiceAccountAuthRequest {
	this := NebiusCloudIntegrationAccountServiceAccountAuthRequest{}
	this.AuthType = authType
	this.ServiceAccountId = serviceAccountId
	return &this
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthRequestWithDefaults instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountServiceAccountAuthRequestWithDefaults() *NebiusCloudIntegrationAccountServiceAccountAuthRequest {
	this := NebiusCloudIntegrationAccountServiceAccountAuthRequest{}
	var authType NebiusCloudIntegrationAccountServiceAccountAuthType = NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetAuthType() NebiusCloudIntegrationAccountServiceAccountAuthType {
	if o == nil {
		var ret NebiusCloudIntegrationAccountServiceAccountAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetAuthTypeOk() (*NebiusCloudIntegrationAccountServiceAccountAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) SetAuthType(v NebiusCloudIntegrationAccountServiceAccountAuthType) {
	o.AuthType = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetPublicKeyId() string {
	if o == nil || o.PublicKeyId == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || o.PublicKeyId == nil {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) HasPublicKeyId() bool {
	return o != nil && o.PublicKeyId != nil
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetServiceAccountId returns the ServiceAccountId field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetServiceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccountId
}

// GetServiceAccountIdOk returns a tuple with the ServiceAccountId field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetServiceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountId, true
}

// SetServiceAccountId sets field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) SetServiceAccountId(v string) {
	o.ServiceAccountId = v
}

// GetStaticApiKey returns the StaticApiKey field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetStaticApiKey() string {
	if o == nil || o.StaticApiKey == nil {
		var ret string
		return ret
	}
	return *o.StaticApiKey
}

// GetStaticApiKeyOk returns a tuple with the StaticApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) GetStaticApiKeyOk() (*string, bool) {
	if o == nil || o.StaticApiKey == nil {
		return nil, false
	}
	return o.StaticApiKey, true
}

// HasStaticApiKey returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) HasStaticApiKey() bool {
	return o != nil && o.StaticApiKey != nil
}

// SetStaticApiKey gets a reference to the given string and assigns it to the StaticApiKey field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) SetStaticApiKey(v string) {
	o.StaticApiKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountServiceAccountAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.PublicKeyId != nil {
		toSerialize["public_key_id"] = o.PublicKeyId
	}
	toSerialize["service_account_id"] = o.ServiceAccountId
	if o.StaticApiKey != nil {
		toSerialize["static_api_key"] = o.StaticApiKey
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType         *NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
		PrivateKey       *string                                              `json:"private_key,omitempty"`
		PublicKeyId      *string                                              `json:"public_key_id,omitempty"`
		ServiceAccountId *string                                              `json:"service_account_id"`
		StaticApiKey     *string                                              `json:"static_api_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.ServiceAccountId == nil {
		return fmt.Errorf("required field service_account_id missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.PrivateKey = all.PrivateKey
	o.PublicKeyId = all.PublicKeyId
	o.ServiceAccountId = *all.ServiceAccountId
	o.StaticApiKey = all.StaticApiKey

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
