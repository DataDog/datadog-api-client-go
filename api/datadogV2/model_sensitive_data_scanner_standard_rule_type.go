// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerStandardRuleType Type of standard rule, based on how it detects sensitive data.
type SensitiveDataScannerStandardRuleType string

// List of SensitiveDataScannerStandardRuleType.
const (
	SENSITIVEDATASCANNERSTANDARDRULETYPE_REGEX_STANDARD_RULE            SensitiveDataScannerStandardRuleType = "regex_standard_rule"
	SENSITIVEDATASCANNERSTANDARDRULETYPE_MACHINE_LEARNING_STANDARD_RULE SensitiveDataScannerStandardRuleType = "machine_learning_standard_rule"
)

var allowedSensitiveDataScannerStandardRuleTypeEnumValues = []SensitiveDataScannerStandardRuleType{
	SENSITIVEDATASCANNERSTANDARDRULETYPE_REGEX_STANDARD_RULE,
	SENSITIVEDATASCANNERSTANDARDRULETYPE_MACHINE_LEARNING_STANDARD_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerStandardRuleType) GetAllowedValues() []SensitiveDataScannerStandardRuleType {
	return allowedSensitiveDataScannerStandardRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerStandardRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerStandardRuleType(value)
	return nil
}

// NewSensitiveDataScannerStandardRuleTypeFromValue returns a pointer to a valid SensitiveDataScannerStandardRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerStandardRuleTypeFromValue(v string) (*SensitiveDataScannerStandardRuleType, error) {
	ev := SensitiveDataScannerStandardRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerStandardRuleType: valid values are %v", v, allowedSensitiveDataScannerStandardRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerStandardRuleType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerStandardRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerStandardRuleType value.
func (v SensitiveDataScannerStandardRuleType) Ptr() *SensitiveDataScannerStandardRuleType {
	return &v
}
