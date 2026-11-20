// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInlineSchema39 Test-only generated schema for the dashboard versioning proof of concept.
type DashboardInlineSchema39 string

// List of DashboardInlineSchema39.
const (
	DASHBOARDINLINESCHEMA39_SUM           DashboardInlineSchema39 = "sum"
	DASHBOARDINLINESCHEMA39_MIN           DashboardInlineSchema39 = "min"
	DASHBOARDINLINESCHEMA39_MAX           DashboardInlineSchema39 = "max"
	DASHBOARDINLINESCHEMA39_AVG           DashboardInlineSchema39 = "avg"
	DASHBOARDINLINESCHEMA39_COUNT         DashboardInlineSchema39 = "count"
	DASHBOARDINLINESCHEMA39_CARDINALITY   DashboardInlineSchema39 = "cardinality"
	DASHBOARDINLINESCHEMA39_MEDIAN        DashboardInlineSchema39 = "median"
	DASHBOARDINLINESCHEMA39_DELTA         DashboardInlineSchema39 = "delta"
	DASHBOARDINLINESCHEMA39_LATEST        DashboardInlineSchema39 = "latest"
	DASHBOARDINLINESCHEMA39_EARLIEST      DashboardInlineSchema39 = "earliest"
	DASHBOARDINLINESCHEMA39_MOST_FREQUENT DashboardInlineSchema39 = "most_frequent"
	DASHBOARDINLINESCHEMA39_ANY           DashboardInlineSchema39 = "any"
)

var allowedDashboardInlineSchema39EnumValues = []DashboardInlineSchema39{
	DASHBOARDINLINESCHEMA39_SUM,
	DASHBOARDINLINESCHEMA39_MIN,
	DASHBOARDINLINESCHEMA39_MAX,
	DASHBOARDINLINESCHEMA39_AVG,
	DASHBOARDINLINESCHEMA39_COUNT,
	DASHBOARDINLINESCHEMA39_CARDINALITY,
	DASHBOARDINLINESCHEMA39_MEDIAN,
	DASHBOARDINLINESCHEMA39_DELTA,
	DASHBOARDINLINESCHEMA39_LATEST,
	DASHBOARDINLINESCHEMA39_EARLIEST,
	DASHBOARDINLINESCHEMA39_MOST_FREQUENT,
	DASHBOARDINLINESCHEMA39_ANY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInlineSchema39) GetAllowedValues() []DashboardInlineSchema39 {
	return allowedDashboardInlineSchema39EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInlineSchema39) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInlineSchema39(value)
	return nil
}

// NewDashboardInlineSchema39FromValue returns a pointer to a valid DashboardInlineSchema39
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInlineSchema39FromValue(v string) (*DashboardInlineSchema39, error) {
	ev := DashboardInlineSchema39(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInlineSchema39: valid values are %v", v, allowedDashboardInlineSchema39EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInlineSchema39) IsValid() bool {
	for _, existing := range allowedDashboardInlineSchema39EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInlineSchema39 value.
func (v DashboardInlineSchema39) Ptr() *DashboardInlineSchema39 {
	return &v
}
