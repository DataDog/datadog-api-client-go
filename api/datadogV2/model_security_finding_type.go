// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingType The type of security finding.
type SecurityFindingType string

// List of SecurityFindingType.
const (
	SECURITYFINDINGTYPE_VULNERABILITY         SecurityFindingType = "vulnerability"
	SECURITYFINDINGTYPE_CONFIGURATION_FINDING SecurityFindingType = "configuration_finding"
)

var allowedSecurityFindingTypeEnumValues = []SecurityFindingType{
	SECURITYFINDINGTYPE_VULNERABILITY,
	SECURITYFINDINGTYPE_CONFIGURATION_FINDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFindingType) GetAllowedValues() []SecurityFindingType {
	return allowedSecurityFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFindingType(value)
	return nil
}

// NewSecurityFindingTypeFromValue returns a pointer to a valid SecurityFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFindingTypeFromValue(v string) (*SecurityFindingType, error) {
	ev := SecurityFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFindingType: valid values are %v", v, allowedSecurityFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFindingType) IsValid() bool {
	for _, existing := range allowedSecurityFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFindingType value.
func (v SecurityFindingType) Ptr() *SecurityFindingType {
	return &v
}
