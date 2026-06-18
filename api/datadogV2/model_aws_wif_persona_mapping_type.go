// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsWifPersonaMappingType Type identifier for an AWS WIF persona mapping.
type AwsWifPersonaMappingType string

// List of AwsWifPersonaMappingType.
const (
	AWSWIFPERSONAMAPPINGTYPE_AWS_WIF_CONFIG AwsWifPersonaMappingType = "aws_wif_config"
)

var allowedAwsWifPersonaMappingTypeEnumValues = []AwsWifPersonaMappingType{
	AWSWIFPERSONAMAPPINGTYPE_AWS_WIF_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsWifPersonaMappingType) GetAllowedValues() []AwsWifPersonaMappingType {
	return allowedAwsWifPersonaMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsWifPersonaMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsWifPersonaMappingType(value)
	return nil
}

// NewAwsWifPersonaMappingTypeFromValue returns a pointer to a valid AwsWifPersonaMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsWifPersonaMappingTypeFromValue(v string) (*AwsWifPersonaMappingType, error) {
	ev := AwsWifPersonaMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsWifPersonaMappingType: valid values are %v", v, allowedAwsWifPersonaMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsWifPersonaMappingType) IsValid() bool {
	for _, existing := range allowedAwsWifPersonaMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsWifPersonaMappingType value.
func (v AwsWifPersonaMappingType) Ptr() *AwsWifPersonaMappingType {
	return &v
}
