// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentSource The source of a segment.
type RumSegmentSource string

// List of RumSegmentSource.
const (
	RUMSEGMENTSOURCE_USER_CREATED RumSegmentSource = "user_created"
	RUMSEGMENTSOURCE_INITIAL      RumSegmentSource = "initial"
)

var allowedRumSegmentSourceEnumValues = []RumSegmentSource{
	RUMSEGMENTSOURCE_USER_CREATED,
	RUMSEGMENTSOURCE_INITIAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumSegmentSource) GetAllowedValues() []RumSegmentSource {
	return allowedRumSegmentSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumSegmentSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumSegmentSource(value)
	return nil
}

// NewRumSegmentSourceFromValue returns a pointer to a valid RumSegmentSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumSegmentSourceFromValue(v string) (*RumSegmentSource, error) {
	ev := RumSegmentSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumSegmentSource: valid values are %v", v, allowedRumSegmentSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumSegmentSource) IsValid() bool {
	for _, existing := range allowedRumSegmentSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumSegmentSource value.
func (v RumSegmentSource) Ptr() *RumSegmentSource {
	return &v
}
