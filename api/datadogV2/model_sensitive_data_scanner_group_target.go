// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupTarget Datadog product whose scanning groups are served by the target-scoped group routes.
type SensitiveDataScannerGroupTarget string

// List of SensitiveDataScannerGroupTarget.
const (
	SENSITIVEDATASCANNERGROUPTARGET_LLM_OBSERVABILITY SensitiveDataScannerGroupTarget = "llm_observability"
)

var allowedSensitiveDataScannerGroupTargetEnumValues = []SensitiveDataScannerGroupTarget{
	SENSITIVEDATASCANNERGROUPTARGET_LLM_OBSERVABILITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerGroupTarget) GetAllowedValues() []SensitiveDataScannerGroupTarget {
	return allowedSensitiveDataScannerGroupTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerGroupTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerGroupTarget(value)
	return nil
}

// NewSensitiveDataScannerGroupTargetFromValue returns a pointer to a valid SensitiveDataScannerGroupTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerGroupTargetFromValue(v string) (*SensitiveDataScannerGroupTarget, error) {
	ev := SensitiveDataScannerGroupTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerGroupTarget: valid values are %v", v, allowedSensitiveDataScannerGroupTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerGroupTarget) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerGroupTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerGroupTarget value.
func (v SensitiveDataScannerGroupTarget) Ptr() *SensitiveDataScannerGroupTarget {
	return &v
}
