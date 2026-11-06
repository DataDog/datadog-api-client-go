// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthProviderAttributesResponse Stored Custom OIDC provider attributes.
type OIDCCloudAuthProviderAttributesResponse struct {
	// Expected token audiences. At least one audience is required. When `org_uuid_source` is `aud`, include datadog/{org_uuid} for the organization creating the provider.
	Audiences []string `json:"audiences,omitempty"`
	// Time the configuration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UUID of the user who created the configuration.
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	// How the signing keys are discovered.
	DiscoveryType *OIDCCloudAuthDiscoveryType `json:"discovery_type,omitempty"`
	// Claim used by legacy identity_pattern mappings and in identity logs.
	IdentityClaimName *string `json:"identity_claim_name,omitempty"`
	// HTTPS issuer URL that must match the token issuer. For oidc_discovery, signing keys are located through the issuer discovery document.
	IssuerUrl *string `json:"issuer_url,omitempty"`
	// HTTPS JSON Web Key Set URL. Required when `discovery_type` is `direct_jwks`.
	JwksUrl *string `json:"jwks_url,omitempty"`
	// Optional name of the token claim containing the Datadog organization UUID.
	OrgUuidClaimName *string `json:"org_uuid_claim_name,omitempty"`
	// Source of the organization UUID in the signed token. The `aud` source accepts a datadog/{org_uuid} audience.
	OrgUuidSource *OIDCCloudAuthOrgUUIDSource `json:"org_uuid_source,omitempty"`
	// Name identifying the provider within the organization.
	ProviderName *string `json:"provider_name,omitempty"`
	// Claim matchers required for every token from this provider, in addition to the persona mapping matchers. Shared issuers can require an exact tenant claim.
	RequiredClaims []OIDCCloudAuthClaimMatcher `json:"required_claims,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthProviderAttributesResponse instantiates a new OIDCCloudAuthProviderAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthProviderAttributesResponse() *OIDCCloudAuthProviderAttributesResponse {
	this := OIDCCloudAuthProviderAttributesResponse{}
	var identityClaimName string = "sub"
	this.IdentityClaimName = &identityClaimName
	var orgUuidSource OIDCCloudAuthOrgUUIDSource = OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID
	this.OrgUuidSource = &orgUuidSource
	return &this
}

// NewOIDCCloudAuthProviderAttributesResponseWithDefaults instantiates a new OIDCCloudAuthProviderAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthProviderAttributesResponseWithDefaults() *OIDCCloudAuthProviderAttributesResponse {
	this := OIDCCloudAuthProviderAttributesResponse{}
	var identityClaimName string = "sub"
	this.IdentityClaimName = &identityClaimName
	var orgUuidSource OIDCCloudAuthOrgUUIDSource = OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID
	this.OrgUuidSource = &orgUuidSource
	return &this
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetAudiences() []string {
	if o == nil || o.Audiences == nil {
		var ret []string
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetAudiencesOk() (*[]string, bool) {
	if o == nil || o.Audiences == nil {
		return nil, false
	}
	return &o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasAudiences() bool {
	return o != nil && o.Audiences != nil
}

// SetAudiences gets a reference to the given []string and assigns it to the Audiences field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetAudiences(v []string) {
	o.Audiences = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetCreatedBy() uuid.UUID {
	if o == nil || o.CreatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given uuid.UUID and assigns it to the CreatedBy field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = &v
}

// GetDiscoveryType returns the DiscoveryType field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetDiscoveryType() OIDCCloudAuthDiscoveryType {
	if o == nil || o.DiscoveryType == nil {
		var ret OIDCCloudAuthDiscoveryType
		return ret
	}
	return *o.DiscoveryType
}

// GetDiscoveryTypeOk returns a tuple with the DiscoveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetDiscoveryTypeOk() (*OIDCCloudAuthDiscoveryType, bool) {
	if o == nil || o.DiscoveryType == nil {
		return nil, false
	}
	return o.DiscoveryType, true
}

// HasDiscoveryType returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasDiscoveryType() bool {
	return o != nil && o.DiscoveryType != nil
}

// SetDiscoveryType gets a reference to the given OIDCCloudAuthDiscoveryType and assigns it to the DiscoveryType field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetDiscoveryType(v OIDCCloudAuthDiscoveryType) {
	o.DiscoveryType = &v
}

// GetIdentityClaimName returns the IdentityClaimName field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetIdentityClaimName() string {
	if o == nil || o.IdentityClaimName == nil {
		var ret string
		return ret
	}
	return *o.IdentityClaimName
}

// GetIdentityClaimNameOk returns a tuple with the IdentityClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetIdentityClaimNameOk() (*string, bool) {
	if o == nil || o.IdentityClaimName == nil {
		return nil, false
	}
	return o.IdentityClaimName, true
}

// HasIdentityClaimName returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasIdentityClaimName() bool {
	return o != nil && o.IdentityClaimName != nil
}

// SetIdentityClaimName gets a reference to the given string and assigns it to the IdentityClaimName field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetIdentityClaimName(v string) {
	o.IdentityClaimName = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasIssuerUrl() bool {
	return o != nil && o.IssuerUrl != nil
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetJwksUrl() string {
	if o == nil || o.JwksUrl == nil {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetJwksUrlOk() (*string, bool) {
	if o == nil || o.JwksUrl == nil {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasJwksUrl() bool {
	return o != nil && o.JwksUrl != nil
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

// GetOrgUuidClaimName returns the OrgUuidClaimName field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetOrgUuidClaimName() string {
	if o == nil || o.OrgUuidClaimName == nil {
		var ret string
		return ret
	}
	return *o.OrgUuidClaimName
}

// GetOrgUuidClaimNameOk returns a tuple with the OrgUuidClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetOrgUuidClaimNameOk() (*string, bool) {
	if o == nil || o.OrgUuidClaimName == nil {
		return nil, false
	}
	return o.OrgUuidClaimName, true
}

// HasOrgUuidClaimName returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasOrgUuidClaimName() bool {
	return o != nil && o.OrgUuidClaimName != nil
}

// SetOrgUuidClaimName gets a reference to the given string and assigns it to the OrgUuidClaimName field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetOrgUuidClaimName(v string) {
	o.OrgUuidClaimName = &v
}

// GetOrgUuidSource returns the OrgUuidSource field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetOrgUuidSource() OIDCCloudAuthOrgUUIDSource {
	if o == nil || o.OrgUuidSource == nil {
		var ret OIDCCloudAuthOrgUUIDSource
		return ret
	}
	return *o.OrgUuidSource
}

// GetOrgUuidSourceOk returns a tuple with the OrgUuidSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetOrgUuidSourceOk() (*OIDCCloudAuthOrgUUIDSource, bool) {
	if o == nil || o.OrgUuidSource == nil {
		return nil, false
	}
	return o.OrgUuidSource, true
}

// HasOrgUuidSource returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasOrgUuidSource() bool {
	return o != nil && o.OrgUuidSource != nil
}

// SetOrgUuidSource gets a reference to the given OIDCCloudAuthOrgUUIDSource and assigns it to the OrgUuidSource field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetOrgUuidSource(v OIDCCloudAuthOrgUUIDSource) {
	o.OrgUuidSource = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasProviderName() bool {
	return o != nil && o.ProviderName != nil
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRequiredClaims returns the RequiredClaims field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderAttributesResponse) GetRequiredClaims() []OIDCCloudAuthClaimMatcher {
	if o == nil || o.RequiredClaims == nil {
		var ret []OIDCCloudAuthClaimMatcher
		return ret
	}
	return o.RequiredClaims
}

// GetRequiredClaimsOk returns a tuple with the RequiredClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) GetRequiredClaimsOk() (*[]OIDCCloudAuthClaimMatcher, bool) {
	if o == nil || o.RequiredClaims == nil {
		return nil, false
	}
	return &o.RequiredClaims, true
}

// HasRequiredClaims returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderAttributesResponse) HasRequiredClaims() bool {
	return o != nil && o.RequiredClaims != nil
}

// SetRequiredClaims gets a reference to the given []OIDCCloudAuthClaimMatcher and assigns it to the RequiredClaims field.
func (o *OIDCCloudAuthProviderAttributesResponse) SetRequiredClaims(v []OIDCCloudAuthClaimMatcher) {
	o.RequiredClaims = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthProviderAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Audiences != nil {
		toSerialize["audiences"] = o.Audiences
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
	if o.DiscoveryType != nil {
		toSerialize["discovery_type"] = o.DiscoveryType
	}
	if o.IdentityClaimName != nil {
		toSerialize["identity_claim_name"] = o.IdentityClaimName
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	if o.JwksUrl != nil {
		toSerialize["jwks_url"] = o.JwksUrl
	}
	if o.OrgUuidClaimName != nil {
		toSerialize["org_uuid_claim_name"] = o.OrgUuidClaimName
	}
	if o.OrgUuidSource != nil {
		toSerialize["org_uuid_source"] = o.OrgUuidSource
	}
	if o.ProviderName != nil {
		toSerialize["provider_name"] = o.ProviderName
	}
	if o.RequiredClaims != nil {
		toSerialize["required_claims"] = o.RequiredClaims
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthProviderAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Audiences         []string                    `json:"audiences,omitempty"`
		CreatedAt         *time.Time                  `json:"created_at,omitempty"`
		CreatedBy         *uuid.UUID                  `json:"created_by,omitempty"`
		DiscoveryType     *OIDCCloudAuthDiscoveryType `json:"discovery_type,omitempty"`
		IdentityClaimName *string                     `json:"identity_claim_name,omitempty"`
		IssuerUrl         *string                     `json:"issuer_url,omitempty"`
		JwksUrl           *string                     `json:"jwks_url,omitempty"`
		OrgUuidClaimName  *string                     `json:"org_uuid_claim_name,omitempty"`
		OrgUuidSource     *OIDCCloudAuthOrgUUIDSource `json:"org_uuid_source,omitempty"`
		ProviderName      *string                     `json:"provider_name,omitempty"`
		RequiredClaims    []OIDCCloudAuthClaimMatcher `json:"required_claims,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audiences", "created_at", "created_by", "discovery_type", "identity_claim_name", "issuer_url", "jwks_url", "org_uuid_claim_name", "org_uuid_source", "provider_name", "required_claims"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Audiences = all.Audiences
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	if all.DiscoveryType != nil && !all.DiscoveryType.IsValid() {
		hasInvalidField = true
	} else {
		o.DiscoveryType = all.DiscoveryType
	}
	o.IdentityClaimName = all.IdentityClaimName
	o.IssuerUrl = all.IssuerUrl
	o.JwksUrl = all.JwksUrl
	o.OrgUuidClaimName = all.OrgUuidClaimName
	if all.OrgUuidSource != nil && !all.OrgUuidSource.IsValid() {
		hasInvalidField = true
	} else {
		o.OrgUuidSource = all.OrgUuidSource
	}
	o.ProviderName = all.ProviderName
	o.RequiredClaims = all.RequiredClaims

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
