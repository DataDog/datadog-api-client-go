// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerMachineLearningRuleType Type of a rule based on a standard machine learning rule.
type SensitiveDataScannerMachineLearningRuleType string

// List of SensitiveDataScannerMachineLearningRuleType.
const (
	SENSITIVEDATASCANNERMACHINELEARNINGRULETYPE_MACHINE_LEARNING_STANDARD_RULE SensitiveDataScannerMachineLearningRuleType = "machine_learning_standard_rule"
)

var allowedSensitiveDataScannerMachineLearningRuleTypeEnumValues = []SensitiveDataScannerMachineLearningRuleType{
	SENSITIVEDATASCANNERMACHINELEARNINGRULETYPE_MACHINE_LEARNING_STANDARD_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerMachineLearningRuleType) GetAllowedValues() []SensitiveDataScannerMachineLearningRuleType {
	return allowedSensitiveDataScannerMachineLearningRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerMachineLearningRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerMachineLearningRuleType(value)
	return nil
}

// NewSensitiveDataScannerMachineLearningRuleTypeFromValue returns a pointer to a valid SensitiveDataScannerMachineLearningRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerMachineLearningRuleTypeFromValue(v string) (*SensitiveDataScannerMachineLearningRuleType, error) {
	ev := SensitiveDataScannerMachineLearningRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerMachineLearningRuleType: valid values are %v", v, allowedSensitiveDataScannerMachineLearningRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerMachineLearningRuleType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerMachineLearningRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerMachineLearningRuleType value.
func (v SensitiveDataScannerMachineLearningRuleType) Ptr() *SensitiveDataScannerMachineLearningRuleType {
	return &v
}
