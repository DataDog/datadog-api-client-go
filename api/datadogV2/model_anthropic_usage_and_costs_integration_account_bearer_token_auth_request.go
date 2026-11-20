// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest Bearer token authentication. Use an Anthropic Admin API key (Platform plan) or an Analytics API key with the `read:analytics` scope (Enterprise plan). An invalid key, or an Analytics key without the `read:analytics` scope, is rejected with a `422` response.
type AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest struct {
	// The authentication method type.
	AuthType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Anthropic Admin API key with permission to read usage and cost data, or an Analytics API key with the `read:analytics` scope. Generate an Admin API key in the Anthropic Console under **Settings > Admin Keys**.
	Token string `json:"token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest instantiates a new AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest(authType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType, token string) *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest {
	this := AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest{}
	this.AuthType = authType
	this.Token = token
	return &this
}

// NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequestWithDefaults instantiates a new AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequestWithDefaults() *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest {
	this := AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest{}
	var authType AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType = ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) GetAuthType() AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) GetAuthTypeOk() (*AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) SetAuthType(v AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) SetToken(v string) {
	o.Token = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType
	toSerialize["token"] = o.Token
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
		Token    *string                                                      `json:"token"`
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
