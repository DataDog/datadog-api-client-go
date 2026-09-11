// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationResponse Authentication configured on the Snowflake integration account.
type SnowflakeIntegrationAccountAuthenticationResponse struct {
	// The authentication method type.
	AuthType SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
	// Name that distinguishes this private key from other keys in Datadog.
	PrivateKeyName string `json:"private_key_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountAuthenticationResponse instantiates a new SnowflakeIntegrationAccountAuthenticationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountAuthenticationResponse(authType SnowflakeIntegrationAccountPrivateKeyAuthType, privateKeyName string) *SnowflakeIntegrationAccountAuthenticationResponse {
	this := SnowflakeIntegrationAccountAuthenticationResponse{}
	this.AuthType = authType
	this.PrivateKeyName = privateKeyName
	return &this
}

// NewSnowflakeIntegrationAccountAuthenticationResponseWithDefaults instantiates a new SnowflakeIntegrationAccountAuthenticationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountAuthenticationResponseWithDefaults() *SnowflakeIntegrationAccountAuthenticationResponse {
	this := SnowflakeIntegrationAccountAuthenticationResponse{}
	var authType SnowflakeIntegrationAccountPrivateKeyAuthType = SNOWFLAKEINTEGRATIONACCOUNTPRIVATEKEYAUTHTYPE_SNOWFLAKE_PRIVATE_KEY
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) GetAuthType() SnowflakeIntegrationAccountPrivateKeyAuthType {
	if o == nil {
		var ret SnowflakeIntegrationAccountPrivateKeyAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) GetAuthTypeOk() (*SnowflakeIntegrationAccountPrivateKeyAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) SetAuthType(v SnowflakeIntegrationAccountPrivateKeyAuthType) {
	o.AuthType = v
}

// GetPrivateKeyName returns the PrivateKeyName field value.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) GetPrivateKeyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrivateKeyName
}

// GetPrivateKeyNameOk returns a tuple with the PrivateKeyName field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) GetPrivateKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyName, true
}

// SetPrivateKeyName sets field value.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) SetPrivateKeyName(v string) {
	o.PrivateKeyName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["private_key_name"] = o.PrivateKeyName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountAuthenticationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType       *SnowflakeIntegrationAccountPrivateKeyAuthType `json:"auth_type"`
		PrivateKeyName *string                                        `json:"private_key_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.PrivateKeyName == nil {
		return fmt.Errorf("required field private_key_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "private_key_name"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.PrivateKeyName = *all.PrivateKeyName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
