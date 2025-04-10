// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProcessSummaryType Type of process summary.
type ProcessSummaryType string

// List of ProcessSummaryType.
const (
	PROCESSSUMMARYTYPE_PROCESS ProcessSummaryType = "process"
)

var allowedProcessSummaryTypeEnumValues = []ProcessSummaryType{
	PROCESSSUMMARYTYPE_PROCESS,
}

// GetAllowedValues returns the list of possible values.
func (v *ProcessSummaryType) GetAllowedValues() []ProcessSummaryType {
	return allowedProcessSummaryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProcessSummaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProcessSummaryType(value)
	return nil
}

// NewProcessSummaryTypeFromValue returns a pointer to a valid ProcessSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProcessSummaryTypeFromValue(v string) (*ProcessSummaryType, error) {
	ev := ProcessSummaryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProcessSummaryType: valid values are %v", v, allowedProcessSummaryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProcessSummaryType) IsValid() bool {
	for _, existing := range allowedProcessSummaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProcessSummaryType value.
func (v ProcessSummaryType) Ptr() *ProcessSummaryType {
	return &v
}
