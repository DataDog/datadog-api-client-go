// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate Bearer token authentication. Use a project-scoped OpenAI API key with Write permission for model capabilities. While `openai-llm-completions` is enabled, a key that cannot call the OpenAI chat completions endpoint is rejected with a `422` response. Only the provided fields are changed. To keep the stored token, omit `token`.
type OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate struct {
	// The authentication method type.
	AuthType OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Secret token used to authenticate with OpenAI.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate instantiates a new OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate(authType OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate {
	this := OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdateWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdateWithDefaults() *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate {
	this := OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate{}
	var authType OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType = OPENAILLMEVALUATIONSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) GetAuthType() OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) GetAuthTypeOk() (*OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) SetAuthType(v OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) MarshalJSON() ([]byte, error) {
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
func (o *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
		Token    *string                                                    `json:"token,omitempty"`
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
