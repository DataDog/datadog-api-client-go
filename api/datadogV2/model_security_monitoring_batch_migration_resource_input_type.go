// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringBatchMigrationResourceInputType The type of the source rule: `spl` for a Splunk SPL search or saved search definition,
// `natural_language` for a plain-English description of the behavior to detect, or `kql`
// for a Microsoft KQL query. SPL inputs must be self-contained, with search macros expanded.
type SecurityMonitoringBatchMigrationResourceInputType string

// List of SecurityMonitoringBatchMigrationResourceInputType.
const (
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_NATURAL_LANGUAGE SecurityMonitoringBatchMigrationResourceInputType = "natural_language"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_SPL              SecurityMonitoringBatchMigrationResourceInputType = "spl"
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_KQL              SecurityMonitoringBatchMigrationResourceInputType = "kql"
)

var allowedSecurityMonitoringBatchMigrationResourceInputTypeEnumValues = []SecurityMonitoringBatchMigrationResourceInputType{
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_NATURAL_LANGUAGE,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_SPL,
	SECURITYMONITORINGBATCHMIGRATIONRESOURCEINPUTTYPE_KQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringBatchMigrationResourceInputType) GetAllowedValues() []SecurityMonitoringBatchMigrationResourceInputType {
	return allowedSecurityMonitoringBatchMigrationResourceInputTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringBatchMigrationResourceInputType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringBatchMigrationResourceInputType(value)
	return nil
}

// NewSecurityMonitoringBatchMigrationResourceInputTypeFromValue returns a pointer to a valid SecurityMonitoringBatchMigrationResourceInputType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringBatchMigrationResourceInputTypeFromValue(v string) (*SecurityMonitoringBatchMigrationResourceInputType, error) {
	ev := SecurityMonitoringBatchMigrationResourceInputType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringBatchMigrationResourceInputType: valid values are %v", v, allowedSecurityMonitoringBatchMigrationResourceInputTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringBatchMigrationResourceInputType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringBatchMigrationResourceInputTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringBatchMigrationResourceInputType value.
func (v SecurityMonitoringBatchMigrationResourceInputType) Ptr() *SecurityMonitoringBatchMigrationResourceInputType {
	return &v
}
