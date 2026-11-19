// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingAttributesResponse Stored HCP Terraform persona mapping attributes.
type TerraformCloudAuthPersonaMappingAttributesResponse struct {
	// Handle of the mapped Datadog account, resolved when the mapping was created.
	AccountIdentifier string `json:"account_identifier"`
	// UUID of the mapped Datadog user or service account.
	AccountUuid string `json:"account_uuid"`
	// Audience to set as the `TFC_WORKLOAD_IDENTITY_AUDIENCE_DATADOG` environment variable in HCP Terraform. Datadog computes this value as `datadog/<organization UUID>`.
	Audience *string `json:"audience,omitempty"`
	// HCP Terraform workload identity token issuer. Empty for mappings created before the issuer was required.
	IssuerUrl *string `json:"issuer_url,omitempty"`
	// Pattern matched against the `sub` claim of HCP Terraform workload identity tokens.
	SubPattern string `json:"sub_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformCloudAuthPersonaMappingAttributesResponse instantiates a new TerraformCloudAuthPersonaMappingAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformCloudAuthPersonaMappingAttributesResponse(accountIdentifier string, accountUuid string, subPattern string) *TerraformCloudAuthPersonaMappingAttributesResponse {
	this := TerraformCloudAuthPersonaMappingAttributesResponse{}
	this.AccountIdentifier = accountIdentifier
	this.AccountUuid = accountUuid
	this.SubPattern = subPattern
	return &this
}

// NewTerraformCloudAuthPersonaMappingAttributesResponseWithDefaults instantiates a new TerraformCloudAuthPersonaMappingAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformCloudAuthPersonaMappingAttributesResponseWithDefaults() *TerraformCloudAuthPersonaMappingAttributesResponse {
	this := TerraformCloudAuthPersonaMappingAttributesResponse{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetAccountUuid returns the AccountUuid field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAccountUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountUuid
}

// GetAccountUuidOk returns a tuple with the AccountUuid field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAccountUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUuid, true
}

// SetAccountUuid sets field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) SetAccountUuid(v string) {
	o.AccountUuid = v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) HasAudience() bool {
	return o != nil && o.Audience != nil
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) HasIssuerUrl() bool {
	return o != nil && o.IssuerUrl != nil
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetSubPattern returns the SubPattern field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetSubPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubPattern
}

// GetSubPatternOk returns a tuple with the SubPattern field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) GetSubPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubPattern, true
}

// SetSubPattern sets field value.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) SetSubPattern(v string) {
	o.SubPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformCloudAuthPersonaMappingAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	toSerialize["account_uuid"] = o.AccountUuid
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	toSerialize["sub_pattern"] = o.SubPattern

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformCloudAuthPersonaMappingAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string `json:"account_identifier"`
		AccountUuid       *string `json:"account_uuid"`
		Audience          *string `json:"audience,omitempty"`
		IssuerUrl         *string `json:"issuer_url,omitempty"`
		SubPattern        *string `json:"sub_pattern"`
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
	if all.SubPattern == nil {
		return fmt.Errorf("required field sub_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_identifier", "account_uuid", "audience", "issuer_url", "sub_pattern"})
	} else {
		return err
	}
	o.AccountIdentifier = *all.AccountIdentifier
	o.AccountUuid = *all.AccountUuid
	o.Audience = all.Audience
	o.IssuerUrl = all.IssuerUrl
	o.SubPattern = *all.SubPattern

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
