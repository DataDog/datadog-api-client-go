// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountBearerTokenAuthUpdate Bearer token authentication. Use a workspace-scoped OpenAI Admin key with the `codex.enterprise.analytics.read` scope. Only the fields provided are changed; omit `token` to keep the stored one.
type OpenAICodexIntegrationAccountBearerTokenAuthUpdate struct {
	// The authentication method type.
	AuthType OpenAICodexIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Secret token used to authenticate with OpenAI Codex.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountBearerTokenAuthUpdate instantiates a new OpenAICodexIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountBearerTokenAuthUpdate(authType OpenAICodexIntegrationAccountBearerTokenAuthType) *OpenAICodexIntegrationAccountBearerTokenAuthUpdate {
	this := OpenAICodexIntegrationAccountBearerTokenAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewOpenAICodexIntegrationAccountBearerTokenAuthUpdateWithDefaults instantiates a new OpenAICodexIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountBearerTokenAuthUpdateWithDefaults() *OpenAICodexIntegrationAccountBearerTokenAuthUpdate {
	this := OpenAICodexIntegrationAccountBearerTokenAuthUpdate{}
	var authType OpenAICodexIntegrationAccountBearerTokenAuthType = OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) GetAuthType() OpenAICodexIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret OpenAICodexIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) GetAuthTypeOk() (*OpenAICodexIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) SetAuthType(v OpenAICodexIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountBearerTokenAuthUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *OpenAICodexIntegrationAccountBearerTokenAuthType `json:"auth_type"`
		Token    *string                                           `json:"token,omitempty"`
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
	o.Token = all.Token

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
