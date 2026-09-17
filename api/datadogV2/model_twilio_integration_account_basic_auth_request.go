// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountBasicAuthRequest Username and password authentication.
type TwilioIntegrationAccountBasicAuthRequest struct {
	// The authentication method type.
	AuthType TwilioIntegrationAccountBasicAuthType `json:"auth_type"`
	// Secret password or private key.
	Password string `json:"password"`
	// Non-secret username or public identifier for the credential pair.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountBasicAuthRequest instantiates a new TwilioIntegrationAccountBasicAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountBasicAuthRequest(authType TwilioIntegrationAccountBasicAuthType, password string, username string) *TwilioIntegrationAccountBasicAuthRequest {
	this := TwilioIntegrationAccountBasicAuthRequest{}
	this.AuthType = authType
	this.Password = password
	this.Username = username
	return &this
}

// NewTwilioIntegrationAccountBasicAuthRequestWithDefaults instantiates a new TwilioIntegrationAccountBasicAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountBasicAuthRequestWithDefaults() *TwilioIntegrationAccountBasicAuthRequest {
	this := TwilioIntegrationAccountBasicAuthRequest{}
	var authType TwilioIntegrationAccountBasicAuthType = TWILIOINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetAuthType() TwilioIntegrationAccountBasicAuthType {
	if o == nil {
		var ret TwilioIntegrationAccountBasicAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetAuthTypeOk() (*TwilioIntegrationAccountBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) SetAuthType(v TwilioIntegrationAccountBasicAuthType) {
	o.AuthType = v
}

// GetPassword returns the Password field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountBasicAuthRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *TwilioIntegrationAccountBasicAuthRequest) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountBasicAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationAccountBasicAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *TwilioIntegrationAccountBasicAuthType `json:"auth_type"`
		Password *string                                `json:"password"`
		Username *string                                `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.Password = *all.Password
	o.Username = *all.Username

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
