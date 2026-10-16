// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationPublishedResourceStatus The outcome of publishing a batch migration resource.
type SecurityMonitoringBatchMigrationPublishedResourceStatus string

// List of SecurityMonitoringBatchMigrationPublishedResourceStatus.
const (
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCESTATUS_PUBLISHED SecurityMonitoringBatchMigrationPublishedResourceStatus = "published"
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCESTATUS_FAILED    SecurityMonitoringBatchMigrationPublishedResourceStatus = "failed"
)

var allowedSecurityMonitoringBatchMigrationPublishedResourceStatusEnumValues = []SecurityMonitoringBatchMigrationPublishedResourceStatus{
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCESTATUS_PUBLISHED,
	SECURITYMONITORINGBATCHMIGRATIONPUBLISHEDRESOURCESTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationPublishedResourceStatus) GetAllowedValues() []SecurityMonitoringBatchMigrationPublishedResourceStatus {
	return allowedSecurityMonitoringBatchMigrationPublishedResourceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationPublishedResourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationPublishedResourceStatus(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationPublishedResourceStatusFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationPublishedResourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationPublishedResourceStatusFromValue(v string) (*SecurityMonitoringBatchMigrationPublishedResourceStatus, error) {
	ev := SecurityMonitoringBatchMigrationPublishedResourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationPublishedResourceStatus: valid values are %v", v, allowedSecurityMonitoringBatchMigrationPublishedResourceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationPublishedResourceStatus) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationPublishedResourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationPublishedResourceStatus value.
func (v SecurityMonitoringBatchMigrationPublishedResourceStatus) Ptr() *SecurityMonitoringBatchMigrationPublishedResourceStatus {
	return &v
}
