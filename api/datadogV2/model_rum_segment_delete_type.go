// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentDeleteType Type of the deleted segment resource.
type RumSegmentDeleteType string

// List of RumSegmentDeleteType.
const (
	RUMSEGMENTDELETETYPE_DELETED_SEGMENT RumSegmentDeleteType = "deleted_segment"
)

var allowedRumSegmentDeleteTypeEnumValues = []RumSegmentDeleteType{
	RUMSEGMENTDELETETYPE_DELETED_SEGMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentDeleteType) GetAllowedValues() []RumSegmentDeleteType {
	return allowedRumSegmentDeleteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentDeleteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentDeleteType(value)
	return nil
}

// NewRumSegmentDeleteTypeFromValue returns a pointer to a valid RumSegmentDeleteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentDeleteTypeFromValue(v string) (*RumSegmentDeleteType, error) {
	ev := RumSegmentDeleteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentDeleteType: valid values are %v", v, allowedRumSegmentDeleteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentDeleteType) IsValid() bool {
	for _, existing := range allowedRumSegmentDeleteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentDeleteType value.
func (v RumSegmentDeleteType) Ptr() *RumSegmentDeleteType {
	return &v
}
