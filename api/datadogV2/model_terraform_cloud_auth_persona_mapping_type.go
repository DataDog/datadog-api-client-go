// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformCloudAuthPersonaMappingType Type identifier for an HCP Terraform persona mapping.
type TerraformCloudAuthPersonaMappingType string

// List of TerraformCloudAuthPersonaMappingType.
const (
	TERRAFORMCLOUDAUTHPERSONAMAPPINGTYPE_TERRAFORM_CLOUD_AUTH_CONFIG TerraformCloudAuthPersonaMappingType = "terraform_cloud_auth_config"
)

var allowedTerraformCloudAuthPersonaMappingTypeEnumValues = []TerraformCloudAuthPersonaMappingType{
	TERRAFORMCLOUDAUTHPERSONAMAPPINGTYPE_TERRAFORM_CLOUD_AUTH_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TerraformCloudAuthPersonaMappingType) GetAllowedValues() []TerraformCloudAuthPersonaMappingType {
	return allowedTerraformCloudAuthPersonaMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TerraformCloudAuthPersonaMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TerraformCloudAuthPersonaMappingType(value)
	return nil
}

// NewTerraformCloudAuthPersonaMappingTypeFromValue returns a pointer to a valid TerraformCloudAuthPersonaMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTerraformCloudAuthPersonaMappingTypeFromValue(v string) (*TerraformCloudAuthPersonaMappingType, error) {
	ev := TerraformCloudAuthPersonaMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TerraformCloudAuthPersonaMappingType: valid values are %v", v, allowedTerraformCloudAuthPersonaMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TerraformCloudAuthPersonaMappingType) IsValid() bool {
	for _, existing := range allowedTerraformCloudAuthPersonaMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TerraformCloudAuthPersonaMappingType value.
func (v TerraformCloudAuthPersonaMappingType) Ptr() *TerraformCloudAuthPersonaMappingType {
	return &v
}
