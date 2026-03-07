// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentResourceType Type of the segment resource.
type RumSegmentResourceType string

// List of RumSegmentResourceType.
const (
	RUMSEGMENTRESOURCETYPE_SEGMENT RumSegmentResourceType = "segment"
)

var allowedRumSegmentResourceTypeEnumValues = []RumSegmentResourceType{
	RUMSEGMENTRESOURCETYPE_SEGMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentResourceType) GetAllowedValues() []RumSegmentResourceType {
	return allowedRumSegmentResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentResourceType(value)
	return nil
}

// NewRumSegmentResourceTypeFromValue returns a pointer to a valid RumSegmentResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentResourceTypeFromValue(v string) (*RumSegmentResourceType, error) {
	ev := RumSegmentResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentResourceType: valid values are %v", v, allowedRumSegmentResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentResourceType) IsValid() bool {
	for _, existing := range allowedRumSegmentResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentResourceType value.
func (v RumSegmentResourceType) Ptr() *RumSegmentResourceType {
	return &v
}
