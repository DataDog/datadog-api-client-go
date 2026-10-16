// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishType The type of resource for a batch migration publish request.
type SecurityMonitoringBatchMigrationPublishType string

// List of SecurityMonitoringBatchMigrationPublishType.
const (
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHTYPE_BATCH_MIGRATION_PUBLISH SecurityMonitoringBatchMigrationPublishType = "batchMigrationPublish"
)

var allowedSecurityMonitoringBatchMigrationPublishTypeEnumValues = []SecurityMonitoringBatchMigrationPublishType{
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHTYPE_BATCH_MIGRATION_PUBLISH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationPublishType) GetAllowedValues() []SecurityMonitoringBatchMigrationPublishType {
	return allowedSecurityMonitoringBatchMigrationPublishTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationPublishType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationPublishType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationPublishTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationPublishType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationPublishTypeFromValue(v string) (*SecurityMonitoringBatchMigrationPublishType, error) {
	ev := SecurityMonitoringBatchMigrationPublishType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationPublishType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationPublishTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationPublishType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationPublishTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationPublishType value.
func (v SecurityMonitoringBatchMigrationPublishType) Ptr() *SecurityMonitoringBatchMigrationPublishType {
	return &v
}
