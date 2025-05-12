// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthCredentials The auth credentials of the user. Consists of a public key fingerprint and private key.
type AuthCredentials struct {
	// The public key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// The `RSA` private key in `PEM` format.
	PrivateKey string `json:"private_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAuthCredentials instantiates a new AuthCredentials object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthCredentials(fingerprint string, privateKey string) *AuthCredentials {
	this := AuthCredentials{}
	this.Fingerprint = fingerprint
	this.PrivateKey = privateKey
	return &this
}

// NewAuthCredentialsWithDefaults instantiates a new AuthCredentials object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthCredentialsWithDefaults() *AuthCredentials {
	this := AuthCredentials{}
	return &this
}

// GetFingerprint returns the Fingerprint field value.
func (o *AuthCredentials) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AuthCredentials) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AuthCredentials) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetPrivateKey returns the PrivateKey field value.
func (o *AuthCredentials) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *AuthCredentials) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *AuthCredentials) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AuthCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["private_key"] = o.PrivateKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AuthCredentials) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fingerprint *string `json:"fingerprint"`
		PrivateKey  *string `json:"private_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.PrivateKey == nil {
		return fmt.Errorf("required field private_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fingerprint", "private_key"})
	} else {
		return err
	}
	o.Fingerprint = *all.Fingerprint
	o.PrivateKey = *all.PrivateKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
