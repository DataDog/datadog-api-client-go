// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationType The type of investigation signal.
type SecurityMonitoringSignalInvestigationType string

// List of SecurityMonitoringSignalInvestigationType.
const (
	SECURITYMONITORINGSIGNALINVESTIGATIONTYPE_INVESTIGATION_SIGNAL SecurityMonitoringSignalInvestigationType = "investigation_signal"
)

var allowedSecurityMonitoringSignalInvestigationTypeEnumValues = []SecurityMonitoringSignalInvestigationType{
	SECURITYMONITORINGSIGNALINVESTIGATIONTYPE_INVESTIGATION_SIGNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalInvestigationType) GetAllowedValues() []SecurityMonitoringSignalInvestigationType {
	return allowedSecurityMonitoringSignalInvestigationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalInvestigationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalInvestigationType(value)
	return nil
}

// NewSecurityMonitoringSignalInvestigationTypeFromValue returns a pointer to a valid SecurityMonitoringSignalInvestigationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalInvestigationTypeFromValue(v string) (*SecurityMonitoringSignalInvestigationType, error) {
	ev := SecurityMonitoringSignalInvestigationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalInvestigationType: valid values are %v", v, allowedSecurityMonitoringSignalInvestigationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalInvestigationType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalInvestigationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalInvestigationType value.
func (v SecurityMonitoringSignalInvestigationType) Ptr() *SecurityMonitoringSignalInvestigationType {
	return &v
}
