// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationRiskLevel Estimated risk of a remediation step or proposed fix.
type RemediationRiskLevel string

// List of RemediationRiskLevel.
const (
	REMEDIATIONRISKLEVEL_LOW    RemediationRiskLevel = "low"
	REMEDIATIONRISKLEVEL_MEDIUM RemediationRiskLevel = "medium"
	REMEDIATIONRISKLEVEL_HIGH   RemediationRiskLevel = "high"
)

var allowedRemediationRiskLevelEnumValues = []RemediationRiskLevel{
	REMEDIATIONRISKLEVEL_LOW,
	REMEDIATIONRISKLEVEL_MEDIUM,
	REMEDIATIONRISKLEVEL_HIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationRiskLevel) GetAllowedValues() []RemediationRiskLevel {
	return allowedRemediationRiskLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationRiskLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationRiskLevel(value)
	return nil
}

// NewRemediationRiskLevelFromValue returns a pointer to a valid RemediationRiskLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationRiskLevelFromValue(v string) (*RemediationRiskLevel, error) {
	ev := RemediationRiskLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationRiskLevel: valid values are %v", v, allowedRemediationRiskLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationRiskLevel) IsValid() bool {
	for _, existing := range allowedRemediationRiskLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationRiskLevel value.
func (v RemediationRiskLevel) Ptr() *RemediationRiskLevel {
	return &v
}
