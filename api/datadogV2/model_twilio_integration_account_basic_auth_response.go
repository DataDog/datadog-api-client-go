// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountBasicAuthResponse The basic authentication method and username configured on the account.
type TwilioIntegrationAccountBasicAuthResponse struct {
	// The authentication method type.
	AuthType TwilioIntegrationAccountBasicAuthType `json:"auth_type"`
	// Non-secret username or public identifier for the credential pair.
	Username string `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountBasicAuthResponse instantiates a new TwilioIntegrationAccountBasicAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountBasicAuthResponse(authType TwilioIntegrationAccountBasicAuthType, username string) *TwilioIntegrationAccountBasicAuthResponse {
	this := TwilioIntegrationAccountBasicAuthResponse{}
	this.AuthType = authType
	this.Username = username
	return &this
}

// NewTwilioIntegrationAccountBasicAuthResponseWithDefaults instantiates a new TwilioIntegrationAccountBasicAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountBasicAuthResponseWithDefaults() *TwilioIntegrationAccountBasicAuthResponse {
	this := TwilioIntegrationAccountBasicAuthResponse{}
	var authType TwilioIntegrationAccountBasicAuthType = TWILIOINTEGRATIONACCOUNTBASICAUTHTYPE_BASIC
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *TwilioIntegrationAccountBasicAuthResponse) GetAuthType() TwilioIntegrationAccountBasicAuthType {
	if o == nil {
		var ret TwilioIntegrationAccountBasicAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountBasicAuthResponse) GetAuthTypeOk() (*TwilioIntegrationAccountBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *TwilioIntegrationAccountBasicAuthResponse) SetAuthType(v TwilioIntegrationAccountBasicAuthType) {
	o.AuthType = v
}

// GetUsername returns the Username field value.
func (o *TwilioIntegrationAccountBasicAuthResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountBasicAuthResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *TwilioIntegrationAccountBasicAuthResponse) SetUsername(v string) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountBasicAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationAccountBasicAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *TwilioIntegrationAccountBasicAuthType `json:"auth_type"`
		Username *string                                `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "username"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.Username = *all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
