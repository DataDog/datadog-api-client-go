// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingCreateAttributes Attributes for creating an HCP Terraform persona mapping.
type TerraformCloudAuthPersonaMappingCreateAttributes struct {
	// Handle or email address of the Datadog user or service account that matching tokens authenticate as. Its permissions must not exceed the caller's.
	AccountIdentifier string `json:"account_identifier"`
	// HCP Terraform workload identity token issuer. Use `https://app.terraform.io` for HCP Terraform in the US or `https://app.eu.terraform.io` for HCP Terraform in Europe.
	IssuerUrl TerraformCloudAuthIssuerURL `json:"issuer_url"`
	// Pattern matched against the `sub` claim of HCP Terraform workload identity tokens. Use `organization:<org>:project:<project>:workspace:<workspace>:run_phase:<plan|apply>` for workspace runs, or `organization:<org>:module:<module>:operation:test_run` for module test runs. Segments are HCP Terraform names, not IDs.
	// Any segment except the organization can be `*` to match any value.
	SubPattern string `json:"sub_pattern"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTerraformCloudAuthPersonaMappingCreateAttributes instantiates a new TerraformCloudAuthPersonaMappingCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformCloudAuthPersonaMappingCreateAttributes(accountIdentifier string, issuerUrl TerraformCloudAuthIssuerURL, subPattern string) *TerraformCloudAuthPersonaMappingCreateAttributes {
	this := TerraformCloudAuthPersonaMappingCreateAttributes{}
	this.AccountIdentifier = accountIdentifier
	this.IssuerUrl = issuerUrl
	this.SubPattern = subPattern
	return &this
}

// NewTerraformCloudAuthPersonaMappingCreateAttributesWithDefaults instantiates a new TerraformCloudAuthPersonaMappingCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformCloudAuthPersonaMappingCreateAttributesWithDefaults() *TerraformCloudAuthPersonaMappingCreateAttributes {
	this := TerraformCloudAuthPersonaMappingCreateAttributes{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetIssuerUrl returns the IssuerUrl field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetIssuerUrl() TerraformCloudAuthIssuerURL {
	if o == nil {
		var ret TerraformCloudAuthIssuerURL
		return ret
	}
	return o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetIssuerUrlOk() (*TerraformCloudAuthIssuerURL, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerUrl, true
}

// SetIssuerUrl sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) SetIssuerUrl(v TerraformCloudAuthIssuerURL) {
	o.IssuerUrl = v
}

// GetSubPattern returns the SubPattern field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetSubPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubPattern
}

// GetSubPatternOk returns a tuple with the SubPattern field value
// and a boolean to check if the value has been set.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) GetSubPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubPattern, true
}

// SetSubPattern sets field value.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) SetSubPattern(v string) {
	o.SubPattern = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformCloudAuthPersonaMappingCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_identifier"] = o.AccountIdentifier
	toSerialize["issuer_url"] = o.IssuerUrl
	toSerialize["sub_pattern"] = o.SubPattern
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformCloudAuthPersonaMappingCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountIdentifier *string                      `json:"account_identifier"`
		IssuerUrl         *TerraformCloudAuthIssuerURL `json:"issuer_url"`
		SubPattern        *string                      `json:"sub_pattern"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountIdentifier == nil {
		return fmt.Errorf("required field account_identifier missing")
	}
	if all.IssuerUrl == nil {
		return fmt.Errorf("required field issuer_url missing")
	}
	if all.SubPattern == nil {
		return fmt.Errorf("required field sub_pattern missing")
	}

	hasInvalidField := false
	o.AccountIdentifier = *all.AccountIdentifier
	if !all.IssuerUrl.IsValid() {
		hasInvalidField = true
	} else {
		o.IssuerUrl = *all.IssuerUrl
	}
	o.SubPattern = *all.SubPattern

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
