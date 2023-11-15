// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTypes Security monitoring rule types.
type SecurityMonitoringRuleTypes string

// List of SecurityMonitoringRuleTypes.
const (
	SECURITYMONITORINGRULETYPES_APPLICATION_SECURITY         SecurityMonitoringRuleTypes = "application_security"
	SECURITYMONITORINGRULETYPES_LOG_DETECTION                SecurityMonitoringRuleTypes = "log_detection"
	SECURITYMONITORINGRULETYPES_CLOUD_CONFIGURATION          SecurityMonitoringRuleTypes = "cloud_configuration"
	SECURITYMONITORINGRULETYPES_INFRASTRUCTURE_CONFIGURATION SecurityMonitoringRuleTypes = "infrastructure_configuration"
	SECURITYMONITORINGRULETYPES_WORKLOAD_SECURITY            SecurityMonitoringRuleTypes = "workload_security"
	SECURITYMONITORINGRULETYPES_SIGNAL_CORRELATION           SecurityMonitoringRuleTypes = "signal_correlation"
	SECURITYMONITORINGRULETYPES_VULNERABILITY_MANAGEMENT     SecurityMonitoringRuleTypes = "vulnerability_management"
)

var allowedSecurityMonitoringRuleTypesEnumValues = []SecurityMonitoringRuleTypes{
	SECURITYMONITORINGRULETYPES_APPLICATION_SECURITY,
	SECURITYMONITORINGRULETYPES_LOG_DETECTION,
	SECURITYMONITORINGRULETYPES_CLOUD_CONFIGURATION,
	SECURITYMONITORINGRULETYPES_INFRASTRUCTURE_CONFIGURATION,
	SECURITYMONITORINGRULETYPES_WORKLOAD_SECURITY,
	SECURITYMONITORINGRULETYPES_SIGNAL_CORRELATION,
	SECURITYMONITORINGRULETYPES_VULNERABILITY_MANAGEMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleTypes) GetAllowedValues() []SecurityMonitoringRuleTypes {
	return allowedSecurityMonitoringRuleTypesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleTypes(value)
	return nil
}

// NewSecurityMonitoringRuleTypesFromValue returns a pointer to a valid SecurityMonitoringRuleTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleTypesFromValue(v string) (*SecurityMonitoringRuleTypes, error) {
	ev := SecurityMonitoringRuleTypes(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleTypes: valid values are %v", v, allowedSecurityMonitoringRuleTypesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleTypes) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleTypes value.
func (v SecurityMonitoringRuleTypes) Ptr() *SecurityMonitoringRuleTypes {
	return &v
}
