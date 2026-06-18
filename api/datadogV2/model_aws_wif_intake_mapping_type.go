// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsWifIntakeMappingType Type identifier for an AWS WIF intake mapping.
type AwsWifIntakeMappingType string

// List of AwsWifIntakeMappingType.
const (
	AWSWIFINTAKEMAPPINGTYPE_AWS_WIF_INTAKE_MAPPING AwsWifIntakeMappingType = "aws_wif_intake_mapping"
)

var allowedAwsWifIntakeMappingTypeEnumValues = []AwsWifIntakeMappingType{
	AWSWIFINTAKEMAPPINGTYPE_AWS_WIF_INTAKE_MAPPING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsWifIntakeMappingType) GetAllowedValues() []AwsWifIntakeMappingType {
	return allowedAwsWifIntakeMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsWifIntakeMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsWifIntakeMappingType(value)
	return nil
}

// NewAwsWifIntakeMappingTypeFromValue returns a pointer to a valid AwsWifIntakeMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsWifIntakeMappingTypeFromValue(v string) (*AwsWifIntakeMappingType, error) {
	ev := AwsWifIntakeMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsWifIntakeMappingType: valid values are %v", v, allowedAwsWifIntakeMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsWifIntakeMappingType) IsValid() bool {
	for _, existing := range allowedAwsWifIntakeMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsWifIntakeMappingType value.
func (v AwsWifIntakeMappingType) Ptr() *AwsWifIntakeMappingType {
	return &v
}
