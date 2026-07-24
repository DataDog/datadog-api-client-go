// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackThreatIntelDetailsType Type for threat intelligence content pack details.
type SecurityMonitoringContentPackThreatIntelDetailsType string

// List of SecurityMonitoringContentPackThreatIntelDetailsType.
const (
	SECURITYMONITORINGCONTENTPACKTHREATINTELDETAILSTYPE_THREAT_INTEL SecurityMonitoringContentPackThreatIntelDetailsType = "threat_intel"
)

var allowedSecurityMonitoringContentPackThreatIntelDetailsTypeEnumValues = []SecurityMonitoringContentPackThreatIntelDetailsType{
	SECURITYMONITORINGCONTENTPACKTHREATINTELDETAILSTYPE_THREAT_INTEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackThreatIntelDetailsType) GetAllowedValues() []SecurityMonitoringContentPackThreatIntelDetailsType {
	return allowedSecurityMonitoringContentPackThreatIntelDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackThreatIntelDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackThreatIntelDetailsType(value)
	return nil
}

// NewSecurityMonitoringContentPackThreatIntelDetailsTypeFromValue returns a pointer to a valid SecurityMonitoringContentPackThreatIntelDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackThreatIntelDetailsTypeFromValue(v string) (*SecurityMonitoringContentPackThreatIntelDetailsType, error) {
	ev := SecurityMonitoringContentPackThreatIntelDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackThreatIntelDetailsType: valid values are %v", v, allowedSecurityMonitoringContentPackThreatIntelDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackThreatIntelDetailsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackThreatIntelDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackThreatIntelDetailsType value.
func (v SecurityMonitoringContentPackThreatIntelDetailsType) Ptr() *SecurityMonitoringContentPackThreatIntelDetailsType {
	return &v
}
