// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentSegmentType The type of a segment based on its data query configuration.
type RumSegmentSegmentType string

// List of RumSegmentSegmentType.
const (
	RUMSEGMENTSEGMENTTYPE_STATIC          RumSegmentSegmentType = "static"
	RUMSEGMENTSEGMENTTYPE_EVENT_PLATFORM  RumSegmentSegmentType = "event_platform"
	RUMSEGMENTSEGMENTTYPE_COMBINATION     RumSegmentSegmentType = "combination"
	RUMSEGMENTSEGMENTTYPE_JOURNEYS        RumSegmentSegmentType = "journeys"
	RUMSEGMENTSEGMENTTYPE_REFERENCE_TABLE RumSegmentSegmentType = "reference_table"
	RUMSEGMENTSEGMENTTYPE_TEMPLATES       RumSegmentSegmentType = "templates"
)

var allowedRumSegmentSegmentTypeEnumValues = []RumSegmentSegmentType{
	RUMSEGMENTSEGMENTTYPE_STATIC,
	RUMSEGMENTSEGMENTTYPE_EVENT_PLATFORM,
	RUMSEGMENTSEGMENTTYPE_COMBINATION,
	RUMSEGMENTSEGMENTTYPE_JOURNEYS,
	RUMSEGMENTSEGMENTTYPE_REFERENCE_TABLE,
	RUMSEGMENTSEGMENTTYPE_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentSegmentType) GetAllowedValues() []RumSegmentSegmentType {
	return allowedRumSegmentSegmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentSegmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentSegmentType(value)
	return nil
}

// NewRumSegmentSegmentTypeFromValue returns a pointer to a valid RumSegmentSegmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentSegmentTypeFromValue(v string) (*RumSegmentSegmentType, error) {
	ev := RumSegmentSegmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentSegmentType: valid values are %v", v, allowedRumSegmentSegmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentSegmentType) IsValid() bool {
	for _, existing := range allowedRumSegmentSegmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentSegmentType value.
func (v RumSegmentSegmentType) Ptr() *RumSegmentSegmentType {
	return &v
}
