// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthPersonaMappingCreateAttributes Attributes for creating a Custom OIDC persona mapping.
type OIDCCloudAuthPersonaMappingCreateAttributes struct {
	// UUID of the Datadog user or service account to authenticate as.
	AccountUuid string `json:"account_uuid"`
	// Every claim condition must match. Supply claim_matchers or identity_pattern, but not both.
	ClaimMatchers []OIDCCloudAuthClaimMatcherCreate `json:"claim_matchers,omitempty"`
	// Legacy glob pattern, limited to 500 UTF-8 bytes when creating a mapping, matched against the provider identity claim. Supply this or claim_matchers, but not both. Prefer claim_matchers with exact mode for a single Kubernetes service account.
	IdentityPattern *string `json:"identity_pattern,omitempty"`
	// UUID of the registered Custom OIDC provider.
	ProviderId uuid.UUID `json:"provider_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthPersonaMappingCreateAttributes instantiates a new OIDCCloudAuthPersonaMappingCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthPersonaMappingCreateAttributes(accountUuid string, providerId uuid.UUID) *OIDCCloudAuthPersonaMappingCreateAttributes {
	this := OIDCCloudAuthPersonaMappingCreateAttributes{}
	this.AccountUuid = accountUuid
	this.ProviderId = providerId
	return &this
}

// NewOIDCCloudAuthPersonaMappingCreateAttributesWithDefaults instantiates a new OIDCCloudAuthPersonaMappingCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthPersonaMappingCreateAttributesWithDefaults() *OIDCCloudAuthPersonaMappingCreateAttributes {
	this := OIDCCloudAuthPersonaMappingCreateAttributes{}
	return &this
}

// GetAccountUuid returns the AccountUuid field value.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetAccountUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetAccountUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUuid, true
}

// SetAccountUuid sets field value.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) SetAccountUuid(v string) {
	o.AccountUuid = v
}

// GetClaimMatchers returns the ClaimMatchers field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetClaimMatchers() []OIDCCloudAuthClaimMatcherCreate {
	if o == nil || o.ClaimMatchers == nil {
		var ret []OIDCCloudAuthClaimMatcherCreate
		return ret
	}
	return o.ClaimMatchers
}

// GetClaimMatchersOk returns a tuple with the ClaimMatchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetClaimMatchersOk() (*[]OIDCCloudAuthClaimMatcherCreate, bool) {
	if o == nil || o.ClaimMatchers == nil {
		return nil, false
	}
	return &o.ClaimMatchers, true
}

// HasClaimMatchers returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) HasClaimMatchers() bool {
	return o != nil && o.ClaimMatchers != nil
}

// SetClaimMatchers gets a reference to the given []OIDCCloudAuthClaimMatcherCreate and assigns it to the ClaimMatchers field.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) SetClaimMatchers(v []OIDCCloudAuthClaimMatcherCreate) {
	o.ClaimMatchers = v
}

// GetIdentityPattern returns the IdentityPattern field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetIdentityPattern() string {
	if o == nil || o.IdentityPattern == nil {
		var ret string
		return ret
	}
	return *o.IdentityPattern
}

// GetIdentityPatternOk returns a tuple with the IdentityPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetIdentityPatternOk() (*string, bool) {
	if o == nil || o.IdentityPattern == nil {
		return nil, false
	}
	return o.IdentityPattern, true
}

// HasIdentityPattern returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) HasIdentityPattern() bool {
	return o != nil && o.IdentityPattern != nil
}

// SetIdentityPattern gets a reference to the given string and assigns it to the IdentityPattern field.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) SetIdentityPattern(v string) {
	o.IdentityPattern = &v
}

// GetProviderId returns the ProviderId field value.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetProviderId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) GetProviderIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) SetProviderId(v uuid.UUID) {
	o.ProviderId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthPersonaMappingCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_uuid"] = o.AccountUuid
	if o.ClaimMatchers != nil {
		toSerialize["claim_matchers"] = o.ClaimMatchers
	}
	if o.IdentityPattern != nil {
		toSerialize["identity_pattern"] = o.IdentityPattern
	}
	toSerialize["provider_id"] = o.ProviderId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthPersonaMappingCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountUuid     *string                           `json:"account_uuid"`
		ClaimMatchers   []OIDCCloudAuthClaimMatcherCreate `json:"claim_matchers,omitempty"`
		IdentityPattern *string                           `json:"identity_pattern,omitempty"`
		ProviderId      *uuid.UUID                        `json:"provider_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountUuid == nil {
		return fmt.Errorf("required field account_uuid missing")
	}
	if all.ProviderId == nil {
		return fmt.Errorf("required field provider_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_uuid", "claim_matchers", "identity_pattern", "provider_id"})
	} else {
		return err
	}
	o.AccountUuid = *all.AccountUuid
	o.ClaimMatchers = all.ClaimMatchers
	o.IdentityPattern = all.IdentityPattern
	o.ProviderId = *all.ProviderId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
