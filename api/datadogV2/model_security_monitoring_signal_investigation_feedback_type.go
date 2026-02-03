// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationFeedbackType The type of feedback.
type SecurityMonitoringSignalInvestigationFeedbackType string

// List of SecurityMonitoringSignalInvestigationFeedbackType.
const (
	SECURITYMONITORINGSIGNALINVESTIGATIONFEEDBACKTYPE_INVESTIGATION_FEEDBACK SecurityMonitoringSignalInvestigationFeedbackType = "investigation_feedback"
)

var allowedSecurityMonitoringSignalInvestigationFeedbackTypeEnumValues = []SecurityMonitoringSignalInvestigationFeedbackType{
	SECURITYMONITORINGSIGNALINVESTIGATIONFEEDBACKTYPE_INVESTIGATION_FEEDBACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalInvestigationFeedbackType) GetAllowedValues() []SecurityMonitoringSignalInvestigationFeedbackType {
	return allowedSecurityMonitoringSignalInvestigationFeedbackTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalInvestigationFeedbackType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalInvestigationFeedbackType(value)
	return nil
}

// NewSecurityMonitoringSignalInvestigationFeedbackTypeFromValue returns a pointer to a valid SecurityMonitoringSignalInvestigationFeedbackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalInvestigationFeedbackTypeFromValue(v string) (*SecurityMonitoringSignalInvestigationFeedbackType, error) {
	ev := SecurityMonitoringSignalInvestigationFeedbackType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalInvestigationFeedbackType: valid values are %v", v, allowedSecurityMonitoringSignalInvestigationFeedbackTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalInvestigationFeedbackType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalInvestigationFeedbackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalInvestigationFeedbackType value.
func (v SecurityMonitoringSignalInvestigationFeedbackType) Ptr() *SecurityMonitoringSignalInvestigationFeedbackType {
	return &v
}
