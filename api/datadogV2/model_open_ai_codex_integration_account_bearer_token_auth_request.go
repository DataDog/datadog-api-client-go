// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountBearerTokenAuthRequest Bearer token authentication. Use a workspace-scoped OpenAI Admin key with the `codex.enterprise.analytics.read` scope.
type OpenAICodexIntegrationAccountBearerTokenAuthRequest struct {
	// The authentication method type.
	AuthType OpenAICodexIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Secret token used to authenticate with OpenAI Codex.
	Token string `json:"token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountBearerTokenAuthRequest instantiates a new OpenAICodexIntegrationAccountBearerTokenAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountBearerTokenAuthRequest(authType OpenAICodexIntegrationAccountBearerTokenAuthType, token string) *OpenAICodexIntegrationAccountBearerTokenAuthRequest {
	this := OpenAICodexIntegrationAccountBearerTokenAuthRequest{}
	this.AuthType = authType
	this.Token = token
	return &this
}

// NewOpenAICodexIntegrationAccountBearerTokenAuthRequestWithDefaults instantiates a new OpenAICodexIntegrationAccountBearerTokenAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountBearerTokenAuthRequestWithDefaults() *OpenAICodexIntegrationAccountBearerTokenAuthRequest {
	this := OpenAICodexIntegrationAccountBearerTokenAuthRequest{}
	var authType OpenAICodexIntegrationAccountBearerTokenAuthType = OPENAICODEXINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) GetAuthType() OpenAICodexIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret OpenAICodexIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) GetAuthTypeOk() (*OpenAICodexIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) SetAuthType(v OpenAICodexIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) SetToken(v string) {
	o.Token = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountBearerTokenAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["token"] = o.Token
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountBearerTokenAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *OpenAICodexIntegrationAccountBearerTokenAuthType `json:"auth_type"`
		Token    *string                                           `json:"token"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	if all.Token == nil {
		return fmt.Errorf("required field token missing")
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}
	o.Token = *all.Token

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
