// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountServiceAccountAuthResponse The service-account key pair authentication method configured on the account.
type NebiusCloudIntegrationAccountServiceAccountAuthResponse struct {
	// The authentication method type.
	AuthType NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
	// ID of the public key half of the authorized key pair. Omitted when unset.
	PublicKeyId *string `json:"public_key_id,omitempty"`
	// ID of the Nebius service account the account authenticates as.
	ServiceAccountId string `json:"service_account_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthResponse instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountServiceAccountAuthResponse(authType NebiusCloudIntegrationAccountServiceAccountAuthType, serviceAccountId string) *NebiusCloudIntegrationAccountServiceAccountAuthResponse {
	this := NebiusCloudIntegrationAccountServiceAccountAuthResponse{}
	this.AuthType = authType
	this.ServiceAccountId = serviceAccountId
	return &this
}

// NewNebiusCloudIntegrationAccountServiceAccountAuthResponseWithDefaults instantiates a new NebiusCloudIntegrationAccountServiceAccountAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountServiceAccountAuthResponseWithDefaults() *NebiusCloudIntegrationAccountServiceAccountAuthResponse {
	this := NebiusCloudIntegrationAccountServiceAccountAuthResponse{}
	var authType NebiusCloudIntegrationAccountServiceAccountAuthType = NEBIUSCLOUDINTEGRATIONACCOUNTSERVICEACCOUNTAUTHTYPE_NEBIUS_SERVICE_ACCOUNT
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetAuthType() NebiusCloudIntegrationAccountServiceAccountAuthType {
	if o == nil {
		var ret NebiusCloudIntegrationAccountServiceAccountAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetAuthTypeOk() (*NebiusCloudIntegrationAccountServiceAccountAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) SetAuthType(v NebiusCloudIntegrationAccountServiceAccountAuthType) {
	o.AuthType = v
}

// GetPublicKeyId returns the PublicKeyId field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetPublicKeyId() string {
	if o == nil || o.PublicKeyId == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyId
}

// GetPublicKeyIdOk returns a tuple with the PublicKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetPublicKeyIdOk() (*string, bool) {
	if o == nil || o.PublicKeyId == nil {
		return nil, false
	}
	return o.PublicKeyId, true
}

// HasPublicKeyId returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) HasPublicKeyId() bool {
	return o != nil && o.PublicKeyId != nil
}

// SetPublicKeyId gets a reference to the given string and assigns it to the PublicKeyId field.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) SetPublicKeyId(v string) {
	o.PublicKeyId = &v
}

// GetServiceAccountId returns the ServiceAccountId field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetServiceAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccountId
}

// GetServiceAccountIdOk returns a tuple with the ServiceAccountId field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) GetServiceAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountId, true
}

// SetServiceAccountId sets field value.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) SetServiceAccountId(v string) {
	o.ServiceAccountId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountServiceAccountAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.PublicKeyId != nil {
		toSerialize["public_key_id"] = o.PublicKeyId
	}
	toSerialize["service_account_id"] = o.ServiceAccountId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationAccountServiceAccountAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType         *NebiusCloudIntegrationAccountServiceAccountAuthType `json:"auth_type"`
		PublicKeyId      *string                                              `json:"public_key_id,omitempty"`
		ServiceAccountId *string                                              `json:"service_account_id"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "public_key_id", "service_account_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.PublicKeyId = all.PublicKeyId
	o.ServiceAccountId = *all.ServiceAccountId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
