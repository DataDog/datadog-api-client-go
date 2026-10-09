// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideSetActionType The action that applies a manual severity override.
type SeverityOverrideSetActionType string

// List of SeverityOverrideSetActionType.
const (
	SEVERITYOVERRIDESETACTIONTYPE_SET SeverityOverrideSetActionType = "set"
)

var allowedSeverityOverrideSetActionTypeEnumValues = []SeverityOverrideSetActionType{
	SEVERITYOVERRIDESETACTIONTYPE_SET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityOverrideSetActionType) GetAllowedValues() []SeverityOverrideSetActionType {
	return allowedSeverityOverrideSetActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityOverrideSetActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityOverrideSetActionType(value)
	return nil
}

// NewSeverityOverrideSetActionTypeFromValue returns a pointer to a valid SeverityOverrideSetActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityOverrideSetActionTypeFromValue(v string) (*SeverityOverrideSetActionType, error) {
	ev := SeverityOverrideSetActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityOverrideSetActionType: valid values are %v", v, allowedSeverityOverrideSetActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityOverrideSetActionType) IsValid() bool {
	for _, existing := range allowedSeverityOverrideSetActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityOverrideSetActionType value.
func (v SeverityOverrideSetActionType) Ptr() *SeverityOverrideSetActionType {
	return &v
}
