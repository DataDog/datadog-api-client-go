// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceType The type of resource for a batch migration resource.
type SecurityMonitoringBatchMigrationResourceType string

// List of SecurityMonitoringBatchMigrationResourceType.
const (
	SECURITYMONITORINGBATCHMIGRATIONRESOURCETYPE_BATCH_MIGRATION_RESOURCE SecurityMonitoringBatchMigrationResourceType = "batchMigrationResource"
)

var allowedSecurityMonitoringBatchMigrationResourceTypeEnumValues = []SecurityMonitoringBatchMigrationResourceType{
	SECURITYMONITORINGBATCHMIGRATIONRESOURCETYPE_BATCH_MIGRATION_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationResourceType) GetAllowedValues() []SecurityMonitoringBatchMigrationResourceType {
	return allowedSecurityMonitoringBatchMigrationResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationResourceType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationResourceTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationResourceTypeFromValue(v string) (*SecurityMonitoringBatchMigrationResourceType, error) {
	ev := SecurityMonitoringBatchMigrationResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationResourceType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationResourceType value.
func (v SecurityMonitoringBatchMigrationResourceType) Ptr() *SecurityMonitoringBatchMigrationResourceType {
	return &v
}
