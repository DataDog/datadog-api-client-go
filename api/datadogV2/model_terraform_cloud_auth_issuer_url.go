// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthIssuerURL HCP Terraform workload identity token issuer. Use `https://app.terraform.io` for HCP Terraform in the US or `https://app.eu.terraform.io` for HCP Terraform in Europe.
type TerraformCloudAuthIssuerURL string

// List of TerraformCloudAuthIssuerURL.
const (
	TERRAFORMCLOUDAUTHISSUERURL_US TerraformCloudAuthIssuerURL = "https://app.terraform.io"
	TERRAFORMCLOUDAUTHISSUERURL_EU TerraformCloudAuthIssuerURL = "https://app.eu.terraform.io"
)

var allowedTerraformCloudAuthIssuerURLEnumValues = []TerraformCloudAuthIssuerURL{
	TERRAFORMCLOUDAUTHISSUERURL_US,
	TERRAFORMCLOUDAUTHISSUERURL_EU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TerraformCloudAuthIssuerURL) GetAllowedValues() []TerraformCloudAuthIssuerURL {
	return allowedTerraformCloudAuthIssuerURLEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TerraformCloudAuthIssuerURL) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TerraformCloudAuthIssuerURL(value)
	return nil
}

// NewTerraformCloudAuthIssuerURLFromValue returns a pointer to a valid TerraformCloudAuthIssuerURL
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTerraformCloudAuthIssuerURLFromValue(v string) (*TerraformCloudAuthIssuerURL, error) {
	ev := TerraformCloudAuthIssuerURL(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TerraformCloudAuthIssuerURL: valid values are %v", v, allowedTerraformCloudAuthIssuerURLEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TerraformCloudAuthIssuerURL) IsValid() bool {
	for _, existing := range allowedTerraformCloudAuthIssuerURLEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TerraformCloudAuthIssuerURL value.
func (v TerraformCloudAuthIssuerURL) Ptr() *TerraformCloudAuthIssuerURL {
	return &v
}
