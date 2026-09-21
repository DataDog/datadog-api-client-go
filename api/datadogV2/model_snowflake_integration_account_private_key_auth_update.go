// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountPrivateKeyAuthUpdate RSA key pair authentication, the only method Snowflake integration accounts support. Only the fields provided are changed; omit `private_key` to keep the stored one.
type SnowflakeIntegrationAccountPrivateKeyAuthUpdate struct {
	// The authentication method type.
	AuthType SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
	// The private key, in PEM format.
	PrivateKey *string `json:"private_key,omitempty"`
	// Name that distinguishes this private key from other keys in Datadog.
	PrivateKeyName *string `json:"private_key_name,omitempty"`
	// Passphrase that decrypts the private key. Provide it only when the key is encrypted. Omit it to keep the stored passphrase, send `null` or an empty string to remove it, or send a value to replace it.
	PrivateKeyPassphrase datadog.NullableString `json:"private_key_passphrase,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountPrivateKeyAuthUpdate instantiates a new SnowflakeIntegrationAccountPrivateKeyAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountPrivateKeyAuthUpdate(authType SnowflakeIntegrationAccountPrivateKeyAuthType) *SnowflakeIntegrationAccountPrivateKeyAuthUpdate {
	this := SnowflakeIntegrationAccountPrivateKeyAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewSnowflakeIntegrationAccountPrivateKeyAuthUpdateWithDefaults instantiates a new SnowflakeIntegrationAccountPrivateKeyAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountPrivateKeyAuthUpdateWithDefaults() *SnowflakeIntegrationAccountPrivateKeyAuthUpdate {
	this := SnowflakeIntegrationAccountPrivateKeyAuthUpdate{}
	var authType SnowflakeIntegrationAccountPrivateKeyAuthType = SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetAuthType() SnowflakeIntegrationAccountPrivateKeyAuthType {
	if o == nil {
		var ret SnowflakeIntegrationAccountPrivateKeyAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetAuthTypeOk() (*SnowflakeIntegrationAccountPrivateKeyAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SetAuthType(v SnowflakeIntegrationAccountPrivateKeyAuthType) {
	o.AuthType = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyName returns the PrivateKeyName field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKeyName() string {
	if o == nil || o.PrivateKeyName == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyName
}

// GetPrivateKeyNameOk returns a tuple with the PrivateKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKeyNameOk() (*string, bool) {
	if o == nil || o.PrivateKeyName == nil {
		return nil, false
	}
	return o.PrivateKeyName, true
}

// HasPrivateKeyName returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) HasPrivateKeyName() bool {
	return o != nil && o.PrivateKeyName != nil
}

// SetPrivateKeyName gets a reference to the given string and assigns it to the PrivateKeyName field.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SetPrivateKeyName(v string) {
	o.PrivateKeyName = &v
}

// GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKeyPassphrase() string {
	if o == nil || o.PrivateKeyPassphrase.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyPassphrase.Get()
}

// GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) GetPrivateKeyPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyPassphrase.Get(), o.PrivateKeyPassphrase.IsSet()
}

// HasPrivateKeyPassphrase returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) HasPrivateKeyPassphrase() bool {
	return o != nil && o.PrivateKeyPassphrase.IsSet()
}

// SetPrivateKeyPassphrase gets a reference to the given datadog.NullableString and assigns it to the PrivateKeyPassphrase field.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SetPrivateKeyPassphrase(v string) {
	o.PrivateKeyPassphrase.Set(&v)
}

// SetPrivateKeyPassphraseNil sets the value for PrivateKeyPassphrase to be an explicit nil.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) SetPrivateKeyPassphraseNil() {
	o.PrivateKeyPassphrase.Set(nil)
}

// UnsetPrivateKeyPassphrase ensures that no value is present for PrivateKeyPassphrase, not even an explicit nil.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) UnsetPrivateKeyPassphrase() {
	o.PrivateKeyPassphrase.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountPrivateKeyAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.PrivateKeyName != nil {
		toSerialize["private_key_name"] = o.PrivateKeyName
	}
	if o.PrivateKeyPassphrase.IsSet() {
		toSerialize["private_key_passphrase"] = o.PrivateKeyPassphrase.Get()
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountPrivateKeyAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType             *SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
		PrivateKey           *string                                        `json:"private_key,omitempty"`
		PrivateKeyName       *string                                        `json:"private_key_name,omitempty"`
		PrivateKeyPassphrase datadog.NullableString                         `json:"private_key_passphrase,omitempty"`
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
	o.PrivateKeyName = all.PrivateKeyName
	o.PrivateKeyPassphrase = all.PrivateKeyPassphrase

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
