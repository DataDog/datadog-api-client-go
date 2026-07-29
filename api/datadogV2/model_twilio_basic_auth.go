// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioBasicAuth API Key & Secret authentication for Twilio.
type TwilioBasicAuth struct {
	// Twilio API Key SID for authentication. Create from Twilio Console > Account > API Keys & Tokens.
	ApiKey string `json:"api_key"`
	// Twilio API Key Secret (token) corresponding to the API Key SID.
	ApiKeyToken string `json:"api_key_token"`
	// Authentication method discriminator.
	Type TwilioBasicAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioBasicAuth instantiates a new TwilioBasicAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioBasicAuth(apiKey string, apiKeyToken string, typeVar TwilioBasicAuthType) *TwilioBasicAuth {
	this := TwilioBasicAuth{}
	this.ApiKey = apiKey
	this.ApiKeyToken = apiKeyToken
	this.Type = typeVar
	return &this
}

// NewTwilioBasicAuthWithDefaults instantiates a new TwilioBasicAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioBasicAuthWithDefaults() *TwilioBasicAuth {
	this := TwilioBasicAuth{}
	var typeVar TwilioBasicAuthType = TWILIOBASICAUTHTYPE_BASIC
	this.Type = typeVar
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *TwilioBasicAuth) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *TwilioBasicAuth) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *TwilioBasicAuth) SetApiKey(v string) {
	o.ApiKey = v
}

// GetApiKeyToken returns the ApiKeyToken field value.
func (o *TwilioBasicAuth) GetApiKeyToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKeyToken
}

// GetApiKeyTokenOk returns a tuple with the ApiKeyToken field value
// and a boolean to check if the value has been set.
func (o *TwilioBasicAuth) GetApiKeyTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyToken, true
}

// SetApiKeyToken sets field value.
func (o *TwilioBasicAuth) SetApiKeyToken(v string) {
	o.ApiKeyToken = v
}

// GetType returns the Type field value.
func (o *TwilioBasicAuth) GetType() TwilioBasicAuthType {
	if o == nil {
		var ret TwilioBasicAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TwilioBasicAuth) GetTypeOk() (*TwilioBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TwilioBasicAuth) SetType(v TwilioBasicAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioBasicAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["api_key_token"] = o.ApiKeyToken
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioBasicAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey      *string              `json:"api_key"`
		ApiKeyToken *string              `json:"api_key_token"`
		Type        *TwilioBasicAuthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.ApiKeyToken == nil {
		return fmt.Errorf("required field api_key_token missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "api_key_token", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = *all.ApiKey
	o.ApiKeyToken = *all.ApiKeyToken
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
