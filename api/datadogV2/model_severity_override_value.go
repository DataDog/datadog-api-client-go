// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideValue Severity to apply to the findings.
// `info` sets the lowest severity the finding type allows.
type SeverityOverrideValue string

// List of SeverityOverrideValue.
const (
	SEVERITYOVERRIDEVALUE_CRITICAL SeverityOverrideValue = "critical"
	SEVERITYOVERRIDEVALUE_HIGH     SeverityOverrideValue = "high"
	SEVERITYOVERRIDEVALUE_MEDIUM   SeverityOverrideValue = "medium"
	SEVERITYOVERRIDEVALUE_LOW      SeverityOverrideValue = "low"
	SEVERITYOVERRIDEVALUE_INFO     SeverityOverrideValue = "info"
)

var allowedSeverityOverrideValueEnumValues = []SeverityOverrideValue{
	SEVERITYOVERRIDEVALUE_CRITICAL,
	SEVERITYOVERRIDEVALUE_HIGH,
	SEVERITYOVERRIDEVALUE_MEDIUM,
	SEVERITYOVERRIDEVALUE_LOW,
	SEVERITYOVERRIDEVALUE_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityOverrideValue) GetAllowedValues() []SeverityOverrideValue {
	return allowedSeverityOverrideValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityOverrideValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityOverrideValue(value)
	return nil
}

// NewSeverityOverrideValueFromValue returns a pointer to a valid SeverityOverrideValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityOverrideValueFromValue(v string) (*SeverityOverrideValue, error) {
	ev := SeverityOverrideValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityOverrideValue: valid values are %v", v, allowedSeverityOverrideValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityOverrideValue) IsValid() bool {
	for _, existing := range allowedSeverityOverrideValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityOverrideValue value.
func (v SeverityOverrideValue) Ptr() *SeverityOverrideValue {
	return &v
}
