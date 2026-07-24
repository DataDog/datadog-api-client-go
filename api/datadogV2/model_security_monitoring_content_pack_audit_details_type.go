// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackAuditDetailsType Type for audit trail content pack details.
type SecurityMonitoringContentPackAuditDetailsType string

// List of SecurityMonitoringContentPackAuditDetailsType.
const (
	SECURITYMONITORINGCONTENTPACKAUDITDETAILSTYPE_AUDIT SecurityMonitoringContentPackAuditDetailsType = "audit"
)

var allowedSecurityMonitoringContentPackAuditDetailsTypeEnumValues = []SecurityMonitoringContentPackAuditDetailsType{
	SECURITYMONITORINGCONTENTPACKAUDITDETAILSTYPE_AUDIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackAuditDetailsType) GetAllowedValues() []SecurityMonitoringContentPackAuditDetailsType {
	return allowedSecurityMonitoringContentPackAuditDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackAuditDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackAuditDetailsType(value)
	return nil
}

// NewSecurityMonitoringContentPackAuditDetailsTypeFromValue returns a pointer to a valid SecurityMonitoringContentPackAuditDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackAuditDetailsTypeFromValue(v string) (*SecurityMonitoringContentPackAuditDetailsType, error) {
	ev := SecurityMonitoringContentPackAuditDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackAuditDetailsType: valid values are %v", v, allowedSecurityMonitoringContentPackAuditDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackAuditDetailsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackAuditDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackAuditDetailsType value.
func (v SecurityMonitoringContentPackAuditDetailsType) Ptr() *SecurityMonitoringContentPackAuditDetailsType {
	return &v
}
