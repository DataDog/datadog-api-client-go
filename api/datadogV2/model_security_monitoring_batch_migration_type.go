// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationType The type of resource for a batch migration.
type SecurityMonitoringBatchMigrationType string

// List of SecurityMonitoringBatchMigrationType.
const (
	SECURITYMONITORINGBATCHMIGRATIONTYPE_BATCH_MIGRATION SecurityMonitoringBatchMigrationType = "batchMigration"
)

var allowedSecurityMonitoringBatchMigrationTypeEnumValues = []SecurityMonitoringBatchMigrationType{
	SECURITYMONITORINGBATCHMIGRATIONTYPE_BATCH_MIGRATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationType) GetAllowedValues() []SecurityMonitoringBatchMigrationType {
	return allowedSecurityMonitoringBatchMigrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationTypeFromValue(v string) (*SecurityMonitoringBatchMigrationType, error) {
	ev := SecurityMonitoringBatchMigrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationType value.
func (v SecurityMonitoringBatchMigrationType) Ptr() *SecurityMonitoringBatchMigrationType {
	return &v
}
