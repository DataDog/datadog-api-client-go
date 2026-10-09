// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideClearActionType The action that removes a manual severity override.
type SeverityOverrideClearActionType string

// List of SeverityOverrideClearActionType.
const (
	SEVERITYOVERRIDECLEARACTIONTYPE_CLEAR SeverityOverrideClearActionType = "clear"
)

var allowedSeverityOverrideClearActionTypeEnumValues = []SeverityOverrideClearActionType{
	SEVERITYOVERRIDECLEARACTIONTYPE_CLEAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityOverrideClearActionType) GetAllowedValues() []SeverityOverrideClearActionType {
	return allowedSeverityOverrideClearActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityOverrideClearActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityOverrideClearActionType(value)
	return nil
}

// NewSeverityOverrideClearActionTypeFromValue returns a pointer to a valid SeverityOverrideClearActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityOverrideClearActionTypeFromValue(v string) (*SeverityOverrideClearActionType, error) {
	ev := SeverityOverrideClearActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityOverrideClearActionType: valid values are %v", v, allowedSeverityOverrideClearActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityOverrideClearActionType) IsValid() bool {
	for _, existing := range allowedSeverityOverrideClearActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityOverrideClearActionType value.
func (v SeverityOverrideClearActionType) Ptr() *SeverityOverrideClearActionType {
	return &v
}
