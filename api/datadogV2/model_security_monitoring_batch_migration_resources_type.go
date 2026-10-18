// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourcesType The type of resource for a request on a set of batch migration resources.
type SecurityMonitoringBatchMigrationResourcesType string

// List of SecurityMonitoringBatchMigrationResourcesType.
const (
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTYPE_BATCH_MIGRATION_RESOURCES SecurityMonitoringBatchMigrationResourcesType = "batchMigrationResources"
)

var allowedSecurityMonitoringBatchMigrationResourcesTypeEnumValues = []SecurityMonitoringBatchMigrationResourcesType{
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTYPE_BATCH_MIGRATION_RESOURCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationResourcesType) GetAllowedValues() []SecurityMonitoringBatchMigrationResourcesType {
	return allowedSecurityMonitoringBatchMigrationResourcesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationResourcesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationResourcesType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationResourcesTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationResourcesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationResourcesTypeFromValue(v string) (*SecurityMonitoringBatchMigrationResourcesType, error) {
	ev := SecurityMonitoringBatchMigrationResourcesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationResourcesType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationResourcesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationResourcesType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationResourcesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationResourcesType value.
func (v SecurityMonitoringBatchMigrationResourcesType) Ptr() *SecurityMonitoringBatchMigrationResourcesType {
	return &v
}
