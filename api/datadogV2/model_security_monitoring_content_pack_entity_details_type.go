// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackEntityDetailsType Type for entity content pack details.
type SecurityMonitoringContentPackEntityDetailsType string

// List of SecurityMonitoringContentPackEntityDetailsType.
const (
	SECURITYMONITORINGCONTENTPACKENTITYDETAILSTYPE_ENTITY SecurityMonitoringContentPackEntityDetailsType = "entity"
)

var allowedSecurityMonitoringContentPackEntityDetailsTypeEnumValues = []SecurityMonitoringContentPackEntityDetailsType{
	SECURITYMONITORINGCONTENTPACKENTITYDETAILSTYPE_ENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackEntityDetailsType) GetAllowedValues() []SecurityMonitoringContentPackEntityDetailsType {
	return allowedSecurityMonitoringContentPackEntityDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackEntityDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackEntityDetailsType(value)
	return nil
}

// NewSecurityMonitoringContentPackEntityDetailsTypeFromValue returns a pointer to a valid SecurityMonitoringContentPackEntityDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackEntityDetailsTypeFromValue(v string) (*SecurityMonitoringContentPackEntityDetailsType, error) {
	ev := SecurityMonitoringContentPackEntityDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackEntityDetailsType: valid values are %v", v, allowedSecurityMonitoringContentPackEntityDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackEntityDetailsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackEntityDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackEntityDetailsType value.
func (v SecurityMonitoringContentPackEntityDetailsType) Ptr() *SecurityMonitoringContentPackEntityDetailsType {
	return &v
}
