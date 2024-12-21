// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Severity The vulnerability severity.
type Severity string

// List of Severity.
const (
	SEVERITY_UNKNOWN  Severity = "Unknown"
	SEVERITY_NONE     Severity = "None"
	SEVERITY_LOW      Severity = "Low"
	SEVERITY_MEDIUM   Severity = "Medium"
	SEVERITY_HIGH     Severity = "High"
	SEVERITY_CRITICAL Severity = "Critical"
)

var allowedSeverityEnumValues = []Severity{
	SEVERITY_UNKNOWN,
	SEVERITY_NONE,
	SEVERITY_LOW,
	SEVERITY_MEDIUM,
	SEVERITY_HIGH,
	SEVERITY_CRITICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Severity) GetAllowedValues() []Severity {
	return allowedSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Severity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Severity(value)
	return nil
}

// NewSeverityFromValue returns a pointer to a valid Severity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityFromValue(v string) (*Severity, error) {
	ev := Severity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Severity: valid values are %v", v, allowedSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Severity) IsValid() bool {
	for _, existing := range allowedSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Severity value.
func (v Severity) Ptr() *Severity {
	return &v
}
