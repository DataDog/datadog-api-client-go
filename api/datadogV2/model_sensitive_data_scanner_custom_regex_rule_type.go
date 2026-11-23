// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerCustomRegexRuleType Type of a custom rule.
type SensitiveDataScannerCustomRegexRuleType string

// List of SensitiveDataScannerCustomRegexRuleType.
const (
	SENSITIVEDATASCANNERCUSTOMREGEXRULETYPE_REGEX_CUSTOM_RULE SensitiveDataScannerCustomRegexRuleType = "regex_custom_rule"
)

var allowedSensitiveDataScannerCustomRegexRuleTypeEnumValues = []SensitiveDataScannerCustomRegexRuleType{
	SENSITIVEDATASCANNERCUSTOMREGEXRULETYPE_REGEX_CUSTOM_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerCustomRegexRuleType) GetAllowedValues() []SensitiveDataScannerCustomRegexRuleType {
	return allowedSensitiveDataScannerCustomRegexRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerCustomRegexRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerCustomRegexRuleType(value)
	return nil
}

// NewSensitiveDataScannerCustomRegexRuleTypeFromValue returns a pointer to a valid SensitiveDataScannerCustomRegexRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerCustomRegexRuleTypeFromValue(v string) (*SensitiveDataScannerCustomRegexRuleType, error) {
	ev := SensitiveDataScannerCustomRegexRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerCustomRegexRuleType: valid values are %v", v, allowedSensitiveDataScannerCustomRegexRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerCustomRegexRuleType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerCustomRegexRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerCustomRegexRuleType value.
func (v SensitiveDataScannerCustomRegexRuleType) Ptr() *SensitiveDataScannerCustomRegexRuleType {
	return &v
}
