// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesColorOrder Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesColorOrder string

// List of TimeseriesColorOrder.
const (
	TIMESERIESCOLORORDER_MONOTONIC TimeseriesColorOrder = "monotonic"
	TIMESERIESCOLORORDER_SHUFFLED  TimeseriesColorOrder = "shuffled"
)

var allowedTimeseriesColorOrderEnumValues = []TimeseriesColorOrder{
	TIMESERIESCOLORORDER_MONOTONIC,
	TIMESERIESCOLORORDER_SHUFFLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesColorOrder) GetAllowedValues() []TimeseriesColorOrder {
	return allowedTimeseriesColorOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesColorOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesColorOrder(value)
	return nil
}

// NewTimeseriesColorOrderFromValue returns a pointer to a valid TimeseriesColorOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesColorOrderFromValue(v string) (*TimeseriesColorOrder, error) {
	ev := TimeseriesColorOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesColorOrder: valid values are %v", v, allowedTimeseriesColorOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesColorOrder) IsValid() bool {
	for _, existing := range allowedTimeseriesColorOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesColorOrder value.
func (v TimeseriesColorOrder) Ptr() *TimeseriesColorOrder {
	return &v
}
