// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountBearerTokenAuthResponse The bearer token authentication method configured on the account.
type GitHubCostsIntegrationAccountBearerTokenAuthResponse struct {
	// The authentication method type.
	AuthType GitHubCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationAccountBearerTokenAuthResponse instantiates a new GitHubCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationAccountBearerTokenAuthResponse(authType GitHubCostsIntegrationAccountBearerTokenAuthType) *GitHubCostsIntegrationAccountBearerTokenAuthResponse {
	this := GitHubCostsIntegrationAccountBearerTokenAuthResponse{}
	this.AuthType = authType
	return &this
}

// NewGitHubCostsIntegrationAccountBearerTokenAuthResponseWithDefaults instantiates a new GitHubCostsIntegrationAccountBearerTokenAuthResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationAccountBearerTokenAuthResponseWithDefaults() *GitHubCostsIntegrationAccountBearerTokenAuthResponse {
	this := GitHubCostsIntegrationAccountBearerTokenAuthResponse{}
	var authType GitHubCostsIntegrationAccountBearerTokenAuthType = GITHUBCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN
	this.AuthType = authType
	return &this
}

// GetAuthType returns the AuthType field value.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthResponse) GetAuthType() GitHubCostsIntegrationAccountBearerTokenAuthType {
	if o == nil {
		var ret GitHubCostsIntegrationAccountBearerTokenAuthType
		return ret
	}
	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthResponse) GetAuthTypeOk() (*GitHubCostsIntegrationAccountBearerTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthResponse) SetAuthType(v GitHubCostsIntegrationAccountBearerTokenAuthType) {
	o.AuthType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationAccountBearerTokenAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_type"] = o.AuthType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationAccountBearerTokenAuthResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthType *GitHubCostsIntegrationAccountBearerTokenAuthType `json:"auth_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthType == nil {
		return fmt.Errorf("required field auth_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthType.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthType = *all.AuthType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
