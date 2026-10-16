// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishedResourceType The type of resource for a published batch migration resource.
type SecurityMonitoringBatchMigrationPublishedResourceType string

// List of SecurityMonitoringBatchMigrationPublishedResourceType.
const (
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCETYPE_BATCH_MIGRATION_PUBLISHED_RESOURCE SecurityMonitoringBatchMigrationPublishedResourceType = "batchMigrationPublishedResource"
)

var allowedSecurityMonitoringBatchMigrationPublishedResourceTypeEnumValues = []SecurityMonitoringBatchMigrationPublishedResourceType{
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCETYPE_BATCH_MIGRATION_PUBLISHED_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationPublishedResourceType) GetAllowedValues() []SecurityMonitoringBatchMigrationPublishedResourceType {
	return allowedSecurityMonitoringBatchMigrationPublishedResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationPublishedResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationPublishedResourceType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationPublishedResourceTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationPublishedResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationPublishedResourceTypeFromValue(v string) (*SecurityMonitoringBatchMigrationPublishedResourceType, error) {
	ev := SecurityMonitoringBatchMigrationPublishedResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationPublishedResourceType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationPublishedResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationPublishedResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationPublishedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationPublishedResourceType value.
func (v SecurityMonitoringBatchMigrationPublishedResourceType) Ptr() *SecurityMonitoringBatchMigrationPublishedResourceType {
	return &v
}
