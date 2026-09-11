// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OIDCDiscoveryDocument OpenID Connect provider metadata.
type OIDCDiscoveryDocument struct {
	// URL of the OAuth2 authorization endpoint.
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// Signing algorithms supported for ID tokens.
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	// URL identifying the OpenID Connect issuer.
	Issuer string `json:"issuer"`
	// URL of the JSON Web Key Set used to verify ID token signatures.
	JwksUri string `json:"jwks_uri"`
	// OAuth2 response types supported by the provider.
	ResponseTypesSupported []string `json:"response_types_supported"`
	// Subject identifier types supported by the provider.
	SubjectTypesSupported []string `json:"subject_types_supported"`
	// URL of the OAuth2 token endpoint.
	TokenEndpoint string `json:"token_endpoint"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOIDCDiscoveryDocument instantiates a new OIDCDiscoveryDocument object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOIDCDiscoveryDocument(authorizationEndpoint string, idTokenSigningAlgValuesSupported []string, issuer string, jwksUri string, responseTypesSupported []string, subjectTypesSupported []string, tokenEndpoint string) *OIDCDiscoveryDocument {
	this := OIDCDiscoveryDocument{}
	this.AuthorizationEndpoint = authorizationEndpoint
	this.IdTokenSigningAlgValuesSupported = idTokenSigningAlgValuesSupported
	this.Issuer = issuer
	this.JwksUri = jwksUri
	this.ResponseTypesSupported = responseTypesSupported
	this.SubjectTypesSupported = subjectTypesSupported
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewOIDCDiscoveryDocumentWithDefaults instantiates a new OIDCDiscoveryDocument object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOIDCDiscoveryDocumentWithDefaults() *OIDCDiscoveryDocument {
	this := OIDCDiscoveryDocument{}
	return &this
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value.
func (o *OIDCDiscoveryDocument) GetAuthorizationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint sets field value.
func (o *OIDCDiscoveryDocument) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = v
}

// GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field value.
func (o *OIDCDiscoveryDocument) GetIdTokenSigningAlgValuesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IdTokenSigningAlgValuesSupported
}

// GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetIdTokenSigningAlgValuesSupportedOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdTokenSigningAlgValuesSupported, true
}

// SetIdTokenSigningAlgValuesSupported sets field value.
func (o *OIDCDiscoveryDocument) SetIdTokenSigningAlgValuesSupported(v []string) {
	o.IdTokenSigningAlgValuesSupported = v
}

// GetIssuer returns the Issuer field value.
func (o *OIDCDiscoveryDocument) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value.
func (o *OIDCDiscoveryDocument) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwksUri returns the JwksUri field value.
func (o *OIDCDiscoveryDocument) GetJwksUri() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetJwksUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksUri, true
}

// SetJwksUri sets field value.
func (o *OIDCDiscoveryDocument) SetJwksUri(v string) {
	o.JwksUri = v
}

// GetResponseTypesSupported returns the ResponseTypesSupported field value.
func (o *OIDCDiscoveryDocument) GetResponseTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResponseTypesSupported
}

// GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetResponseTypesSupportedOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTypesSupported, true
}

// SetResponseTypesSupported sets field value.
func (o *OIDCDiscoveryDocument) SetResponseTypesSupported(v []string) {
	o.ResponseTypesSupported = v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value.
func (o *OIDCDiscoveryDocument) GetSubjectTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetSubjectTypesSupportedOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypesSupported, true
}

// SetSubjectTypesSupported sets field value.
func (o *OIDCDiscoveryDocument) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = v
}

// GetTokenEndpoint returns the TokenEndpoint field value.
func (o *OIDCDiscoveryDocument) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OIDCDiscoveryDocument) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value.
func (o *OIDCDiscoveryDocument) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OIDCDiscoveryDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	toSerialize["id_token_signing_alg_values_supported"] = o.IdTokenSigningAlgValuesSupported
	toSerialize["issuer"] = o.Issuer
	toSerialize["jwks_uri"] = o.JwksUri
	toSerialize["response_types_supported"] = o.ResponseTypesSupported
	toSerialize["subject_types_supported"] = o.SubjectTypesSupported
	toSerialize["token_endpoint"] = o.TokenEndpoint

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OIDCDiscoveryDocument) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthorizationEndpoint            *string   `json:"authorization_endpoint"`
		IdTokenSigningAlgValuesSupported *[]string `json:"id_token_signing_alg_values_supported"`
		Issuer                           *string   `json:"issuer"`
		JwksUri                          *string   `json:"jwks_uri"`
		ResponseTypesSupported           *[]string `json:"response_types_supported"`
		SubjectTypesSupported            *[]string `json:"subject_types_supported"`
		TokenEndpoint                    *string   `json:"token_endpoint"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthorizationEndpoint == nil {
		return fmt.Errorf("required field authorization_endpoint missing")
	}
	if all.IdTokenSigningAlgValuesSupported == nil {
		return fmt.Errorf("required field id_token_signing_alg_values_supported missing")
	}
	if all.Issuer == nil {
		return fmt.Errorf("required field issuer missing")
	}
	if all.JwksUri == nil {
		return fmt.Errorf("required field jwks_uri missing")
	}
	if all.ResponseTypesSupported == nil {
		return fmt.Errorf("required field response_types_supported missing")
	}
	if all.SubjectTypesSupported == nil {
		return fmt.Errorf("required field subject_types_supported missing")
	}
	if all.TokenEndpoint == nil {
		return fmt.Errorf("required field token_endpoint missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authorization_endpoint", "id_token_signing_alg_values_supported", "issuer", "jwks_uri", "response_types_supported", "subject_types_supported", "token_endpoint"})
	} else {
		return err
	}
	o.AuthorizationEndpoint = *all.AuthorizationEndpoint
	o.IdTokenSigningAlgValuesSupported = *all.IdTokenSigningAlgValuesSupported
	o.Issuer = *all.Issuer
	o.JwksUri = *all.JwksUri
	o.ResponseTypesSupported = *all.ResponseTypesSupported
	o.SubjectTypesSupported = *all.SubjectTypesSupported
	o.TokenEndpoint = *all.TokenEndpoint

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
