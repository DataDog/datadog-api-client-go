// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceRegisterTokenResponseAttributes Attributes for the registered token.
type OnPremManagementServiceRegisterTokenResponseAttributes struct {
	// The token string.
	TokenString string `json:"token_string"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnPremManagementServiceRegisterTokenResponseAttributes instantiates a new OnPremManagementServiceRegisterTokenResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnPremManagementServiceRegisterTokenResponseAttributes(tokenString string) *OnPremManagementServiceRegisterTokenResponseAttributes {
	this := OnPremManagementServiceRegisterTokenResponseAttributes{}
	this.TokenString = tokenString
	return &this
}

// NewOnPremManagementServiceRegisterTokenResponseAttributesWithDefaults instantiates a new OnPremManagementServiceRegisterTokenResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnPremManagementServiceRegisterTokenResponseAttributesWithDefaults() *OnPremManagementServiceRegisterTokenResponseAttributes {
	this := OnPremManagementServiceRegisterTokenResponseAttributes{}
	return &this
}

// GetTokenString returns the TokenString field value.
func (o *OnPremManagementServiceRegisterTokenResponseAttributes) GetTokenString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TokenString
}

// GetTokenStringOk returns a tuple with the TokenString field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenResponseAttributes) GetTokenStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenString, true
}

// SetTokenString sets field value.
func (o *OnPremManagementServiceRegisterTokenResponseAttributes) SetTokenString(v string) {
	o.TokenString = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnPremManagementServiceRegisterTokenResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["token_string"] = o.TokenString

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnPremManagementServiceRegisterTokenResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TokenString *string `json:"token_string"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TokenString == nil {
		return fmt.Errorf("required field token_string missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"token_string"})
	} else {
		return err
	}
	o.TokenString = *all.TokenString

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
