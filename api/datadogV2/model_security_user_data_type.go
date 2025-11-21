// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityUserDataType Users resource type.
type SecurityUserDataType string

// List of SecurityUserDataType.
const (
	SECURITYUSERDATATYPE_USERS SecurityUserDataType = "users"
)

var allowedSecurityUserDataTypeEnumValues = []SecurityUserDataType{
	SECURITYUSERDATATYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityUserDataType) GetAllowedValues() []SecurityUserDataType {
	return allowedSecurityUserDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityUserDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityUserDataType(value)
	return nil
}

// NewSecurityUserDataTypeFromValue returns a pointer to a valid SecurityUserDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityUserDataTypeFromValue(v string) (*SecurityUserDataType, error) {
	ev := SecurityUserDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityUserDataType: valid values are %v", v, allowedSecurityUserDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityUserDataType) IsValid() bool {
	for _, existing := range allowedSecurityUserDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityUserDataType value.
func (v SecurityUserDataType) Ptr() *SecurityUserDataType {
	return &v
}
