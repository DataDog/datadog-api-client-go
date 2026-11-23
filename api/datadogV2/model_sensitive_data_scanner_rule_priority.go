// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerRulePriority Severity of the issues found by the rule.
type SensitiveDataScannerRulePriority string

// List of SensitiveDataScannerRulePriority.
const (
	SENSITIVEDATASCANNERRULEPRIORITY_CRITICAL SensitiveDataScannerRulePriority = "critical"
	SENSITIVEDATASCANNERRULEPRIORITY_HIGH     SensitiveDataScannerRulePriority = "high"
	SENSITIVEDATASCANNERRULEPRIORITY_MEDIUM   SensitiveDataScannerRulePriority = "medium"
	SENSITIVEDATASCANNERRULEPRIORITY_LOW      SensitiveDataScannerRulePriority = "low"
	SENSITIVEDATASCANNERRULEPRIORITY_INFO     SensitiveDataScannerRulePriority = "info"
)

var allowedSensitiveDataScannerRulePriorityEnumValues = []SensitiveDataScannerRulePriority{
	SENSITIVEDATASCANNERRULEPRIORITY_CRITICAL,
	SENSITIVEDATASCANNERRULEPRIORITY_HIGH,
	SENSITIVEDATASCANNERRULEPRIORITY_MEDIUM,
	SENSITIVEDATASCANNERRULEPRIORITY_LOW,
	SENSITIVEDATASCANNERRULEPRIORITY_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerRulePriority) GetAllowedValues() []SensitiveDataScannerRulePriority {
	return allowedSensitiveDataScannerRulePriorityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerRulePriority) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerRulePriority(value)
	return nil
}

// NewSensitiveDataScannerRulePriorityFromValue returns a pointer to a valid SensitiveDataScannerRulePriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerRulePriorityFromValue(v string) (*SensitiveDataScannerRulePriority, error) {
	ev := SensitiveDataScannerRulePriority(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerRulePriority: valid values are %v", v, allowedSensitiveDataScannerRulePriorityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerRulePriority) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerRulePriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerRulePriority value.
func (v SensitiveDataScannerRulePriority) Ptr() *SensitiveDataScannerRulePriority {
	return &v
}
