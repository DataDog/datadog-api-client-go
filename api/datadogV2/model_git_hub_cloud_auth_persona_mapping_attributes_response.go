// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCloudAuthPersonaMappingAttributesResponse Attributes for GitHub cloud authentication persona mapping response.
type GitHubCloudAuthPersonaMappingAttributesResponse struct {
	// Datadog account identifier (email or handle) mapped to the GitHub principal.
	AccountIdentifier string `json:"account_identifier"`
	// Datadog account UUID.
	AccountUuid string `json:"account_uuid"`
	// GitHub Actions OIDC claims to match against. Each field is a regular expression.
	// The `sub` claim is required; all other claims are optional. A token matches only when
	// all provided patterns match simultaneously (AND semantics).
	ClaimMatchers GitHubOIDCClaimPatterns `json:"claim_matchers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGitHubCloudAuthPersonaMappingAttributesResponse instantiates a new GitHubCloudAuthPersonaMappingAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCloudAuthPersonaMappingAttributesResponse(accountIdentifier string, accountUuid string, claimMatchers GitHubOIDCClaimPatterns) *GitHubCloudAuthPersonaMappingAttributesResponse {
	this := GitHubCloudAuthPersonaMappingAttributesResponse{}
	this.AccountIdentifier = accountIdentifier
	this.AccountUuid = accountUuid
	this.ClaimMatchers = claimMatchers
	return &this
}

// NewGitHubCloudAuthPersonaMappingAttributesResponseWithDefaults instantiates a new GitHubCloudAuthPersonaMappingAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCloudAuthPersonaMappingAttributesResponseWithDefaults() *GitHubCloudAuthPersonaMappingAttributesResponse {
	this := GitHubCloudAuthPersonaMappingAttributesResponse{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetAccountUuid returns the AccountUuid field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetAccountUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value
// and a boolean to check if the value has been set.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetAccountUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUuid, true
}

// SetAccountUuid sets field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) SetAccountUuid(v string) {
	o.AccountUuid = v
}

// GetClaimMatchers returns the ClaimMatchers field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetClaimMatchers() GitHubOIDCClaimPatterns {
	if o == nil {
		var ret GitHubOIDCClaimPatterns
		return ret
	}
	return o.ClaimMatchers
}

// GetClaimMatchersOk returns a tuple with the ClaimMatchers field value
// and a boolean to check if the value has been set.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) GetClaimMatchersOk() (*GitHubOIDCClaimPatterns, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimMatchers, true
}

// SetClaimMatchers sets field value.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) SetClaimMatchers(v GitHubOIDCClaimPatterns) {
	o.ClaimMatchers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCloudAuthPersonaMappingAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	toSerialize["account_uuid"] = o.AccountUuid
	toSerialize["claim_matchers"] = o.ClaimMatchers
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCloudAuthPersonaMappingAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string                  `json:"account_identifier"`
		AccountUuid       *string                  `json:"account_uuid"`
		ClaimMatchers     *GitHubOIDCClaimPatterns `json:"claim_matchers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountIdentifier == nil {
		return fmt.Errorf("required field account_identifier missing")
	}
	if all.AccountUuid == nil {
		return fmt.Errorf("required field account_uuid missing")
	}
	if all.ClaimMatchers == nil {
		return fmt.Errorf("required field claim_matchers missing")
	}

	hasInvalidField := false
	o.AccountIdentifier = *all.AccountIdentifier
	o.AccountUuid = *all.AccountUuid
	if all.ClaimMatchers.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClaimMatchers = *all.ClaimMatchers

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
