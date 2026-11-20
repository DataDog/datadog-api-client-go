// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupTargetType Type of target whose scanning groups are served by the target-scoped group routes. `llm_observability` is Agent Observability.
type SensitiveDataScannerGroupTargetType string

// List of SensitiveDataScannerGroupTargetType.
const (
	SENSITIVEDATASCANNERGROUPTARGETTYPE_LLM_OBSERVABILITY SensitiveDataScannerGroupTargetType = "llm_observability"
)

var allowedSensitiveDataScannerGroupTargetTypeEnumValues = []SensitiveDataScannerGroupTargetType{
	SENSITIVEDATASCANNERGROUPTARGETTYPE_LLM_OBSERVABILITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerGroupTargetType) GetAllowedValues() []SensitiveDataScannerGroupTargetType {
	return allowedSensitiveDataScannerGroupTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerGroupTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerGroupTargetType(value)
	return nil
}

// NewSensitiveDataScannerGroupTargetTypeFromValue returns a pointer to a valid SensitiveDataScannerGroupTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerGroupTargetTypeFromValue(v string) (*SensitiveDataScannerGroupTargetType, error) {
	ev := SensitiveDataScannerGroupTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerGroupTargetType: valid values are %v", v, allowedSensitiveDataScannerGroupTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerGroupTargetType) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerGroupTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerGroupTargetType value.
func (v SensitiveDataScannerGroupTargetType) Ptr() *SensitiveDataScannerGroupTargetType {
	return &v
}
