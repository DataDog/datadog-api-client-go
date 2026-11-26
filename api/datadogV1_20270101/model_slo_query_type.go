// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloQueryType Test-only generated schema for the dashboard versioning proof of concept.
type SloQueryType string

// List of SloQueryType.
const (
	SLOQUERYTYPE_MONITOR    SloQueryType = "monitor"
	SLOQUERYTYPE_METRIC     SloQueryType = "metric"
	SLOQUERYTYPE_TIME_SLICE SloQueryType = "time_slice"
)

var allowedSloQueryTypeEnumValues = []SloQueryType{
	SLOQUERYTYPE_MONITOR,
	SLOQUERYTYPE_METRIC,
	SLOQUERYTYPE_TIME_SLICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloQueryType) GetAllowedValues() []SloQueryType {
	return allowedSloQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloQueryType(value)
	return nil
}

// NewSloQueryTypeFromValue returns a pointer to a valid SloQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloQueryTypeFromValue(v string) (*SloQueryType, error) {
	ev := SloQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloQueryType: valid values are %v", v, allowedSloQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloQueryType) IsValid() bool {
	for _, existing := range allowedSloQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloQueryType value.
func (v SloQueryType) Ptr() *SloQueryType {
	return &v
}
