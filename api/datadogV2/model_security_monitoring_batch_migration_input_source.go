// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationInputSource The origin of a batch migration. Always `batch_migration` for batch migrations
// created through this API.
type SecurityMonitoringBatchMigrationInputSource string

// List of SecurityMonitoringBatchMigrationInputSource.
const (
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_BATCH_MIGRATION SecurityMonitoringBatchMigrationInputSource = "batch_migration"
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_THREAT_HUNTING  SecurityMonitoringBatchMigrationInputSource = "threat_hunting"
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_USER            SecurityMonitoringBatchMigrationInputSource = "user"
)

var allowedSecurityMonitoringBatchMigrationInputSourceEnumValues = []SecurityMonitoringBatchMigrationInputSource{
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_BATCH_MIGRATION,
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_THREAT_HUNTING,
	SECURITYMONITORINGBATCHMIGRATIONINPUTSOURCE_USER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationInputSource) GetAllowedValues() []SecurityMonitoringBatchMigrationInputSource {
	return allowedSecurityMonitoringBatchMigrationInputSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationInputSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationInputSource(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationInputSourceFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationInputSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationInputSourceFromValue(v string) (*SecurityMonitoringBatchMigrationInputSource, error) {
	ev := SecurityMonitoringBatchMigrationInputSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationInputSource: valid values are %v", v, allowedSecurityMonitoringBatchMigrationInputSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationInputSource) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationInputSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationInputSource value.
func (v SecurityMonitoringBatchMigrationInputSource) Ptr() *SecurityMonitoringBatchMigrationInputSource {
	return &v
}
