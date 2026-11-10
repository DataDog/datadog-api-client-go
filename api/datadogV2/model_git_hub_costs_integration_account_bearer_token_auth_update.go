// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountBearerTokenAuthUpdate Bearer token authentication. For an enterprise, use a personal access token (classic) from an Enterprise Admin with the `manage_billing:enterprise` and `read:org` scopes. For an organization, use a fine-grained personal access token created by an organization owner, with the organization as resource owner and read-only `Administration` permission. Only the fields provided are changed; omit `token` to keep the stored one.
type GitHubCostsIntegrationAccountBearerTokenAuthUpdate struct {
	// The authentication method type.
	AuthType GitHubCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// Secret token used to authenticate with GitHub.
	Token *string `json:"token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountBearerTokenAuthUpdate instantiates a new GitHubCostsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountBearerTokenAuthUpdate(authType GitHubCostsIntegrationAccountBearerTokenAuthType) *GitHubCostsIntegrationAccountBearerTokenAuthUpdate {
	this := GitHubCostsIntegrationAccountBearerTokenAuthUpdate{}
	this.AuthType = authType
	return &this
}

// NewGitHubCostsIntegrationAccountBearerTokenAuthUpdateWithDefaults instantiates a new GitHubCostsIntegrationAccountBearerTokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountBearerTokenAuthUpdateWithDefaults() *GitHubCostsIntegrationAccountBearerTokenAuthUpdate {
	this := GitHubCostsIntegrationAccountBearerTokenAuthUpdate{}
	var authType GitHubCostsIntegrationAccountBearerTokenAuthType = GITHUBCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) GetAuthType() GitHubCostsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret GitHubCostsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) GetAuthTypeOk() (*GitHubCostsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) SetAuthType(v GitHubCostsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) HasToken() bool {
	return o != nil && o.Token != nil
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) SetToken(v string) {
	o.Token = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountBearerTokenAuthUpdate) MarshalJSON() ([]byte, error) {
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
func (o *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *GitHubCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
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
