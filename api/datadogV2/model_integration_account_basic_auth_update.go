// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAccountBasicAuthUpdate Username and password authentication. Only the fields provided are changed; omit `password` to keep the stored one.
type IntegrationAccountBasicAuthUpdate struct {
	// The authentication method type.
	AuthType IntegrationAccountBasicAuthType `json:"auth_type"`
	// Secret password or private key.
	Password *string `json:"password,omitempty"`
	// Non-secret username or public identifier for the credential pair.
	Username *string `json:"username,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationAccountBasicAuthUpdate instantiates a new IntegrationAccountBasicAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationAccountBasicAuthUpdate(authType IntegrationAccountBasicAuthType) *IntegrationAccountBasicAuthUpdate {
	this := IntegrationAccountBasicAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewIntegrationAccountBasicAuthUpdateWithDefaults instantiates a new IntegrationAccountBasicAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationAccountBasicAuthUpdateWithDefaults() *IntegrationAccountBasicAuthUpdate {
	this := IntegrationAccountBasicAuthUpdate{}
	var authType IntegrationAccountBasicAuthType = INTEGRATIONACCOUNTBASICAUTHTYPE_BASIC
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *IntegrationAccountBasicAuthUpdate) GetAuthType() IntegrationAccountBasicAuthType {
	if o == nil {
		var ret IntegrationAccountBasicAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *IntegrationAccountBasicAuthUpdate) GetAuthTypeOk() (*IntegrationAccountBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *IntegrationAccountBasicAuthUpdate) SetAuthType(v IntegrationAccountBasicAuthType) {
	o.AuthType = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IntegrationAccountBasicAuthUpdate) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountBasicAuthUpdate) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IntegrationAccountBasicAuthUpdate) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IntegrationAccountBasicAuthUpdate) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IntegrationAccountBasicAuthUpdate) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationAccountBasicAuthUpdate) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IntegrationAccountBasicAuthUpdate) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IntegrationAccountBasicAuthUpdate) SetUsername(v string) {
	o.Username = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationAccountBasicAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationAccountBasicAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *IntegrationAccountBasicAuthType `json:"auth_type"`
		Password *string                          `json:"password,omitempty"`
		Username *string                          `json:"username,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type", "password", "username"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.Password = all.Password
	o.Username = all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
