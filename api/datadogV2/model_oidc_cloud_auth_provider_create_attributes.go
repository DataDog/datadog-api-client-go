// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCCloudAuthProviderCreateAttributes Attributes for creating a Custom OIDC provider.
type OIDCCloudAuthProviderCreateAttributes struct {
	// Expected token audiences. At least one audience is required. When `org_uuid_source` is `aud`, include datadog/{org_uuid} for the organization creating the provider.
	Audiences []string `json:"audiences"`
	// How the signing keys are discovered.
	DiscoveryType OIDCCloudAuthDiscoveryType `json:"discovery_type"`
	// Claim used by legacy identity_pattern mappings and in identity logs.
	IdentityClaimName *string `json:"identity_claim_name,omitempty"`
	// HTTPS issuer URL that must match the token issuer. For oidc_discovery, signing keys are located through the issuer discovery document.
	IssuerUrl string `json:"issuer_url"`
	// HTTPS JSON Web Key Set URL. Required when `discovery_type` is `direct_jwks`.
	JwksUrl *string `json:"jwks_url,omitempty"`
	// Optional name of the token claim containing the Datadog organization UUID.
	OrgUuidClaimName *string `json:"org_uuid_claim_name,omitempty"`
	// Source of the organization UUID in the signed token. The `aud` source accepts a datadog/{org_uuid} audience.
	OrgUuidSource *OIDCCloudAuthOrgUUIDSource `json:"org_uuid_source,omitempty"`
	// Name identifying the provider within the organization.
	ProviderName string `json:"provider_name"`
	// Claim matchers required for every token from this provider, in addition to the persona mapping matchers. Shared issuers can require an exact tenant claim.
	RequiredClaims []OIDCCloudAuthClaimMatcherCreate `json:"required_claims,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCCloudAuthProviderCreateAttributes instantiates a new OIDCCloudAuthProviderCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCCloudAuthProviderCreateAttributes(audiences []string, discoveryType OIDCCloudAuthDiscoveryType, issuerUrl string, providerName string) *OIDCCloudAuthProviderCreateAttributes {
	this := OIDCCloudAuthProviderCreateAttributes{}
	this.Audiences = audiences
	this.DiscoveryType = discoveryType
	var identityClaimName string = "sub"
	this.IdentityClaimName = &identityClaimName
	this.IssuerUrl = issuerUrl
	var orgUuidSource OIDCCloudAuthOrgUUIDSource = OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID
	this.OrgUuidSource = &orgUuidSource
	this.ProviderName = providerName
	return &this
}

// NewOIDCCloudAuthProviderCreateAttributesWithDefaults instantiates a new OIDCCloudAuthProviderCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCCloudAuthProviderCreateAttributesWithDefaults() *OIDCCloudAuthProviderCreateAttributes {
	this := OIDCCloudAuthProviderCreateAttributes{}
	var identityClaimName string = "sub"
	this.IdentityClaimName = &identityClaimName
	var orgUuidSource OIDCCloudAuthOrgUUIDSource = OIDCCLOUDAUTHORGUUIDSOURCE_DD_ORG_UUID
	this.OrgUuidSource = &orgUuidSource
	return &this
}

// GetAudiences returns the Audiences field value.
func (o *OIDCCloudAuthProviderCreateAttributes) GetAudiences() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetAudiencesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audiences, true
}

// SetAudiences sets field value.
func (o *OIDCCloudAuthProviderCreateAttributes) SetAudiences(v []string) {
	o.Audiences = v
}

// GetDiscoveryType returns the DiscoveryType field value.
func (o *OIDCCloudAuthProviderCreateAttributes) GetDiscoveryType() OIDCCloudAuthDiscoveryType {
	if o == nil {
		var ret OIDCCloudAuthDiscoveryType
		return ret
	}
	return o.DiscoveryType
}

// GetDiscoveryTypeOk returns a tuple with the DiscoveryType field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetDiscoveryTypeOk() (*OIDCCloudAuthDiscoveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryType, true
}

// SetDiscoveryType sets field value.
func (o *OIDCCloudAuthProviderCreateAttributes) SetDiscoveryType(v OIDCCloudAuthDiscoveryType) {
	o.DiscoveryType = v
}

// GetIdentityClaimName returns the IdentityClaimName field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderCreateAttributes) GetIdentityClaimName() string {
	if o == nil || o.IdentityClaimName == nil {
		var ret string
		return ret
	}
	return *o.IdentityClaimName
}

// GetIdentityClaimNameOk returns a tuple with the IdentityClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetIdentityClaimNameOk() (*string, bool) {
	if o == nil || o.IdentityClaimName == nil {
		return nil, false
	}
	return o.IdentityClaimName, true
}

// HasIdentityClaimName returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) HasIdentityClaimName() bool {
	return o != nil && o.IdentityClaimName != nil
}

// SetIdentityClaimName gets a reference to the given string and assigns it to the IdentityClaimName field.
func (o *OIDCCloudAuthProviderCreateAttributes) SetIdentityClaimName(v string) {
	o.IdentityClaimName = &v
}

// GetIssuerUrl returns the IssuerUrl field value.
func (o *OIDCCloudAuthProviderCreateAttributes) GetIssuerUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetIssuerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerUrl, true
}

// SetIssuerUrl sets field value.
func (o *OIDCCloudAuthProviderCreateAttributes) SetIssuerUrl(v string) {
	o.IssuerUrl = v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderCreateAttributes) GetJwksUrl() string {
	if o == nil || o.JwksUrl == nil {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetJwksUrlOk() (*string, bool) {
	if o == nil || o.JwksUrl == nil {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) HasJwksUrl() bool {
	return o != nil && o.JwksUrl != nil
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *OIDCCloudAuthProviderCreateAttributes) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

// GetOrgUuidClaimName returns the OrgUuidClaimName field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderCreateAttributes) GetOrgUuidClaimName() string {
	if o == nil || o.OrgUuidClaimName == nil {
		var ret string
		return ret
	}
	return *o.OrgUuidClaimName
}

// GetOrgUuidClaimNameOk returns a tuple with the OrgUuidClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetOrgUuidClaimNameOk() (*string, bool) {
	if o == nil || o.OrgUuidClaimName == nil {
		return nil, false
	}
	return o.OrgUuidClaimName, true
}

// HasOrgUuidClaimName returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) HasOrgUuidClaimName() bool {
	return o != nil && o.OrgUuidClaimName != nil
}

// SetOrgUuidClaimName gets a reference to the given string and assigns it to the OrgUuidClaimName field.
func (o *OIDCCloudAuthProviderCreateAttributes) SetOrgUuidClaimName(v string) {
	o.OrgUuidClaimName = &v
}

// GetOrgUuidSource returns the OrgUuidSource field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderCreateAttributes) GetOrgUuidSource() OIDCCloudAuthOrgUUIDSource {
	if o == nil || o.OrgUuidSource == nil {
		var ret OIDCCloudAuthOrgUUIDSource
		return ret
	}
	return *o.OrgUuidSource
}

// GetOrgUuidSourceOk returns a tuple with the OrgUuidSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetOrgUuidSourceOk() (*OIDCCloudAuthOrgUUIDSource, bool) {
	if o == nil || o.OrgUuidSource == nil {
		return nil, false
	}
	return o.OrgUuidSource, true
}

// HasOrgUuidSource returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) HasOrgUuidSource() bool {
	return o != nil && o.OrgUuidSource != nil
}

// SetOrgUuidSource gets a reference to the given OIDCCloudAuthOrgUUIDSource and assigns it to the OrgUuidSource field.
func (o *OIDCCloudAuthProviderCreateAttributes) SetOrgUuidSource(v OIDCCloudAuthOrgUUIDSource) {
	o.OrgUuidSource = &v
}

// GetProviderName returns the ProviderName field value.
func (o *OIDCCloudAuthProviderCreateAttributes) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value.
func (o *OIDCCloudAuthProviderCreateAttributes) SetProviderName(v string) {
	o.ProviderName = v
}

// GetRequiredClaims returns the RequiredClaims field value if set, zero value otherwise.
func (o *OIDCCloudAuthProviderCreateAttributes) GetRequiredClaims() []OIDCCloudAuthClaimMatcherCreate {
	if o == nil || o.RequiredClaims == nil {
		var ret []OIDCCloudAuthClaimMatcherCreate
		return ret
	}
	return o.RequiredClaims
}

// GetRequiredClaimsOk returns a tuple with the RequiredClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) GetRequiredClaimsOk() (*[]OIDCCloudAuthClaimMatcherCreate, bool) {
	if o == nil || o.RequiredClaims == nil {
		return nil, false
	}
	return &o.RequiredClaims, true
}

// HasRequiredClaims returns a boolean if a field has been set.
func (o *OIDCCloudAuthProviderCreateAttributes) HasRequiredClaims() bool {
	return o != nil && o.RequiredClaims != nil
}

// SetRequiredClaims gets a reference to the given []OIDCCloudAuthClaimMatcherCreate and assigns it to the RequiredClaims field.
func (o *OIDCCloudAuthProviderCreateAttributes) SetRequiredClaims(v []OIDCCloudAuthClaimMatcherCreate) {
	o.RequiredClaims = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCCloudAuthProviderCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["audiences"] = o.Audiences
	toSerialize["discovery_type"] = o.DiscoveryType
	if o.IdentityClaimName != nil {
		toSerialize["identity_claim_name"] = o.IdentityClaimName
	}
	toSerialize["issuer_url"] = o.IssuerUrl
	if o.JwksUrl != nil {
		toSerialize["jwks_url"] = o.JwksUrl
	}
	if o.OrgUuidClaimName != nil {
		toSerialize["org_uuid_claim_name"] = o.OrgUuidClaimName
	}
	if o.OrgUuidSource != nil {
		toSerialize["org_uuid_source"] = o.OrgUuidSource
	}
	toSerialize["provider_name"] = o.ProviderName
	if o.RequiredClaims != nil {
		toSerialize["required_claims"] = o.RequiredClaims
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCCloudAuthProviderCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Audiences         *[]string                         `json:"audiences"`
		DiscoveryType     *OIDCCloudAuthDiscoveryType       `json:"discovery_type"`
		IdentityClaimName *string                           `json:"identity_claim_name,omitempty"`
		IssuerUrl         *string                           `json:"issuer_url"`
		JwksUrl           *string                           `json:"jwks_url,omitempty"`
		OrgUuidClaimName  *string                           `json:"org_uuid_claim_name,omitempty"`
		OrgUuidSource     *OIDCCloudAuthOrgUUIDSource       `json:"org_uuid_source,omitempty"`
		ProviderName      *string                           `json:"provider_name"`
		RequiredClaims    []OIDCCloudAuthClaimMatcherCreate `json:"required_claims,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Audiences == nil {
		return fmt.Errorf("required field audiences missing")
	}
	if all.DiscoveryType == nil {
		return fmt.Errorf("required field discovery_type missing")
	}
	if all.IssuerUrl == nil {
		return fmt.Errorf("required field issuer_url missing")
	}
	if all.ProviderName == nil {
		return fmt.Errorf("required field provider_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audiences", "discovery_type", "identity_claim_name", "issuer_url", "jwks_url", "org_uuid_claim_name", "org_uuid_source", "provider_name", "required_claims"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Audiences = *all.Audiences
	if !all.DiscoveryType.IsValid() {
		hasInvalidField = true
	} else {
		o.DiscoveryType = *all.DiscoveryType
	}
	o.IdentityClaimName = all.IdentityClaimName
	o.IssuerUrl = *all.IssuerUrl
	o.JwksUrl = all.JwksUrl
	o.OrgUuidClaimName = all.OrgUuidClaimName
	if all.OrgUuidSource != nil && !all.OrgUuidSource.IsValid() {
		hasInvalidField = true
	} else {
		o.OrgUuidSource = all.OrgUuidSource
	}
	o.ProviderName = *all.ProviderName
	o.RequiredClaims = all.RequiredClaims

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
