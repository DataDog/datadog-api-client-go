// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthPersonaMappingAttributesResponse Stored Custom OIDC persona mapping attributes.
type OIDCCloudAuthPersonaMappingAttributesResponse struct {
	// Handle resolved when the mapping was created; not updated if the account handle changes.
	AccountIdentifier *string `json:"account_identifier,omitempty"`
	// UUID of the Datadog user or service account to authenticate as.
	AccountUuid *string `json:"account_uuid,omitempty"`
	// Every claim condition must match. Supply claim_matchers or identity_pattern, but not both.
	ClaimMatchers []OIDCCloudAuthClaimMatcher `json:"claim_matchers,omitempty"`
	// Time the configuration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UUID of the user who created the configuration.
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	// Legacy glob pattern matched against the provider identity claim. Supply this or claim_matchers, but not both. Prefer claim_matchers with exact mode for a single Kubernetes service account.
	IdentityPattern *string `json:"identity_pattern,omitempty"`
	// UUID of the registered Custom OIDC provider.
	ProviderId *uuid.UUID `json:"provider_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthPersonaMappingAttributesResponse instantiates a new OIDCCloudAuthPersonaMappingAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthPersonaMappingAttributesResponse() *OIDCCloudAuthPersonaMappingAttributesResponse {
	this := OIDCCloudAuthPersonaMappingAttributesResponse{}
	return &this
}

// NewOIDCCloudAuthPersonaMappingAttributesResponseWithDefaults instantiates a new OIDCCloudAuthPersonaMappingAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthPersonaMappingAttributesResponseWithDefaults() *OIDCCloudAuthPersonaMappingAttributesResponse {
	this := OIDCCloudAuthPersonaMappingAttributesResponse{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifier() string {
	if o == nil || o.AccountIdentifier == nil {
		var ret string
		return ret
	}
	return *o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifierOk() (*string, bool) {
	if o == nil || o.AccountIdentifier == nil {
		return nil, false
	}
	return o.AccountIdentifier, true
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasAccountIdentifier() bool {
	return o != nil && o.AccountIdentifier != nil
}

// SetAccountIdentifier gets a reference to the given string and assigns it to the AccountIdentifier field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetAccountIdentifier(v string) {
	o.AccountIdentifier = &v
}

// GetAccountUuid returns the AccountUuid field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetAccountUuid() string {
	if o == nil || o.AccountUuid == nil {
		var ret string
		return ret
	}
	return *o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetAccountUuidOk() (*string, bool) {
	if o == nil || o.AccountUuid == nil {
		return nil, false
	}
	return o.AccountUuid, true
}

// HasAccountUuid returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasAccountUuid() bool {
	return o != nil && o.AccountUuid != nil
}

// SetAccountUuid gets a reference to the given string and assigns it to the AccountUuid field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetAccountUuid(v string) {
	o.AccountUuid = &v
}

// GetClaimMatchers returns the ClaimMatchers field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetClaimMatchers() []OIDCCloudAuthClaimMatcher {
	if o == nil || o.ClaimMatchers == nil {
		var ret []OIDCCloudAuthClaimMatcher
		return ret
	}
	return o.ClaimMatchers
}

// GetClaimMatchersOk returns a tuple with the ClaimMatchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetClaimMatchersOk() (*[]OIDCCloudAuthClaimMatcher, bool) {
	if o == nil || o.ClaimMatchers == nil {
		return nil, false
	}
	return &o.ClaimMatchers, true
}

// HasClaimMatchers returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasClaimMatchers() bool {
	return o != nil && o.ClaimMatchers != nil
}

// SetClaimMatchers gets a reference to the given []OIDCCloudAuthClaimMatcher and assigns it to the ClaimMatchers field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetClaimMatchers(v []OIDCCloudAuthClaimMatcher) {
	o.ClaimMatchers = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetCreatedBy() uuid.UUID {
	if o == nil || o.CreatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given uuid.UUID and assigns it to the CreatedBy field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = &v
}

// GetIdentityPattern returns the IdentityPattern field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetIdentityPattern() string {
	if o == nil || o.IdentityPattern == nil {
		var ret string
		return ret
	}
	return *o.IdentityPattern
}

// GetIdentityPatternOk returns a tuple with the IdentityPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetIdentityPatternOk() (*string, bool) {
	if o == nil || o.IdentityPattern == nil {
		return nil, false
	}
	return o.IdentityPattern, true
}

// HasIdentityPattern returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasIdentityPattern() bool {
	return o != nil && o.IdentityPattern != nil
}

// SetIdentityPattern gets a reference to the given string and assigns it to the IdentityPattern field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetIdentityPattern(v string) {
	o.IdentityPattern = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetProviderId() uuid.UUID {
	if o == nil || o.ProviderId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) GetProviderIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) HasProviderId() bool {
	return o != nil && o.ProviderId != nil
}

// SetProviderId gets a reference to the given uuid.UUID and assigns it to the ProviderId field.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) SetProviderId(v uuid.UUID) {
	o.ProviderId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthPersonaMappingAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountIdentifier != nil {
		toSerialize["account_identifier"] = o.AccountIdentifier
	}
	if o.AccountUuid != nil {
		toSerialize["account_uuid"] = o.AccountUuid
	}
	if o.ClaimMatchers != nil {
		toSerialize["claim_matchers"] = o.ClaimMatchers
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.IdentityPattern != nil {
		toSerialize["identity_pattern"] = o.IdentityPattern
	}
	if o.ProviderId != nil {
		toSerialize["provider_id"] = o.ProviderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthPersonaMappingAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string                     `json:"account_identifier,omitempty"`
		AccountUuid       *string                     `json:"account_uuid,omitempty"`
		ClaimMatchers     []OIDCCloudAuthClaimMatcher `json:"claim_matchers,omitempty"`
		CreatedAt         *time.Time                  `json:"created_at,omitempty"`
		CreatedBy         *uuid.UUID                  `json:"created_by,omitempty"`
		IdentityPattern   *string                     `json:"identity_pattern,omitempty"`
		ProviderId        *uuid.UUID                  `json:"provider_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_identifier", "account_uuid", "claim_matchers", "created_at", "created_by", "identity_pattern", "provider_id"})
	} else {
		return err
	}
	o.AccountIdentifier = all.AccountIdentifier
	o.AccountUuid = all.AccountUuid
	o.ClaimMatchers = all.ClaimMatchers
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.IdentityPattern = all.IdentityPattern
	o.ProviderId = all.ProviderId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
