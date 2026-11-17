// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountServiceAccountAuthUpdate Service-account key pair authentication, the only method Nebius Cloud integration accounts support. Only the fields provided are changed; omit `private_key` and `static_api_key` to keep the stored ones.
type NebiusCloudIntegrationAccountServiceAccountAuthUpdate struct {
	// The authentication method type.
	AuthType NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
	// The private key of the authorized key pair, in PEM format. Omit it to keep the stored one; an empty string is rejected with a `422` response.
	PrivateKey *string `json:"private_key,omitempty"`
	// ID of the public key half of the authorized key pair. Omit it to keep the stored public key ID, send `null` or an empty string to remove it, or send a value to replace it. Removing it is rejected with a `422` response while log, resource, or service-quota collection is enabled.
	PublicKeyId datadog.NullableString `json:"public_key_id,omitempty"`
	// Non-secret ID of the Nebius service account the account authenticates as. It cannot be cleared; sending `null` or an empty string is rejected with a `422` response.
	ServiceAccountId *string `json:"service_account_id,omitempty"`
	// Nebius static API key with the OBSERVABILITY scope. Omit it to keep the stored one; an empty string is rejected with a `422` response.
	StaticApiKey *string `json:"static_api_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthUpdate instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountServiceAccountAuthUpdate(authType NebiusCloudIntegrationAccountServiceAccountAuthType) *NebiusCloudIntegrationAccountServiceAccountAuthUpdate {
	this := NebiusCloudIntegrationAccountServiceAccountAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthUpdateWithDefaults instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountServiceAccountAuthUpdateWithDefaults() *NebiusCloudIntegrationAccountServiceAccountAuthUpdate {
	this := NebiusCloudIntegrationAccountServiceAccountAuthUpdate{}
	var authType NebiusCloudIntegrationAccountServiceAccountAuthType = NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetAuthType() NebiusCloudIntegrationAccountServiceAccountAuthType {
	if o == nil {
		var ret NebiusCloudIntegrationAccountServiceAccountAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetAuthTypeOk() (*NebiusCloudIntegrationAccountServiceAccountAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetAuthType(v NebiusCloudIntegrationAccountServiceAccountAuthType) {
	o.AuthType = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetPublicKeyId() string {
	if o == nil || o.PublicKeyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyId.Get()
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetPublicKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKeyId.Get(), o.PublicKeyId.IsSet()
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) HasPublicKeyId() bool {
	return o != nil && o.PublicKeyId.IsSet()
}

// SetPublicKeyId gets a reference to the given datadog.NullableString and assigns it to the PublicKeyId field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetPublicKeyId(v string) {
	o.PublicKeyId.Set(&v)
}

// SetPublicKeyIdNil sets the value for PublicKeyId to be an explicit nil.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetPublicKeyIdNil() {
	o.PublicKeyId.Set(nil)
}

// UnsetPublicKeyId ensures that no value is present for PublicKeyId, not even an explicit nil.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) UnsetPublicKeyId() {
	o.PublicKeyId.Unset()
}

// GetServiceAccountId returns the ServiceAccountId field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetServiceAccountId() string {
	if o == nil || o.ServiceAccountId == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountId
}

// GetServiceAccountIdOk returns a tuple with the ServiceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetServiceAccountIdOk() (*string, bool) {
	if o == nil || o.ServiceAccountId == nil {
		return nil, false
	}
	return o.ServiceAccountId, true
}

// HasServiceAccountId returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) HasServiceAccountId() bool {
	return o != nil && o.ServiceAccountId != nil
}

// SetServiceAccountId gets a reference to the given string and assigns it to the ServiceAccountId field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetServiceAccountId(v string) {
	o.ServiceAccountId = &v
}

// GetStaticApiKey returns the StaticApiKey field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetStaticApiKey() string {
	if o == nil || o.StaticApiKey == nil {
		var ret string
		return ret
	}
	return *o.StaticApiKey
}

// GetStaticApiKeyOk returns a tuple with the StaticApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) GetStaticApiKeyOk() (*string, bool) {
	if o == nil || o.StaticApiKey == nil {
		return nil, false
	}
	return o.StaticApiKey, true
}

// HasStaticApiKey returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) HasStaticApiKey() bool {
	return o != nil && o.StaticApiKey != nil
}

// SetStaticApiKey gets a reference to the given string and assigns it to the StaticApiKey field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) SetStaticApiKey(v string) {
	o.StaticApiKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountServiceAccountAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.PublicKeyId.IsSet() {
		toSerialize["public_key_id"] = o.PublicKeyId.Get()
	}
	if o.ServiceAccountId != nil {
		toSerialize["service_account_id"] = o.ServiceAccountId
	}
	if o.StaticApiKey != nil {
		toSerialize["static_api_key"] = o.StaticApiKey
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType         *NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
		PrivateKey       *string                                              `json:"private_key,omitempty"`
		PublicKeyId      datadog.NullableString                               `json:"public_key_id,omitempty"`
		ServiceAccountId *string                                              `json:"service_account_id,omitempty"`
		StaticApiKey     *string                                              `json:"static_api_key,omitempty"`
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
	o.PrivateKey = all.PrivateKey
	o.PublicKeyId = all.PublicKeyId
	o.ServiceAccountId = all.ServiceAccountId
	o.StaticApiKey = all.StaticApiKey

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
