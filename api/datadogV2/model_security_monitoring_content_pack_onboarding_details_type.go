// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackOnboardingDetailsType Type for onboarding content pack details.
type SecurityMonitoringContentPackOnboardingDetailsType string

// List of SecurityMonitoringContentPackOnboardingDetailsType.
const (
	SECURITYMONITORINGCONTENTPACKONBOARDINGDETAILSTYPE_ONBOARDING SecurityMonitoringContentPackOnboardingDetailsType = "onboarding"
)

var allowedSecurityMonitoringContentPackOnboardingDetailsTypeEnumValues = []SecurityMonitoringContentPackOnboardingDetailsType{
	SECURITYMONITORINGCONTENTPACKONBOARDINGDETAILSTYPE_ONBOARDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackOnboardingDetailsType) GetAllowedValues() []SecurityMonitoringContentPackOnboardingDetailsType {
	return allowedSecurityMonitoringContentPackOnboardingDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackOnboardingDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackOnboardingDetailsType(value)
	return nil
}

// NewSecurityMonitoringContentPackOnboardingDetailsTypeFromValue returns a pointer to a valid SecurityMonitoringContentPackOnboardingDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackOnboardingDetailsTypeFromValue(v string) (*SecurityMonitoringContentPackOnboardingDetailsType, error) {
	ev := SecurityMonitoringContentPackOnboardingDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackOnboardingDetailsType: valid values are %v", v, allowedSecurityMonitoringContentPackOnboardingDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackOnboardingDetailsType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackOnboardingDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackOnboardingDetailsType value.
func (v SecurityMonitoringContentPackOnboardingDetailsType) Ptr() *SecurityMonitoringContentPackOnboardingDetailsType {
	return &v
}
