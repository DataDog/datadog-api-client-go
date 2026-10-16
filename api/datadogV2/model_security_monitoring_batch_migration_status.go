// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationStatus The status of a batch migration. A batch migration is created as `draft` and moves to
// `in_progress` when started. It moves to `completed` once no resource is `pending` or
// `in_progress`, or to `canceled` when canceled.
type SecurityMonitoringBatchMigrationStatus string

// List of SecurityMonitoringBatchMigrationStatus.
const (
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_DRAFT       SecurityMonitoringBatchMigrationStatus = "draft"
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_IN_PROGRESS SecurityMonitoringBatchMigrationStatus = "in_progress"
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_CANCELED    SecurityMonitoringBatchMigrationStatus = "canceled"
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_COMPLETED   SecurityMonitoringBatchMigrationStatus = "completed"
)

var allowedSecurityMonitoringBatchMigrationStatusEnumValues = []SecurityMonitoringBatchMigrationStatus{
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_DRAFT,
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_IN_PROGRESS,
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_CANCELED,
	SECURITYMONITORINGBATCHMIGRATIONSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationStatus) GetAllowedValues() []SecurityMonitoringBatchMigrationStatus {
	return allowedSecurityMonitoringBatchMigrationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationStatus(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationStatusFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationStatusFromValue(v string) (*SecurityMonitoringBatchMigrationStatus, error) {
	ev := SecurityMonitoringBatchMigrationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationStatus: valid values are %v", v, allowedSecurityMonitoringBatchMigrationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationStatus) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationStatus value.
func (v SecurityMonitoringBatchMigrationStatus) Ptr() *SecurityMonitoringBatchMigrationStatus {
	return &v
}
