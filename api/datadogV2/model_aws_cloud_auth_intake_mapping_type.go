// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthIntakeMappingType Type identifier for AWS cloud authentication intake mapping
type AWSCloudAuthIntakeMappingType string

// List of AWSCloudAuthIntakeMappingType.
const (
	AWSCLOUDAUTHINTAKEMAPPINGTYPE_AWS_CLOUD_AUTH_INTAKE_MAPPING AWSCloudAuthIntakeMappingType = "aws_cloud_auth_intake_mapping"
)

var allowedAWSCloudAuthIntakeMappingTypeEnumValues = []AWSCloudAuthIntakeMappingType{
	AWSCLOUDAUTHINTAKEMAPPINGTYPE_AWS_CLOUD_AUTH_INTAKE_MAPPING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSCloudAuthIntakeMappingType) GetAllowedValues() []AWSCloudAuthIntakeMappingType {
	return allowedAWSCloudAuthIntakeMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSCloudAuthIntakeMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSCloudAuthIntakeMappingType(value)
	return nil
}

// NewAWSCloudAuthIntakeMappingTypeFromValue returns a pointer to a valid AWSCloudAuthIntakeMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSCloudAuthIntakeMappingTypeFromValue(v string) (*AWSCloudAuthIntakeMappingType, error) {
	ev := AWSCloudAuthIntakeMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSCloudAuthIntakeMappingType: valid values are %v", v, allowedAWSCloudAuthIntakeMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSCloudAuthIntakeMappingType) IsValid() bool {
	for _, existing := range allowedAWSCloudAuthIntakeMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSCloudAuthIntakeMappingType value.
func (v AWSCloudAuthIntakeMappingType) Ptr() *AWSCloudAuthIntakeMappingType {
	return &v
}
