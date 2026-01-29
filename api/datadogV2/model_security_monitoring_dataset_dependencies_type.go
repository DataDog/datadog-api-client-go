// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDependenciesType The type of the response.
type SecurityMonitoringDatasetDependenciesType string

// List of SecurityMonitoringDatasetDependenciesType.
const (
	SECURITYMONITORINGDATASETDEPENDENCIESTYPE_SECURITY_MONITORING_DATASET_DEPENDENCIES SecurityMonitoringDatasetDependenciesType = "security_monitoring_dataset_dependencies"
)

var allowedSecurityMonitoringDatasetDependenciesTypeEnumValues = []SecurityMonitoringDatasetDependenciesType{
	SECURITYMONITORINGDATASETDEPENDENCIESTYPE_SECURITY_MONITORING_DATASET_DEPENDENCIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringDatasetDependenciesType) GetAllowedValues() []SecurityMonitoringDatasetDependenciesType {
	return allowedSecurityMonitoringDatasetDependenciesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringDatasetDependenciesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringDatasetDependenciesType(value)
	return nil
}

// NewSecurityMonitoringDatasetDependenciesTypeFromValue returns a pointer to a valid SecurityMonitoringDatasetDependenciesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringDatasetDependenciesTypeFromValue(v string) (*SecurityMonitoringDatasetDependenciesType, error) {
	ev := SecurityMonitoringDatasetDependenciesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringDatasetDependenciesType: valid values are %v", v, allowedSecurityMonitoringDatasetDependenciesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringDatasetDependenciesType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringDatasetDependenciesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringDatasetDependenciesType value.
func (v SecurityMonitoringDatasetDependenciesType) Ptr() *SecurityMonitoringDatasetDependenciesType {
	return &v
}
