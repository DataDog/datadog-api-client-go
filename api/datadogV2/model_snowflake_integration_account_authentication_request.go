// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationRequest RSA key pair authentication, the only method Snowflake integration accounts support. Generate an RSA key pair and assign the public key to the Snowflake user named in `settings.username`. Because an update replaces this object as a whole, every required field must be sent again on each update, even when only one of them is changing.
type SnowflakeIntegrationAccountAuthenticationRequest struct {
	// The authentication method type.
	AuthType SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
	// The private key, in PEM format.
	PrivateKey string `json:"private_key"`
	// Name that distinguishes this private key from other keys in Datadog.
	PrivateKeyName string `json:"private_key_name"`
	// Passphrase that decrypts the private key. Provide it only when the key is encrypted.
	PrivateKeyPassphrase *string `json:"private_key_passphrase,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountAuthenticationRequest instantiates a new SnowflakeIntegrationAccountAuthenticationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountAuthenticationRequest(authType SnowflakeIntegrationAccountPrivateKeyAuthType, privateKey string, privateKeyName string) *SnowflakeIntegrationAccountAuthenticationRequest {
	this := SnowflakeIntegrationAccountAuthenticationRequest{}
	this.AuthType = authType
	this.PrivateKey = privateKey
	this.PrivateKeyName = privateKeyName
	return &this
}

// NewSnowflakeIntegrationAccountAuthenticationRequestWithDefaults instantiates a new SnowflakeIntegrationAccountAuthenticationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountAuthenticationRequestWithDefaults() *SnowflakeIntegrationAccountAuthenticationRequest {
	this := SnowflakeIntegrationAccountAuthenticationRequest{}
	var authType SnowflakeIntegrationAccountPrivateKeyAuthType = SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetAuthType() SnowflakeIntegrationAccountPrivateKeyAuthType {
	if o == nil {
		var ret SnowflakeIntegrationAccountPrivateKeyAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetAuthTypeOk() (*SnowflakeIntegrationAccountPrivateKeyAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) SetAuthType(v SnowflakeIntegrationAccountPrivateKeyAuthType) {
	o.AuthType = v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPrivateKeyName returns the PrivateKeyName field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKeyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKeyName
}

// GetPrivateKeyNameOk returns a tuple with the PrivateKeyName field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyName, true
}

// SetPrivateKeyName sets field value.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) SetPrivateKeyName(v string) {
	o.PrivateKeyName = v
}

// GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKeyPassphrase() string {
	if o == nil || o.PrivateKeyPassphrase == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassphrase
}

// GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) GetPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil || o.PrivateKeyPassphrase == nil {
		return nil, false
	}
	return o.PrivateKeyPassphrase, true
}

// HasPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) HasPrivateKeyPassphrase() bool {
	return o != nil && o.PrivateKeyPassphrase != nil
}

// SetPrivateKeyPassphrase gets a reference to the given string and assigns it to the PrivateKeyPassphrase field.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) SetPrivateKeyPassphrase(v string) {
	o.PrivateKeyPassphrase = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["private_key_name"] = o.PrivateKeyName
	if o.PrivateKeyPassphrase != nil {
		toSerialize["private_key_passphrase"] = o.PrivateKeyPassphrase
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountAuthenticationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType             *SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
		PrivateKey           *string                                        `json:"private_key"`
		PrivateKeyName       *string                                        `json:"private_key_name"`
		PrivateKeyPassphrase *string                                        `json:"private_key_passphrase,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	if all.PrivateKeyName == nil {
		return fmt.Errorf("required field private_key_name missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.PrivateKey = *all.PrivateKey
	o.PrivateKeyName = *all.PrivateKeyName
	o.PrivateKeyPassphrase = all.PrivateKeyPassphrase

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
