// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceStatus The status of a batch migration resource. `completed` resources have a translation in
// `output` that is ready to be published. `duplicate` is reserved for future duplicate rule
// detection.
type SecurityMonitoringBatchMigrationResourceStatus string

// List of SecurityMonitoringBatchMigrationResourceStatus.
const (
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PENDING     SecurityMonitoringBatchMigrationResourceStatus = "pending"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_IN_PROGRESS SecurityMonitoringBatchMigrationResourceStatus = "in_progress"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_COMPLETED   SecurityMonitoringBatchMigrationResourceStatus = "completed"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_CANCELED    SecurityMonitoringBatchMigrationResourceStatus = "canceled"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_FAILED      SecurityMonitoringBatchMigrationResourceStatus = "failed"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PUBLISHING  SecurityMonitoringBatchMigrationResourceStatus = "publishing"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PUBLISHED   SecurityMonitoringBatchMigrationResourceStatus = "published"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_DUPLICATE   SecurityMonitoringBatchMigrationResourceStatus = "duplicate"
)

var allowedSecurityMonitoringBatchMigrationResourceStatusEnumValues = []SecurityMonitoringBatchMigrationResourceStatus{
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PENDING,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_IN_PROGRESS,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_COMPLETED,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_CANCELED,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_FAILED,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PUBLISHING,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_PUBLISHED,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCESTATUS_DUPLICATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationResourceStatus) GetAllowedValues() []SecurityMonitoringBatchMigrationResourceStatus {
	return allowedSecurityMonitoringBatchMigrationResourceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationResourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationResourceStatus(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationResourceStatusFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationResourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationResourceStatusFromValue(v string) (*SecurityMonitoringBatchMigrationResourceStatus, error) {
	ev := SecurityMonitoringBatchMigrationResourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationResourceStatus: valid values are %v", v, allowedSecurityMonitoringBatchMigrationResourceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationResourceStatus) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationResourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationResourceStatus value.
func (v SecurityMonitoringBatchMigrationResourceStatus) Ptr() *SecurityMonitoringBatchMigrationResourceStatus {
	return &v
}
