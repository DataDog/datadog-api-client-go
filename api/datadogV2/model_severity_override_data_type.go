// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideDataType Severity override resource type.
type SeverityOverrideDataType string

// List of SeverityOverrideDataType.
const (
	SEVERITYOVERRIDEDATATYPE_SEVERITY_OVERRIDE SeverityOverrideDataType = "severity_override"
)

var allowedSeverityOverrideDataTypeEnumValues = []SeverityOverrideDataType{
	SEVERITYOVERRIDEDATATYPE_SEVERITY_OVERRIDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityOverrideDataType) GetAllowedValues() []SeverityOverrideDataType {
	return allowedSeverityOverrideDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityOverrideDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityOverrideDataType(value)
	return nil
}

// NewSeverityOverrideDataTypeFromValue returns a pointer to a valid SeverityOverrideDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityOverrideDataTypeFromValue(v string) (*SeverityOverrideDataType, error) {
	ev := SeverityOverrideDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityOverrideDataType: valid values are %v", v, allowedSeverityOverrideDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityOverrideDataType) IsValid() bool {
	for _, existing := range allowedSeverityOverrideDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityOverrideDataType value.
func (v SeverityOverrideDataType) Ptr() *SeverityOverrideDataType {
	return &v
}
